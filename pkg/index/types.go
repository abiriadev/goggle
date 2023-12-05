package index

type FuncDef struct {
	Name string
	Args []string
	Ret  string
}

type Index struct {
	Items []FuncDef
}

func MergeIndex(idxes []Index) Index {
	newi := Index{make([]FuncDef, 0)}

	for _, idx := range idxes {
		newi.Items = append(newi.Items, idx.Items...)
	}

	return newi
}
