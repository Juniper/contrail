package replication

import (
	"context"
	"encoding/binary"
	"errors"
	"testing"
	"time"

	pkglog "github.com/Juniper/contrail/pkg/log"
	"github.com/jackc/pgx"
	"github.com/kyleconroy/pgoutput"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type oner interface {
	On(string, ...interface{}) *mock.Call
}

func TestPostgresWatcherWatch(t *testing.T) {
	slot, publication, snapshot, lsn := "test-sub", "test-pub", "snapshot-id", uint64(2778)
	var ctx context.Context
	cancel := func() {}

	tests := []struct {
		name            string
		initMock        func(oner)
		expectedMessage pgoutput.Message
		watchError      bool
	}{
		{
			name: "should return error when GetReplicationSlot fails",
			initMock: func(o oner) {
				o.On("GetReplicationSlot", mock.Anything, mock.Anything, mock.Anything).Return(uint64(0), "", assert.AnError).Once()
			},
			watchError: true,
		},
		{
			name: "should return error when RenewPublication fails",
			initMock: func(o oner) {
				o.On("GetReplicationSlot", mock.Anything, mock.Anything, mock.Anything).Return(lsn, snapshot, nil).Once()
				o.On("RenewPublication", mock.Anything, publication).Return(assert.AnError).Once()
			},
			watchError: true,
		},
		{
			name: "should return error when StartReplication fails",
			initMock: func(o oner) {
				o.On("GetReplicationSlot", mock.Anything, mock.Anything, mock.Anything).Return(lsn, snapshot, nil).Once()
				o.On("RenewPublication", mock.Anything, publication).Return(nil).Once()
				o.On("DoInTransactionSnapshot", mock.Anything, snapshot, mock.Anything).Return(nil).Once()
				o.On("StartReplication", slot, publication, uint64(0)).Return(assert.AnError).Once()
			},
			watchError: true,
		},
		{
			name: "should return error when WaitForReplicationMessage returns unknown error",
			initMock: func(o oner) {
				o.On("GetReplicationSlot", mock.Anything, mock.Anything, mock.Anything).Return(lsn, snapshot, nil).Once()
				o.On("RenewPublication", mock.Anything, publication).Return(nil).Once()
				o.On("DoInTransactionSnapshot", mock.Anything, snapshot, mock.Anything).Return(nil).Once()
				o.On("StartReplication", slot, publication, uint64(0)).Return(nil).Once()
				o.On("WaitForReplicationMessage", mock.Anything).Return(nil, assert.AnError).Once()
			},
			watchError: true,
		},
		{
			name: "should stop on WaitForReplicationMessage when context cancelled",
			initMock: func(o oner) {
				o.On("GetReplicationSlot", mock.Anything, mock.Anything, mock.Anything).Return(lsn, snapshot, nil).Once()
				o.On("RenewPublication", mock.Anything, publication).Return(nil).Once()
				o.On("DoInTransactionSnapshot", mock.Anything, snapshot, mock.Anything).Return(nil).Once()
				o.On("StartReplication", slot, publication, uint64(0)).Return(nil).Once()
				o.On("WaitForReplicationMessage", mock.Anything).Run(func(mock.Arguments) {
					cancel()
				}).Return((*pgx.ReplicationMessage)(nil), nil).Once()
				o.On("Close").Return(nil).Once()
			},
		},
		{
			name: "should continue when WaitForReplicationMessage returns context deadline",
			initMock: func(o oner) {
				o.On("GetReplicationSlot", mock.Anything, mock.Anything, mock.Anything).Return(lsn, snapshot, nil).Once()
				o.On("RenewPublication", mock.Anything, publication).Return(nil).Once()
				o.On("DoInTransactionSnapshot", mock.Anything, snapshot, mock.Anything).Return(nil).Once()
				o.On("StartReplication", slot, publication, uint64(0)).Return(nil).Once()
				o.On("WaitForReplicationMessage", mock.Anything).Return(nil, context.DeadlineExceeded).Twice()
				o.On("WaitForReplicationMessage", mock.Anything).Run(func(mock.Arguments) {
					cancel()
				}).Return((*pgx.ReplicationMessage)(nil), nil).Once()
				o.On("Close").Return(nil).Once()
			},
			watchError: false,
		},
		{
			name: "should pass to handler received WAL message",
			initMock: func(o oner) {
				o.On("GetReplicationSlot", mock.Anything, mock.Anything, mock.Anything).Return(lsn, snapshot, nil).Once()
				o.On("RenewPublication", mock.Anything, publication).Return(nil).Once()
				o.On("DoInTransactionSnapshot", mock.Anything, snapshot, mock.Anything).Return(nil).Once()
				o.On("StartReplication", slot, publication, uint64(0)).Return(nil).Once()
				o.On("WaitForReplicationMessage", mock.Anything).Return(
					&pgx.ReplicationMessage{WalMessage: &pgx.WalMessage{WalData: getBeginData(pgoutput.Begin{})}},
					nil,
				).Twice()
				o.On("WaitForReplicationMessage", mock.Anything).Run(func(mock.Arguments) {
					cancel()
				}).Return((*pgx.ReplicationMessage)(nil), nil).Once()
				o.On("Close").Return(nil).Once()
			},
			expectedMessage: pgoutput.Begin{
				Timestamp: time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC),
			},
			watchError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx, cancel = context.WithCancel(context.Background())
			var receivedMsg pgoutput.Message
			h := func(_ context.Context, m pgoutput.Message) error {
				receivedMsg = m
				return nil
			}
			m := &mockPostgresWatcherConnection{}
			if tt.initMock != nil {
				tt.initMock(m)
			}
			w := givenPostgresWatcher(slot, publication, m, h)

			err := w.Watch(ctx)

			if tt.watchError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, tt.expectedMessage, receivedMsg)
			m.AssertExpectations(t)
			cancel()
		})
	}
}

func TestPostgresWatcherContextCancellation(t *testing.T) {
	// given
	slot, publication, snapshot, lsn := "test-sub", "test-pub", "snapshot-id", uint64(2778)

	closeErr := errors.New("some closing error")

	m := &mockPostgresWatcherConnection{}
	m.On("GetReplicationSlot", mock.Anything, mock.Anything, mock.Anything).Return(lsn, snapshot, nil).Once()
	m.On("RenewPublication", mock.Anything, publication).Return(nil).Once()
	m.On("DoInTransactionSnapshot", mock.Anything, snapshot, mock.Anything).Return(nil).Once()
	m.On("StartReplication", slot, publication, uint64(0)).Return(nil).Once()
	m.On("Close").Return(closeErr).Once()

	w := givenPostgresWatcher(slot, publication, m, nil)

	ctx, cancel := context.WithCancel(context.Background())

	// when
	cancel()
	err := w.Watch(ctx)

	// then
	assert.Equal(t, closeErr, err)
	m.AssertExpectations(t)
}

func TestPostgresWatcherClose(t *testing.T) {
	canceled := false
	cancel := func() {
		canceled = true
	}
	w := givenPostgresWatcher("", "", &mockPostgresWatcherConnection{}, nil)
	w.cancel = cancel

	w.Close()

	assert.True(t, canceled)
}

func getBeginData(m pgoutput.Begin) []byte {
	b := make([]byte, 21)
	b[0] = 'B'
	binary.BigEndian.PutUint64(b[1:], m.LSN)
	binary.BigEndian.PutUint64(b[9:], 0)
	binary.BigEndian.PutUint32(b[17:], uint32(m.XID))
	return b
}

func givenPostgresWatcher(
	slot, publication string,
	conn postgresWatcherConnection,
	handler Handler,
) *PostgresWatcher {
	return &PostgresWatcher{
		conf: PostgresSubscriptionConfig{
			StatusTimeout: time.Second,
			Slot:          slot,
			Publication:   publication,
		},
		conn:    conn,
		handler: handler,
		log:     pkglog.NewLogger("postgres-watcher"),
	}
}

type mockPostgresWatcherConnection struct {
	mock.Mock
}

func (m *mockPostgresWatcherConnection) Close() error {
	args := m.MethodCalled("Close")
	return args.Error(0)
}

func (m *mockPostgresWatcherConnection) GetReplicationSlot(
	name string,
) (lastLSN uint64, snapshotName string, err error) {
	args := m.MethodCalled("GetReplicationSlot", name)
	return args.Get(0).(uint64), args.String(1), args.Error(2)
}

func (m *mockPostgresWatcherConnection) RenewPublication(ctx context.Context, name string) error {
	args := m.MethodCalled("RenewPublication", ctx, name)
	return args.Error(0)
}

func (m *mockPostgresWatcherConnection) StartReplication(slot string, publication string, startLSN uint64) error {
	args := m.MethodCalled("StartReplication", slot, publication, startLSN)
	return args.Error(0)
}

func (m *mockPostgresWatcherConnection) WaitForReplicationMessage(
	ctx context.Context,
) (*pgx.ReplicationMessage, error) {
	args := m.MethodCalled("WaitForReplicationMessage", ctx)
	if args.Error(1) != nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*pgx.ReplicationMessage), args.Error(1)
}

func (m *mockPostgresWatcherConnection) SendStatus(lastLSN uint64) error {
	args := m.MethodCalled("SendStatus", lastLSN)
	return args.Error(0)
}

func (m *mockPostgresWatcherConnection) DoInTransactionSnapshot(
	ctx context.Context, snapshotName string, do func(context.Context) error,
) error {
	args := m.MethodCalled("DoInTransactionSnapshot", ctx, snapshotName, do)
	return args.Error(0)
}

func TestWatchFailsWhenCanalStartFails(t *testing.T) {
	w := NewMySQLWatcher(&failingCanalStub{})
	err := w.Watch(context.Background())
	assert.Error(t, err)
}

func TestWatchStartsCanal(t *testing.T) {
	s := &succeedingCanalSpy{}
	w := NewMySQLWatcher(s)

	err := w.Watch(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, 1, s.startCounter)
	assert.Equal(t, 0, s.closeCounter)
}

func TestStopClosesCanal(t *testing.T) {
	s := &succeedingCanalSpy{}
	w := NewMySQLWatcher(s)

	w.Close()

	assert.Equal(t, 0, s.startCounter)
	assert.Equal(t, 1, s.closeCounter)
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
