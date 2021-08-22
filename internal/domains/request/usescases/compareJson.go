package usescases

import (
	"encoding/json"
	"io"
	"reflect"
)

type CompareJSON struct {
}

func (c *CompareJSON) Compare(a, b io.Reader) (bool, error) {
	var j, j2 interface{}
	d := json.NewDecoder(a)
	if err := d.Decode(&j); err != nil {
		return false, err
	}
	d = json.NewDecoder(b)
	if err := d.Decode(&j2); err != nil {
		return false, err
	}
	return reflect.DeepEqual(j2, j), nil
}
