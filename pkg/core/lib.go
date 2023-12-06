package core

import (
	"github.com/samber/mo"
)

type FuncDef struct {
	Package string   `json:"pkg"`
	Name    string   `json:"name"`
	Args    []string `json:"args"`
	Return  string   `json:"ret"`
	Summary string   `json:"sum"`
	Link    string   `json:"link"`
}

type MethodDef struct {
	Package  string   `json:"pkg"`
	Name     string   `json:"name"`
	Receiver string   `json:"recv"`
	Args     []string `json:"args"`
	Return   string   `json:"ret"`
	Summary  string   `json:"sum"`
	Link     string   `json:"link"`
}

func (this *FuncDef) ToResult(sim Similarity) ResultItem {
	return ResultItem{
		Similarity: sim,
		Sig:        this.Signature(),
		Summary:    this.Summary,
		Link:       this.Link,
	}
}

type Item = mo.Either[FuncDef, MethodDef]

type ResultItem struct {
	Similarity Similarity
	Sig        string
	Summary    string
	Link       string
}

type ResultSet struct {
	Results []ResultItem
}

func NewResultSet() ResultSet {
	return ResultSet{make([]ResultItem, 0)}
}
