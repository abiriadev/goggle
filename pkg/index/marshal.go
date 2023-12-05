package index

import (
	"encoding/json"
	"io"
	"os"
)

func (index *Index) MarshalJson() ([]byte, error) {
	return json.Marshal(index)
}

func (index *Index) MarshalJsonTo(w io.Writer) error {
	return json.NewEncoder(w).Encode(index)
}

func (index *Index) Save(filename string) error {
	w, e := os.Create(filename)
	if e != nil {
		return e
	}

	return index.MarshalJsonTo(w)
}
