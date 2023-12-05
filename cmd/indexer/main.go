package main

import (
	"flag"

	"github.com/abiriadev/goggle/pkg/index"
)

func main() {
	indexDestFileName := flag.String("o", "./index.json", "path to save index file")

	flag.Parse()

	f := flag.Arg(0)

	if f == "" {
		panic("provide path to target package")
	}

	index, e := index.IndexPackage(f)
	if e != nil {
		panic(e)
	}

	index.Save(*indexDestFileName)
}
