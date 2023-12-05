package index

import (
	"html/template"
	"strings"
)

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

func (this *FuncDef) String() string {
	t := template.Must(template.New("").Parse(`{{.Name}}({{range .Args}}{{.}}, {{end}}) {{.Ret}}`))
	var buf strings.Builder
	t.Execute(&buf, this)
	return buf.String()
}
