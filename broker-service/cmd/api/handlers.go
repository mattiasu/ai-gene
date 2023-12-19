package main

import (
	"net/http"

	"github.com/tsawler/toolbox"
)

type RequestPayload struct {
	Action string      `json:"action"`
	Init   InitPayload `json:"init,omitempty"`
	Log    LogPayload  `json:"log,omitempty"`
}

type LogPayload struct {
	Name string `json:"name"`
	Data string `json:"data"`
}

type InitPayload struct {
	Name string `json:"name"`
	Data string `json:"data"`
}

func (app *Config) Broker(w http.ResponseWriter, r *http.Request) {
	var tools toolbox.Tools

	payload := toolbox.JSONResponse{
		Error:   false,
		Message: "Hit the broker endpoint",
	}

	_ = tools.WriteJSON(w, http.StatusAccepted, payload)
}

func (app *Config) GetQuote(w http.ResponseWriter, r *http.Request) {
	var tools toolbox.Tools

	payload := toolbox.JSONResponse{
		Error:   false,
		Message: "Hit the get quote endpoint",
	}

	_ = tools.WriteJSON(w, http.StatusAccepted, payload)
}
