package utils

import (
	"testing"
)

func TestToJson(t *testing.T) {
	var tests = []struct {
		data     map[string]interface{}
		expected string
	}{{
		data:     map[string]interface{}{"message": "hello"},
		expected: `{"message":"hello"}`,
	}, {
		data:     map[string]interface{}{"message": 100},
		expected: `{"message":100}`,
	}}

	for _, test := range tests {
		result, err := NewResultTransformer(test.data).ToJSON()
		if err != nil {
			t.Fatal(err)
		}
		if test.expected != result {
			t.Error(test.expected, "!=", result)
		}
	}
}
