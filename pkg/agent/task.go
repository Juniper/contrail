package agent

import (
	"fmt"

	"github.com/flosch/pongo2"
	"github.com/pkg/errors"
	yaml "gopkg.in/yaml.v2"
)

type task struct {
	SchemaIDs     []string  `yaml:"schema_ids"`
	Commands      []string  `yaml:"commands"`
	Common        []handler `yaml:"common"`
	OnCreate      []handler `yaml:"on_create"`
	OnUpdate      []handler `yaml:"on_update"`
	OnDelete      []handler `yaml:"on_delete"`
	OutputPath    string    `yaml:"output_path"`
	WorkDirectory string    `yaml:"work_directory"`
	agent         *Agent
}

func (task *task) runHandlers(handlers []handler, context map[string]interface{}) error {
	for index, handler := range handlers {
		err := task.runHandler(handler, context)
		if err != nil {
			return errors.Wrap(err, fmt.Sprintf("[line %d]", index))
		}
	}
	return nil
}

func (task *task) runHandler(handler handler, context map[string]interface{}) error {
	for id := range handler {
		if handlerFunc, ok := globalTaskHandler[id]; ok {
			output, err := handlerFunc(handler, task, context)
			if err != nil {
				byteHandler, _ := yaml.Marshal(handler) // nolint: errcheck
				return errors.Wrap(err, fmt.Sprintf("%voutput:%s", string(byteHandler), output))
			}
			register, ok := handler["register"]
			if ok {
				context[register.(string)] = output //nolint: errcheck
			}
			return nil
		}
	}
	return fmt.Errorf("task handler not found: %v", handler)
}

func applyTemplate(rawTemplate interface{}, context map[string]interface{}) (string, error) {
	if rawTemplate == nil {
		return "", nil
	}
	templateString, ok := rawTemplate.(string) //nolint: errcheck
	if !ok {
		return "", errors.New("invalid template string")
	}
	template, err := pongo2.FromString(templateString)
	if err != nil {
		return "", err
	}
	return template.Execute(context)
}

func (task *task) init(agent *Agent) {
	task.agent = agent
}

func (task *task) action(action string, resource map[string]interface{}) error {
	context := pongo2.Context{
		"resource": resource,
		"action":   action,
		"input":    task.agent.config,
	}
	err := task.runHandlers(task.Common, context)
	if err != nil {
		return err
	}
	switch action {
	case actionCreate:
		return task.runHandlers(task.OnCreate, context)
	case actionUpdate:
		return task.runHandlers(task.OnUpdate, context)
	case actionDelete:
		return task.runHandlers(task.OnDelete, context)
	}
	return nil
}
