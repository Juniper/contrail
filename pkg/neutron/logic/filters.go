package logic

// filtersHas checks if one or more keys are present in filters.
// Will return true if at least one key has been defined and all keys are present and not empty.
func filtersHas(filters Filters, keys ... string) bool {
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
