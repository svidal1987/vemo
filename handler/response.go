package handler

import (
	"encoding/json"
	"net/http"
)

type responseError struct {
	Message string `json:"message"`
}

func InternalServerError(w http.ResponseWriter) {
	writeResponse(http.StatusInternalServerError, getMessageError("internal server error"), w)
}

func BadRquest(w http.ResponseWriter, message string) {
	writeResponse(http.StatusBadRequest, getMessageError("bad request: "+message), w)
}

func NotFound(w http.ResponseWriter) {
	writeResponse(http.StatusNotFound, getMessageError("not found"), w)
}

func ResponseOk(w http.ResponseWriter, jsonBytes []byte) {
	writeResponse(http.StatusOK, jsonBytes, w)
}

func writeResponse(status int, jsonBytes []byte, w http.ResponseWriter) {
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(status)
	w.Write(jsonBytes)
}

func getMessageError(message string) []byte {
	response, _ := json.Marshal(responseError{Message: message})
	return response
}
