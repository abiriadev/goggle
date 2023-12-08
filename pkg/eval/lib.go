package eval

import (
	"cmp"
	"slices"

	"github.com/abiriadev/goggle/pkg/core"
	"github.com/abiriadev/goggle/pkg/index"
	"github.com/abiriadev/goggle/pkg/query"
	"github.com/hbollon/go-edlib"
)

func Lev(a, b string) core.Similarity {
	s, err := edlib.StringsSimilarity(a, b, edlib.Levenshtein)
	if err != nil {
		// unreachable error
		panic(err)
	}

	return core.Similarity(1 - s)
}

func AccSim(sims []core.Similarity) core.Similarity {
	acc := core.Equivalent

	for _, sim := range sims {
		acc += sim
	}

	return acc / core.Similarity(len(sims))
}

func CombSim(a, b core.Similarity) core.Similarity {
	return (a + b) / 2
}

func EvaluateName(ident string, query string) core.Similarity {
	if query == "" {
		return core.Equivalent
	} else {
		return Lev(query, ident)
	}
}

func EvaluateArgs(args []core.Arg, query query.Query) core.Similarity {
	if len(query.Args) != len(args) {
		return core.Different
	}

	sarr := make([]core.Similarity, len(args))

	for i, arg := range query.Args {
		if args[i].Type != arg {
			return core.Different
		}

		sarr = append(sarr,
			EvaluateName(arg, args[i].Name),
		)
	}

	return AccSim(sarr)
}

func EvaluateFunc(item *core.FuncDef, query query.Query) core.Similarity {
	if query.Ret == item.Return {
		argsSim := EvaluateArgs(item.Args, query)

		namesSim := EvaluateName(item.Name, query.Name)

		return CombSim(argsSim, namesSim)
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
