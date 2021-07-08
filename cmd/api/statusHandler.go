package main

import (
	"encoding/json"
	"net/http"
)

func (app *application) statusHandler(w http.ResponseWriter, r *http.Request) {
	status := AppStatus{
		Status:      "Available",
		Environment: app.config.env,
		Version:     version,
	}
	js, err := json.MarshalIndent(status, "", "")
	if err != nil {
		app.logger.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(js)
	if err != nil {
		app.logger.Println(err)
	}
}
