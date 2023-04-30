package file

import (
	"fmt"
	"io/fs"
	"os/user"
	"syscall"
)

const (
	tFormat string = "Jan 02 15:04"
)

type Info struct {
	UID     string
	GID     string
	ModDate string
	Name    string
	Size    int64
}

func newInfo(path string, d fs.DirEntry, info fs.FileInfo) (Info, error){
	sys := info.Sys()
	if sys == nil {
		return Info{}, fmt.Errorf("failed to get file sys")
	}
	sysInfo, ok := sys.(*syscall.Stat_t)
	if !ok {
		return Info{}, fmt.Errorf("failed to get sys stats")
	}

	uid, err := user.LookupId(fmt.Sprint(sysInfo.Uid))
	if err != nil {
		return Info{}, fmt.Errorf("failed to lookup uid: %w", err)
	}

	gid, err := user.LookupGroupId(fmt.Sprint(sysInfo.Gid))
	if err != nil {
		return Info{}, fmt.Errorf("failed to lookup gid: %w", err)
	}

	return Info{
		UID:     uid.Username,
		GID:     gid.Name,
		ModDate: info.ModTime().Format(tFormat),
		Name:    info.Name(),
		Size:    info.Size(),
	}, nil
}

func GetFiles(f fs.FS) ([]Info, error) {
	var files []Info
	err := fs.WalkDir(f, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		
		info, err := fs.Stat(f, path)
		if err != nil {
			return fmt.Errorf("failed to get fs stats: %w", err)
		}

		switch {
		case d.IsDir() && path != ".":
			currentFile, err := newInfo(path, d, info)
			if err != nil {
				return fmt.Errorf("failed to create new info: %w", err)
			}

			files = append(files, currentFile)

			return fs.SkipDir
		case path == ".":
			return nil
		}

		currentFile, err := newInfo(path, d, info)
		if err != nil {
			return fmt.Errorf("failed to create new info: %w", err)
		}

		files = append(files, currentFile)

		return nil
	})
	if err != nil {
		return nil, err
	}
	return files, nil
}
