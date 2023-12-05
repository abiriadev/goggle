package main

import (
	"flag"

	"github.com/abiriadev/goggle/pkg/index"
)

func main() {
	indexDestFileName := flag.String("o", "./index.json", "path to save index file")

	flag.Parse()

	target := flag.Args()
	if len(target) == 0 {
		target = append(target, "std")
	}

	index, err := index.NewIndexer().IndexPackages(target)
	if err != nil {
		panic(err)
	}

	index.Save(*indexDestFileName)
}
