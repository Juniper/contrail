package agent

import (
	"errors"
	"fmt"
	"reflect"
	"time"

	pkglog "github.com/Juniper/contrail/pkg/log"
	"github.com/sirupsen/logrus"
)

const (
	actionCreate = "create"
	actionUpdate = "update"
	actionDelete = "delete"
)

type watcher interface {
	watch()
}

type pollingWatcher struct {
	conf      *task
	agent     *Agent
	SchemaID  string
	resources map[string]map[string]interface{}
	log       *logrus.Entry
}

func (w *pollingWatcher) Action(actionType string, resource map[string]interface{}) error {
	return w.conf.action(actionType, resource)
}

// nolint: gocyclo
func (w *pollingWatcher) Sync() error {
	resources := w.resources
	//TODO(nati) Proper stop support using channel
	var list []interface{}
	w.log.Debug("Polling data")
	_, err := w.agent.APIServer.Read(w.agent.schemas[w.SchemaID].PluralPath, &list)
	if err != nil {
		return err
	}
	idsExistsInServer := map[string]bool{}
	for _, rawResource := range list {
		resource, ok := rawResource.(map[string]interface{})
		if !ok {
			return fmt.Errorf("invalid resource type")
		}
		id, ok := resource["id"].(string)
		if !ok {
			return fmt.Errorf("invalid resource type")
		}
		idsExistsInServer[id] = true
		existingResource, exists := resources[resource["id"].(string)]
		if exists {
			if !reflect.DeepEqual(resource, existingResource) {
				err = w.Action(actionUpdate, resource)
				if err != nil {
					w.log.WithError(err).Error()
				}
				resources[id] = resource
			}
			continue
		}
		// store the exact schemaID instead of pattern
		resource["schema_id"] = w.SchemaID
		err = w.Action(actionCreate, resource)
		if err != nil {
			w.log.WithError(err).Error()
		}
		resources[id] = resource
	}
	// Handling for deletion
	for id, existingResource := range resources {
		if _, ok := idsExistsInServer[id]; ok {
			continue
		}
		err = w.Action(actionDelete, existingResource)
		if err != nil {
			w.log.WithError(err).Error()
		}
		delete(resources, id)
	}
	return nil
}

func (w *pollingWatcher) watch() {
	//TODO(nati) proper error handing
	//TODO(nati) support parallel execution and lock
	for {
		time.Sleep(time.Second)
		err := w.Sync()
		if err != nil {
			w.agent.APIServer.Login() // nolint: errcheck
		}
	}
}

func newPollingWatcher(agent *Agent, task *task, schemaID string) (*pollingWatcher, error) {
	return &pollingWatcher{
		agent:     agent,
		conf:      task,
		SchemaID:  schemaID,
		resources: map[string]map[string]interface{}{},
		log:       pkglog.NewLogger("polling-watcher"),
	}, nil
}

func newWatcher(agent *Agent, task *task, schemaID string) (watcher, error) {
	switch agent.config.Watcher {
	case "polling":
		return newPollingWatcher(agent, task, schemaID)
	}
	//TODO(nati) Support binlog based gohan watch API
	return nil, errors.New("unsupported watcher type")
}
