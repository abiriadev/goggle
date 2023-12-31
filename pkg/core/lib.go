package core

import (
	"github.com/samber/mo"
)

type Arg struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type FuncDef struct {
	Package     string `json:"pkg"`
	PackageMame string `json:"pkg_name"`
	Name        string `json:"name"`
	Args        []Arg  `json:"args"`
	Return      string `json:"ret"`
	Summary     string `json:"sum"`
	Link        string `json:"link"`
}

type MethodDef struct {
	Package     string `json:"pkg"`
	PackageMame string `json:"pkg_name"`
	Name        string `json:"name"`
	Receiver    string `json:"recv"`
	Args        []Arg  `json:"args"`
	Return      string `json:"ret"`
	Summary     string `json:"sum"`
	Link        string `json:"link"`
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
	Similarity Similarity `json:"sim"`
	Sig        string     `json:"sig"`
	Summary    string     `json:"summary"`
	Link       string     `json:"link"`
}

type ResultSet struct {
	Results []ResultItem `json:"items"`
}

func NewResultSet() ResultSet {
	return ResultSet{make([]ResultItem, 0)}
}
