package api

import (
	"encoding/json"
	"net/http"
)

// reciever func to WriteJson
// func to send response back to the client.....
func (*Config) writeJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v) // the .Encode converts the data to json dat type

}
