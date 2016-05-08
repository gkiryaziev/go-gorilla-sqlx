package utils

import (
	"encoding/json"
)

// ResultTransformer structure
type ResultTransformer struct {
	value interface{}
}

// NewResultTransformer constructor
func NewResultTransformer(value interface{}) *ResultTransformer {
	return &ResultTransformer{value}
}

// Set value
func (rt *ResultTransformer) Set(value interface{}) {
	rt.value = value
}

// Get value
func (rt *ResultTransformer) Get() interface{} {
	return rt.value
}

// ToJSON return json
func (rt *ResultTransformer) ToJSON() (string, error) {

	json, err := json.Marshal(rt.value)
	if err != nil {
		return "", err
	}

	return string(json), nil
}
