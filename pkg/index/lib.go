package index

import (
	"fmt"
	"go/ast"
	"go/doc"
	"go/parser"
	"go/token"
)

func IndexPackage(path string) (Index, error) {
	fs := token.NewFileSet()
	astmap, err := parser.ParseDir(fs, path, nil, parser.ParseComments)

	if err != nil {
		return Index{}, err
	}

	index := make([]FuncDef, 0)

	for pn, p := range astmap {
		fmt.Println("package:", pn)

		d := doc.New(p, pn, 0)

		for _, f := range d.Funcs {
			r := f.Decl.Type.Results
			if r == nil || r.NumFields() != 1 {
				continue
			}

			v, b := r.List[0].Type.(*ast.Ident)
			if !b {
				continue
			}

			index = append(index, FuncDef{
				Name: f.Name,
				Ret:  v.Name,
			})
		}

	}

	return Index{index}, nil
}