package server

import (
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
)

func NewID() string {
	return uuid.New().String()
}

type Error struct {
	ID            string                 `json:"id"`
	DetailedError string                 `json:"detailed_error"`
	Where         string                 `json:"-"`
	Params        map[string]interface{} `json:"params"`
	StatusCode    int                    `json:"status_code"`
	RequestID     string                 `json:"request_id"`
}

func NewError(where, id, dErr string, Params map[string]interface{}, statusCode int) *Error {
	return &Error{
		Where:         where,
		ID:            id,
		DetailedError: dErr,
		Params:        Params,
		StatusCode:    statusCode,
	}
}

func (e *Error) Error() string {
	return fmt.Sprintf("where:%s id: %s msg:%s", e.Where, e.ID, e.DetailedError)
}

func (e *Error) ToJSON() string {
	b, _ := json.Marshal(e)
	return string(b)
}
