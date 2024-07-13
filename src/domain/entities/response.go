package entities

type ResponseMessage struct {
	Message string `json:"message"`
}

type ResponseModel struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Status  int         `json:"status,omitempty"`
}

type ResponseBool struct {
	Message string `json:"message"`
	IsTrue  bool   `json:"istrue,omitempty"`
}
