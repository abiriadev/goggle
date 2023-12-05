package core

import (
	"encoding/json"
	"io"
)

func (rs *ResultSet) MarshalJson() ([]byte, error) {
	return json.Marshal(rs)
}

func (rs *ResultSet) MarshalJsonTo(w io.Writer) error {
	return json.NewEncoder(w).Encode(rs)
}
