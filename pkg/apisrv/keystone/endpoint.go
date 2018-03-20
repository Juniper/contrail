package keystone

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"sync"
)

// EndpointStore is used to store cluster endpoints in-memory
type EndpointStore struct {
	store *sync.Map
}

type keystoneEndpoint struct {
	privateURL string
	publicURL  string
}

type clusterEndpoint struct {
	dbConn        *sql.DB
	endpointStore *EndpointStore
	clusterID     string
	token         string
	ctx           context.Context
}

//MakeEndpointStore is used to make a in memory endpoint store.
func MakeEndpointStore() *EndpointStore {
	return &EndpointStore{
		store: new(sync.Map),
	}
}

func (ce *clusterEndpoint) readEndpoint() (*string, *string, error) {
	q := []string{"SELECT public_url, private_url",
		"FROM endpoint WHERE name=keystone"}
	var err error
	var rows *sql.Rows
	if ce.clusterID == "any" {
		rows, err = ce.dbConn.Query(strings.Join(q, " "))
	} else {
		q = append(q, "AND parent_uuid=?")
		rows, err = ce.dbConn.Query(strings.Join(q, " "), "keystone", ce.clusterID)
	}
	if err != nil {
		return nil, nil, err
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, nil, err
	}
	rowCount := 0
	var publicURL string
	var privateURL string
	for rows.Next() {
		rowCount++
		if err := rows.Scan(&publicURL, &privateURL); err != nil {
			return nil, nil, err
		}
	}
	if rowCount > 1 {
		return nil, nil,
			fmt.Errorf("Multimatch of keystone endpoint for cluster: %s",
				ce.clusterID)
	}
	return &publicURL, &privateURL, nil
}

func (ce *clusterEndpoint) getEndpoint() (*keystoneEndpoint, error) {
	e, ok := ce.endpointStore.store.Load(ce.clusterID)
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
	ce.endpointStore.store.Store(ce.clusterID, ep)
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
