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

// func filterJSON(data map[string]interface{}, fields string, nodes []string) map[string]interface{} {
// 	if fields == "" && len(nodes) == 0 {
// 		// Returns the original JSON if there are no fields or nodes to filter
// 		return data
// 	}

// 	filteredData := make(map[string]interface{})

// 	for key, value := range data {
// 		if fields != "" && contains(strings.Split(fields, ","), key) {
// 			// Removes the property if it is in the field list
// 			continue
// 		}

// 		switch v := value.(type) {
// 		case map[string]interface{}:
// 			// Recursively filter properties from nested JSON
// 			filteredValue := filterJSON(v, fields, nodes)
// 			if len(filteredValue) > 0 {
// 				// Keep property only if nested JSON is not empty
// 				filteredData[key] = filteredValue
// 			}
// 		default:
// 			// Keep property if not a nested map
// 			filteredData[key] = value
// 		}
// 	}

// 	// Remove nodes from JSON
// 	for _, node := range nodes {
// 		delete(filteredData, node)
// 	}

// 	return filteredData
// }

// func contains(slice []string, item string) bool {
// 	for _, s := range slice {
// 		if s == item {
// 			return true
// 		}
// 	}
// 	return false
// }
