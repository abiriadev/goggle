package pkg

import (
	"fmt"
	"go/parser"
	"go/token"
)

func parse(path string) (Index, error) {
	fs := token.NewFileSet()

	astmap, err := parser.ParseDir(fs, path, nil, parser.ParseComments)

	if err != nil {
		return Index{}, err
	}

	for f, p := range astmap {
		fmt.Println(f, p.Name)
	}

	return Index{}, nil
}
