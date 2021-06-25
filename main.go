package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const port = ":4000"

type URLDescription struct {
	URL         string `json:"url"`
	Method      string `json:"method"`
	Description string `json:"description"`
	Payload     string `json:"payload,omitempty"`
}

func documentation(rw http.ResponseWriter, r *http.Request) {
	data := []URLDescription{
		{
			URL:         "/",
			Method:      "GET",
			Description: "See Documentation",
		},
		{
			URL:         "/blocks",
			Method:      "POST",
			Description: "Add A Block",
			Payload:     "data:string",
		},
	}
	rw.Header().Add("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(data)
}

func main() {
	http.HandleFunc("/", documentation)
	fmt.Printf("Listening on http://localhost%s\n", port)
	http.ListenAndServe(port, nil)
}
