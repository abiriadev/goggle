package core

import (
	"fmt"
	"strings"

	"github.com/repeale/fp-go"
)

type ToSignature interface {
	Signature() string
}

func args(args []Arg) string {
	return fmt.Sprintf("(%s)", strings.Join(fp.Map(func(arg Arg) string {
		return fmt.Sprintf("%s %s", arg.Name, arg.Type)
	})(args), ", "))
}

func (f *FuncDef) Signature() string {
	return fmt.Sprintf("func %s.%s%s %s", f.PackageMame, f.Name, args(f.Args), f.Return)
}

func (ri *ResultItem) Signature() string {
	return ri.Sig
}

func (ri *ResultItem) SigWithComment() string {
	return fmt.Sprintf("%s\t// %s", ri.Sig, ri.Summary)
}
