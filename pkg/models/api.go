package models

import (
	"strconv"
	"strings"

	"github.com/labstack/echo"
)

// Block of constants consumed by client code.
const (
	FiltersKey      = "filters"
	PageMarkerKey   = "pageMarker"
	PageLimitKey    = "pageLimit"
	DetailKey       = "detail"
	CountKey        = "is_count"
	SharedKey       = "shared"
	ExcludeHRefsKey = "exclude_hrefs"
	ParentTypeKey   = "parent_type"
	ParentFQNameKey = "parent_fq_name_str"
	ParentUUIDsKey  = "parent_id"
	BackrefUUIDsKey = "back_ref_id"
	ObjectUUIDsKey  = "obj_uuids"
	FieldsKey       = "fields"
)

func parsePositiveNumber(query string, defaultValue int) int {
	i, err := strconv.Atoi(query)
	if err != nil {
		return defaultValue
	}
	if i < 0 {
		return defaultValue
	}
	return i
}

func parseBool(query string) bool {
	return strings.ToLower(query) == "true"
}

func parseStringList(query string) []string {
	if query == "" {
		return nil
	}
	return strings.Split(query, ",")
}

//ParseFQName parse string representation of FQName.
func ParseFQName(fqNameString string) []string {
	if fqNameString == "" {
		return nil
	}
	return strings.Split(fqNameString, ":")
}

//GetListSpec makes ListSpec from Query Parameters
func GetListSpec(c echo.Context) *ListSpec {
	filters := ParseFilter(c.QueryParam(FiltersKey))
	pageMarker := parsePositiveNumber(c.QueryParam(PageMarkerKey), 0)
	pageLimit := parsePositiveNumber(c.QueryParam(PageLimitKey), 100)
	detail := parseBool(c.QueryParam(DetailKey))
	count := parseBool(c.QueryParam(CountKey))
	shared := parseBool(c.QueryParam(SharedKey))
	excludeHrefs := parseBool(c.QueryParam(ExcludeHRefsKey))
	parentType := c.QueryParam(ParentTypeKey)
	parentFQName := ParseFQName(c.QueryParam(ParentFQNameKey))
	parentUUIDs := parseStringList(c.QueryParam(ParentUUIDsKey))
	backrefUUIDs := parseStringList(c.QueryParam(BackrefUUIDsKey))
	objectUUIDs := parseStringList(c.QueryParam(ObjectUUIDsKey))
	fields := parseStringList(c.QueryParam(FieldsKey))
	return &ListSpec{
		Filters:      filters,
		Fields:       fields,
		ParentType:   parentType,
		ParentFQName: parentFQName,
		Limit:        pageLimit,
		Offset:       pageMarker,
		Detail:       detail,
		Count:        count,
		ExcludeHrefs: excludeHrefs,
		Shared:       shared,
		ParentUUIDs:  parentUUIDs,
		BackRefUUIDs: backrefUUIDs,
		ObjectUUIDs:  objectUUIDs,
	}
}

//AppendFilter return a filter for specific key.
func AppendFilter(filters []*Filter, key string, values ...string) []*Filter {
	var filter *Filter
	if len(values) == 0 {
		return filters
	}
	for _, f := range filters {
		if f.Key == key {
			filter = f
			break
		}
	}
	if filter == nil {
		filter = &Filter{
			Key:    key,
			Values: []string{},
		}
		filters = append(filters, filter)
	}
	filter.Values = append(filter.Values, values...)
	return filters
}

//ParseFilter makes Filter from comma separated string.
//Eg. check==a,check==b,name==Bob
func ParseFilter(filterString string) []*Filter {
	filters := []*Filter{}
	if filterString == "" {
		return filters
	}
	parts := strings.Split(filterString, ",")
	for _, part := range parts {
		keyValue := strings.Split(part, "==")
		if len(keyValue) != 2 {
			continue
		}
		key := keyValue[0]
		value := keyValue[1]
		filters = AppendFilter(filters, key, value)
	}
	return filters
}
