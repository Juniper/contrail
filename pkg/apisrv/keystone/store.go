package keystone

import (
	"fmt"
	"sync"
	"time"

	kscommon "github.com/Juniper/contrail/pkg/keystone"
	uuid "github.com/satori/go.uuid"
)

//Store is used to provide a persistence layer for tokens.
type Store interface {
	CreateToken(*kscommon.User, *kscommon.Project) (string, *kscommon.Token)
	ValidateToken(string) (*kscommon.Token, bool)
	RetrieveToken(string) (*kscommon.Token, error)
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
func (store *InMemoryStore) CreateToken(user *kscommon.User, project *kscommon.Project) (string, *kscommon.Token) {
	tokenID := uuid.NewV4().String()
	token := &kscommon.Token{
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
func (store *InMemoryStore) ValidateToken(tokenID string) (*kscommon.Token, bool) {
	i, ok := store.store.Load(tokenID)
	if !ok {
		return nil, false
	}
	token, ok := i.(*kscommon.Token)
	return token, ok
}

//RetrieveToken is used to retrieve a token, and return a token body.
func (store *InMemoryStore) RetrieveToken(tokenID string) (*kscommon.Token, error) {
	i, ok := store.store.Load(tokenID)
	if !ok {
		return nil, fmt.Errorf("token not found")
	}
	token, ok := i.(*kscommon.Token)
	var err error
	if !ok {
		err = fmt.Errorf("can't convert stored token (%v) back to it's type", i)
	}
	return token, err
}
