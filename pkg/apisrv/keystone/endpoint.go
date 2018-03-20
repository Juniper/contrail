package keystone

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/spf13/viper"

	"github.com/Juniper/contrail/pkg/common"
	"github.com/Juniper/contrail/pkg/db"
	"github.com/Juniper/contrail/pkg/models"
)

type keystoneEndpoint struct {
	privateURL string
	publicURL  string
}

type clusterEndpoint struct {
	dbConn        *sql.DB
	endpointStore *common.EndpointStore
	clusterID     string
	token         string
	ctx           context.Context
}

func (ce *clusterEndpoint) readEndpoint() (*string, *string, error) {
	dbService := db.NewService(ce.dbConn, viper.GetString("database.dialect"))
	ctx := common.NoAuth(context.Background())
	values := []string{"keystone"}
	filters := []*models.Filter{&models.Filter{
		Key:    "name",
		Values: values,
	}}
	var spec models.ListSpec
	if ce.clusterID == "any" {
		spec = models.ListSpec{
			Filters: filters,
		}
	} else {
		parents := []string{ce.clusterID}
		fields := []string{"public_url", "private_url"}
		spec = models.ListSpec{
			Filters:     filters,
			ParentUUIDs: parents,
			Fields:      fields,
		}
	}
	request := &models.ListEndpointRequest{Spec: &spec}
	response, err := dbService.ListEndpoint(ctx, request)
	if err != nil {
		return nil, nil, err
	}
	endpoints := response.Endpoints
	if len(endpoints) == 1 {
		return &endpoints[0].PublicURL, &endpoints[0].PrivateURL, nil
	} else if len(endpoints) > 1 {
		return nil, nil,
			fmt.Errorf("Multimatch of keystone endpoint for cluster: %s",
				ce.clusterID)
	}
	return nil, nil, nil
}

func (ce *clusterEndpoint) getEndpoint() (*keystoneEndpoint, error) {
	e, ok := ce.endpointStore.Data.Load(ce.clusterID)
	if !ok {
		return nil, fmt.Errorf("keystone endpoint for cluster: %s not found", ce.clusterID)
	}
	ep, ok := e.(*keystoneEndpoint)
	return ep, nil
}

func (ce *clusterEndpoint) syncEndpoint() (*keystoneEndpoint, error) {
	// Read from db
	public, private, err := ce.readEndpoint()
	if err != nil {
		return nil, err
	}
	ep := &keystoneEndpoint{
		privateURL: *private,
		publicURL:  *public,
	}
	// sync the endpoint to store
	ce.endpointStore.Data.Store(ce.clusterID, ep)
	return ep, nil
}

func (ce *clusterEndpoint) authenticate() (context.Context, error) {
	var ok error
	var newCtx context.Context
	ep, err := ce.getEndpoint()
	if err != nil {
		// endpoint not in memory, sync and authenticate
		ep, err = ce.syncEndpoint()
		if err != nil {
			return nil, err
		}
		auth := newAuth(ep.privateURL, true) // TODO:(ijohnson) add insecure in schema
		newCtx, ok = authenticate(ce.ctx, auth, ce.token)
	} else {
		// endpoint found in memory, try authenticate
		auth := newAuth(ep.privateURL, true) // TODO:(ijohnson) add insecure in schema
		newCtx, err = authenticate(ce.ctx, auth, ce.token)
		if err != nil {
			// sync endpoint and re try authenticate
			ep, err = ce.syncEndpoint()
			if err != nil {
				return nil, err
			}
			newCtx, ok = authenticate(ce.ctx, auth, ce.token)
		}
	}
	return newCtx, ok
}
