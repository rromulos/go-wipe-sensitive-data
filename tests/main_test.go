package tests

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/rromulos/go-clean-sensitive-data/internal"
)

func TestFilterJSON(t *testing.T) {
	// Test 1: Empty fields, no nodes to filter
	data := map[string]interface{}{
		"name":    "John Doe",
		"email":   "johndoe@example.com",
		"age":     30,
		"address": map[string]interface{}{"street": "123 Main St", "city": "New York"},
	}
	filteredData := internal.FilterJSON(data, "", nil)
	if !reflect.DeepEqual(filteredData, data) {
		t.Errorf("The filtered JSON is not the same as the original JSON. Expected: %v, Obtained: %v", data, filteredData)
	}

	// Test 2: Fields to filter
	filteredData = internal.FilterJSON(data, "name,email", nil)
	expectedData := map[string]interface{}{
		"age":     30,
		"address": map[string]interface{}{"street": "123 Main St", "city": "New York"},
	}
	if !reflect.DeepEqual(filteredData, expectedData) {
		t.Errorf("The filtered JSON does not match the expected result. Expected: %v, Obtained: %v", expectedData, filteredData)
	}

	// Test 3: Nodes to filter
	filteredData = internal.FilterJSON(data, "", []string{"address"})
	expectedData = map[string]interface{}{
		"name":  "John Doe",
		"email": "johndoe@example.com",
		"age":   30,
	}
	if !reflect.DeepEqual(filteredData, expectedData) {
		t.Errorf("The filtered JSON does not match the expected result. Expected: %v, Obtained: %v", expectedData, filteredData)
	}

	// Test 4: Fields and nodes to filter
	filteredData = internal.FilterJSON(data, "name", []string{"address"})
	expectedData = map[string]interface{}{
		"email": "johndoe@example.com",
		"age":   30,
	}
	if !reflect.DeepEqual(filteredData, expectedData) {
		t.Errorf("The filtered JSON does not match the expected result. expected: %v, Obtained: %v", expectedData, filteredData)
	}

	// Test 5: Empty JSON
	emptyData := map[string]interface{}{}
	filteredData = internal.FilterJSON(emptyData, "", nil)
	if !reflect.DeepEqual(filteredData, emptyData) {
		t.Errorf("The filtered JSON does not match the expected result. expect: %v, Obtained: %v", emptyData, filteredData)
	}

	// Test 6: JSON with only one node to filter
	dataWithSingleNode := map[string]interface{}{
		"address": map[string]interface{}{"street": "123 Main St", "city": "New York"},
	}
	filteredData = internal.FilterJSON(dataWithSingleNode, "", []string{"address"})
	expectedData = map[string]interface{}{}
	if !reflect.DeepEqual(filteredData, expectedData) {
		t.Errorf("The filtered JSON does not match the expected result. expect: %v, Obtained: %v", expectedData, filteredData)
	}
}

func TestFilterJSON_EmptyFieldsAndNodes(t *testing.T) {
	// Test 7: Empty fields and nodes
	data := map[string]interface{}{
		"name":    "John Doe",
		"email":   "johndoe@example.com",
		"age":     30,
		"address": map[string]interface{}{"street": "123 Main St", "city": "New York"},
	}
	filteredData := internal.FilterJSON(data, "", []string{})
	if !reflect.DeepEqual(filteredData, data) {
		t.Errorf("The filtered JSON does not match the expected result. expect: %v, Obtained: %v", data, filteredData)
	}
}

func TestFilterJSON_NestedJSON(t *testing.T) {
	// Test 8: Nested JSON with fields and nodes to filter
	data := map[string]interface{}{
		"name":    "John Doe",
		"email":   "johndoe@example.com",
		"age":     30,
		"address": map[string]interface{}{"street": "123 Main St", "city": "New York"},
		"contacts": []map[string]interface{}{
			{"type": "phone", "value": "123456789"},
			{"type": "email", "value": "contact@example.com"},
		},
	}
	filteredData := internal.FilterJSON(data, "name,email", []string{"address", "contacts"})
	expectedData := map[string]interface{}{
		"age": 30,
	}
	if !reflect.DeepEqual(filteredData, expectedData) {
		t.Errorf("The filtered JSON does not match the expected result. expect: %v, Obtained: %v", expectedData, filteredData)
	}
}

func TestFilterJSON_InvalidJSON(t *testing.T) {
	// Test 9: Invalid JSON
	invalidJSON := []byte(`{"name": "John Doe", "email": "johndoe@example.com", "age": 30, "address": {"street": "123 Main St", "city": "New York"}`)

	var invalidJSONConverted map[string]interface{}
	json.Unmarshal(invalidJSON, &invalidJSONConverted)

	filteredData := internal.FilterJSON(invalidJSONConverted, "", nil)
	if filteredData != nil {
		t.Error("Invalid JSON should not return any results")
	}
}

func TestFilterJSON_NoDataParameter(t *testing.T) {
	// Test 10: No "date" parameter
	filteredData := internal.FilterJSON(nil, "", nil)
	if filteredData != nil {
		t.Error("The function should return nil when the 'data' parameter is not provided")
	}
}

func TestFilterJSON_UnsupportedDataType(t *testing.T) {
	// Test 11: Unsupported data type (integer)
	data := map[string]interface{}{
		"name":    "John Doe",
		"age":     30,
		"address": map[string]interface{}{"street": "123 Main St", "city": "New York"},
		"score":   90,
	}
	filteredData := internal.FilterJSON(data, "score", nil)
	expectedData := map[string]interface{}{
		"name":    "John Doe",
		"age":     30,
		"address": map[string]interface{}{"street": "123 Main St", "city": "New York"},
	}
	if !reflect.DeepEqual(filteredData, expectedData) {
		t.Errorf("The filtered JSON does not match the expected result. expect: %v, Obtained: %v", expectedData, filteredData)
	}
}

func TestFilterJSON_EmptyJSON(t *testing.T) {
	// Test 12: Empty JSON
	filteredData := internal.FilterJSON(nil, "", nil)
	if filteredData != nil {
		t.Error("The function must return nil when the JSON is empty")
	}
}
