package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/TulgaCG/lsgo/pkg/file"
)

func init() {
	flag.Parse()
}

func main() {
	for _, arg := range flag.Args() {
		if strings.Contains(arg, "~/") {
			strings.Trim(arg, "~/")
			arg = filepath.Join(os.Getenv("HOME"), arg)
		}

		f := os.DirFS(arg)
		files, err := file.GetFiles(f)
		if err != nil {
			fmt.Fprintf(os.Stderr, "failed to get files from the path %s: %v", arg, err)
		}

		for _, f := range files {
			fmt.Fprintln(os.Stdout, f.UID, f.GID, f.Size, f.ModDate, f.Name)
		}
	}
}
