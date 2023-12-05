package index

import (
	"go/ast"
	"go/doc"

	"golang.org/x/tools/go/packages"
)

type Indexer struct{}

func NewIndexer() Indexer {
	return Indexer{}
}

func (indexer Indexer) IndexPackages(pkgsToIndex []string) (Index, error) {
	pkgs, err := packages.Load(
		&packages.Config{
			Mode: packages.NeedName |
				packages.NeedTypes |
				packages.NeedSyntax,
		},
		pkgsToIndex...,
	)
	if err != nil {
		return Index{}, nil
	}

	index := make([]FuncDef, 0)

	for _, pkg := range pkgs {
		d, err := doc.NewFromFiles(pkg.Fset, pkg.Syntax, pkg.PkgPath)
		if err != nil {
			continue
		}

		for _, f := range d.Funcs {
			r := f.Decl.Type.Results
			if r == nil || r.NumFields() != 1 {
				continue
			}

			v, b := r.List[0].Type.(*ast.Ident)
			if !b {
				continue
			}

			args := make([]string, 0)

			pl := f.Decl.Type.Params.List
			for _, arg := range pl {
				if arg != nil {
					v, b := r.List[0].Type.(*ast.Ident)
					if !b {
						continue
					}

					args = append(args, v.Name)
				}
			}

			index = append(index, FuncDef{
				Name: f.Name,
				Args: args,
				Ret:  v.Name,
			})
		}
	}

	return Index{index}, nil
}
