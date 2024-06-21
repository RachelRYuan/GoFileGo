package utils

import (
	"encoding/json"
)

// ObjectToJson serializes an object to JSON and then deserializes it into a specified data type.
func ObjectToJson[T any](object interface{}, data *T) error {
	jsonData, err := json.Marshal(object)
	if err != nil {
		return err
	}

	err = json.Unmarshal(jsonData, data)
	return err
}

// StringToEntity deserializes a JSON string into a specified data type.
func StringToEntity[T any](value interface{}, object *T) error {
	jsonData, err := json.Marshal(value)
	if err != nil {
		return err
	}

	err = json.Unmarshal(jsonData, object)
	return err
}
