package apisrv

import (
	"fmt"
	"net"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/Juniper/contrail/pkg/testutil"
	"github.com/flosch/pongo2"
	"github.com/labstack/echo"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v2"

	apicommon "github.com/Juniper/contrail/pkg/apisrv/common"
	"github.com/Juniper/contrail/pkg/apisrv/keystone"

	"github.com/Juniper/contrail/pkg/common"
	log "github.com/sirupsen/logrus"
)

// TestServer is httptest.Server instance
var TestServer *httptest.Server

// APIServer is test API Server instance
var APIServer *Server

// LogFatalIfErr logs the err during function call
func LogFatalIfErr(f func() error) {
	if err := f(); err != nil {
		log.Fatal(err)
	}
}

// SetupAndRunTest does test setup and run tests for
// all supported db types.
func SetupAndRunTest(m *testing.M) {
	err := common.InitConfig()
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

//CreateTestProject in keystone.
func CreateTestProject(s *Server, testID string) {
	assignment := s.Keystone.Assignment.(*keystone.StaticAssignment)
	assignment.Projects[testID] = &keystone.Project{
		Domain: assignment.Domains["default"],
		ID:     testID,
		Name:   testID,
	}

	assignment.Users[testID] = &keystone.User{
		Domain:   assignment.Domains["default"],
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

//Task has API request and expected response.
type Task struct {
	Name    string      `yaml:"name,omitempty"`
	Client  string      `yaml:"client,omitempty"`
	Request *Request    `yaml:"request,omitempty"`
	Expect  interface{} `yaml:"expect,omitempty"`
}

//TestScenario has a list of tasks.
type TestScenario struct {
	Name        string              `yaml:"name,omitempty"`
	Description string              `yaml:"description,omitempty"`
	Tables      []string            `yaml:"tables,omitempty"`
	Clients     map[string]*Client  `yaml:"clients,omitempty"`
	Cleanup     []map[string]string `yaml:"cleanup,omitempty"`
	Workflow    []*Task             `yaml:"workflow,omitempty"`
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

// RunTestScenario runs test from loaded test scenario
func RunTestScenario(t *testing.T, testScenario *TestScenario) {
	clients := map[string]*Client{}

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
	for _, cleanTask := range testScenario.Cleanup {
		fmt.Println(cleanTask)
		clientID := cleanTask["client"]
		if clientID == "" {
			clientID = "default"
		}
		client := clients[clientID]
		// delete existing resources.
		log.Debugf("[Clean task] Path: %s, TestScenario: %s\n", cleanTask["path"], testScenario.Name)
		response, err := client.Delete(cleanTask["path"], nil) // nolint
		if err != nil && response.StatusCode != 404 {
			log.Debug(err)
		}
	}
	for _, task := range testScenario.Workflow {
		log.Debugf("[Task] Name: %s, TestScenario: %s\n", task.Name, testScenario.Name)
		task.Request.Data = common.YAMLtoJSONCompat(task.Request.Data)
		clientID := "default"
		if task.Client != "" {
			clientID = task.Client
		}
		client := clients[clientID]
		_, err := client.DoRequest(task.Request)
		assert.NoError(t, err, fmt.Sprintf("task %v failed", task))
		task.Expect = common.YAMLtoJSONCompat(task.Expect)
		ok := common.AssertEqual(t, task.Expect, task.Request.Output, fmt.Sprintf("task %v failed", task))
		if !ok {
			break
		}
	}
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
