package common

type generalRes struct {
	Status  bool        `json:"status"`
	Message string      `json:"message,omitempty"`
	Payload interface{} `json:"payload,omitempty"`
	Result  interface{} `json:"result,omitempty"`
}

func ResponseOKSuccess(message string, payload, result interface{}) *generalRes {
	return &generalRes{true, message, payload, result}
}

func BadResponseFailure(message string, payload, result interface{}) *generalRes {
	return &generalRes{false, message, payload, result}
}

type Dictionary map[string]interface{}
