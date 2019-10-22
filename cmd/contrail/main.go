package main

import (
	"github.com/Juniper/contrail/pkg/cmd/contrail"
	"github.com/Juniper/contrail/pkg/logutil"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	echoRedirect := echo.New()
	echoRedirect.Pre(middleware.HTTPSRedirect())

	err := contrail.Contrail.Execute()
	if err != nil {
		logutil.FatalWithStackTrace(err)
	}
}
