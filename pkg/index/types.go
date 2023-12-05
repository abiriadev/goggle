package index

import (
	"html/template"
	"strings"

	"github.com/repeale/fp-go"
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

// The central struct for all kind of indexing and searching
type Index struct {
	FuncItems   []FuncDef
	MethodItems []MethodDef
}

func NewIndex() Index {
	return Index{make([]FuncDef, 0), make([]MethodDef, 0)}
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
			fp.Map(func(idx Index) []FuncDef {
				return idx.FuncItems
			})(idxes),
		),
		MethodItems: ConcatSlice(
			fp.Map(func(idx Index) []MethodDef {
				return idx.MethodItems
			})(idxes),
		),
	}
}

func (this *FuncDef) String() string {
	t := template.Must(
		template.New("").Parse(`func {{.Name}}({{range .Args}}{{.}}, {{end}}) {{.Ret}}`),
	)
	var buf strings.Builder
	t.Execute(&buf, this)
	return buf.String()
}
