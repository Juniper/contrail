package logic

// haveKeys checks if one or more keys are present in filters.
// Will return true if at least one key has been defined and all keys are present and not empty.
func (f Filters) haveKeys(keys ...string) bool {
	if len(keys) == 0 {
		return false
	}

	for _, key := range keys {
		filter, ok := f[key]
		if !ok || len(filter) == 0 {
			return false
		}
	}

	return true
}

// checkValue check equality of values in filters struct under specific key and provided sequence of strings
func (filters Filters) checkValue(key string, values ...string) bool {
	if !filters.haveKeys(key) {
		return true
	}
	if len(filters[key]) != len(values) {
		return false
	}

	for i, v := range values {
		if filters[key][i] != v {
			return false
		}
	}

	return true
}
