package types

import (
	"encoding/json"
	"fmt"
	"io"
)

// JSONObject is a custom GraphQL scalar type for any JSON object.
type JSONObject map[string]interface{}

// MarshalJSON for the JSONObject scalar
func (j JSONObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(j)
}

// UnmarshalJSON for the JSONObject scalar
func (j *JSONObject) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &j)
}

// UnmarshalGQL parses a GraphQL input for JSONObject
func (j *JSONObject) UnmarshalGQL(v interface{}) error {
	switch value := v.(type) {
	case map[string]interface{}:
		*j = JSONObject(value)
		return nil
	default:
		return fmt.Errorf("invalid type for JSONObject")
	}
}

// MarshalGQL writes a JSONObject as GraphQL JSON
func (j JSONObject) MarshalGQL(w io.Writer) {
	json.NewEncoder(w).Encode(j)
}