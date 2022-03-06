package helper

import "encoding/json"

var JsonMarshalFunc = Marshal

func ExecuteMarshal(v interface{}) ([]byte, error) {
	return JsonMarshalFunc(v)
}

func Marshal(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}