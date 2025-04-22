package model

import (
	"encoding/json"
)

type Status string
type ResponseType string

const (
	Success Status = "SUCCESS"
	Error   Status = "ERROR"
)

const (
	TypeBattInfo ResponseType = "BATT"
	TypeChat     ResponseType = "CHAT"
	TypeAppt     ResponseType = "APPT"
)

type Response struct {
	Status Status       `json:"status"`
	Type   ResponseType `json:"type,omitempty"`
	Data   any          `json:"data"`
}

// Returns JSON bytes for a success response
func ResponseSuccess(data any, respType ResponseType) ([]byte, error) {
	resp := Response{
		Status: Success,
		Type:   respType,
		Data:   data,
	}
	return json.Marshal(resp)
}

// Returns JSON bytes for an error response
func ResponseError(message string) ([]byte, error) {
	resp := Response{
		Status: Error,
		Data:   message,
	}
	return json.Marshal(resp)
}
