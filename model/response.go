package model

import (
	"encoding/json"
)

type ResponseType string

const (
	TypeBattInfo ResponseType = "BATT"
	TypeChat     ResponseType = "CHAT"
	TypeAppt     ResponseType = "APPT"
)

type Response struct {
	Status string       `json:"status"`
	Type   ResponseType `json:"type,omitempty"`
	Data   any          `json:"data"`
}

// Returns JSON bytes for a success response
func ResponseSuccess(data any, respType ResponseType) ([]byte, error) {
	resp := Response{
		Status: "success",
		Type:   respType,
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
