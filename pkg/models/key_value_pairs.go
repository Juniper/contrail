package models

import "github.com/pkg/errors"

// GetValue checks gets the value from key-value pair Slice by key.
func (kvps *KeyValuePairs) GetValue(key string) (string, error) {
	for _, kvp := range kvps.GetKeyValuePair() {
		if kvp.GetKey() == key {
			return kvp.GetValue(), nil
		}
	}
	return "", errors.Errorf("key '%s' doesn't exists", key)
}
