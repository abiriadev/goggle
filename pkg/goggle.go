package pkg

import "go/types"

type FuncDef struct {
	name string
	ret  types.Type
}

type Index struct {
	items []FuncDef
}
