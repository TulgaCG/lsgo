package file

import (
	"fmt"
	"io/fs"
	"os"
	"os/user"
	"path/filepath"
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

func (fi *Info) GetInfo() string {
	return fmt.Sprintf("%s %s %v %s %s", fi.UID, fi.GID, fi.Size, fi.ModDate, fi.Name)
}

func newInfo(info fs.FileInfo) (Info, error) {
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

		switch {
		case path == ".":
			return nil
		case d.IsDir():
			info, err := fs.Stat(f, path)
			if err != nil {
				return fmt.Errorf("failed to get dir fs stats: %w", err)
			}

			currentFile, err := newInfo(info)
			if err != nil {
				return fmt.Errorf("failed to create new info: %w", err)
			}

			files = append(files, currentFile)

			return fs.SkipDir
		case d.Type() == os.ModeSymlink:
			symPath := filepath.Join(fmt.Sprint(f), path)
			info, err := os.Lstat(symPath)
			if err != nil {
				return fmt.Errorf("failed to get stats of symlink: %w", err)
			}

			currentFile, err := newInfo(info)
			if err != nil {
				return fmt.Errorf("failed to create new info: %w", err)
			}
			currentFile.Name = fmt.Sprintf("%s -> %s", currentFile.Name, symPath)

			files = append(files, currentFile)
		default:
			info, err := fs.Stat(f, path)
			if err != nil {
				return fmt.Errorf("failed to get def fs stats: %w", err)
			}

			currentFile, err := newInfo(info)
			if err != nil {
				return fmt.Errorf("failed to create new info: %w", err)
			}
	
			files = append(files, currentFile)
		}

		return nil
	})
	if err != nil {
		return nil, err
	}
	return files, nil
}
