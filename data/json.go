package data

import (
	"encoding/json"
	"io"
)

func ToJSON(i interface{}, r io.Writer) error {
	e := json.NewEncoder(r)

	return e.Encode(i)
}

// Decodes a JSON to product
func FromJSON(i interface{}, r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(i)
}
