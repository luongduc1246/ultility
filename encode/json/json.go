package json

import "encoding/json"

type JsonEncode struct{}

func (js JsonEncode) Marshal(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

func (js JsonEncode) Unmarshal(v []byte, o interface{}) error {
	return json.Unmarshal(v, o)
}
