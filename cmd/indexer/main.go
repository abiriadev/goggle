package main

import (
	"flag"

	"github.com/abiriadev/goggle/pkg/index"
)

func main() {
	indexDestFileName := flag.String("o", "", "path to save index file")
	format := flag.String("f", "gob", "index format")

	if *format != "json" && *format != "gob" {
		panic("invalid format: only `json` and `gob` are supported")
	}

	flag.Parse()

	target := flag.Args()
	if len(target) == 0 {
		target = append(target, "std")
	}

	index, err := index.NewIndexer().IndexPackages(target)
	if err != nil {
		panic(err)
	}

	if *format == "json" {
		if *indexDestFileName == "" {
			*indexDestFileName = "index.json"
		}

		err = index.SaveJson(*indexDestFileName)
	} else {
		if *indexDestFileName == "" {
			*indexDestFileName = "index.gob"
		}

		err = index.Save(*indexDestFileName)
	}

	if err != nil {
		panic(err)
	}
}
