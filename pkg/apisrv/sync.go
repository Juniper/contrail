package apisrv

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

//SyncRequest contains list of sync.
type SyncRequest struct {
	Resources []*SyncResource `json:"resources"`
}

//SyncResource contains one of sync resource.
type SyncResource struct {
	uuid string
	Type string                 `json:"type"`
	Body map[string]interface{} `json:"body"`
}

//UUID returns uuid of the resource
func (r *SyncResource) UUID() string {
	if r.uuid != "" {
		return r.uuid
	}
	uuid, _ := r.Body["uuid"].(string)
	r.uuid = uuid
	return r.uuid
}

func syncCreateOrUpdate(c echo.Context) error {
	var errors []error
	var syncRequest SyncRequest
	if err := c.Bind(&syncRequest); err != nil {
		log.Debug(err)
		return echo.NewHTTPError(http.StatusBadRequest, "malformed sync request")
	}
	resources := syncRequest.Resources
	doneIndex := 0
	var resource *SyncResource
	for doneIndex, resource = range resources {
		handler := dispatchResourceHandler(c, resource)
		err := handler.CreateOrUpdate(resource)
		if err != nil {
			log.Debug(err)
			errors = append(errors, err)
			break
		}
	}
	if len(errors) != 0 {
		log.Debug("trying rollback on sync API")
		for i := doneIndex; i >= 0; i-- {
			resource = resources[i]
			handler := dispatchResourceHandler(c, resource)
			err := handler.Delete(resource)
			if err != nil {
				log.Debug(err)
				errors = append(errors, err)
			}
		}
	}
	if len(errors) != 0 {
		return fmt.Errorf("sync failed %v", errors)
	}
	return c.JSON(http.StatusCreated, syncRequest)
}

//syncDelete deletes resources.
func syncDelete(c echo.Context) error {
	var errors []error
	var syncRequest SyncRequest
	if err := c.Bind(&syncRequest); err != nil {
		log.Debug(err)
		return echo.NewHTTPError(http.StatusBadRequest, "malformed sync request")
	}
	resources := syncRequest.Resources
	for i := len(resources) - 1; i >= 0; i-- {
		resource := resources[i]
		handler := dispatchResourceHandler(c, resource)
		err := handler.Delete(resource)
		if err != nil {
			log.Debug(err)
			errors = append(errors, err)
		}
	}
	if len(errors) != 0 {
		return fmt.Errorf("sync delete failed %v", errors)
	}
	return c.JSON(http.StatusNoContent, nil)
}

//ResourceHandler handle southbound resource management.
type ResourceHandler interface {
	CreateOrUpdate(*SyncResource) error
	Delete(*SyncResource) error
}

func dispatchResourceHandler(c echo.Context, resource *SyncResource) ResourceHandler {
	tokenString := c.Request().Header.Get("X-Auth-Token")
	contrailBackend := viper.GetString("sync.contrail")
	client := NewClientForProxy(contrailBackend, tokenString)
	return &ContrailHandler{
		client: client,
	}
}

//ContrailHandler handle contrail VNC objects.
type ContrailHandler struct {
	client *Client
}

//CreateOrUpdate tries to create or update resources.
func (h *ContrailHandler) CreateOrUpdate(resource *SyncResource) error {
	uuid := resource.UUID()
	path := resource.Type + "/" + uuid
	body := map[string]interface{}{
		resource.Type: resource.Body,
	}
	var readOutput map[string]interface{}
	var output map[string]interface{}
	_, err := h.client.Read(path, &readOutput)
	log.Debug(err)
	if err == nil {
		_, err = h.client.Create(resource.Type+"s", body, &output)
	} else {
		_, err = h.client.Update(path, body, &output)
	}
	outputResource, _ := output[resource.Type].(map[string]interface{})
	log.Debug("output resource", output)
	resource.Body = outputResource
	if err != nil {
		return err
	}
	return nil
}

//Delete tries to delete resoruces
func (h *ContrailHandler) Delete(resource *SyncResource) error {
	uuid := resource.UUID()
	path := resource.Type + "/" + uuid
	_, err := h.client.Delete(path, nil)
	return err
}
