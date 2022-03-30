package parse

import (
	"github.com/ghodss/yaml"
)

type Person struct {
	Name string `json:"name"` // Affects YAML field names too.
	Age  int    `json:"age"`
}

func Marshal(p interface{}) ([]byte, error) {
	return yaml.Marshal(p)
}

func Unmarshal(y []byte, p interface{}) error {
	return yaml.Unmarshal(y, p)
}
