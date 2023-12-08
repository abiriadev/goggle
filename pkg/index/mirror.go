package index

import (
	"encoding/json"
	"io"
	"net/http"
	"time"
)

func ResolveFullIndex() ([]ModuleIndex, error) {
	midxes := make([]ModuleIndex, 0)

	for {
		r, err := FetchModuleIndex()
		if err != nil {
			return nil, err
		}

		m, err := ParseModuleIndex(r)
		if err != nil {
			return nil, err
		}

		midxes = append(midxes, m...)

		if len(m) != 0 {
			IncTimeStamp(m[len(m)])
		}
	}
}

func IncTimeStamp(midxes ModuleIndex) time.Time {
	return midxes.Timestamp.Add(time.Nanosecond)
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
	Timestamp time.Time
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
