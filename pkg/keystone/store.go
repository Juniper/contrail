package keystone

import (
	"fmt"
	"sync"
	"time"

	"github.com/Juniper/asf/pkg/keystone"

	uuid "github.com/satori/go.uuid"
)

//Store is used to provide a persistence layer for tokens.
type Store interface {
	CreateToken(*keystone.User, *keystone.Project) (string, *keystone.Token)
	ValidateToken(string) (*keystone.Token, bool)
	RetrieveToken(string) (*keystone.Token, error)
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
func (store *InMemoryStore) CreateToken(user *keystone.User, project *keystone.Project) (string, *keystone.Token) {
	tokenID := uuid.NewV4().String()
	token := &keystone.Token{
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
func (store *InMemoryStore) ValidateToken(tokenID string) (*keystone.Token, bool) {
	i, ok := store.store.Load(tokenID)
	if !ok {
		return nil, false
	}
	token, ok := i.(*keystone.Token)
	return token, ok
}

//RetrieveToken is used to retrieve a token, and return a token body.
func (store *InMemoryStore) RetrieveToken(tokenID string) (*keystone.Token, error) {
	i, ok := store.store.Load(tokenID)
	if !ok {
		return nil, fmt.Errorf("token not found")
	}
	token, ok := i.(*keystone.Token)
	var err error
	if !ok {
		err = fmt.Errorf("can't convert stored token (%v) back to it's type", i)
	}
	return token, err
}
