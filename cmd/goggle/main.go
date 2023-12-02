package main

import (
	"flag"

	"github.com/abiriadev/goggle/pkg"
)

func pathArg() string {
	flag.Parse()
	f := flag.Arg(0)

	if f == "" {
		panic("provide path to target package")
	}

	return f
}

func main() {
	f := pathArg()

	pkg.Parse(f)
}
