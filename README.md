# lsgo
Catgo is a simple command-line interface (CLI) app written in Go. It is a project that I developed to practice my Golang skills. The aim of Catgo is to replicate the basic functionality of the standard Unix cat utility, which reads files sequentially and writes them to standard output (stdout).

lsgo is a CLI app written in Go. The aim of lsgo is to replicate the basic functionality of the standard Unix `ls` utility, which lists file informations in given path to stdout.

## CLI Usage
To use lsgo, install the main package and run it directly in your terminal. Alternatively, you can install the catgo package and use it in your own Go projects.

Here's how to install CLI with `go` command:
```bash
$ go install github.com/TulgaCG/lsgo/cmd/lsgo
```

After installation you can use lsgo in your terminal with the following command: 
```bash
$ lsgo <flag> <path>
```

Example:
```bash
$ lsgo ~/go/bin
```
Output:
```bash
asmfmt catgo cmd dlv errcheck fillstruct godef goimports golangci-lint gomodifytags gopls gorename gotags guru iferr impl keyify lsgo motion revive staticcheck 

```

### Flags

```
Usage of lsgo:
  -a	do not ignore entries starting with
  -l	use a long listing format
```

Here's an example of how to use the list flag:
```bash
$ lsgo -l ~/go/pkg
```
Output:
```bash
tlgcngd tlgcngd 4096 Apr 24 01:26 mod
tlgcngd tlgcngd 4096 Apr 24 01:07 sumdb
```

## Package Usage
To use lsgo package in your own projects, you can use following command:
```bash
$ go get github.com/TulgaCG/lsgo/pkg/file
```
Here's an example of how to use the lsgo package in your Go code:
```go
package main

import (
	"os"

	"github.com/TulgaCG/lsgo/pkg/file"
)

func main() {
	arrayOfPaths := []string{
		"~/",
		"~/go/bin",
	}

	listOpts := file.ListOpts{
		List: true,
		ShowHidden: false,
	}

	file.List(os.Stdout, arrayOfPaths, listOpts)
}
```
This code will read informations of the files at the path $HOME, $HOME/go/bin, and output the the information to the stdout using long list formation.

Alternatively you can use `file.GetInfo` function to get informations of the files as an `Info` array.

Here's an example:
```go
	home, ok := os.LookupEnv("HOME")
	if !ok {
		log.Fatal("failed to get home directory")
	}
	path := filepath.Join(home, "go/bin")
	
	f := os.DirFS(path)
	arrayOfInfos, err := file.GetInfo(f)
	if err != nil {
		log.Fatal("failed to get file infos %w", err)
	}
```

## License
lsgo is licensed under the Apache-2.0 License. Please see the [LICENSE](https://github.com/TulgaCG/lsgo/blob/main/LICENSE) file for more information.

## Contributing
As this project is solely for my own learning purposes, I am not currently accepting contributions from external sources. However, if you have any suggestions or feedback, feel free to open an issue on the project's [GitHub page](https://github.com/TulgaCG/lsgo).