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

func UnmarshalJson(data []byte) (Index, error) {
	var idx Index
	return idx, json.Unmarshal(data, &idx)
}

func UnmarshalJsonFrom(r io.Reader) (Index, error) {
	var idx Index
	return idx, json.NewDecoder(r).Decode(&idx)
}

func Load(filename string) (Index, error) {
	r, e := os.Open(filename)
	if e != nil {
		return Index{}, e
	}

	return UnmarshalJsonFrom(r)
}
