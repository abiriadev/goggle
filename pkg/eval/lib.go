package eval

import (
	"cmp"
	"reflect"
	"slices"

	"github.com/abiriadev/goggle/pkg/core"
	"github.com/abiriadev/goggle/pkg/index"
	"github.com/abiriadev/goggle/pkg/query"
	"github.com/hbollon/go-edlib"
)

func EvaluateArgs(args []core.Arg, query query.Query) core.Similarity {
	if len(query.Args) != len(args) {
		return core.Different
	}

	sarr := make([]core.Similarity, 0)

	for i, arg := range query.Args {
		if args[i].Type != arg {
			return core.Different
		}

		sarr = append(sarr,
			core.Similarity(
				Lev(arg, args[i].Name),
			),
		)
	}
}

func EvaluateFunc(item *core.FuncDef, query query.Query) core.Similarity {
	if query.Ret == item.Return && reflect.DeepEqual(query.Args, item.Args) {
		if query.Name == "" {
			return core.Equivalent
		} else {
			d, err := edlib.StringsSimilarity(query.Name, item.Name, edlib.Levenshtein)
			if err != nil {
				// unreachable error
				panic(err)
			}

			return core.Similarity(1 - d)
		}
	}
	return core.Different
}

// TODO: complete this function later
func EvaluateMethod(item *core.MethodDef, query query.Query) core.Similarity {
	return core.Different
}

func limitSlice[T any](slice []T, limit int) []T {
	if len(slice) < limit {
		return slice
	} else {
		return slice[:limit]
	}
}

func Query(index *index.Index, query query.Query, limit int) core.ResultSet {
	rs := core.NewResultSet()

	for _, item := range index.FuncItems {
		sim := EvaluateFunc(&item, query)

		rs.Results = append(rs.Results, item.ToResult(sim))
	}

	slices.SortFunc(rs.Results, func(a, b core.ResultItem) int {
		return cmp.Compare(a.Similarity, b.Similarity)
	})

	rs.Results = limitSlice(rs.Results, limit)

	return rs
}
