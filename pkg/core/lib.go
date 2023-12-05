package core

import (
	"github.com/samber/mo"
)

type FuncDef struct {
	Pkg     string
	Name    string
	Args    []string
	Ret     string
	Summary string
	DocLink string
}

type MethodDef struct {
	Pkg      string
	Name     string
	Receiver string
	Args     []string
	Ret      string
	Summary  string
	DocLink  string
}

type Item = mo.Either[FuncDef, MethodDef]

type Similarity float64

type ResultItem struct {
	Similarity Similarity
	Sig        string
	Summary    string
	Link       string
}

type ResultSet struct {
	Results []ResultItem
}
