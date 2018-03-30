package db

import (
	"database/sql/driver"
	"errors"

	"encoding/json"
)

type jsonValue struct {
	data interface{}
}

func JSON(v interface{}) *jsonValue {
	return &jsonValue{
		data: v,
	}
}

func (m *jsonValue) Value() (driver.Value, error) {

	return json.Marshal(m.data)
}

// Scan implements the sql.Scanner interface, ungzipping the value coming off
// the wire and storing the raw result in the GzippedText.
func (m *jsonValue) Scan(src interface{}) error {
	var source []byte
	switch src.(type) {
	case string:
		source = []byte(src.(string))
	case []byte:
		source = src.([]byte)
	default:
		return errors.New("Incompatible type for jsonValue")
	}

	return json.Unmarshal(source, m.data)
}
