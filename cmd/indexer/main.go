package main

import (
	"flag"

	"github.com/abiriadev/goggle/pkg/index"
	"github.com/kr/pretty"
	"golang.org/x/tools/go/packages"
)

func main() {
	indexDestFileName := flag.String("o", "./index.json", "path to save index file")

	flag.Parse()

	cfg := &packages.Config{Mode: packages.NeedFiles | packages.NeedSyntax}

	target := flag.Args()
	if len(target) == 0 {
		target = append(target, "std")
	}

	pkgs, err := packages.Load(cfg, target...)
	if err != nil {
		panic(err)
	}
	if packages.PrintErrors(pkgs) > 0 {
		panic(err)
	}

	idxes := make([]index.Index, 0)
	for _, pkg := range pkgs {
		pretty.Println(pkg)

		index, e := index.IndexPackage(".")
		if e != nil {
			panic(e)
		}

		idxes = append(idxes, index)
	}

	index := index.MergeIndex(idxes)

	index.Save(*indexDestFileName)
}
