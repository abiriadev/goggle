package goggle

import (
	"github.com/abiriadev/goggle/pkg/core"
	"github.com/abiriadev/goggle/pkg/eval"
	"github.com/abiriadev/goggle/pkg/index"
	"github.com/abiriadev/goggle/pkg/query"
	"github.com/alecthomas/participle/v2"
)

// reexports:

// ResultSet represents a query result ordered by accuracy
type ResultSet = core.ResultSet

// ResultItem represents a single item that matches to the query
type ResultItem = core.ResultItem

// Similarity represents how the results are accurate
type Similarity = core.Similarity

// The main entry to the Goggle search engine
type Goggle struct {
	index index.Index
	qp    participle.Parser[query.Query]
}

// Initialize new Goggle search engine by loading index from the path
func Load(indexFile string) (Goggle, error) {
	index, err := index.Load(indexFile)
	if err != nil {
		return Goggle{}, err
	}

	qp, err := query.QueryParser()
	if err != nil {
		return Goggle{}, err
	}

	return Goggle{index, *qp}, nil
}

// Query Goggle with a given query
func (g *Goggle) Query(query string) (core.ResultSet, error) {
	q, err := g.qp.ParseString("", query)
	if err != nil {
		return core.ResultSet{}, err
	}

	return eval.Query(&g.index, *q), nil
}
