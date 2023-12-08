package index

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"time"
)

// maximum index server limit
const LIMIT = 2000

func ResolveFullIndex() ([]ModuleIndex, error) {
	midxes := make([]ModuleIndex, 0)
	var inc *time.Time

	for {
		r, err := FetchModuleIndex(inc, LIMIT)
		if err != nil {
			return nil, err
		}

		m, err := ParseModuleIndex(r)
		if err != nil {
			return nil, err
		}

		midxes = append(midxes, m...)

		if len(m) != 0 {
			*inc = IncTimeStamp(m[len(m)])
		}
	}

	return midxes, nil
}

func IncTimeStamp(midxes ModuleIndex) time.Time {
	return midxes.Timestamp.Add(time.Nanosecond)
}

func FetchModuleIndex(since *time.Time, limit int) (io.Reader, error) {
	req, err := http.NewRequest("GET", "https://index.golang.org/index", nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	if since != nil {
		q.Add("since", since.Format(time.RFC3339))
	}
	q.Add("limit", strconv.Itoa(limit))
	req.URL.RawQuery = q.Encode()

	if err != nil {
		return nil, err
	}

	client := &http.Client{}

	res, err := client.Do(req)
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
