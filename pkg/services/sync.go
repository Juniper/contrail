package services

import (
	"strings"

	"github.com/Juniper/contrail/pkg/common"
)

type data interface{}

//RESTResource represents REST resource request.
type RESTResource struct {
	Kind string `json:"kind"`
	Data data   `json:"data"`
}

//RESTSyncRequest has multiple rest requests.
type RESTSyncRequest struct {
	Resources []*RESTResource `json:"resources"`
}

func (resource RESTResource) data() map[string]interface{} {
	return resource.Data.(map[string]interface{})
}

func (resource RESTResource) uuid() string {
	return (resource.data())["uuid"].(string)
}

func (resource RESTResource) depends() ([]string, error) {
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

func visitResource(uuid string, sorted []*RESTResource,
	resourceMap map[string]*RESTResource, stateGraph map[string]state) ([]*RESTResource, error) {
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
func (request *RESTSyncRequest) Sort() (err error) {
	sorted := []*RESTResource{}
	stateGraph := map[string]state{}
	resourceMap := map[string]*RESTResource{}
	for _, resource := range request.Resources {
		uuid := resource.uuid()
		stateGraph[uuid] = notVisited
		resourceMap[uuid] = resource
	}
	foundNotVisited := true
	for foundNotVisited {
		foundNotVisited = false
		for uuid, state := range stateGraph {
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
