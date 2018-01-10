package common

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"

	"github.com/labstack/echo"
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
		fmt.Println(api.Path())
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

func parseFlag(query string) bool {
	return query != ""
}

func parseStringList(query string) []string {
	if query == "" {
		return nil
	}
	return strings.Split(query, ",")
}

//GetListSpec makes ListSpec from Query Parameters
func GetListSpec(c echo.Context) *ListSpec {
	filter := ParseFilter(c.QueryParam("filters"))
	pageMarker := parsePositiveNumber(c.QueryParam("pageMarker"), 0)
	pageLimit := parsePositiveNumber(c.QueryParam("pageLimit"), 100)
	detail := parseBool(c.QueryParam("detail"))
	count := parseBool(c.QueryParam("is_count"))
	shared := parseBool(c.QueryParam("shared"))
	excludeHrefs := parseFlag(c.QueryParam("exclude_hrefs"))
	parentType := c.QueryParam("parent_type")
	parentFQName := ParseFQName(c.QueryParam("parent_fq_name_str"))
	parentUUIDs := parseStringList(c.QueryParam("parent_id"))
	backrefUUIDs := parseStringList(c.QueryParam("back_ref_id"))
	objectUUIDs := parseStringList(c.QueryParam("obj_uuids"))
	fields := parseStringList(c.QueryParam("fields"))
	return &ListSpec{
		Filter:       filter,
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
		BackrefUUIDs: backrefUUIDs,
		ObjectUUIDs:  objectUUIDs,
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
