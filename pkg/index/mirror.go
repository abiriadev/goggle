package index

import (
	"encoding/json"
	"io"
	"net/http"
)

func ResolveFullIndex() {

}

func FetchModuleIndex() (io.Reader, error) {
	res, err := http.Get("https://index.golang.org/index")
	if err != nil {
		return nil, err
	}

	return res.Body, nil
}

type ModuleIndex struct {
	Path      string
	Version   string
	Timestamp string
}

func ParseModuleIndex(r io.Reader) ([]ModuleIndex, error) {
	dec := json.NewDecoder(r)

	midxes := make([]ModuleIndex, 0)

	for {
		var v ModuleIndex
		err := dec.Decode(&v)
		if err != nil {
			if err == io.EOF {
				break
			}

			return nil, nil
		}

		midxes = append(midxes, v)
	}

	return midxes, nil
}
