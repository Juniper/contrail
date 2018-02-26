package keystone

import (
	"sync"
	"time"

	uuid "github.com/satori/go.uuid"
)

//Store is used to provide a persistence layer for tokens.
type Store interface {
	CreateToken(*User, *Project) (string, *Token)
	ValidateToken(string) (*Token, bool)
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
func (store *InMemoryStore) CreateToken(user *User, project *Project) (string, *Token) {
	tokenID := uuid.NewV4().String()
	token := &Token{
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
func (store *InMemoryStore) ValidateToken(tokenID string) (*Token, bool) {
	i, ok := store.store.Load(tokenID)
	if !ok {
		return nil, false
	}
	token, ok := i.(*Token)
	return token, ok
}
