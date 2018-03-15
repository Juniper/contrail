package etcdlock

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
	"google.golang.org/grpc"
)

type LockerSuite struct {
	suite.Suite
	ctx    context.Context
	locker *Locker
}

func (s *LockerSuite) SetupSuite() {
	s.ctx = context.Background()
	locker, err := NewLocker(LockerOptions{
		Address:     "127.0.0.1:2379",
		DialOptions: []grpc.DialOption{grpc.WithInsecure()},
	})
	s.Nil(err)
	s.locker = locker
}

func (s *LockerSuite) TestNewLocker() {
	locker, err := NewLocker(LockerOptions{
		Address:     "127.0.0.1:2379",
		DialOptions: []grpc.DialOption{grpc.WithInsecure()},
	})

	s.Nil(err)
	s.Equal(locker.etcdKeyPrefix, defaultEtcdKeyPrefix)
}

func (s *LockerSuite) TestLockNewKey() {
	lock, err := s.locker.Lock(s.ctx, "test_lock_new_key", 3*time.Second)

	s.Nil(err)
	s.Contains(string(lock.keyName), "test_lock_new_key")
}

func (s *LockerSuite) TestLockEmptyKey() {
	_, err := s.locker.Lock(s.ctx, "", 3*time.Second)

	s.Equal(err.Error(), ErrEmptyKey.Error())
}

func (s *LockerSuite) TestLockLockedKey() {
	start := time.Now()

	_, err := s.locker.Lock(s.ctx, "test_lock_locked_key", 3*time.Second)
	s.Nil(err)

	_, err = s.locker.Lock(s.ctx, "test_lock_locked_key", 3*time.Second)
	s.Nil(err)

	s.True(time.Now().Sub(start) > 3*time.Second)
	s.True(time.Now().Sub(start) < 4*time.Second)
}

func (s *LockerSuite) TestIsLockedEmptyKey() {
	_, err := s.locker.IsLocked(s.ctx, "")

	s.Equal(err.Error(), ErrEmptyKey.Error())
}

func (s *LockerSuite) TestIsLocked() {
	isLocked, err := s.locker.IsLocked(s.ctx, "test_is_locked_1")

	s.Nil(err)
	s.False(isLocked)

	_, err = s.locker.Lock(s.ctx, "test_is_locked_1", 3*time.Second)
	s.Nil(err)

	_, err = s.locker.Lock(s.ctx, "test_is_locked_2", 10*time.Second)
	s.Nil(err)

	isLocked, err = s.locker.IsLocked(s.ctx, "test_is_locked_1")

	s.Nil(err)
	s.True(isLocked)

	isLocked, err = s.locker.IsLocked(s.ctx, "test_is_locked_1")

	s.Nil(err)
	s.True(isLocked)

	time.Sleep(4 * time.Second)

	isLocked, err = s.locker.IsLocked(s.ctx, "test_is_locked_1")

	s.Nil(err)
	s.False(isLocked)

	isLocked, err = s.locker.IsLocked(s.ctx, "test_is_locked_2")

	s.Nil(err)
	s.True(isLocked)
}

func TestLocker(t *testing.T) {
	suite.Run(t, new(LockerSuite))
}
