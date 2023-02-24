package data

import (
	"encoding/json"
	"io"
)

func ToJSON(data interface{}, rw io.Writer) error {
	e := json.NewEncoder(rw)
	return e.Encode(data)
}

func FromJSON(data interface{}, r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(data)
}
