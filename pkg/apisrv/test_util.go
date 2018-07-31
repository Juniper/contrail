package apisrv

import (
	"fmt"
	"net"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/flosch/pongo2"
	"github.com/labstack/echo"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v2"

	"github.com/Juniper/contrail/pkg/apisrv/client"
	apicommon "github.com/Juniper/contrail/pkg/apisrv/common"
	"github.com/Juniper/contrail/pkg/apisrv/keystone"
	"github.com/Juniper/contrail/pkg/common"
	"github.com/Juniper/contrail/pkg/testutil"
	log "github.com/sirupsen/logrus"
)

const (
	defaultClientID = "default"
	defaultDomainID = "default"
)

// TestServer is httptest.Server instance
var TestServer *httptest.Server

// APIServer is test API Server instance
var APIServer *Server

// SetupAndRunTest does test setup and run tests for
// all supported db types.
func SetupAndRunTest(m *testing.M) {
	err := initViperConfig()
	if err != nil {
		log.Fatal(err)
	}
	common.SetLogLevel()

	for _, iConfig := range viper.GetStringMap("test_database") {
		config := common.InterfaceToInterfaceMap(iConfig)
		viper.Set("database.type", config["type"])
		viper.Set("database.host", config["host"])
		viper.Set("database.user", config["user"])
		viper.Set("database.name", config["name"])
		viper.Set("database.password", config["password"])
		viper.Set("database.dialect", config["dialect"])
		RunTestForDB(m, viper.GetString("database.type"))
	}
}

func initViperConfig() error {
	viper.SetConfigName("contrail")
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

// RunTestForDB runs tests for all supported DB
func RunTestForDB(m *testing.M, dbType string) {
	server, testServer := LaunchTestAPIServer()
	defer testServer.Close()
	defer LogFatalIfErr(server.Close)

	log.WithField("dbType", dbType).Info("Starting tests for DB")
	code := m.Run()
	log.WithField("dbType", dbType).Info("Finished tests for DB")
	if code != 0 {
		os.Exit(code)
	}
}

//LaunchTestAPIServer used to launch test API Server.
func LaunchTestAPIServer() (*Server, *httptest.Server) {
	var err error
	APIServer, err = NewServer()
	if err != nil {
		log.Fatal(err)
	}

	TestServer = testutil.NewTestHTTPServer(APIServer.Echo)

	viper.Set("keystone.authurl", TestServer.URL+"/keystone/v3")
	err = APIServer.Init()
	if err != nil {
		log.Fatal(err)
	}

	return APIServer, TestServer
}

//AddKeystoneProjectAndUser adds Keystone project and user in Server internal state.
func AddKeystoneProjectAndUser(s *Server, testID string) {
	assignment := s.Keystone.Assignment.(*keystone.StaticAssignment)
	assignment.Projects[testID] = &keystone.Project{
		Domain: assignment.Domains[defaultDomainID],
		ID:     testID,
		Name:   testID,
	}

	assignment.Users[testID] = &keystone.User{
		Domain:   assignment.Domains[defaultDomainID],
		ID:       testID,
		Name:     testID,
		Password: testID,
		Roles: []*keystone.Role{
			{
				ID:      "member",
				Name:    "Member",
				Project: assignment.Projects[testID],
			},
		},
	}
}

// ForceProxyUpdate requests an immediate update of endpoints and waits for its completion.
func (s *Server) ForceProxyUpdate() {
	s.Proxy.forceUpdate()
}

//Task has API request and expected response.
type Task struct {
	Name    string          `yaml:"name,omitempty"`
	Client  string          `yaml:"client,omitempty"`
	Request *client.Request `yaml:"request,omitempty"`
	Expect  interface{}     `yaml:"expect,omitempty"`
}

//TestScenario has a list of tasks.
type TestScenario struct {
	Name        string                  `yaml:"name,omitempty"`
	Description string                  `yaml:"description,omitempty"`
	Tables      []string                `yaml:"tables,omitempty"`
	Clients     map[string]*client.HTTP `yaml:"clients,omitempty"`
	Cleanup     []map[string]string     `yaml:"cleanup,omitempty"`
	Workflow    []*Task                 `yaml:"workflow,omitempty"`
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
func RunCleanTestScenario(t *testing.T, testScenario *TestScenario) {
	log.Info("Running clean test scenario: ", testScenario.Name)
	clients := prepareClients(t, testScenario)
	tracked := runTestScenario(t, testScenario, clients)
	cleanupTrackedResources(tracked, clients)
}

// RunDirtyTestScenario runs test scenario from loaded yaml file, leaves all resources after scenario
func RunDirtyTestScenario(t *testing.T, testScenario *TestScenario) func() {
	log.Info("Running *DIRTY* test scenario: ", testScenario.Name)
	clients := prepareClients(t, testScenario)
	tracked := runTestScenario(t, testScenario, clients)
	cleanupFunc := func() {
		cleanupTrackedResources(tracked, clients)
	}
	return cleanupFunc
}

func cleanupTrackedResources(tracked []trackedResource, clients map[string]*client.HTTP) {
	log.Infof("There are %v resources to clean (in clean tests should be ZERO)", len(tracked))
	for i, tr := range tracked {
		log.Warnf("POST clean up resource %v / %v: %v {clien:t %v}", i+1, len(tracked), tr.Path, tr.Client)
		response, err := clients[tr.Client].EnsureDeleted(tr.Path, nil)
		if err != nil {
			log.Errorf("Error deleting dirty resource: %v, for url path '%v' with client %v", err, tr.Path, tr.Client)
			continue // It is desired to loop over all resources even with errors
		}
		if response.StatusCode != 404 {
			log.Warnf("DIRTY test scenario: left resource with path '%v'", tr.Path)
		}
	}
}

func prepareClients(t *testing.T, testScenario *TestScenario) clientsList {
	clients := clientsList{}

	for key, client := range testScenario.Clients {
		//Rewrite endpoint for test server
		client.Endpoint = TestServer.URL
		client.AuthURL = TestServer.URL + "/keystone/v3"
		client.InSecure = true
		client.Init()

		clients[key] = client

		err := clients[key].Login()
		assert.NoError(t, err, "client failed to login")
	}
	return clients
}

func runTestScenario(t *testing.T, testScenario *TestScenario, clients clientsList) (tracked []trackedResource) {
	for _, cleanTask := range testScenario.Cleanup {
		log.Debugf("CLEAN TASK -> %v", cleanTask)
		clientID := cleanTask["client"]
		if clientID == "" {
			clientID = defaultClientID
		}
		client := clients[clientID]
		// delete existing resources.
		log.Debugf("[Clean task] Path: %s, TestScenario: %s\n", cleanTask["path"], testScenario.Name)
		response, err := client.EnsureDeleted(cleanTask["path"], nil) // nolint
		if err != nil && response.StatusCode != 404 {
			log.Debug(err)
		}
	}
	log.Info("CLEANUP COMPETE! Starting test sequence...")
	for _, task := range testScenario.Workflow {
		log.Debugf("[Task] Name: %s, TestScenario: %s\n", task.Name, testScenario.Name)
		task.Request.Data = common.YAMLtoJSONCompat(task.Request.Data)
		clientID := defaultClientID
		if task.Client != "" {
			clientID = task.Client
		}
		client, ok := clients[clientID]
		if !assert.True(t, ok,
			"Client '%v' not defined in test scenario '%v' task '%v'", clientID, testScenario.Name, task) {
			break
		}
		response, err := client.DoRequest(task.Request)
		tracked = handleTestResponse(task, response.StatusCode, err, tracked)
		assert.NoError(t, err, fmt.Sprintf("In test scenario '%v' task '%v' failed", testScenario.Name, task))

		task.Expect = common.YAMLtoJSONCompat(task.Expect)
		ok = testutil.AssertEqual(t, task.Expect, task.Request.Output,
			fmt.Sprintf("In test scenario '%v' task' %v' failed", testScenario.Name, task))
		if !ok {
			log.Errorf("Assertion error was: %+v", err)
			break
		}
	}
	// Reverse the order in tracked array so delete of nested resources is possible
	// https://github.com/golang/go/wiki/SliceTricks#reversing
	for left, right := 0, len(tracked)-1; left < right; left, right = left+1, right-1 {
		tracked[left], tracked[right] = tracked[right], tracked[left]
	}
	log.Info("TEST SEQUENCE COMPLETE!")
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
	log.Infof("Tracked requests: %+v", tracked)
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

// LogFatalIfErr logs the err during function call
func LogFatalIfErr(f func() error) {
	if err := f(); err != nil {
		log.Fatal(err)
	}
}
