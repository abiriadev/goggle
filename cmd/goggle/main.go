package main

import (
	"flag"
	"fmt"

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

	index, e := pkg.Parse(f)

	if e != nil {
		panic(e)
	}

	for _, fd := range index.Items {
		fmt.Println(fd.Name, "->", fd.Ret)
	}
}
