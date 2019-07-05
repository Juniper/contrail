package client

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"path/filepath"
	"strconv"
	"time"

	"github.com/Juniper/contrail/pkg/fileutil"
	"github.com/Juniper/contrail/pkg/keystone"
	"github.com/Juniper/contrail/pkg/logutil"
	"github.com/Juniper/contrail/pkg/models/basemodels"
	"github.com/Juniper/contrail/pkg/schema"
	"github.com/Juniper/contrail/pkg/services"
	"github.com/Juniper/contrail/pkg/services/baseservices"
	"github.com/flosch/pongo2"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	yaml "gopkg.in/yaml.v2"
)

const (
	retryMax         = 5
	serverSchemaFile = "schema.json"
)

// CLI represents API Server's command line interface.
type CLI struct {
	HTTP

	schemaRoot string
	log        *logrus.Entry
}

// NewCLIByViper returns new logged in CLI client using Viper configuration.
func NewCLIByViper() (*CLI, error) {
	return NewCLI(
		&HTTPConfig{
			ID:       viper.GetString("client.id"),
			Password: viper.GetString("client.password"),
			Endpoint: viper.GetString("client.endpoint"),
			AuthURL:  viper.GetString("keystone.authurl"),
			Scope: keystone.NewScope(
				viper.GetString("client.domain_id"),
				viper.GetString("client.domain_name"),
				viper.GetString("client.project_id"),
				viper.GetString("client.project_name"),
			),
			Insecure: viper.GetBool("insecure"),
		},
		viper.GetString("client.schema_root"),
	)
}

// NewCLI returns new logged in CLI Client.
func NewCLI(c *HTTPConfig, schemaRoot string) (*CLI, error) {
	client := NewHTTP(c)

	_, err := client.Login(context.Background())
	if err != nil {
		return nil, err
	}

	return &CLI{
		HTTP:       *client,
		schemaRoot: schemaRoot,
		log:        logutil.NewLogger("cli"),
	}, nil
}

// ShowResource shows resource with given schemaID and UUID.
func (c *CLI) ShowResource(schemaID, uuid string) (string, error) {
	if schemaID == "" || uuid == "" {
		return c.showHelp(schemaID, showHelpTemplate)
	}

	var response map[string]interface{}
	_, err := c.Read(context.Background(), urlPath(schemaID, uuid), &response)
	if err != nil {
		return "", err
	}

	data, ok := response[basemodels.SchemaIDToKind(schemaID)].(map[string]interface{})
	if !ok {
		return "", errors.Errorf(
			"type assertion to map[string]interface{} failed on response data: %v",
			response[basemodels.SchemaIDToKind(schemaID)],
		)
	}
	event, err := services.NewEvent(services.EventOption{
		Kind: schemaID,
		Data: data,
	})
	if err != nil {
		return "", err
	}
	eventList := &services.EventList{Events: []*services.Event{event}}
	output, err := yaml.Marshal(eventList)
	if err != nil {
		return "", err
	}
	return string(output), nil
}

const showHelpTemplate = `Show command possible usages:
{% for schema in schemas %}contrail show {{ schema.ID }} $UUID
{% endfor %}`

// ListParameters contains parameters for list command.
type ListParameters struct {
	Filters      string
	PageLimit    int64
	PageMarker   string
	Detail       bool
	Count        bool
	Shared       bool
	ExcludeHRefs bool
	ParentFQName string
	ParentType   string
	ParentUUIDs  string
	BackrefUUIDs string
	// TODO(Daniel): handle RefUUIDs
	ObjectUUIDs string
	Fields      string
}

// ListResources lists resources with given schemaID using filters.
func (c *CLI) ListResources(schemaID string, lp *ListParameters) (string, error) {
	if schemaID == "" {
		return c.showHelp("", listHelpTemplate)
	}

	var response map[string][]interface{}
	if _, err := c.ReadWithQuery(
		context.Background(),
		pluralPath(schemaID),
		queryParameters(lp),
		&response,
	); err != nil {
		return "", err
	}

	el, err := makeEventList(schemaID, response, lp.Detail)
	if err != nil {
		return "", err
	}

	output, err := yaml.Marshal(el)
	if err != nil {
		return "", err
	}
	return string(output), nil
}

const listHelpTemplate = `List command possible usages:
{% for schema in schemas %}contrail list {{ schema.ID }}
{% endfor %}`

func pluralPath(schemaID string) string {
	return "/" + basemodels.SchemaIDToKind(schemaID) + "s"
}

func queryParameters(lp *ListParameters) url.Values {
	values := url.Values{}
	for k, v := range map[string]string{
		baseservices.FiltersKey:      lp.Filters,
		baseservices.PageLimitKey:    strconv.FormatInt(lp.PageLimit, 10),
		baseservices.PageMarkerKey:   lp.PageMarker,
		baseservices.DetailKey:       strconv.FormatBool(lp.Detail),
		baseservices.CountKey:        strconv.FormatBool(lp.Count),
		baseservices.SharedKey:       strconv.FormatBool(lp.Shared),
		baseservices.ExcludeHRefsKey: strconv.FormatBool(lp.ExcludeHRefs),
		baseservices.ParentFQNameKey: lp.ParentFQName,
		baseservices.ParentTypeKey:   lp.ParentType,
		baseservices.ParentUUIDsKey:  lp.ParentUUIDs,
		baseservices.BackrefUUIDsKey: lp.BackrefUUIDs,
		// TODO(Daniel): handle RefUUIDs
		baseservices.ObjectUUIDsKey: lp.ObjectUUIDs,
		baseservices.FieldsKey:      lp.Fields,
	} {
		if !isZeroValue(v) {
			values.Set(k, v)
		}
	}
	return values
}

func isZeroValue(value interface{}) bool {
	return value == "" || value == 0 || value == false
}

func makeEventList(schemaID string, response map[string][]interface{}, detail bool) (*services.EventList, error) {
	var el *services.EventList
	var err error
	if detail {
		el, err = makeEventListFromDetailedResponse(schemaID, response)
	} else {
		el, err = makeEventListFromStandardResponse(schemaID, response)
	}

	return el, err
}

func makeEventListFromDetailedResponse(schemaID string, response map[string][]interface{}) (*services.EventList, error) {
	var el services.EventList
	for _, list := range response {
		for _, externalObjectIf := range list {
			externalObject, ok := externalObjectIf.(map[string]interface{})
			if !ok {
				return nil, errors.Errorf("detailed response contains invalid data: %v", externalObjectIf)
			}

			for _, object := range externalObject {
				e, err := makeEvent(schemaID, object)
				if err != nil {
					return nil, err
				}

				el.Events = append(el.Events, e)
			}
		}
	}
	return &el, nil
}

func makeEventListFromStandardResponse(schemaID string, response map[string][]interface{}) (*services.EventList, error) {
	var el services.EventList
	for _, list := range response {
		for _, object := range list {
			e, err := makeEvent(schemaID, object)
			if err != nil {
				return nil, err
			}

			el.Events = append(el.Events, e)
		}
	}
	return &el, nil
}

func makeEvent(schemaID string, object interface{}) (*services.Event, error) {
	m, ok := object.(map[string]interface{})
	if !ok {
		return nil, errors.Errorf("response contains a resource that is not a JSON object: %v", object)
	}

	e, err := services.NewEvent(services.EventOption{
		Kind: schemaID,
		Data: m,
	})
	if err != nil {
		return nil, errors.Errorf("failed to create event - skipping: %v", err)
	}

	return e, nil
}

// SyncResources synchronizes state of resources specified in given file.
func (c *CLI) SyncResources(filePath string) (string, error) {
	request, err := readResources(filePath)
	if err != nil {
		return "", err
	}
	response := []*services.Event{}
	_, err = c.Create(context.Background(), "/sync", request, &response)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	output, err := yaml.Marshal(&services.EventList{
		Events: response})
	if err != nil {
		return "", err
	}
	return string(output), nil
}

// SetResourceParameter sets parameter value of resource with given schemaID na UUID.
func (c *CLI) SetResourceParameter(schemaID, uuid, yamlString string) (string, error) {
	if schemaID == "" || uuid == "" {
		return c.showHelp(schemaID, setHelpTemplate)
	}

	var data map[string]interface{}
	if err := yaml.Unmarshal([]byte(yamlString), &data); err != nil {
		return "", err
	}

	data["uuid"] = uuid
	_, err := c.Update(context.Background(), urlPath(schemaID, uuid), map[string]interface{}{
		basemodels.SchemaIDToKind(schemaID): fileutil.YAMLtoJSONCompat(data),
	}, nil)
	if err != nil {
		return "", err
	}

	return c.ShowResource(schemaID, uuid)
}

const setHelpTemplate = `Set command possible usages:
{% for schema in schemas %}contrail set {{ schema.ID }} $UUID $YAML
{% endfor %}`

// DeleteResource deletes resource with given schemaID and UUID.
func (c *CLI) DeleteResource(schemaID, uuid string) (string, error) {
	if schemaID == "" || uuid == "" {
		return c.showHelp(schemaID, removeHelpTemplate)
	}

	response, err := c.EnsureDeleted(context.Background(), urlPath(schemaID, uuid), nil)
	if err != nil {
		return "", err
	}
	if response.StatusCode == http.StatusNotFound {
		c.log.WithField("path", urlPath(schemaID, uuid)).Debug("Not found")
	}

	return "", nil
}

const removeHelpTemplate = `Remove command possible usages:
{% for schema in schemas %}contrail rm {{ schema.ID }} $UUID
{% endfor %}`

// DeleteResources deletes multiple resources specified in given file.
func (c *CLI) DeleteResources(filePath string) (string, error) {
	request, err := readResources(filePath)
	if err != nil {
		return "", nil
	}

	ctx := context.Background()
	for i := len(request.Events) - 1; i >= 0; i-- {
		r := request.Events[i].GetResource()
		response, dErr := c.EnsureDeleted(
			ctx,
			urlPath(r.Kind(), r.GetUUID()),
			nil,
		)
		if dErr != nil {
			return "", dErr
		}
		if response.StatusCode == http.StatusNotFound {
			c.log.WithField("path", urlPath(r.Kind(), r.GetUUID())).Info("Not found")
		}
	}

	return "", nil
}

// readResources decodes single or array of input data from YAML.
func readResources(file string) (*services.EventList, error) {
	request := &services.EventList{}
	err := fileutil.LoadFile(file, request)
	return request, err
}

func urlPath(schemaID, uuid string) string {
	return "/" + basemodels.SchemaIDToKind(schemaID) + "/" + uuid
}

// ShowSchema returns schema with with given schemaID.
func (c *CLI) ShowSchema(schemaID string) (string, error) {
	return c.showHelp(schemaID, schemaTemplate)
}

const schemaTemplate = `
{% for schema in schemas %}
# {{ schema.Title }} {{ schema.Description }}
- kind: {{ schema.ID }}
  data: {% for key, value in schema.JSONSchema.Properties %}
    {{ key }}: {{ value.Default }} # {{ value.Title }} ({{ value.Type }}) {% endfor %}
{% endfor %}`

func (c *CLI) showHelp(schemaID string, template string) (string, error) {
	serverSchema := filepath.Join(c.schemaRoot, serverSchemaFile)
	api, err := c.fetchServerAPI(serverSchema)
	if err != nil {
		return "", err
	}
	schemas := api.Schemas
	if schemaID != "" {
		s := api.SchemaByID(schemaID)
		if s == nil {
			return "", fmt.Errorf("schema %s not found", schemaID)
		}
		schemas = []*schema.Schema{s}
	}
	tpl, err := pongo2.FromString(template)
	if err != nil {
		return "", err
	}
	o, err := tpl.Execute(pongo2.Context{"schemas": schemas})
	if err != nil {
		return "", err
	}
	return o, nil
}

func (c *CLI) fetchServerAPI(serverSchema string) (*schema.API, error) {
	var api schema.API
	ctx := context.Background()
	for i := 0; i < retryMax; i++ {
		_, err := c.Read(ctx, serverSchema, &api)
		if err == nil {
			break
		}
		logrus.WithError(err).Warn("Failed to connect API Server - reconnecting")
		time.Sleep(time.Second)
	}
	return &api, nil
}
