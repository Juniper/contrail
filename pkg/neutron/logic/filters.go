package logic

// isPresentInFilters checks if keys are present in filters
// returns true when key is not present in filters or
// key is present in filters and value parameter matches filter's value list
func isPresentInFilters(filters Filters, key string, value string) bool {
	filter, ok := filters[key]
	if !ok {
		return true
	}

	if key == "tenant_id" {
		value = contrailUUIDToNeutronID(value)
	}

	for _, f := range filter {
		if f == value {
			return true
		}
	}

	return false
}
