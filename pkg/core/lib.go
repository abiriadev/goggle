package core

import (
	"strings"
	"text/template"

	"github.com/samber/mo"
)

type ToSignature interface {
	Signature() string
}

type FuncDef struct {
	Package string   `json:"pkg"`
	Name    string   `json:"name"`
	Args    []string `json:"args"`
	Return  string   `json:"ret"`
	Summary string   `json:"sum"`
	Link    string   `json:"link"`
}

func (this *FuncDef) Signature() string {
	t := template.Must(
		template.New("").Parse(`func {{.Name}}({{range .Args}}{{.}}, {{end}}) {{.Ret}}`),
	)
	var buf strings.Builder
	t.Execute(&buf, this)
	return buf.String()
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
		// TODO: use proper sig
		Sig:     this.Name,
		Summary: this.Summary,
		Link:    this.Link,
	}
}

type Item = mo.Either[FuncDef, MethodDef]

type ResultItem struct {
	Similarity Similarity
	Sig        string
	Summary    string
	Link       string
}

func (ri *ResultItem) Signature() string {
	return ri.Sig
}

type ResultSet struct {
	Results []ResultItem
}

func NewResultSet() ResultSet {
	return ResultSet{make([]ResultItem, 0)}
}
