package replication

import (
	"context"
	"testing"

	"github.com/jackc/pgx"
	"github.com/kyleconroy/pgoutput"
	"github.com/stretchr/testify/assert"
)

func TestSubscriptionWatcherWatch(t *testing.T) {
	started := false
	s := mockStarter(func(ctx context.Context, c *pgx.ReplicationConn, h pgoutput.Handler) error {
		started = true
		return nil
	})

	w := NewSubscriptionWatcher(nil, s, NewPgoutputEventHandler(nil))

	err := w.Watch(context.Background())

	assert.NoError(t, err)

	assert.True(t, started)
	assert.NotNil(t, w.cancel)
	w.cancel()
}

func TestSubscriptionWatcherClose(t *testing.T) {
	canceled := false
	cancel := func() {
		canceled = true
	}
	w := NewSubscriptionWatcher(nil, nil, nil)
	w.cancel = cancel

	w.Close()

	assert.True(t, canceled)
}

type mockStarter func(context.Context, *pgx.ReplicationConn, pgoutput.Handler) error

func (s mockStarter) Start(ctx context.Context, c *pgx.ReplicationConn, h pgoutput.Handler) error {
	return s(ctx, c, h)
}

func TestBinlogWatcherIsNoopByDefault(t *testing.T) {
	w := givenBinlogWatcher(nil)

	err := w.Watch(context.Background())
	assert.NoError(t, err)

	w.Close()
}

func TestWatchFailsWhenCanalStartFails(t *testing.T) {
	w := givenBinlogWatcher(&failingCanalStub{})
	err := w.Watch(context.Background())
	assert.Error(t, err)
}

func TestWatchStartsCanal(t *testing.T) {
	s := &succeedingCanalSpy{}
	w := givenBinlogWatcher(s)

	err := w.Watch(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, 1, s.startCounter)
	assert.Equal(t, 0, s.closeCounter)
}

func TestStopClosesCanal(t *testing.T) {
	s := &succeedingCanalSpy{}
	w := givenBinlogWatcher(s)

	w.Close()

	assert.Equal(t, 0, s.startCounter)
	assert.Equal(t, 1, s.closeCounter)
}

func givenBinlogWatcher(c abstractCanal) *BinlogWatcher {
	return NewBinlogWatcher(c)
}

type failingCanalStub struct{}

func (c *failingCanalStub) Run() error { return assert.AnError }

func (c *failingCanalStub) Close() {}

type succeedingCanalSpy struct {
	startCounter int
	closeCounter int
}

func (c *succeedingCanalSpy) Run() error {
	c.startCounter++
	return nil
}

func (c *succeedingCanalSpy) Close() {
	c.closeCounter++
}
