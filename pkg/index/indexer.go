package index

import (
	"fmt"
	"go/ast"
	"go/doc"
	"strings"

	"github.com/abiriadev/goggle/pkg/core"
	"github.com/repeale/fp-go"
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

	index := NewIndex()

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

			args := make([]core.Arg, 0)

			// analyzeFuncDecl(f.Decl)

			pl := f.Decl.Type.Params.List
			for _, arg := range pl {
				if arg != nil {
					a := arg.Names
					if len(a) != 1 {
						continue
					}

					a1 := a[0]

					t, isIdent := arg.Type.(*ast.Ident)
					if !isIdent {
						continue
					}

					args = append(args, core.Arg{
						Name: a1.Name,
						Type: t.Name,
					})
				}
			}

			var docLinkRoute string
			if docLinkRoute = d.ImportPath; isVendored(d.ImportPath) {
				docLinkRoute = strings.TrimPrefix(d.ImportPath, "vendor/")
			}

			index.FuncItems = append(index.FuncItems, core.FuncDef{
				Package:     d.ImportPath,
				PackageMame: pkg.Name,
				Name:        f.Name,
				Args:        args,
				Return:      v.Name,
				Summary:     d.Synopsis(f.Doc),
				Link:        fmt.Sprintf("https://pkg.go.dev/%s#%s", docLinkRoute, f.Name),
			})
		}
	}

	return index, nil
}

func analyzeFuncDecl(f *ast.FuncDecl) {
	// WARN: nilable
	// name := f.Name.Name

	// WARN: nilable
	ft := f.Type

	for _, field := range fp.Filter(func(n *ast.Field) bool {
		return n != nil
	})(ft.Params.List) {
		fty := field.Type

		fmt.Printf("fty: %T\n", fty)
	}
}

func isVendored(importPath string) bool {
	return strings.HasPrefix(importPath, "vendor/")
	// || strings.Contains(importPath, "/vendor/")
}
