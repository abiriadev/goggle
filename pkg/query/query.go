package query

import "github.com/alecthomas/participle/v2"

type Query struct {
	Name string   `"func"? @Ident?`
	Args []string `"(" ( @Ident ( "," @Ident )* )? ")"`
	Ret  string   `@Ident?`
}

func QueryParser() (*participle.Parser[Query], error) {
	p, err := participle.Build[Query]()
	return p, err
}
