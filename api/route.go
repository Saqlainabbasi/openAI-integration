// routes.go
package api

import (
	"net/http"
)

type Config struct {
	Router *http.ServeMux
}

func NewConfig() *Config {
	return &Config{
		Router: http.NewServeMux(),
	}
}

func (app *Config) Routes() {

	app.Router.HandleFunc("/generate-text", app.TextHandler)
	// app.Router.HandleFunc("/complete-code", app.CodeHandler)
}
