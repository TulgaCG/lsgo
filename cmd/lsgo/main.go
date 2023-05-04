package main

import (
	"flag"
	"os"

	"github.com/TulgaCG/lsgo/pkg/file"
)

var opts file.ListOpts

func init() {
	flag.BoolVar(&opts.ShowHidden, "a", false, "do not ignore entries starting with")
	flag.BoolVar(&opts.List, "l", false, "use a long listing format")
	flag.Parse()
}

func main() {
	file.List(os.Stdout, opts, flag.Args()...)
}
