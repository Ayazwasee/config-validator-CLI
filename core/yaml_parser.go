package core

import (
	"fmt"
	"io"
	"os"

	"gopkg.in/yaml.v2"
)

func ParseYAMLFile(path string) (map[string]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	var raw map[string]interface{}
	if err := yaml.Unmarshal(data, &raw); err != nil {
		return nil, err
	}

	// Convert interface{} to string map
	env := make(map[string]string)
	for k, v := range raw {
		env[k] = toString(v)
	}

	return env, nil
}

// helper to safely convert types to string
func toString(v interface{}) string {
	switch val := v.(type) {
	case string:
		return val
	case int, int64, float64, bool:
		return fmt.Sprintf("%v", val)
	default:
		return ""
	}
}
