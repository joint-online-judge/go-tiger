package json

import (
	"encoding/json"
	// "github.com/bytedance/sonic"
)

// sonic disabled since it does not support Go 1.20 so far

func Marshal(val interface{}) ([]byte, error) {
	// return sonic.Marshal(val)
	return json.Marshal(val)
}

func Unmarshal(buf []byte, val interface{}) error {
	// return sonic.Unmarshal(buf, val)
	return json.Unmarshal(buf, val)
}
