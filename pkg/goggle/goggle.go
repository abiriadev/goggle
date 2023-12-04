package goggle

import "github.com/abiriadev/goggle/pkg/query"

type FuncDef struct {
	Name string
	Ret  string
}

type Index struct {
	Items []FuncDef
}

func (index *Index) Query(q query.Query) []FuncDef {
	rl := make([]FuncDef, 0)

	for _, fd := range index.Items {
		if q.Ret == fd.Ret {
			rl = append(rl, fd)
		}
	}

	return rl
}
