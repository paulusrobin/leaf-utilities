package leafViper

import (
	"github.com/paulusrobin/leaf-utilities/encoding/json"
	"strings"
)

func getDirectory(path string) string {
	splits := strings.Split(path, "/")
	return strings.Join(splits[:len(splits)-1], "/")
}

func getFile(path string) (string, error) {
	splits := strings.Split(path, "/")
	last := splits[len(splits)-1]

	return last, nil
}

func marshallMap(m map[string]interface{}) map[string]interface{} {
	for k, v := range m {
		switch c := v.(type) {
		case map[string]interface{}:
			b, _ := json.Marshal(c)
			m[k] = b
		default:
			continue
		}
	}

	return m
}
