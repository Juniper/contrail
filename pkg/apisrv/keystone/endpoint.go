package keystone

import (
	"context"
	"fmt"

	"github.com/Juniper/contrail/pkg/common"
	"github.com/Juniper/contrail/pkg/db"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/serviceif"

	log "github.com/sirupsen/logrus"
)

const (
	nameField            = "name"
	keystoneEndpointName = "keystone"
	publicURLField       = "public_url"
	privateURLField      = "private_url"
)

type keystoneEndpoint struct {
	privateURL string
	publicURL  string
}

type clusterEndpoint struct {
	dbService     serviceif.Service
	endpointStore *common.EndpointStore
	clusterID     string
	token         string
	context       context.Context
}

func (c *clusterEndpoint) isValidEndpointFields(fields []string) bool {
	validFields := make(map[string]bool)
	for _, field := range fields {
		for _, validField := range db.EndpointFields {
			if field == validField {
				validFields[field] = true
				break
			}
		}
		if _, ok := validFields[field]; !ok {
			log.Fatalf("%s is not a valid endpoint field", field)
		}
	}
	if len(fields) != len(validFields) {
		return false
	}
	return true
}

func (c *clusterEndpoint) readEndpointURL() (string, string, error) {
	ok := c.isValidEndpointFields([]string{nameField, publicURLField, privateURLField})
	if !ok {
		return "", "",
			fmt.Errorf("Invalid endpoint filends used in query")

	}
	ctx := common.NoAuth(context.Background())
	values := []string{keystoneEndpointName}
	filters := []*models.Filter{{
		Key:    nameField,
		Values: values,
	}}
	var spec models.ListSpec
	fields := []string{publicURLField, privateURLField}
	if c.clusterID == "" {
		spec = models.ListSpec{
			Limit:   100,
			Filters: filters,
			Fields:  fields,
		}
	} else {
		parents := []string{c.clusterID}
		spec = models.ListSpec{
			Limit:       100,
			Filters:     filters,
			ParentUUIDs: parents,
			Fields:      fields,
		}
	}
	request := &models.ListEndpointRequest{Spec: &spec}
	log.Debugf("Listing keystone endpoint of cluster:%s", c.clusterID)
	response, err := c.dbService.ListEndpoint(ctx, request)
	if err != nil {
		return "", "", err
	}
	endpoints := response.Endpoints
	if len(endpoints) == 1 {
		c.clusterID = endpoints[0].ParentUUID
		return endpoints[0].PublicURL, endpoints[0].PrivateURL, nil
	} else if len(endpoints) > 1 {
		return "", "",
			fmt.Errorf("Multimatch of keystone endpoint for cluster:%s",
				c.clusterID)
	}
	return "", "",
		fmt.Errorf("Keystone endpoint not found for cluster:%s", c.clusterID)
}

func (c *clusterEndpoint) readEndpoint() (*keystoneEndpoint, error) {
	// Read url from db
	log.Debugf("Reading cluster:%s keystone endpoint", c.clusterID)
	public, private, err := c.readEndpointURL()
	if err != nil {
		return nil, err
	}
	// clusterID should be updated during sync
	// check if a cluster is created
	if c.clusterID == "" {
		// No cluster found
		return nil, common.ErrorNotFound
	}
	e := &keystoneEndpoint{}
	if public != "" {
		e.publicURL = public
	}
	if private != "" {
		e.privateURL = private
	}
	return e, nil
}

func (c *clusterEndpoint) getEndpoint() (*keystoneEndpoint, error) {
	endpoint, ok := c.endpointStore.Data.Load(c.clusterID)
	if !ok {
		return nil, fmt.Errorf("keystone endpoint for cluster:%s not found", c.clusterID)
	}
	e, ok := endpoint.(*keystoneEndpoint)
	return e, nil
}

func (c *clusterEndpoint) syncEndpoint() (*keystoneEndpoint, error) {
	// Read from db
	e, err := c.readEndpoint()
	if err != nil {
		return nil, err
	}
	// sync the endpoint to store
	log.Debugf("syncing cluster:%s keystone endpoint in-memory", c.clusterID)
	c.endpointStore.Data.Store(c.clusterID, e)
	return e, nil
}

func (c *clusterEndpoint) authenticate() (context.Context, error) {
	var ok error
	var newCtx context.Context
	e, err := c.getEndpoint()
	if err != nil {
		// endpoint not in memory, sync and authenticate
		log.Debugf("cluster:%s keystone endpoint not in-memory", c.clusterID)
		e, err = c.syncEndpoint()
		if err != nil {
			return nil, err
		}
		auth := newAuth(e.privateURL, true) // TODO:(ijohnson) add insecure in schema
		log.Debugf("authenticate using cluster:%s keystone endpoint", c.clusterID)
		newCtx, ok = authenticate(c.context, auth, c.token)
	} else {
		// endpoint found in memory, try authenticate
		log.Debugf("cluster:%s keystone endpoint found in-memory", c.clusterID)
		auth := newAuth(e.privateURL, true) // TODO:(ijohnson) add insecure in schema
		log.Debugf("authenticate using cluster:%s keystone endpoint", c.clusterID)
		newCtx, err = authenticate(c.context, auth, c.token)
		if err != nil {
			// sync endpoint and re try authenticate
			log.Debugf("auth failed, re-try after endpoint sync")
			e, err = c.syncEndpoint()
			if err != nil {
				return nil, err
			}
			log.Debugf("authenticate using cluster:%s keystone endpoint", c.clusterID)
			newCtx, ok = authenticate(c.context, auth, c.token)
		}
	}
	return newCtx, ok
}
