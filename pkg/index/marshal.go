package index

import (
	"encoding/gob"
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

func (index *Index) SaveJson(filename string) error {
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

func LoadJson(filename string) (Index, error) {
	r, e := os.Open(filename)
	if e != nil {
		return Index{}, e
	}

	return UnmarshalJsonFrom(r)
}

func (index *Index) Save(filename string) error {
	w, err := os.Create(filename)
	if err != nil {
		return err
	}

	gob.Register(Index{})

	enc := gob.NewEncoder(w)

	return enc.Encode(index)
}

func Load(filename string) (Index, error) {
	r, err := os.Open(filename)
	if err != nil {
		return Index{}, err
	}

	gob.Register(Index{})

	dec := gob.NewDecoder(r)

	var index Index
	err = dec.Decode(&index)

	return index, err
}
