package entity

import "fuux/internal/entity/error"

type Response struct {
	Error   int         `json:"error,omitempty"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func ResponseError(err *error.Error) *Response {
	return &Response{
		Error:   err.Code,
		Message: err.Message,
	}
}

func SuccessResponse() *Response {
	return &Response{
		Error:   0,
		Message: "Success",
	}
}
