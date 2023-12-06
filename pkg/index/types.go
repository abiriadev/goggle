package index

import (
	"github.com/abiriadev/goggle/pkg/core"
	"github.com/repeale/fp-go"
)

// The central struct for all kind of indexing and searching
type Index struct {
	FuncItems   []core.FuncDef
	MethodItems []core.MethodDef
}

func NewIndex() Index {
	return Index{make([]core.FuncDef, 0), make([]core.MethodDef, 0)}
}

func ConcatSliceTo[T any](slices [][]T, dest []T) []T {
	for _, sl := range slices {
		dest = append(dest, sl...)
	}

	return dest
}

func ConcatSlice[T any](slices [][]T) []T {
	r := make([]T, 0)

	for _, sl := range slices {
		r = append(r, sl...)
	}

	return r
}

func MergeIndex(idxes []Index) Index {
	return Index{
		FuncItems: ConcatSlice(
			fp.Map(func(idx Index) []core.FuncDef {
				return idx.FuncItems
			})(idxes),
		),
		MethodItems: ConcatSlice(
			fp.Map(func(idx Index) []core.MethodDef {
				return idx.MethodItems
			})(idxes),
		),
	}
}
