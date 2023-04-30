package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/TulgaCG/lsgo/pkg/file"
)

func main() {
	f := os.DirFS(filepath.Join(os.Getenv("HOME"), "github"))
	files, err := file.GetFiles(f)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
	
	for _, f := range files {
		fmt.Println(f.UID, f.GID, f.Size, f.ModDate, f.Name)
	}
}
