package services

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Juniper/contrail/pkg/common"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/labstack/echo"
)

// MetaData represents resource meta data.
type MetaData struct {
	UUID   string
	FQName []string
	Type   string
}

type metadataGetter interface {
	GetMetaData(ctx context.Context, uuid string, fqName []string) (*MetaData, error)
}

// ContrailService handles API requests.
type ContrailService struct {
	BaseService

	MetadataGetter metadataGetter
	TypeValidator  *models.TypeValidator
}

// restSync handles Sync request.
func (service *ContrailService) restSync(c echo.Context) error {
	events := &EventList{}
	if err := c.Bind(events); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid JSON format: %v", err))
	}

	if err := events.Sort(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Sorting events failed: %v", err))
	}

	responses, err := events.Process(c.Request().Context(), service)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusCreated, responses.Events)
}
