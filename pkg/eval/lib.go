package eval

import (
	"reflect"

	"github.com/abiriadev/goggle/pkg/core"
	"github.com/abiriadev/goggle/pkg/index"
	"github.com/abiriadev/goggle/pkg/query"
)

func EvaluateFunc(item *core.FuncDef, query query.Query) core.Similarity {
	if query.Ret == item.Ret && reflect.DeepEqual(query.Args, item.Args) {
		return core.Equivalent
	}
	return core.Different
}

func EvaluateMethod(item *core.MethodDef, query query.Query) core.Similarity

func Query(index *index.Index, query query.Query) core.ResultSet {
	rs := core.NewResultSet()

	return rs
}
