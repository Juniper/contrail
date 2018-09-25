package keystone

import (
	"fmt"
	"sync"
	"time"

	types "github.com/Juniper/contrail/pkg/common/keystone"
	uuid "github.com/satori/go.uuid"
)

//Store is used to provide a persistence layer for tokens.
type Store interface {
	CreateToken(*types.User, *types.Project) (string, *types.Token)
	ValidateToken(string) (*types.Token, bool)
	RetrieveToken(string) (*types.Token, error)
}

//InMemoryStore is an implementation of Store based on in-memory synced map.
type InMemoryStore struct {
	store  *sync.Map
	expire time.Duration
}

//MakeInMemoryStore is used to make a in memory store.
func MakeInMemoryStore(expire time.Duration) *InMemoryStore {
	return &InMemoryStore{
		store:  new(sync.Map),
		expire: expire,
	}
}

//CreateToken is used to create a token for a user.
//This method also persists a token.
func (store *InMemoryStore) CreateToken(user *types.User, project *types.Project) (string, *types.Token) {
	tokenID := uuid.NewV4().String()
	token := &types.Token{
		IssuedAt:  time.Now(),
		ExpiresAt: time.Now().Add(store.expire),
		User:      user,
		Domain:    user.Domain,
		Project:   project,
		Roles:     user.Roles,
	}
	store.store.Store(tokenID, token)
	return tokenID, token
}

//ValidateToken is used to validate a token, and return a token body.
func (store *InMemoryStore) ValidateToken(tokenID string) (*types.Token, bool) {
	i, ok := store.store.Load(tokenID)
	if !ok {
		return nil, false
	}
	token, ok := i.(*types.Token)
	return token, ok
}

//RetrieveToken is used to retrive a token, and return a token body.
func (store *InMemoryStore) RetrieveToken(tokenID string) (*types.Token, error) {
	i, ok := store.store.Load(tokenID)
	if !ok {
		return nil, fmt.Errorf("token not found")
	}
	token, ok := i.(*types.Token)
	var err error
	if !ok {
		err = fmt.Errorf("can't convert stored token (%v) back to it's type", i)
	}
	return token, err
}
