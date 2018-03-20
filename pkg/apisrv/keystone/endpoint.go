package keystone

import (
	"context"
	"fmt"

	"github.com/Juniper/contrail/pkg/common"
	"github.com/Juniper/contrail/pkg/db"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/serviceif"

	apicommon "github.com/Juniper/contrail/pkg/apisrv/common"
	log "github.com/sirupsen/logrus"
)

const (
	nameField            = "name"
	keystoneEndpointName = "keystone"
	publicURLField       = "public_url"
	privateURLField      = "private_url"
	scope                = "private"
	limit                = 100
)

type clusterEndpoint struct {
	dbService     serviceif.Service
	endpointStore *apicommon.EndpointStore
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
			log.Errorf("%s is not a valid endpoint field", field)
		}
	}
	if len(fields) != len(validFields) {
		return false
	}
	return true
}

func (c *clusterEndpoint) readKeystoneEndpoints() []*models.Endpoint {
	ok := c.isValidEndpointFields([]string{nameField})
	if !ok {
		log.Fatal("Invalid endpoint fields used in query")

	}
	ctx := common.NoAuth(context.Background())
	values := []string{keystoneEndpointName}
	filters := []*models.Filter{{
		Key:    nameField,
		Values: values,
	}}
	var spec models.ListSpec
	if c.clusterID == "" {
		spec = models.ListSpec{
			Limit:   limit,
			Filters: filters,
		}
	} else {
		parents := []string{c.clusterID}
		spec = models.ListSpec{
			Limit:       limit,
			Filters:     filters,
			ParentUUIDs: parents,
		}
	}
	request := &models.ListEndpointRequest{Spec: &spec}
	log.Debugf("Listing keystone endpoint of cluster:%s", c.clusterID)
	response, err := c.dbService.ListEndpoint(ctx, request)
	if err != nil {
		log.Errorf("DB read failed: %s", err)
		return nil
	}
	return response.Endpoints
}

func (c *clusterEndpoint) getEndpoint() (authURL string) {
	targetStore := c.endpointStore.Read(c.clusterID)
	if targetStore == nil {
		log.Debugf("Endpoint target store not found for cluster: %s", c.clusterID)
		return ""
	}
	e := targetStore.Next(scope)
	if e == "" {
		log.Debugf("Endpoint target empty for cluster: %s", c.clusterID)
		return ""
	}
	return e
}

func (c *clusterEndpoint) syncEndpoint() (authURL string) {
	// Read from db
	endpoints := c.readKeystoneEndpoints()
	if endpoints == nil {
		return ""
	}
	// sync the endpoint to store
	log.Debugf("syncing cluster:%s keystone endpoint in-memory", c.clusterID)
	targetStore := c.endpointStore.Read(c.clusterID)
	if targetStore == nil {
		targetStore = apicommon.MakeTargetStore()
	}
	for _, e := range endpoints {
		targetStore.Write(e.UUID, e)
	}
	c.endpointStore.Write(c.clusterID, targetStore)
	return targetStore.Next(scope)
}

func (c *clusterEndpoint) authenticate() (context.Context, error) {
	var err error
	var newCtx context.Context
	var authURL string
	authURL = c.getEndpoint()
	if authURL == "" {
		// endpoint not in memory, sync and authenticate
		log.Debugf("cluster:%s keystone endpoint not in-memory", c.clusterID)
		authURL = c.syncEndpoint()
		if authURL == "" {
			return nil, fmt.Errorf("Endpoint not created for cluster: %s", c.clusterID)
		}
		auth := newAuth(authURL, true) // TODO:(ijohnson) add insecure in schema
		log.Debugf("authenticate using cluster:%s keystone endpoint", c.clusterID)
		newCtx, err = authenticate(c.context, auth, c.token)
	} else {
		// endpoint found in memory, try authenticate
		log.Debugf("cluster:%s keystone endpoint found in-memory", c.clusterID)
		auth := newAuth(authURL, true) // TODO:(ijohnson) add insecure in schema
		log.Debugf("authenticate using cluster:%s keystone endpoint", c.clusterID)
		newCtx, err = authenticate(c.context, auth, c.token)
		if err != nil {
			// sync endpoint and re try authenticate
			log.Debugf("auth failed, re-try after endpoint sync")
			authURL = c.syncEndpoint()
			if authURL == "" {
				return nil, fmt.Errorf("Endpoint deleted for cluster: %s", c.clusterID)
			}
			log.Debugf("authenticate using cluster:%s keystone endpoint", c.clusterID)
			newCtx, err = authenticate(c.context, auth, c.token)
		}
	}
	return newCtx, err
}
