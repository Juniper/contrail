package services

import (
	"strings"

	"github.com/Juniper/contrail/pkg/common"
	"github.com/Juniper/contrail/pkg/serviceif"
	log "github.com/sirupsen/logrus"
	context "golang.org/x/net/context"
)

//Resource events.
const (
	EventCreate = "CREATE"
	EventUpdate = "UPDATE"
	EventDelete = "DELETE"
)

//ResourceEvent represents REST resource event.
type ResourceEvent struct {
	Kind string `json:"kind"`
	//Sadly, we can't use Resource interface here to make JSON/YAML pacakge work.
	Data      interface{} `json:"data"`
	Operation string      `json:"operation"`
}

//Resource is a generic resource interface.
type Resource interface {
	GetUUID() string
	Depends() []string
	ToMap() map[string]interface{}
}

//ResourceList has multiple rest requests.
type ResourceList struct {
	Resources []*ResourceEvent `json:"resources"`
}

//ToMap returns map format of data.
func (r ResourceEvent) ToMap() map[string]interface{} {
	m, ok := r.Data.(map[string]interface{})
	if ok {
		return m
	}
	resource, ok := r.Data.(Resource)
	if ok {
		return resource.ToMap()
	}
	return nil
}

//GetUUID returns uuid of resource.
func (r ResourceEvent) GetUUID() string {
	resource, ok := r.Data.(Resource)
	if ok {
		return resource.GetUUID()
	}
	return (r.ToMap())["uuid"].(string)
}

//Depends returns dependencies.
func (r ResourceEvent) Depends() []string {
	resource, ok := r.Data.(Resource)
	if ok {
		return resource.Depends()
	}
	depends := []string{}
	for key, value := range r.ToMap() {
		if strings.HasSuffix(key, "_refs") {
			refs, ok := value.([]interface{})
			if !ok {
				return nil
			}
			for _, ref := range refs {
				m, ok := ref.(map[string]interface{})
				if !ok {
					return nil
				}
				refUUID, ok := m["uuid"].(string)
				if !ok {
					return nil
				}
				depends = append(depends, refUUID)
			}
		} else if key == "parent_uuid" {
			refUUID, ok := value.(string)
			if !ok {
				return nil
			}
			depends = append(depends, refUUID)
		}
	}
	return depends
}

type state int

const (
	notVisited state = iota
	visited
	temporaryVisited
)

//reorder request using Tarjan's algorithm
func visitResource(uuid string, sorted []*ResourceEvent,
	resourceMap map[string]*ResourceEvent, stateGraph map[string]state) (sortedList []*ResourceEvent, err error) {
	if stateGraph[uuid] == temporaryVisited {
		return nil, common.ErrorBadRequest("dependency loop found in sync request.")
	}
	if stateGraph[uuid] == visited {
		return sorted, nil
	}
	stateGraph[uuid] = temporaryVisited
	resource := resourceMap[uuid]
	depends := resource.Depends()
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
func (r *ResourceList) Sort() (err error) {
	sorted := []*ResourceEvent{}
	stateGraph := map[string]state{}
	resourceMap := map[string]*ResourceEvent{}
	for _, resource := range r.Resources {
		uuid := resource.GetUUID()
		stateGraph[uuid] = notVisited
		resourceMap[uuid] = resource
	}
	foundNotVisited := true
	for foundNotVisited {
		foundNotVisited = false
		for _, resource := range r.Resources {
			uuid := resource.GetUUID()
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
	r.Resources = sorted
	return nil
}

//YAMLCompat translate yaml data
func (r *ResourceList) YAMLCompat() {
	for _, resource := range r.Resources {
		resource.Data = common.YAMLtoJSONCompat(resource.Data)
	}
}

//Process applies events for each service.
func (r *ResourceList) Process(ctx context.Context, service serviceif.Service) ([]*ResourceEvent, error) {
	responses := []*ResourceEvent{}
	for _, resource := range r.Resources {
		response, err := resource.Process(ctx, service)
		if err != nil {
			log.Debug(resource, err)
			return nil, err
		}
		responses = append(responses, response)
	}
	return responses, nil
}
