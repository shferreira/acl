package ciel

import (
	"encoding/json"
	// "github.com/Knetic/govaluate"
)

type Ciel struct {
}

func Compile(rules string) (*Ciel, error) {
	var ruleset map[string]interface{}
	err := json.Unmarshal([]byte(rules), &ruleset)
	if err != nil {
		return nil, err
	}

	walk(ruleset, "", func(path string, value string) {
		print(path + " = " + value + "\n")
	})

	return &Ciel{}, nil
}

func (r *Ciel) Check(query, path string) (bool, error) {
	return false, nil
}

func walk(rules map[string]interface{}, path string, visitor func(string, string)) {
	for k, v := range rules {
		currPath := path + "/" + k
		switch v2 := v.(type) {
		case map[string]interface{}:
			walk(v2, currPath, visitor)
		case string:
			visitor(currPath, v2)
		}
	}
}
