package common

import (
	"database/sql"

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
