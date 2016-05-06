package utils

import (
	"encoding/json"
)

type ResultTransformer struct {
	value interface{}
}

// NewResultTransformer constructor
func NewResultTransformer(value interface{}) *ResultTransformer {
	return &ResultTransformer{value}
}

// Set value
func (this *ResultTransformer) Set(value interface{}) {
	this.value = value
}

// Get value
func (this *ResultTransformer) Get() interface{} {
	return this.value
}

// ToJson return json
func (this *ResultTransformer) ToJson() (string, error) {

	json, err := json.MarshalIndent(this.value, "", "  ")
	if err != nil {
		return "", err
	}

	return string(json), nil
}
