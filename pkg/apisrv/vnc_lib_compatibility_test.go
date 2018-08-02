package apisrv_test

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"net"
	"net/http"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"

	"github.com/Juniper/contrail/pkg/db"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
	"github.com/Juniper/contrail/pkg/testutil/integration"
)

func TestVncLibCompatibility(t *testing.T) {
	s := integration.NewRunningAPIServer(t, "../../..", db.DriverPostgreSQL, false)
	defer s.Close(t)

	h := integration.NewHTTPAPIClient(t, s.URL())

	projID := "vnc-compat-proj-uuid"
	data := &services.CreateProjectRequest{
		Project: &models.Project{
			UUID:       projID,
			ParentType: integration.DomainType,
			ParentUUID: integration.DefaultDomainUUID,
			FQName:     []string{integration.DefaultDomainID, integration.AdminProjectID, projID},
			Quota:      &models.QuotaType{},
		},
	}

	{
		var responseData interface{}
		expected := []int{http.StatusCreated}
		resp, err := h.Do(echo.POST, "/projects", data, &responseData, expected)
		if assert.NoError(t, err, "create project failed\n") {
			defer resp.Body.Close()
			h.DeleteProject(t, projID)
		}
	}

	c := &http.Client{
		Transport: &http.Transport{
			Dial:            (&net.Dialer{}).Dial,
			TLSClientConfig: &tls.Config{InsecureSkipVerify: h.InSecure},
		},
	}
	dataJSON, _ := json.Marshal(data)
	request, _ := http.NewRequest(echo.POST, s.URL()+"/projects", bytes.NewBuffer(dataJSON))

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("X-Auth-Token", h.AuthToken)
	request.Header.Set("X-Contrail-Useragent", "nonempty")

	resp, err := c.Do(request)
	if assert.NoError(t, err, "issuing HTTP request failed") {
		defer resp.Body.Close()
		h.DeleteProject(t, projID)
		assert.Equal(t, http.StatusOK, resp.StatusCode)
	}
}
