package main

import (
	"fmt"
	"net/http"

	"github.com/Saqlainabbasi/openAI-integration/api"
)

func main() {
	//initialize the routes config
	app := api.NewConfig()
	// Use the router
	app.Routes()

	fmt.Println("server listining on port 5000")
	http.ListenAndServe(":5000", app.Router)
}
