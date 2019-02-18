package models

import "github.com/pkg/errors"

func (kvps *KeyValuePairs) KeyExists(key string) bool {
	for _, kvp := range kvps.GetKeyValuePair() {
		if kvp.GetKey() == key {
			return true
		}
	}
	return false
}

func (kvps *KeyValuePairs) GetValue(key string) (string, error) {
	for _, kvp := range kvps.GetKeyValuePair() {
		if kvp.GetKey() == key {
			return kvp.GetValue(), nil
		}
	}
	return "", errors.Errorf("key '%s' doesn't exists", key)
}
