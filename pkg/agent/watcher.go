package agent

import (
	"context"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	"github.com/Juniper/contrail/pkg/logutil"
)

const (
	actionCreate = "create"
	actionUpdate = "update"
	actionDelete = "delete"
)

type watcher interface {
	watch(ctx context.Context)
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
func (w *pollingWatcher) Sync(ctx context.Context) error {
	resources := w.resources
	//TODO(nati) Proper stop support using channel
	var list map[string][]interface{}
	w.log.Debug("Polling data")
	resourcePath := w.agent.schemas[w.SchemaID].PluralPath
	_, err := w.agent.APIServer.Read(ctx, resourcePath, &list)
	if err != nil {
		return err
	}
	idsExistsInServer := map[string]bool{}
	for _, rawResource := range list[strings.TrimLeft(resourcePath, "/")] {
		resource, ok := rawResource.(map[string]interface{})
		if !ok {
			return fmt.Errorf("invalid resource type")
		}
		id, ok := resource["uuid"].(string)
		if !ok {
			return fmt.Errorf("invalid resource type")
		}
		idsExistsInServer[id] = true
		existingResource, exists := resources[id]
		if exists {
			/* set the exact schemaID in the read resource
			   before comparing it with the existing resource
			*/
			resource["schema_id"] = w.SchemaID
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

func (w *pollingWatcher) watch(ctx context.Context) {
	//TODO(nati) proper error handing
	//TODO(nati) support parallel execution and lock
	for {
		time.Sleep(time.Second)
		err := w.Sync(ctx)
		if err != nil {
			_, _ := w.agent.APIServer.Login(ctx) // nolint: errcheck
		}
	}
}

func newPollingWatcher(agent *Agent, task *task, schemaID string) (*pollingWatcher, error) {
	return &pollingWatcher{
		agent:     agent,
		conf:      task,
		SchemaID:  schemaID,
		resources: map[string]map[string]interface{}{},
		log:       logutil.NewLogger("polling-watcher"),
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
