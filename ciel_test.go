package ciel

import (
	"testing"
)

func AssertRule(rules string, query string, path string, expected bool, t *testing.T) {
	compiled, err := Compile(rules)
	if err != nil {
		panic(err)
	}
	result, err := compiled.Check(query, path)
	if err != nil {
		panic(err)
	}
	if result != expected {
		t.Error("To", query, path, "should be", expected)
	}
}

func TestShouldForbidAccess(t *testing.T) {
	rules := `{
		"rules": {
			".read": "false",
			".write": "false"
		}
	}`

	AssertRule(rules, "read", "/", false, t)
	AssertRule(rules, "write", "/", false, t)
}

func TestShouldAllowFullAccess(t *testing.T) {
	rules := `{
		"rules": {
			".read": "true",
			".write": "true"
		}
	}`

	AssertRule(rules, "read", "/", true, t)
	AssertRule(rules, "write", "/", true, t)
}

func TestShould(t *testing.T) {
	rules := `{
		"rules": {
			"parents": {
				"$parent": {
					"$child": {
						".validate": "newData.hasChildren(['age'])",
						"age": {
							".validate": "newData.isNumber()"
						},
						"$other": {
							".validate": "false"
						}
					}
				}
			}
		}
	}`

	AssertRule(rules, "validate", "parents/1/2", false, t)
}
