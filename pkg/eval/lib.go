package eval

import (
	"github.com/abiriadev/goggle/pkg/core"
	"github.com/abiriadev/goggle/pkg/index"
	"github.com/abiriadev/goggle/pkg/query"
)

func EvaluateFunc(item *core.FuncDef, query query.Query) core.Similarity

func EvaluateMethod(item *core.MethodDef, query query.Query) core.Similarity

func Query(index *index.Index, query query.Query) core.ResultSet
