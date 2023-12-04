package pkg

type Query struct {
	Ret string
}

func (index *Index) Query(q Query) []FuncDef {
	rl := make([]FuncDef, 0)

	for _, fd := range index.Items {
		if q.Ret == fd.Ret {
			rl = append(rl, fd)
		}
	}

	return rl
}
