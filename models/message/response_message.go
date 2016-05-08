package message

// ResponseMessage struct
type ResponseMessage struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Info    string `json:"info"`
}
