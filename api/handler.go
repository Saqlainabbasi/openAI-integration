package api

import (
	"net/http"
)

// handler func to handle the incoming request
func (app *Config) TextHandler(w http.ResponseWriter, r *http.Request) {
	//get the promt from the req
	prompt := r.URL.Query().Get("prompt")
	//call the service
	resp, err := app.GetTextCompletion(prompt)
	//check for err
	if err != nil {
		app.writeJSON(w, http.StatusNotFound, err)
	}
	//send response back..
	app.writeJSON(w, http.StatusOK, resp)
}
