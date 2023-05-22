package main

import (
	"encoding/json"
	"net/http"
)

type jsonResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

func (app *Config) Broker(responseWriter http.ResponseWriter, request *http.Request) {
	payload := jsonResponse{
		Error:   false,
		Message: "feijoada",
	}

	out, _ := json.MarshalIndent(payload, "", "\t")
	responseWriter.Header().Set("Content-type", "application/json")
	responseWriter.WriteHeader(http.StatusAccepted)
	responseWriter.Write(out)
}
