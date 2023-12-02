package pkg

import (
	"fmt"
	"go/doc"
	"go/parser"
	"go/token"
)

func Parse(path string) (Index, error) {
	fs := token.NewFileSet()

	astmap, err := parser.ParseDir(fs, path, nil, parser.ParseComments)

	if err != nil {
		return Index{}, err
	}

	for pn, p := range astmap {
		fmt.Println("package:", pn)

		d := doc.New(p, pn, 0)

		for _, f := range d.Funcs {
			fmt.Println("\t", f.Name)
		}
	}

	return Index{}, nil
}
