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
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v2"

	"github.com/Juniper/contrail/pkg/apisrv"
	"github.com/Juniper/contrail/pkg/apisrv/client"
	apicommon "github.com/Juniper/contrail/pkg/apisrv/common"
	"github.com/Juniper/contrail/pkg/apisrv/keystone"
	"github.com/Juniper/contrail/pkg/cast"
	"github.com/Juniper/contrail/pkg/fileutil"
	kscommon "github.com/Juniper/contrail/pkg/keystone"
	"github.com/Juniper/contrail/pkg/logging"
	"github.com/Juniper/contrail/pkg/sync"
	"github.com/Juniper/contrail/pkg/testutil"
	"github.com/Juniper/contrail/pkg/testutil/integration/etcd"
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
			log.Fatal(err)
		}
		defer testutil.LogFatalIfError(cancelEtcdEventProducer)

		if viper.GetBool("sync.enabled") {
			sync, err := sync.NewService()
			if err != nil {
				log.Fatalf("Error initializing integration Sync: %v", err)
			}
			errChan := RunConcurrently(sync)

			defer CloseFatalIfError(sync, errChan)
		}

		if srv, err := NewRunningServer(&APIServerConfig{
			DBDriver:           dbType,
			RepoRootPath:       "../../..",
			EnableEtcdNotifier: true,
			CacheDB:            cacheDB,
		}); err != nil {
			log.Fatalf("Error initializing integration APIServer: %v", err)
		} else {
			*s = srv
		}
		defer testutil.LogFatalIfError((*s).Close)

		if code := m.Run(); code != 0 {
			os.Exit(code)
		}
	})
}

// WithTestDBs does test setup and run tests for
// all supported db types.
func WithTestDBs(f func(dbType string)) {
	err := initViperConfig()
	if err != nil {
		log.Fatal(err)
	}
	logging.SetLogLevel()
	testDBs := viper.GetStringMap("test_database")
	if len(testDBs) == 0 {
		log.Fatal("Test suite expected test database definitions under 'test_database' key")
	}

	for _, iConfig := range testDBs {
		config := cast.InterfaceToInterfaceMap(iConfig)
		dbType, ok := config["type"].(string)
		if !ok {
			log.Error("Failed to read dbType: %v (%T)", dbType, dbType)
		}
		viper.Set("database.type", dbType)
		viper.Set("database.host", config["host"])
		viper.Set("database.user", config["user"])
		viper.Set("database.name", config["name"])
		viper.Set("database.password", config["password"])
		viper.Set("database.dialect", config["dialect"])

		if val, ok := config["use_sync"]; ok && val != "true" {
			viper.Set("server.notify_etcd", false)
			viper.Set("sync.enabled", true)
		} else {
			viper.Set("server.notify_etcd", true)
			viper.Set("sync.enabled", false)
		}

		log.WithField("dbType", dbType).Info("Starting tests for DB")
		f(dbType)
		log.WithField("dbType", dbType).Info("Finished tests for DB")
	}
}

func initViperConfig() error {
	viper.SetConfigName("test_config")
	viper.SetConfigType("yml")
	viper.AddConfigPath(".")
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
type Event = map[string]interface{}

// Watchers map contains slices of events that should be emitted on
// etcd key matching the map key.
type Watchers = map[string][]Event

//Task has API request and expected response.
type Task struct {
	Name     string          `yaml:"name,omitempty"`
	Client   string          `yaml:"client,omitempty"`
	Request  *client.Request `yaml:"request,omitempty"`
	Expect   interface{}     `yaml:"expect,omitempty"`
	Watchers Watchers        `yaml:"watchers,omitempty"`
}

//TestScenario has a list of tasks.
type TestScenario struct {
	Name                  string                  `yaml:"name,omitempty"`
	Description           string                  `yaml:"description,omitempty"`
	IntentCompilerEnabled bool                    `yaml:"intent_compiler_enabled,omitempty"`
	Tables                []string                `yaml:"tables,omitempty"`
	Clients               map[string]*client.HTTP `yaml:"clients,omitempty"`
	Cleanup               []map[string]string     `yaml:"cleanup,omitempty"`
	Workflow              []*Task                 `yaml:"workflow,omitempty"`
	Watchers              Watchers                `yaml:"watchers,omitempty"`
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
	return yaml.Unmarshal(content, testScenario)
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
	log.Debug("Running clean test scenario: ", testScenario.Name)
	ctx := context.Background()
	checkWatchers := StartWatchers(t, testScenario.Watchers)
	stopIC := startIntentCompiler(t, testScenario, server)
	defer stopIC()

	clients := prepareClients(ctx, t, testScenario, server)
	tracked := runTestScenario(ctx, t, testScenario, clients)
	cleanupTrackedResources(ctx, tracked, clients)

	checkWatchers(t)
}

// RunDirtyTestScenario runs test scenario from loaded yaml file, leaves all resources after scenario
func RunDirtyTestScenario(t *testing.T, testScenario *TestScenario, server *APIServer) func() {
	log.Debug("Running *DIRTY* test scenario: ", testScenario.Name)
	ctx := context.Background()
	clients := prepareClients(ctx, t, testScenario, server)
	tracked := runTestScenario(ctx, t, testScenario, clients)
	cleanupFunc := func() {
		cleanupTrackedResources(ctx, tracked, clients)
	}
	return cleanupFunc
}

func cleanupTrackedResources(ctx context.Context, tracked []trackedResource, clients map[string]*client.HTTP) {
	for _, tr := range tracked {
		response, err := clients[tr.Client].EnsureDeleted(ctx, tr.Path, nil)
		if err != nil {
			log.Errorf("Ignored Error deleting dirty resource: %v, for url path '%v' with client %v", err, tr.Path, tr.Client)
			continue // It is desired to loop over all resources even with errors
		}
		if response.StatusCode != 404 {
			log.Warnf("DIRTY test scenario: left resource with path '%v'", tr.Path)
		}
	}
}

// StartWatchers checks if events emitted to etcd match those given in watchers dict.
func StartWatchers(t *testing.T, watchers Watchers) func(t *testing.T) {
	checks := []func(t *testing.T){}

	ec := integrationetcd.NewEtcdClient(t)
	for key := range watchers {
		events := watchers[key]
		collect := ec.WatchKeyN(key, len(events), collectTimeout, clientv3.WithPrefix())

		checks = append(checks, createWatchChecker(collect, key, events))
	}

	return func(t *testing.T) {
		defer ec.Close(t)
		for _, c := range checks {
			c(t)
		}
	}
}

func createWatchChecker(collect func() []string, key string, events []Event) func(t *testing.T) {
	return func(t *testing.T) {
		collected := collect()
		assert.Equal(
			t, len(events), len(collected), "etcd emitted not enough events on %s\n", key,
		)
		for i, e := range events[:len(collected)] {
			c := collected[i]
			var data interface{} = map[string]interface{}{}
			if len(c) > 0 {
				err := json.Unmarshal([]byte(c), &data)
				assert.NoError(t, err)
			}
			testutil.AssertEqual(
				t, e, data, "etcd event not equal for %s[%v]:\n%s", key, i,
			)
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
		//Rewrite endpoint for test server
		client.Endpoint = server.TestServer.URL
		client.AuthURL = server.TestServer.URL + "/keystone/v3"
		client.InSecure = true
		client.Init()

		clients[key] = client

		err := clients[key].Login(ctx)
		assert.NoError(t, err, "client failed to login")
	}
	return clients
}

func runTestScenario(ctx context.Context,
	t *testing.T, testScenario *TestScenario, clients clientsList) (tracked []trackedResource) {
	for _, cleanTask := range testScenario.Cleanup {
		clientID := cleanTask["client"]
		if clientID == "" {
			clientID = defaultClientID
		}
		client := clients[clientID]
		// delete existing resources.
		log.Debugf("[Clean task] Path: %s, TestScenario: %s", cleanTask["path"], testScenario.Name)
		response, err := client.EnsureDeleted(ctx, cleanTask["path"], nil) // nolint
		if err != nil && response.StatusCode != 404 {
			log.Debug(err)
		}
	}
	for _, task := range testScenario.Workflow {
		log.Infof("[Task] Name: %s, TestScenario: %s", task.Name, testScenario.Name)
		checkWatchers := StartWatchers(t, task.Watchers)

		task.Request.Data = fileutil.YAMLtoJSONCompat(task.Request.Data)
		clientID := defaultClientID
		if task.Client != "" {
			clientID = task.Client
		}
		client, ok := clients[clientID]
		if !assert.True(t, ok,
			"Client '%v' not defined in test scenario '%v' task '%v'", clientID, testScenario.Name, task) {
			break
		}
		response, err := client.DoRequest(ctx, task.Request)
		assert.NoError(t, err, fmt.Sprintf("In test scenario '%v' task '%v' failed", testScenario.Name, task))
		tracked = handleTestResponse(task, response.StatusCode, err, tracked)

		task.Expect = fileutil.YAMLtoJSONCompat(task.Expect)
		ok = testutil.AssertEqual(t, task.Expect, task.Request.Output,
			fmt.Sprintf("In test scenario '%v' task' %v' failed", testScenario.Name, task))
		checkWatchers(t)
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
	if task.Request.Output != nil && task.Request.Method == "POST" && code == 200 && rerr == nil {
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
		log.Warn("Not handled SYNC request - yet!")
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
