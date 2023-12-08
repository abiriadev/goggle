package index

import (
	"encoding/json"
	"io"
)

func ResolveFullIndex() {

}

type ModuleIndex struct {
	Path      string
	Version   string
	Timestamp string
}

func ParseIndex(r io.Reader) ([]ModuleIndex, error) {
	dec := json.NewDecoder(r)

	midxes := make([]ModuleIndex, 0)

	var v ModuleIndex
	err := dec.Decode(&v)
	if err != nil {
		return nil, nil
	}

	midxes = append(midxes, v)

	return midxes, nil
}
