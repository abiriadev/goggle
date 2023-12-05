package index

import (
	"fmt"
	"go/ast"
	"go/doc"
	"strings"

	"github.com/abiriadev/goggle/pkg/core"
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

	index := make([]core.FuncDef, 0)

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

			var docLinkRoute string
			if docLinkRoute = d.ImportPath; isVendored(d.ImportPath) {
				docLinkRoute = strings.TrimPrefix(d.ImportPath, "vendor/")
			}

			index = append(index, core.FuncDef{
				Pkg:     d.ImportPath,
				Name:    f.Name,
				Args:    args,
				Ret:     v.Name,
				Summary: f.Doc,
				DocLink: fmt.Sprintf("https://pkg.go.dev/%s#%s", docLinkRoute, f.Name),
			})
		}
	}

	return Index{index}, nil
}

func isVendored(importPath string) bool {
	return strings.HasPrefix(importPath, "vendor/")
	// || strings.Contains(importPath, "/vendor/")
}
