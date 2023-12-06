package core

import (
	"strings"
	"text/template"

	"github.com/samber/mo"
)

type FuncDef struct {
	Pkg     string
	Name    string
	Args    []string
	Ret     string
	Summary string
	DocLink string
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
	Pkg      string
	Name     string
	Receiver string
	Args     []string
	Ret      string
	Summary  string
	DocLink  string
}

func (this *FuncDef) ToResult(sim Similarity) ResultItem {
	return ResultItem{
		Similarity: sim,
		// TODO: use proper sig
		Sig:     this.Name,
		Summary: this.Summary,
		Link:    this.DocLink,
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
