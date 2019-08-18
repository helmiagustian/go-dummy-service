package response

import (
	"encoding/json"
	"net/http"
)

// SuccessBody holds data for success response
type SuccessBody struct {
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"message,omitempty"`
	Meta    MetaInfo    `json:"meta"`
}

// MetaInfo holds meta data
type MetaInfo struct {
	HTTPStatus int `json:"http_status"`
}

func (s *SuccessBody) Statuscode() int {
	return s.Meta.HTTPStatus
}

// ErrorBody holds data for error response
type ErrorBody struct {
	Errors []ErrorInfo `json:"errors"`
	Meta   MetaInfo    `json:"meta"`
}

// ErrorInfo holds error detail
type ErrorInfo struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func (e *ErrorBody) Statuscode() int {
	return e.Meta.HTTPStatus
}

func SuccessDefault() *SuccessBody {
	return &SuccessBody{
		Message: "Success",
		Meta: MetaInfo{
			HTTPStatus: 200,
		},
	}
}

func ErrorDefault() *ErrorBody {
	return &ErrorBody{
		Meta: MetaInfo{
			HTTPStatus: 500,
		},
	}
}

type Resulter interface {
	Statuscode() int
}

func Write(w http.ResponseWriter, result Resulter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(result.Statuscode())
	return json.NewEncoder(w).Encode(result)
}
