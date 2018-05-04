package apisrv

import (
	"crypto/tls"
	"fmt"
	"net"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/flosch/pongo2"
	"github.com/labstack/echo"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	yaml "gopkg.in/yaml.v2"

	apicommon "github.com/Juniper/contrail/pkg/apisrv/common"
	"github.com/Juniper/contrail/pkg/apisrv/keystone"

	"github.com/Juniper/contrail/pkg/common"
	log "github.com/sirupsen/logrus"
)

// TestServer is http test server
var TestServer *httptest.Server

// APIServer is echo server
var APIServer *Server

// LogFatalIfErr logs the err during function call
func LogFatalIfErr(f func() error) {
	if err := f(); err != nil {
		log.Fatal(err)
	}
}

//LaunchTestAPIServer used to launch test api server
func LaunchTestAPIServer() (*Server, *httptest.Server) {
	var err error
	APIServer, err = NewServer()
	if err != nil {
		log.Fatal(err)
	}
	TestServer = httptest.NewUnstartedServer(APIServer.Echo)
	TestServer.TLS = new(tls.Config)
	TestServer.TLS.NextProtos = append(TestServer.TLS.NextProtos, "h2")
	TestServer.StartTLS()

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
		log.Debug(cleanTask["path"])
		_, err := client.Delete(cleanTask["path"], nil) // nolint
		if err != nil {
			log.Debug(err)
		}
	}
	for _, task := range testScenario.Workflow {
		log.Debug("[Step] ", task.Name)
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
