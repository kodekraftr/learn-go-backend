package main

import (
	"net/http"
)

func (app *application) healthCheckerHandler(w http.ResponseWriter, r *http.Request) {

	data := map[string]string{
		"status":  "working",
		"version": "0.0.1",
	}

	err := writeJson(w, http.StatusAccepted, data)

	if err != nil {
		writeJsonError(w, http.StatusInternalServerError, err.Error())
	}
}
