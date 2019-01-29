package baseservices

import (
	"net/url"
	"strconv"
	"strings"

	"github.com/labstack/echo"

	"github.com/Juniper/contrail/pkg/models/basemodels"
)

// Block of constants consumed by client code.
const (
	FiltersKey         = "filters"
	PageMarkerKey      = "page_marker"
	PageLimitKey       = "page_limit"
	DetailKey          = "detail"
	CountKey           = "count"
	SharedKey          = "shared"
	ExcludeHRefsKey    = "exclude_hrefs"
	ParentTypeKey      = "parent_type"
	ParentFQNameKey    = "parent_fq_name_str"
	ParentUUIDsKey     = "parent_id"
	BackrefUUIDsKey    = "back_ref_id"
	ObjectUUIDsKey     = "obj_uuids"
	FieldsKey          = "fields"
	ExcludeChildrenKey = "exclude_children"
	ExcludeBackRefsKey = "exclude_back_refs"
)

func parsePositiveNumber(query string, defaultValue int64) int64 {
	i, err := strconv.Atoi(query)
	if err != nil {
		return defaultValue
	}
	if i < 0 {
		return defaultValue
	}
	return int64(i)
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

// GetListSpec makes ListSpec from Query Parameters
func GetListSpec(c echo.Context) *ListSpec {
	filters := ParseFilter(c.QueryParam(FiltersKey))
	pageMarker := c.QueryParam(PageMarkerKey)
	pageLimit := parsePositiveNumber(c.QueryParam(PageLimitKey), 0)
	detail := parseBool(c.QueryParam(DetailKey))
	count := parseBool(c.QueryParam(CountKey))
	shared := parseBool(c.QueryParam(SharedKey))
	excludeHrefs := parseBool(c.QueryParam(ExcludeHRefsKey))
	parentType := c.QueryParam(ParentTypeKey)
	parentFQName := basemodels.ParseFQName(c.QueryParam(ParentFQNameKey))
	parentUUIDs := parseStringList(c.QueryParam(ParentUUIDsKey))
	backrefUUIDs := parseStringList(c.QueryParam(BackrefUUIDsKey))
	objectUUIDs := parseStringList(c.QueryParam(ObjectUUIDsKey))
	excludeChildren := parseBool(c.QueryParam(ExcludeChildrenKey))
	excludeBackRefs := parseBool(c.QueryParam(ExcludeBackRefsKey))
	fields := parseStringList(c.QueryParam(FieldsKey))
	return &ListSpec{
		Filters:         filters,
		Fields:          fields,
		ParentType:      parentType,
		ParentFQName:    parentFQName,
		Limit:           pageLimit,
		Marker:          pageMarker,
		Detail:          detail,
		Count:           count,
		ExcludeHrefs:    excludeHrefs,
		Shared:          shared,
		ParentUUIDs:     parentUUIDs,
		BackRefUUIDs:    backrefUUIDs,
		ObjectUUIDs:     objectUUIDs,
		ExcludeChildren: excludeChildren,
		ExcludeBackRefs: excludeBackRefs,
	}
}

// URLQuery returns URL query strings.
func (s *ListSpec) URLQuery() url.Values {
	if s == nil {
		return nil
	}
	query := url.Values{}
	addQuery(query, FiltersKey, EncodeFilter(s.Filters))
	if s.Marker != "" {
		addQuery(query, PageMarkerKey, s.Marker)
	}
	if s.Limit > 0 {
		addQuery(query, PageLimitKey, strconv.FormatInt(s.Limit, 10))
	}
	addQueryBool(query, DetailKey, s.Detail)
	addQueryBool(query, CountKey, s.Count)
	addQueryBool(query, SharedKey, s.Shared)
	addQueryBool(query, ExcludeHRefsKey, s.ExcludeHrefs)
	addQuery(query, ParentTypeKey, s.ParentType)
	addQuery(query, ParentFQNameKey, basemodels.FQNameToString(s.ParentFQName))
	addQuery(query, ParentUUIDsKey, encodeStringList(s.ParentUUIDs))
	addQuery(query, BackrefUUIDsKey, encodeStringList(s.BackRefUUIDs))
	addQuery(query, ObjectUUIDsKey, encodeStringList(s.ObjectUUIDs))
	addQuery(query, FieldsKey, encodeStringList(s.Fields))
	addQueryBool(query, ExcludeChildrenKey, s.ExcludeChildren)
	addQueryBool(query, ExcludeBackRefsKey, s.ExcludeBackRefs)
	return query
}

func addQuery(query url.Values, key, value string) {
	if value != "" {
		query.Add(key, value)
	}
}

func addQueryBool(query url.Values, key string, value bool) {
	if value {
		query.Add(key, strconv.FormatBool(value))
	}
}

// AppendFilter return a filter for specific key.
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

// QueryString returns string for query string.
func (f *Filter) QueryString() string {
	var sl []string
	for _, value := range f.Values {
		sl = append(sl, f.Key+"=="+value)
	}
	return encodeStringList(sl)
}

// ParseFilter makes Filter from comma separated string.
// Eg. check==a,check==b,name==Bob
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

// EncodeFilter encodes filter to string.
func EncodeFilter(filters []*Filter) string {
	var sl []string
	for _, filter := range filters {
		sl = append(sl, filter.QueryString())
	}
	return encodeStringList(sl)
}

func encodeStringList(s []string) string {
	return strings.Join(s, ",")
}
