package common

import (
	"database/sql"
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

//RESTAPI defines handlers for REST API calls.
type RESTAPI interface {
	Path() string
	LongPath() string
	SetDB(db *sql.DB)
	Create(c echo.Context) error
	Update(c echo.Context) error
	Delete(c echo.Context) error
	List(c echo.Context) error
	Show(c echo.Context) error
}

//Filter is used to filter API response.
type Filter map[string][]string

//AppendValues appends filter values for key.
func (filter Filter) AppendValues(key string, values []string) {
	if filter == nil {
		return
	}
	if values == nil {
		return
	}
	f, ok := filter[key]
	if !ok {
		f = []string{}
	}
	filter[key] = append(f, values...)
}

var apiRegistory = map[string]RESTAPI{}

//RegisterAPI to add new API for API Registory
func RegisterAPI(api RESTAPI) {
	apiRegistory[api.Path()] = api
}

//Routes registers routes
func Routes(e *echo.Echo) {
	for _, api := range apiRegistory {
		e.POST(api.Path(), api.Create)
		e.PUT(api.LongPath(), api.Update)
		e.DELETE(api.LongPath(), api.Delete)
		e.GET(api.Path(), api.List)
		e.GET(api.LongPath(), api.Show)
	}
}

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

//GetListSpec makes ListSpec from Query Parameters
func GetListSpec(c echo.Context) *ListSpec {
	filter := ParseFilter(c.QueryParam(FiltersKey))
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
		Filter:          filter,
		RequestedFields: fields,
		ParentType:      parentType,
		ParentFQName:    parentFQName,
		Limit:           pageLimit,
		Offset:          pageMarker,
		Detail:          detail,
		Count:           count,
		ExcludeHrefs:    excludeHrefs,
		Shared:          shared,
		ParentUUIDs:     parentUUIDs,
		BackRefUUIDs:    backrefUUIDs,
		ObjectUUIDs:     objectUUIDs,
	}
}

//ParseFilter makes Filter from comma separated string.
//Eg. check==a,check==b,name==Bob
func ParseFilter(filterString string) Filter {
	filter := Filter{}
	if filterString == "" {
		return filter
	}
	parts := strings.Split(filterString, ",")
	for _, part := range parts {
		keyValue := strings.Split(part, "==")
		if len(keyValue) != 2 {
			continue
		}
		key := keyValue[0]
		value := keyValue[1]
		filterForKey, ok := filter[key]
		if !ok {
			filterForKey = []string{}
		}
		filterForKey = append(filterForKey, value)
		filter[key] = filterForKey
	}
	return filter
}
