package integration

// TODO(Michał): Split this file and refactor some functions.

import (
	"context"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/coreos/etcd/clientv3"
	"github.com/flosch/pongo2"
	"github.com/labstack/echo"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cast"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	yaml "gopkg.in/yaml.v2"

	"github.com/Juniper/contrail/pkg/apisrv"
	"github.com/Juniper/contrail/pkg/apisrv/client"
	apicommon "github.com/Juniper/contrail/pkg/apisrv/common"
	"github.com/Juniper/contrail/pkg/apisrv/keystone"
	"github.com/Juniper/contrail/pkg/fileutil"
	"github.com/Juniper/contrail/pkg/format"
	kscommon "github.com/Juniper/contrail/pkg/keystone"
	"github.com/Juniper/contrail/pkg/logutil"
	"github.com/Juniper/contrail/pkg/models/basemodels"
	"github.com/Juniper/contrail/pkg/services/baseservices"
	"github.com/Juniper/contrail/pkg/sync"
	"github.com/Juniper/contrail/pkg/testutil"
	integrationetcd "github.com/Juniper/contrail/pkg/testutil/integration/etcd"
)

const (
	collectTimeout = 5 * time.Second
)

// TestMain is a function that can be called inside package specific TestMain
// to enable integration testing capabilities.
func TestMain(m *testing.M, s **APIServer) {
	WithTestDBs(func(dbType string) {
		cacheDB, cancelEtcdEventProducer, err := RunCacheDB()
		if err != nil {
			logrus.WithError(err).Fatal("Failed to run Cache DB")
		}
		defer testutil.LogFatalIfError(cancelEtcdEventProducer)

		if viper.GetBool("sync.enabled") {
			sync, err := sync.NewService()
			if err != nil {
				logrus.WithError(err).Fatal("Failed to initialize Sync")
			}
			errChan := RunConcurrently(sync)
			defer CloseFatalIfError(sync, errChan)
			<-sync.DumpDone()
		}

		if srv, err := NewRunningServer(&APIServerConfig{
			DBDriver:           dbType,
			RepoRootPath:       "../../..",
			EnableEtcdNotifier: true,
			CacheDB:            cacheDB,
		}); err != nil {
			logrus.WithError(err).Fatal("Failed to initialize API Server")
		} else {
			*s = srv
		}
		defer testutil.LogFatalIfError((*s).Close)

		if code := m.Run(); code != 0 {
			os.Exit(code)
		}
	})
}

// RunTest invokes integration test located in "tests" directory
func RunTest(t *testing.T, name string, server *APIServer) {
	testScenario, err := LoadTest(fmt.Sprintf("./tests/%s.yml", format.CamelToSnake(name)), nil)
	assert.NoError(t, err, "failed to load test data")
	RunCleanTestScenario(t, testScenario, server)
}

// WithTestDBs does test setup and run tests for
// all supported db types.
func WithTestDBs(f func(dbType string)) {
	if err := initViperConfig(); err != nil {
		logrus.WithError(err).Fatal("Failed to initialize Viper config")
	}
	if err := logutil.Configure(viper.GetString("log_level")); err != nil {
		logrus.WithError(err).Fatal()
	}

	testDBs := viper.GetStringMap("test_database")
	if len(testDBs) == 0 {
		logrus.Fatal("Test suite expected test database definitions under 'test_database' key")
	}

	for _, iConfig := range testDBs {
		config := format.InterfaceToInterfaceMap(iConfig)
		dbType, ok := config["type"].(string)
		if !ok {
			logrus.WithFields(logrus.Fields{
				"db-type": dbType,
			}).Error("Failed to read test_database.type value")
		}
		viper.Set("database.type", dbType)
		viper.Set("database.host", config["host"])
		viper.Set("database.user", config["user"])
		viper.Set("database.name", config["name"])
		viper.Set("database.password", config["password"])
		viper.Set("database.dialect", config["dialect"])

		if val, ok := config["use_sync"]; ok && cast.ToBool(val) == true {
			viper.Set("server.notify_etcd", false)
			viper.Set("sync.enabled", true)
		} else {
			viper.Set("server.notify_etcd", true)
			viper.Set("sync.enabled", false)
		}

		logrus.WithField("db-type", dbType).Info("Starting tests with DB")
		f(dbType)
		logrus.WithField("db-type", dbType).Info("Finished tests with DB")
	}
}

func initViperConfig() error {
	viper.SetConfigName("test_config")
	viper.SetConfigType("yml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("../../../../sample")
	viper.AddConfigPath("../../../sample")
	viper.AddConfigPath("../../sample")
	viper.AddConfigPath("../sample")
	viper.AddConfigPath("./sample")
	viper.AddConfigPath("./test_data")
	viper.SetEnvPrefix("contrail")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()
	return viper.ReadInConfig()
}

const (
	defaultClientID = "default"
	defaultDomainID = "default"
)

//AddKeystoneProjectAndUser adds Keystone project and user in Server internal state.
func AddKeystoneProjectAndUser(s *apisrv.Server, testID string) {
	assignment := s.Keystone.Assignment.(*keystone.StaticAssignment) // nolint: errcheck
	assignment.Projects[testID] = &kscommon.Project{
		Domain: assignment.Domains[defaultDomainID],
		ID:     testID,
		Name:   testID,
	}

	assignment.Users[testID] = &kscommon.User{
		Domain:   assignment.Domains[defaultDomainID],
		ID:       testID,
		Name:     testID,
		Password: testID,
		Roles: []*kscommon.Role{
			{
				ID:      "member",
				Name:    "Member",
				Project: assignment.Projects[testID],
			},
		},
	}
}

// Event represents event received from etcd watch.
type Event struct {
	Data     map[string]interface{} `yaml:"data,omitempty"`
	SyncOnly bool                   `yaml:"sync_only,omitempty"`
}

func convertEventsIntoMapList(events []Event) []map[string]interface{} {
	m := make([]map[string]interface{}, len(events))
	for i := range events {
		m[i] = events[i].Data
	}
	return m
}

// Watchers map contains slices of events that should be emitted on
// etcd key matching the map key.
type Watchers map[string][]Event

// Waiters map contains slices of events that have to be emitted
// during single task.
type Waiters map[string][]Event

//Task has API request and expected response.
type Task struct {
	Name     string          `yaml:"name,omitempty"`
	Client   string          `yaml:"client,omitempty"`
	Request  *client.Request `yaml:"request,omitempty"`
	Expect   interface{}     `yaml:"expect,omitempty"`
	Watchers Watchers        `yaml:"watchers,omitempty"`
	Waiters  Waiters         `yaml:"await,omitempty"`
}

// CleanTask defines clean task
type CleanTask struct {
	Client string   `yaml:"client,omitempty"`
	Path   string   `yaml:"path,omitempty"`
	FQName []string `yaml:"fq_name,omitempty"`
	Kind   string   `yaml:"kind,omitempty"`
}

//TestScenario has a list of tasks.
type TestScenario struct {
	Name                  string                  `yaml:"name,omitempty"`
	Description           string                  `yaml:"description,omitempty"`
	IntentCompilerEnabled bool                    `yaml:"intent_compiler_enabled,omitempty"`
	Tables                []string                `yaml:"tables,omitempty"`
	Clients               map[string]*client.HTTP `yaml:"clients,omitempty"`
	CleanTasks            []CleanTask             `yaml:"cleanup,omitempty"`
	Workflow              []*Task                 `yaml:"workflow,omitempty"`
	Watchers              Watchers                `yaml:"watchers,omitempty"`
	TestData              interface{}             `yaml:"test_data,omitempty"`
}

//LoadTest load testscenario.
func LoadTest(file string, ctx map[string]interface{}) (*TestScenario, error) {
	var testScenario TestScenario
	err := LoadTestScenario(&testScenario, file, ctx)
	return &testScenario, err
}

//LoadTestScenario load testscenario.
func LoadTestScenario(testScenario *TestScenario, file string, ctx map[string]interface{}) error {
	template, err := pongo2.FromFile(file)
	if err != nil {
		return errors.Wrap(err, "failed to read test data template")
	}

	content, err := template.ExecuteBytes(ctx)
	if err != nil {
		return errors.Wrap(err, "failed to apply test data template")
	}
	return yaml.UnmarshalStrict(content, testScenario)
}

type trackedResource struct {
	Path   string
	Client string
}

type clientsList map[string]*client.HTTP

// RunCleanTestScenario runs test scenario from loaded yaml file, expects no resources leftovers
func RunCleanTestScenario(
	t *testing.T,
	testScenario *TestScenario,
	server *APIServer,
) {
	logrus.WithField("test-scenario", testScenario.Name).Debug("Running clean test scenario")
	ctx := context.Background()
	checkWatchers := StartWatchers(t, testScenario.Name, testScenario.Watchers)
	stopIC := startIntentCompiler(t, testScenario, server)
	defer stopIC()

	clients := prepareClients(ctx, t, testScenario, server)
	tracked := runTestScenario(ctx, t, testScenario, clients, server.APIServer.DBService)
	cleanupTrackedResources(ctx, tracked, clients)

	checkWatchers(t)
}

// RunDirtyTestScenario runs test scenario from loaded yaml file, leaves all resources after scenario
func RunDirtyTestScenario(t *testing.T, testScenario *TestScenario, server *APIServer) func() {
	logrus.WithField("test-scenario", testScenario.Name).Debug("Running dirty test scenario")
	ctx := context.Background()
	clients := prepareClients(ctx, t, testScenario, server)
	tracked := runTestScenario(ctx, t, testScenario, clients, server.APIServer.DBService)
	cleanupFunc := func() {
		cleanupTrackedResources(ctx, tracked, clients)
	}
	return cleanupFunc
}

func cleanupTrackedResources(ctx context.Context, tracked []trackedResource, clients map[string]*client.HTTP) {
	for _, tr := range tracked {
		response, err := clients[tr.Client].EnsureDeleted(ctx, tr.Path, nil)
		if err != nil {
			logrus.WithError(err).WithFields(logrus.Fields{
				"url-path":    tr.Path,
				"http-client": tr.Client,
			}).Error("Deleting dirty resource failed - ignoring")
		}
		if response.StatusCode == http.StatusOK {
			logrus.WithFields(logrus.Fields{
				"url-path":    tr.Path,
				"http-client": tr.Client,
			}).Warn("Test scenario has not deleted resource but should have deleted - test scenario is dirty")
		}
	}
}

// StartWaiters checks if there are emitted events described before.
func StartWaiters(t *testing.T, task string, waiters Waiters) func(t *testing.T) {
	checks := []func(t *testing.T){}

	ec := integrationetcd.NewEtcdClient(t)

	for key := range waiters {
		events := waiters[key]
		getMissingEvents := ec.WaitForEvents(
			key, convertEventsIntoMapList(events), collectTimeout, clientv3.WithPrefix(),
		)
		checks = append(checks, createWaiterChecker(task, getMissingEvents))
	}

	return func(t *testing.T) {
		defer ec.Close(t)
		for _, c := range checks {
			c(t)
		}
	}
}

func createWaiterChecker(
	task string, getMissingEvents func() ([]map[string]interface{}, error),
) func(t *testing.T) {
	return func(t *testing.T) {
		notFound, err := getMissingEvents()
		assert.Equal(t, nil, err, "waiter for task %v got an error while collecting events: %v", task, err)
		assert.Equal(t, 0, len(notFound), "etcd didn't emitted events %v for task %v", notFound, task)
	}
}

// StartWatchers checks if events emitted to etcd match those given in watchers dict.
func StartWatchers(t *testing.T, task string, watchers Watchers, opts ...clientv3.OpOption) func(t *testing.T) {
	checks := []func(t *testing.T){}

	ec := integrationetcd.NewEtcdClient(t)

	syncEnabled := viper.GetBool("sync.enabled")
	for key := range watchers {
		events := filterEvents(watchers[key], func(e Event) bool {
			if e.SyncOnly {
				return syncEnabled
			}
			return true
		})
		collect := ec.WatchKeyN(key, len(events), collectTimeout, append(opts, clientv3.WithPrefix())...)
		checks = append(checks, createWatchChecker(task, collect, key, events))
	}

	return func(t *testing.T) {
		defer ec.Close(t)
		for _, c := range checks {
			c(t)
		}
	}
}

func filterEvents(evs []Event, pred func(Event) bool) []Event {
	result := make([]Event, 0, len(evs))
	for _, e := range evs {
		if pred(e) {
			result = append(result, e)
		}
	}
	return result
}

func createWatchChecker(task string, collect func() []string, key string, events []Event) func(t *testing.T) {
	return func(t *testing.T) {
		collected := collect()
		eventCount := len(events)
		assert.Equal(
			t, eventCount, len(collected),
			"etcd emitted not enough events on %s(got %v, expected %v)\n",
			key, collected, events,
		)

		for i, e := range events[:len(collected)] {
			c := collected[i]
			var data interface{} = map[string]interface{}{}
			if len(c) > 0 {
				err := json.Unmarshal([]byte(c), &data)
				assert.NoError(t, err)
			}
			testutil.AssertEqual(t, e.Data, data, fmt.Sprintf("task: %s\netcd event not equal for %s[%v]", task, key, i))
		}
	}
}

func startIntentCompiler(
	t *testing.T,
	testScenario *TestScenario,
	server *APIServer,
) context.CancelFunc {
	if testScenario.IntentCompilerEnabled {
		etcdClient := integrationetcd.NewEtcdClient(t)
		etcdClient.Clear(t)

		return RunIntentCompilationService(t, server.TestServer.URL)
	}
	return func() {}
}

func prepareClients(ctx context.Context, t *testing.T, testScenario *TestScenario, server *APIServer) clientsList {
	clients := clientsList{}

	for key, client := range testScenario.Clients {
		client.AuthURL = server.TestServer.URL + "/keystone/v3"
		client.Endpoint = server.TestServer.URL
		client.InSecure = true
		client.Debug = true

		client.Init()

		clients[key] = client

		err := clients[key].Login(ctx)
		assert.NoError(t, err, fmt.Sprintf("client %q failed to login", client.ID))
	}
	return clients
}

func runTestScenario(
	ctx context.Context,
	t *testing.T,
	testScenario *TestScenario,
	clients clientsList,
	m baseservices.MetadataGetter,
) (tracked []trackedResource) {
	for _, cleanTask := range testScenario.CleanTasks {
		logrus.WithFields(logrus.Fields{
			"test-scenario": testScenario.Name,
			"clean-task":    cleanTask,
		}).Debug("Deleting existing resources before test scenario workflow")
		err := performCleanup(ctx, cleanTask, getClientByID(cleanTask.Client, clients), m)
		if err != nil {
			logrus.WithError(err).WithFields(logrus.Fields{
				"test-scenario": testScenario.Name,
				"clean-task":    cleanTask,
			}).Error("Failed to delete existing resource before running workflow - ignoring")
		}
	}
	for _, task := range testScenario.Workflow {
		logrus.WithFields(logrus.Fields{
			"test-scenario": testScenario.Name,
			"task":          task.Name,
		}).Info("Starting task")
		checkWatchers := StartWatchers(t, task.Name, task.Watchers)
		checkWaiters := StartWaiters(t, testScenario.Name, task.Waiters)
		task.Request.Data = fileutil.YAMLtoJSONCompat(task.Request.Data)
		clientID := defaultClientID
		if task.Client != "" {
			clientID = task.Client
		}
		client, ok := clients[clientID]
		if !assert.True(t, ok,
			"Client %q not defined in test scenario %q task %q", clientID, testScenario.Name, task.Name) {
			break
		}
		response, err := client.DoRequest(ctx, task.Request)
		assert.NoError(t, err, fmt.Sprintf("In test scenario %q task %q failed", testScenario.Name, task.Name))
		tracked = handleTestResponse(task, response.StatusCode, err, tracked)

		task.Expect = fileutil.YAMLtoJSONCompat(task.Expect)
		ok = testutil.AssertEqual(t, task.Expect, task.Request.Output,
			fmt.Sprintf("In test scenario %q task %q failed", testScenario.Name, task.Name))
		checkWatchers(t)
		checkWaiters(t)
		if !ok {
			break
		}
	}
	// Reverse the order in tracked array so delete of nested resources is possible
	// https://github.com/golang/go/wiki/SliceTricks#reversing
	for left, right := 0, len(tracked)-1; left < right; left, right = left+1, right-1 {
		tracked[left], tracked[right] = tracked[right], tracked[left]
	}
	return tracked
}

func performCleanup(
	ctx context.Context,
	cleanTask CleanTask,
	client *client.HTTP,
	m baseservices.MetadataGetter,
) error {
	switch {
	case client == nil:
		return fmt.Errorf("failed to delete resource, got nil http client")
	case cleanTask.Path != "":
		return cleanPath(ctx, cleanTask.Path, client)
	case cleanTask.Kind != "" && cleanTask.FQName != nil:
		return cleanByFQNameAndKind(ctx, cleanTask.FQName, cleanTask.Kind, client, m)
	default:
		return fmt.Errorf("invalid clean task %v", cleanTask)
	}
}

func getClientByID(clientID string, clients clientsList) *client.HTTP {
	if clientID == "" {
		clientID = defaultClientID
	}
	return clients[clientID]
}

func cleanPath(ctx context.Context, path string, client *client.HTTP) error {
	response, err := client.EnsureDeleted(ctx, path, nil)
	if err != nil && response.StatusCode != http.StatusNotFound {
		return fmt.Errorf("failed to delete resource, got status code %v", response.StatusCode)
	}
	return nil
}

func cleanByFQNameAndKind(
	ctx context.Context,
	fqName []string,
	kind string,
	client *client.HTTP,
	m baseservices.MetadataGetter,
) error {
	metadata, err := m.GetMetadata(ctx, basemodels.Metadata{
		Type:   kind,
		FQName: fqName,
	})
	if err != nil {
		return errors.Wrapf(err, "failed to fetch uuid for %s with fqName %s", kind, fqName)
	}
	return cleanPath(ctx, "/"+kind+"/"+metadata.UUID, client)
}

func extractResourcePathFromJSON(data interface{}) (path string) {
	if data, ok := data.(map[string]interface{}); ok {
		var uuid interface{}
		if uuid, ok = data["uuid"]; !ok {
			return path
		}
		if suuid, ok := uuid.(string); ok {
			path = "/" + suuid
		}
	}
	return path
}

func extractSyncOperation(syncOp map[string]interface{}, client string) []trackedResource {
	resources := []trackedResource{}
	var operIf, kindIf interface{}
	var ok bool
	if kindIf, ok = syncOp["kind"]; !ok {
		return nil
	}
	if operIf, ok = syncOp["operation"]; !ok {
		return nil
	}
	var oper, kind string
	if oper, ok = operIf.(string); !ok {
		return nil
	}
	if oper != "CREATE" {
		return nil
	}
	if kind, ok = kindIf.(string); !ok {
		return nil
	}
	if dataIf, ok := syncOp["data"]; ok {
		if path := extractResourcePathFromJSON(dataIf); path != "" {
			return append(resources, trackedResource{Path: "/" + kind + path, Client: client})
		}
	}

	return resources
}

func handleTestResponse(task *Task, code int, rerr error, tracked []trackedResource) []trackedResource {
	if task.Request.Output != nil && task.Request.Method == "POST" && code == http.StatusOK && rerr == nil {
		clientID := defaultClientID
		if task.Client != "" {
			clientID = task.Client
		}
		tracked = trackResponse(task.Request.Output, clientID, tracked)
	}
	return tracked
}

func trackResponse(respDataIf interface{}, clientID string, tracked []trackedResource) []trackedResource {
	switch respData := respDataIf.(type) {
	case []interface{}:
		logrus.Warn("Not handled SYNC request - yet!")
		for _, syncOpIf := range respData {
			if syncOp, ok := syncOpIf.(map[string]interface{}); ok {
				tracked = append(tracked, extractSyncOperation(syncOp, clientID)...)
			}
		}
	case map[string]interface{}:
		for k, v := range respData {
			if path := extractResourcePathFromJSON(v); path != "" {
				tracked = append(tracked, trackedResource{Path: "/" + k + path, Client: clientID})
			}
		}
	}
	return tracked
}

func newWellKnownListener(serve string) net.Listener {
	if serve != "" {
		l, err := net.Listen("tcp", serve)
		if err != nil {
			panic(fmt.Sprintf("httptest: failed to listen on %v: %v", serve, err))
		}
		return l
	}
	l, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		if l, err = net.Listen("tcp6", "[::1]:0"); err != nil {
			panic(fmt.Sprintf("httptest: failed to listen on a port: %v", err))
		}
	}
	return l
}

// NewWellKnownServer returns a new server with given port
func NewWellKnownServer(serve string, handler http.Handler) *httptest.Server {
	return &httptest.Server{
		Listener: newWellKnownListener(serve),
		Config:   &http.Server{Handler: handler},
	}
}

// MockServerWithKeystone mocks keystone server
func MockServerWithKeystone(serve, keystoneAuthURL string) *httptest.Server {
	// Echo instance
	e := echo.New()
	keystoneClient := keystone.NewKeystoneClient(keystoneAuthURL, true)
	endpointStore := apicommon.MakeEndpointStore()
	k, err := keystone.Init(e, endpointStore, keystoneClient)
	if err != nil {
		return nil
	}

	// Routes
	e.POST("/v3/auth/tokens", k.CreateTokenAPI)
	e.GET("/v3/auth/tokens", k.ValidateTokenAPI)
	e.GET("/v3/auth/projects", k.GetProjectAPI)
	mockServer := NewWellKnownServer(serve, e)
	mockServer.Start()
	return mockServer
}
