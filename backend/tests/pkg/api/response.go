package api

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Code    int    `json:"code"`
	Error   *Error `json:"error"`
	Payload *any   `json:"payload"`
}

type Error struct {
	Code         string   `json:"code"`
	Message      string   `json:"message"`
	NestedErrors *[]Error `json:"nested_errors,omitempty"`
}

func WriteErrorManual(w http.ResponseWriter, code int, e Error) {
	bytes, err := json.Marshal(&Response{
		Code:    code,
		Error:   &e,
		Payload: nil,
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(code)
	if _, err := w.Write(bytes); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func WriteError(w http.ResponseWriter, code int, ecode string, msg string) {
	bytes, err := json.Marshal(&Response{
		Code: code,
		Error: &Error{
			Code:    ecode,
			Message: msg,
		},
		Payload: nil,
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(code)
	if _, err := w.Write(bytes); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func WriteErrorMessage(w http.ResponseWriter, code int, ecode string, msg string) {
	bytes, err := json.Marshal(&Response{
		Code: code,
		Error: &Error{
			Code:    ecode,
			Message: msg,
		},
		Payload: nil,
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(code)
	if _, err := w.Write(bytes); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func WriteResponse(w http.ResponseWriter, code int, response any) {
	bytes, err := json.Marshal(&Response{
		Code:    code,
		Error:   nil,
		Payload: &response,
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(code)
	if _, err := w.Write(bytes); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}