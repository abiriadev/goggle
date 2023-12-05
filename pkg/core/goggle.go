package core

import (
	"github.com/abiriadev/goggle/pkg/index"
	"github.com/abiriadev/goggle/pkg/query"
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

type ResultItem struct {
	Summary string
	Sig     string
	Link    string
}

type ResultSet struct {
	Results []ResultItem
}

type Similarity float64

func EvaluateFunc(item *FuncDef, query query.Query) Similarity

func EvaluateMethod(item *MethodDef, query query.Query) Similarity

func Query(index *index.Index, query query.Query) ResultSet
