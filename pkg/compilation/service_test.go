package compilation

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/coreos/etcd/mvcc/mvccpb"
	"github.com/gogo/protobuf/types"
	"github.com/golang/mock/gomock"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/Juniper/contrail/pkg/compilation/watch"
	"github.com/Juniper/contrail/pkg/db/etcd"
	"github.com/Juniper/contrail/pkg/testutil/integration"
	"github.com/Juniper/contrail/pkg/compilationif"
	"github.com/Juniper/contrail/pkg/services/mock"
	"github.com/Juniper/contrail/pkg/services"
	"github.com/Juniper/contrail/pkg/models"
)

const (
	jobQueueSize           = 5
	testMessageIndexString = "test-message-index-string"
	testTimeout            = 5 * time.Second
	testEtcdPath           = "contrail-test"
)

func setTestConfig() {
	viper.Set("etcd.endpoints", integration.EtcdEndpoint)
	viper.Set("etcd.path", testEtcdPath)
	viper.Set("etcd.grpc_insecure", true)
	viper.Set("compilation.msg_index_string", testMessageIndexString)
	viper.Set("compilation.number_of_workers", 0)
	viper.Set("compilation.max_job_queue_len", 0)
	viper.Set("compilation.msg_queue_lock_time", 100)
}

func TestIntentCompilationSeviceHandlesMessage(t *testing.T) {
	etcdClient := integration.NewEtcdClient(t)
	_, err := etcdClient.Delete(context.Background(), testMessageIndexString)
	assert.NoError(t, err)

	setTestConfig()

	// create intent compiler
	routineOne := newBlockingStore()
	runIntentCompiler(t, routineOne, "one")

	<-routineOne.StartedWatch

	spyOnJobChannel()

	// store some messages in etcd first
	key, value := "/"+testEtcdPath+"/test", "value"
	putResponse, err := etcdClient.Put(context.Background(), key, value)
	require.NoError(t, err)

	// receive message
	receivedIndex := routineOne.AllowGet()
	assert.Equal(t, "0", receivedIndex)

	routineOne.AllowPut()
	routineOne.WaitForTransaction()
	val, _ := etcdClient.GetString(t, testMessageIndexString)
	assert.Equal(t, fmt.Sprint(putResponse.Header.Revision), val)

	select {
	case <-time.After(testTimeout):
		t.Fatal("test timeout: no job scheduled by intent compiler")
	case j := <-watch.JobQueue:
		assert.Equal(t, putResponse.Header.Revision, j.JobID)
		assert.Empty(t, watch.JobQueue)
	}
}

func TestIntentCompilationSeviceConcurrency(t *testing.T) {
	etcdClient := integration.NewEtcdClient(t)
	_, err := etcdClient.Delete(context.Background(), testMessageIndexString)
	assert.NoError(t, err)

	setTestConfig()

	// create two intent compilers
	routineOne := newBlockingStore()
	runIntentCompiler(t, routineOne, "one")

	routineTwo := newBlockingStore()
	runIntentCompiler(t, routineTwo, "two")

	// wait for them to initialize
	<-routineOne.StartedWatch
	<-routineTwo.StartedWatch

	spyOnJobChannel()

	// store some messages in etcd first
	resourceKey := "/" + testEtcdPath + "/test"
	firstResp, err := etcdClient.Put(context.Background(), resourceKey, "value")
	require.NoError(t, err)

	secondResp, err := etcdClient.Put(context.Background(), resourceKey, "another value")
	require.NoError(t, err)

	_, err = etcdClient.Put(context.Background(), resourceKey,
		"final irrelevant message that should be read but not handled")
	require.NoError(t, err)

	firstRevision, secondRevision := fmt.Sprint(firstResp.Header.Revision), fmt.Sprint(secondResp.Header.Revision)

	// first message
	// allow read for routineOne
	receivedIndex := routineOne.AllowGet()
	assert.Equal(t, "0", receivedIndex)

	// wild routineTwo appeared!
	receivedIndex = routineTwo.AllowGet()
	assert.Equal(t, "0", receivedIndex)

	// routineTwo updates index
	routineTwo.AllowPut()
	routineTwo.WaitForTransaction()
	val, rev := etcdClient.GetString(t, testMessageIndexString)
	assert.Equal(t, firstRevision, val)

	// routineOne tries to update index with old value, but fails
	routineOne.AllowPut()
	etcdClient.ExpectValue(t, testMessageIndexString, val, rev)

	// routineOne retries getting index and then ignores the message
	receivedIndex = routineOne.AllowGet()
	assert.Equal(t, firstRevision, receivedIndex)

	// second message
	// routineTwo handles next message normally
	receivedIndex = routineTwo.AllowGet()
	assert.Equal(t, firstRevision, receivedIndex)

	// routineOne tries to get second message too
	receivedIndex = routineOne.AllowGet()
	assert.Equal(t, firstRevision, receivedIndex)

	// routines try to put indexes again, but only Two succeeds
	routineTwo.AllowPut()
	routineTwo.WaitForTransaction()
	val, rev = etcdClient.GetString(t, testMessageIndexString)
	assert.Equal(t, secondRevision, val)
	routineOne.AllowPut()
	etcdClient.ExpectValue(t, testMessageIndexString, val, rev)

	// routineOne retries getting index and then ignores the message
	receivedIndex = routineOne.AllowGet()
	assert.Equal(t, secondRevision, receivedIndex)

	// last message
	// routines attempt to read index
	receivedIndex = routineOne.AllowGet()
	assert.Equal(t, secondRevision, receivedIndex)
	receivedIndex = routineTwo.AllowGet()
	assert.Equal(t, secondRevision, receivedIndex)
	// now both routines are stuck during handling of last message

	jobs := collectJobs()
	assert.Len(t, jobs, 2, "job queue should contain two jobs")
	if len(jobs) > 2 {
		assert.Equal(t, firstResp.Header.Revision, jobs[0].JobID)
		assert.Equal(t, secondResp.Header.Revision, jobs[1].JobID)
	}
	t.Log(jobs)
}

func runIntentCompiler(t *testing.T, b *blockingStore, name string) {
	ics, err := NewIntentCompilationService()
	require.NoError(t, err)

	b.RegisterIn(ics)
	ics.log = ics.log.WithField("routine", name)
	go func() {
		err = ics.Run(context.Background())
		assert.NoError(t, err)
	}()
}

func TestIntentCompilerHandlesDependentResources(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	compiler := compilationif.NewCompilationService()
	mockService := servicesmock.NewMockService(mockCtrl)
	services.Chain(compiler, mockService)

	network := &models.VirtualNetwork{
		UUID: "Virtual-Network-1",
	}
	policy := &models.NetworkPolicy{
		UUID: "Network-policy-1",
	}

	ref := &models.VirtualNetworkNetworkPolicyRef{
		UUID: policy.UUID,
		To: []string{"default-domain", "default-project", network.UUID},
	}
	network.NetworkPolicyRefs = append(network.NetworkPolicyRefs, ref)
	policy.VirtualNetworkBackRefs = append(policy.VirtualNetworkBackRefs, network)

	mockService.EXPECT().CreateVirtualNetwork(gomock.Not(gomock.Nil()),
		&services.CreateVirtualNetworkRequest{network},
	).Return(&services.CreateVirtualNetworkResponse{network}, nil)

	mockService.EXPECT().UpdateNetworkPolicy(gomock.Not(gomock.Nil()),
		&services.UpdateNetworkPolicyRequest{
			NetworkPolicy: policy,
			FieldMask: types.FieldMask{Paths: []string{}},
	}).Return(&services.UpdateNetworkPolicyResponse{policy}, nil)

	networkJSON, err := json.Marshal(network)
	require.NoError(t, err, "Marshaling the network should succeed")

	err = compiler.HandleResource(context.Background(), int32(mvccpb.PUT), "virtual_network",
		network.UUID, string(networkJSON))
	assert.NoError(t, err)
}

func TestIntentCompilerRunsNextService(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	compiler := compilationif.NewCompilationService()
	mockService := servicesmock.NewMockService(mockCtrl)
	services.Chain(compiler, mockService)

	mockService.EXPECT().CreateVirtualNetwork(gomock.Not(gomock.Nil()), &services.CreateVirtualNetworkRequest{
		VirtualNetwork: &models.VirtualNetwork{},
	}).Return(&services.CreateVirtualNetworkResponse{}, nil)

	err := compiler.HandleResource(context.Background(), int32(mvccpb.PUT), "virtual_network",
		"test_uuid", "{}")
	assert.NoError(t, err)
}

func TestIntentCompilerFailsForBadJSON(t *testing.T) {
	compiler := compilationif.NewCompilationService()
	err := compiler.HandleResource(context.Background(), int32(mvccpb.PUT), "virtual_network", "test_uuid", "")
	assert.Error(t, err)
}

func TestIntentCompilerFailsForUnknownResource(t *testing.T) {
	compiler := compilationif.NewCompilationService()
	err := compiler.HandleResource(context.Background(), int32(mvccpb.PUT), "anything", "", "")
	assert.Error(t, err)
}

func spyOnJobChannel() {
	close(watch.JobQueue)           // close the channel to cancel dispatcher watch
	watch.WatcherInit(jobQueueSize) // create new channel to spy on
}

func collectJobs() (result []watch.JobRequest) {
	for {
		select {
		case j := <-watch.JobQueue:
			result = append(result, j)
		default:
			return result
		}
	}
}

type blockingStore struct {
	store

	StartedWatch chan struct{}

	StartGet    chan struct{}
	FinishedGet chan string

	StartPut    chan struct{}
	FinishedPut chan string

	TransactionInProgress chan struct{}
}

func newBlockingStore() *blockingStore {
	b := &blockingStore{
		StartedWatch:          make(chan struct{}),
		StartPut:              make(chan struct{}, 1),
		FinishedPut:           make(chan string, 1),
		StartGet:              make(chan struct{}, 1),
		FinishedGet:           make(chan string, 1),
		TransactionInProgress: make(chan struct{}),
	}
	close(b.TransactionInProgress)
	return b
}

func (s *blockingStore) AllowGet() (receivedIndex string) {
	s.StartGet <- struct{}{}
	return <-s.FinishedGet
}

func (s *blockingStore) AllowPut() {
	s.StartPut <- struct{}{}
	<-s.FinishedPut
}

func (s *blockingStore) WaitForTransaction() {
	<-s.TransactionInProgress
}

func (s *blockingStore) RegisterIn(ics *IntentCompilationService) {
	s.store = ics.Store
	ics.Store = s
}

func (s *blockingStore) WatchRecursive(
	ctx context.Context, keyPattern string, afterIndex int64,
) chan etcd.Event {
	c := s.store.WatchRecursive(ctx, keyPattern, afterIndex)
	close(s.StartedWatch)
	return c
}

func (s *blockingStore) InTransaction(ctx context.Context, do func(ctx context.Context) error) error {
	s.TransactionInProgress = make(chan struct{})
	wrappedDo := func(ctx context.Context) error {
		txn := blockingTxn{
			Txn:         etcd.GetTxn(ctx),
			StartGet:    s.StartGet,
			FinishedGet: s.FinishedGet,
			StartPut:    s.StartPut,
			FinishedPut: s.FinishedPut,
		}
		ctx = etcd.WithTxn(ctx, txn)
		return do(ctx)
	}

	err := s.store.InTransaction(ctx, wrappedDo)
	close(s.TransactionInProgress)
	return err
}

type blockingTxn struct {
	etcd.Txn

	StartGet    chan struct{}
	FinishedGet chan string

	StartPut    chan struct{}
	FinishedPut chan string
}

func (b blockingTxn) Get(key string) []byte {
	<-b.StartGet
	val := b.Txn.Get(key)
	b.FinishedGet <- string(val)
	return val
}

func (b blockingTxn) Put(key string, val []byte) {
	<-b.StartPut
	b.Txn.Put(key, val)
	b.FinishedPut <- string(val)
}
