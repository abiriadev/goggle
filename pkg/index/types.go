package index

type FuncDef struct {
	Name string
	Args []string
	Ret  string
}

type Index struct {
	items []FuncDef
}

func (index *Index) Items() []FuncDef {
	return index.items
}

func MergeIndex(idxes []Index) Index {
	newi := Index{make([]FuncDef, 0)}

	for _, idx := range idxes {
		newi.items = append(newi.items, idx.items...)
	}

	return newi
}
