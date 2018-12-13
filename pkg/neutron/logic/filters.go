package logic

// filtersHas checks if one or more keys are present in filters.
// Will return true if at least one key has been defined and all keys are present and not empty.
func filtersHas(filters Filters, keys ...string) bool {
	if len(keys) == 0 {
		return false
	}

	for _, key := range keys {
		filter, ok := filters[key]
		if !ok || len(filter) == 0 {
			return false
		}
	}

	return true
}

// checkFilterValue check equality of values in filters struct under specific key and provided sequence of strings
func checkFilterValue(filters Filters, key string, values ...string) bool {
	if !filtersHas(filters, key) {
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
