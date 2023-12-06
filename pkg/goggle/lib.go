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

// Configuration for Goggle
type Config struct {
	Limit int
}

// The main entry to the Goggle search engine
type Goggle struct {
	index index.Index
	qp    participle.Parser[query.Query]
	cfg   Config
}

func NewGoggle(cfg Config) (Goggle, error) {
	qp, err := query.QueryParser()
	if err != nil {
		return Goggle{}, err
	}

	return Goggle{index.NewIndex(), *qp, cfg}, nil
}

// Initialize new Goggle search engine by loading index from the path
func (g *Goggle) Load(indexFile string) error {
	index, err := index.Load(indexFile)
	if err != nil {
		return err
	}

	g.index = index
	return nil
}

// Query Goggle with a given query
func (g *Goggle) Query(query string) (core.ResultSet, error) {
	q, err := g.qp.ParseString("", query)
	if err != nil {
		return core.ResultSet{}, err
	}

	return eval.Query(&g.index, *q, g.cfg.Limit), nil
}
