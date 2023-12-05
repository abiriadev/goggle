package index

import (
	"fmt"
	"reflect"

	"github.com/abiriadev/goggle/pkg/query"
)

func (idx *Index) Query(q query.Query) []FuncDef {
	rl := make([]FuncDef, 0)

	fmt.Println(q.Args)

	for _, fd := range idx.Items {
		if q.Ret == fd.Ret && reflect.DeepEqual(q.Args, fd.Args) {
			rl = append(rl, fd)
		}
	}

	return rl
}
