package utils

import (
	"encoding/json"
)

type ResultTransformer struct {
	value interface{}
}

// constructor for ResultTransformer struct
func NewResultTransformer(value interface{}) *ResultTransformer {
	return &ResultTransformer{value}
}

func (this *ResultTransformer) Set(value interface{}) {
	this.value = value
}

func (this *ResultTransformer) Get() interface{} {
	return this.value
}

func (this *ResultTransformer) ToJson() (string, error) {

	json, err := json.MarshalIndent(this.value, "", "  ")
	if err != nil {
		return "", err
	}

	return string(json), nil
}
