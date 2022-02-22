package json

import (
	jsoniter "github.com/json-iterator/go"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

func Valid(data []byte) bool {
	return json.Valid(data)
}

func SetConfig(config jsoniter.Config) {
	json = config.Froze()
}
