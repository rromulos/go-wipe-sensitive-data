package internal

import (
	"strings"
)

func FilterJSON(data map[string]interface{}, fields string, nodes []string) map[string]interface{} {
	if fields == "" && len(nodes) == 0 {
		// Returns the original JSON if there are no fields or nodes to filter
		return data
	}

	filteredData := make(map[string]interface{})

	for key, value := range data {
		if fields != "" && contains(strings.Split(fields, ","), key) {
			// Removes the property if it is in the field list
			continue
		}

		switch v := value.(type) {
		case map[string]interface{}:
			// Recursively filter properties from nested JSON
			filteredValue := FilterJSON(v, fields, nodes)
			if len(filteredValue) > 0 {
				// Keep property only if nested JSON is not empty
				filteredData[key] = filteredValue
			}
		default:
			// Keep property if not a nested map
			filteredData[key] = value
		}
	}

	// Remove nodes from JSON
	for _, node := range nodes {
		delete(filteredData, node)
	}

	return filteredData
}

func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}
