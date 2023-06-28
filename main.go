package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/rromulos/go-clean-sensitive-data/internal"
)

type RequestData struct {
	Fields string                 `json:"fields"`
	Nodes  []string               `json:"nodes"`
	Data   map[string]interface{} `json:"data"`
}

func main() {
	http.HandleFunc("/filter", filterHandler)
	log.Fatal(http.ListenAndServe(":7777", nil))
}

func filterHandler(w http.ResponseWriter, r *http.Request) {
	// Check if the http method is POST
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Read the body of the request
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}

	// Decode the request body into a struct
	var requestData RequestData
	if err := json.Unmarshal(body, &requestData); err != nil {
		http.Error(w, "invalid JSON", http.StatusBadRequest)
		return
	}

	// Checks if the "data" field was provided
	if requestData.Data == nil {
		http.Error(w, "'data' field not found", http.StatusBadRequest)
		return
	}

	// Filter JSON
	filteredData := internal.FilterJSON(requestData.Data, requestData.Fields, requestData.Nodes)

	// Converts the filtered JSON back to []byte
	filteredBytes, err := json.Marshal(filteredData)
	if err != nil {
		http.Error(w, "Error converting filtered JSON", http.StatusInternalServerError)
		return
	}

	// Sets the content-type of the response to application/json
	w.Header().Set("Content-Type", "application/json")

	// Write the response
	w.Write(filteredBytes)
}
