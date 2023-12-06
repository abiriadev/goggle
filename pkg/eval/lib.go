package eval

import (
	"cmp"
	"reflect"
	"slices"

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

// TODO: complete this function later
func EvaluateMethod(item *core.MethodDef, query query.Query) core.Similarity {
	return core.Different
}

func Query(index *index.Index, query query.Query) core.ResultSet {
	rs := core.NewResultSet()

	for _, item := range index.FuncItems {
		sim := EvaluateFunc(&item, query)

		rs.Results = append(rs.Results, item.ToResult(sim))
	}

	slices.SortFunc(rs.Results, func(a, b core.ResultItem) int {
		return cmp.Compare(a.Similarity, b.Similarity)
	})

	return rs
}
