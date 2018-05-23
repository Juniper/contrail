package services

import (
	"strings"

	"github.com/Juniper/contrail/pkg/common"
)

type data interface{}

//Resource represents REST resource request.
type Resource struct {
	Kind string `json:"kind"`
	Data data   `json:"data"`
}

//ResourceList has multiple rest requests.
type ResourceList struct {
	Resources []*Resource `json:"resources"`
}

func (resource Resource) data() map[string]interface{} {
	return resource.Data.(map[string]interface{})
}

func (resource Resource) uuid() string {
	return (resource.data())["uuid"].(string)
}

func (resource Resource) depends() ([]string, error) {
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

func visitResource(uuid string, sorted []*Resource,
	resourceMap map[string]*Resource, stateGraph map[string]state) ([]*Resource, error) {
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
	sorted := []*Resource{}
	stateGraph := map[string]state{}
	resourceMap := map[string]*Resource{}
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
