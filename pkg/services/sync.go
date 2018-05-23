package services

import (
	"net/http"
	"strings"

	"github.com/Juniper/contrail/pkg/common"
	"github.com/Juniper/contrail/pkg/serviceif"
	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"
	context "golang.org/x/net/context"
)

const (
	create = "CREATE"
	update = "UPDATE"
	delete = "DELETE"
)

// nolint
type ContrailService struct {
	serviceif.BaseService
}

//ResourceEvent represents REST resource event.
type ResourceEvent struct {
	Kind      string      `json:"kind"`
	Data      interface{} `json:"data"`
	Operation string      `json:"operation"`
}

//ResourceList has multiple rest requests.
type ResourceList struct {
	Resources []*ResourceEvent `json:"resources"`
}

func (resource ResourceEvent) data() map[string]interface{} {
	return resource.Data.(map[string]interface{})
}

func (resource ResourceEvent) uuid() string {
	return (resource.data())["uuid"].(string)
}

func (resource ResourceEvent) depends() ([]string, error) {
	depends := []string{}
	for key, value := range resource.data() {
		if strings.HasSuffix(key, "_refs") {
			refs, ok := value.([]interface{})
			if !ok {
				return nil, common.ErrorBadRequest("malformed sync request")
			}
			for _, ref := range refs {
				m, ok := ref.(map[string]interface{})
				if !ok {
					return nil, common.ErrorBadRequest("malformed sync request")
				}
				refUUID, ok := m["uuid"].(string)
				if !ok {
					return nil, common.ErrorBadRequest("malformed sync request")
				}
				depends = append(depends, refUUID)
			}
		} else if key == "parent_uuid" {
			refUUID, ok := value.(string)
			if !ok {
				return nil, common.ErrorBadRequest("malformed sync request")
			}
			depends = append(depends, refUUID)
		}
	}
	return depends, nil
}

//reorder request using Tarjan's algorithm
type state int

const (
	notVisited state = iota
	visited
	temporaryVisited
)

func visitResource(uuid string, sorted []*ResourceEvent,
	resourceMap map[string]*ResourceEvent, stateGraph map[string]state) ([]*ResourceEvent, error) {
	if stateGraph[uuid] == temporaryVisited {
		return nil, common.ErrorBadRequest("dependency loop found in sync request.")
	}
	if stateGraph[uuid] == visited {
		return sorted, nil
	}
	stateGraph[uuid] = temporaryVisited
	resource := resourceMap[uuid]
	depends, err := resource.depends()
	if err != nil {
		return nil, err
	}
	for _, refUUID := range depends {
		sorted, err = visitResource(refUUID, sorted, resourceMap, stateGraph)
		if err != nil {
			return nil, err
		}
		break
	}
	stateGraph[uuid] = visited
	sorted = append(sorted, resource)
	return sorted, nil
}

//Sort sorts request by dependency using Tarjan's algorithm
func (request *ResourceList) Sort() (err error) {
	sorted := []*ResourceEvent{}
	stateGraph := map[string]state{}
	resourceMap := map[string]*ResourceEvent{}
	for _, resource := range request.Resources {
		uuid := resource.uuid()
		stateGraph[uuid] = notVisited
		resourceMap[uuid] = resource
	}
	foundNotVisited := true
	for foundNotVisited {
		foundNotVisited = false
		for _, resource := range request.Resources {
			uuid := resource.uuid()
			state := stateGraph[uuid]
			if state == notVisited {
				sorted, err = visitResource(uuid, sorted, resourceMap, stateGraph)
				if err != nil {
					return err
				}
				foundNotVisited = true
				break
			}
		}
	}
	request.Resources = sorted
	return nil
}

//YAMLCompat translate yaml data
func (request *ResourceList) YAMLCompat() {
	for _, resource := range request.Resources {
		resource.Data = common.YAMLtoJSONCompat(resource.Data)
	}
}

//Process applies events for each service.
func (request *ResourceList) Process(ctx context.Context, service serviceif.Service) ([]*ResourceEvent, error) {
	responses := []*ResourceEvent{}
	for _, resource := range request.Resources {
		response, err := resource.Process(ctx, service)
		if err != nil {
			log.Debug(resource, err)
			return nil, err
		}
		responses = append(responses, response)
	}
	return responses, nil
}

//RESTSync handle a bluk Create REST service.
func (service *ContrailService) RESTSync(c echo.Context) error {
	requestData := &ResourceList{}
	if err := c.Bind(requestData); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "{{ schema.ID }}",
		}).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	responses, err := requestData.Process(ctx, service)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, responses)
}
