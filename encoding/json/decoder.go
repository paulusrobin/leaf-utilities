package json

func UnmarshalFromString(str string, v interface{}) error {
	return json.UnmarshalFromString(str, v)
}

func Unmarshal(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}

