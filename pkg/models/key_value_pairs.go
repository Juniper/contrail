package models

import (
	"github.com/Juniper/asf/pkg/services"
)

// FromDBKeyValuePairs creates slice of KeyValuePairs from db.KeyValuePairs format.
func FromDBKeyValuePairs(kvs []*services.KeyValuePair) []*KeyValuePair {
	result := make([]*KeyValuePair, len(kvs))
	for i, kv := range kvs {
		result[i].Key = kv.Key
		result[i].Value = kv.Value
	}
	return result
}

// GetValue search for specified key and returns its value.
func (kvps *KeyValuePairs) GetValue(key string) string {
	for _, kvp := range kvps.GetKeyValuePair() {
		if kvp.GetKey() == key {
			return kvp.GetValue()
		}
	}
	return ""
}
