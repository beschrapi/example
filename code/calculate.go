package main

import (
	"encoding/json"
	"net/http"

	"github.com/beschrapi/pkg"
)

// Run is the main function to handle the request
// Here you can write your own logic to handle the request
func Run(w http.ResponseWriter, r *http.Request, app *pkg.App) {
	type Calculation struct {
		A int `json:"a"`
		B int `json:"b"`
	}

	var data Calculation
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	result := data.A + data.B
	app.Json(w, result) // Just a helper function to write JSON response
	// You can also use the plain way to write JSON response
	// w.Header().Set("Content-Type", "application/json")
	// json.NewEncoder(w).Encode(result)
}
