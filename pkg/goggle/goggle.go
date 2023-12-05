package goggle

import (
	"github.com/abiriadev/goggle/pkg/index"
	"github.com/abiriadev/goggle/pkg/query"
	"github.com/samber/mo"
)

type Item = mo.Either[index.FuncDef, index.MethodDef]

type ResultItem struct {
	Summary string
	Sig     string
	Link    string
}

type ResultSet struct {
	Results []ResultItem
}

type Similarity float64

func EvaluateFunc(item *index.FuncDef, query query.Query) Similarity

func EvaluateMethod(item *index.MethodDef, query query.Query) Similarity

func Query(index *index.Index, query query.Query) ResultSet
