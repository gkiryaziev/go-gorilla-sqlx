package models

// Header struct
type Header struct {
	Status string      `json:"status"`
	Count  int         `json:"count"`
	Data   interface{} `json:"data"`
}
