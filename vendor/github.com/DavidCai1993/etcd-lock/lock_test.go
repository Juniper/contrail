package etcdlock

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
	"google.golang.org/grpc"
)

type LockSuite struct {
	suite.Suite
	ctx    context.Context
	locker *Locker
}

func (s *LockSuite) SetupSuite() {
	s.ctx = context.Background()
	locker, err := NewLocker(LockerOptions{
		Address:     "127.0.0.1:2379",
		DialOptions: []grpc.DialOption{grpc.WithInsecure()},
	})
	s.Nil(err)
	s.locker = locker
}

func (s *LockSuite) TestUnlock() {
	start := time.Now()

	lock, err := s.locker.Lock(s.ctx, "test_unlock", 5*time.Second)
	s.Nil(err)

	err = lock.Unlock(s.ctx)
	s.Nil(err)

	_, err = s.locker.Lock(s.ctx, "test_unlock", 5*time.Second)
	s.Nil(err)

	s.True(time.Now().Sub(start) < time.Second)
}

func TestLock(t *testing.T) {
	suite.Run(t, new(LockSuite))
}
