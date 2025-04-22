package model

import "encoding/json"

type Response struct {
	Status string `json:"status"`
	Data   any    `json:"data"`
}

// Returns JSON bytes for a success response
func ResponseSuccess(data any) ([]byte, error) {
	resp := Response{
		Status: "success",
		Data:   data,
	}
	return json.Marshal(resp)
}

// Returns JSON bytes for an error response
func ResponseError(message string) ([]byte, error) {
	resp := Response{
		Status: "error",
		Data:   message,
	}
	return json.Marshal(resp)
}
