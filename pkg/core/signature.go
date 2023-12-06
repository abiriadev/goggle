package core

import (
	"fmt"
	"strings"
)

type ToSignature interface {
	Signature() string
}

func args(args []string) string {
	return fmt.Sprintf("(%s)", strings.Join(args, ", "))
}

func (f *FuncDef) Signature() string {
	return fmt.Sprintf("func %s%s %s", f.Name, args(f.Args), f.Return)
}

func (ri *ResultItem) Signature() string {
	return ri.Sig
}

func (ri *ResultItem) SigWithComment() string {
	return fmt.Sprintf("%s\t// %s", ri.Sig, ri.Summary)
}
