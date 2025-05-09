package utils

import (
	"encoding/json"
	"fmt"
)

// MapToStruct converts a generic map to a typed struct
func MapToStruct(data map[string]interface{}, result interface{}) error {
	bytes, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("failed to marshal map: %w", err)
	}

	if err := json.Unmarshal(bytes, result); err != nil {
		return fmt.Errorf("failed to unmarshal into struct: %w", err)
	}

	return nil
}
