package apisrv

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/flosch/pongo2"
	"github.com/labstack/echo"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"

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

// SetupAndRunTest does test setup and run tests for
// all supported db types.
func SetupAndRunTest(m *testing.M) {
	err := common.InitConfig()
	if err != nil {
		log.Fatal(err)
	}
	common.SetLogLevel()
	dbConfig := viper.GetStringMap("test_database")
	for _, iConfig := range dbConfig {
		config := common.InterfaceToInterfaceMap(iConfig)
		viper.Set("database.type", config["type"])
		viper.Set("database.connection", config["connection"])
		viper.Set("database.dialect", config["dialect"])
		RunTestForDB(m)
	}
}

// RunTestForDB runs tests for all supported DB
func RunTestForDB(m *testing.M) {
	server, testServer := LaunchTestAPIServer()
	defer testServer.Close()
	defer LogFatalIfErr(server.Close)
	log.Info("starting test")
	code := m.Run()
	log.Info("finished test")
	if code != 0 {
		os.Exit(code)
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
	Name    string      `yaml:"name"`
	Client  string      `yaml:"client"`
	Request *Request    `yaml:"request"`
	Expect  interface{} `yaml:"expect"`
}

//TestScenario has a list of tasks.
type TestScenario struct {
	Name        string              `yaml:"name"`
	Description string              `yaml:"description"`
	Tables      []string            `yaml:"tables"`
	Clients     map[string]*Client  `yaml:"clients"`
	Cleanup     []map[string]string `yaml:"cleanup"`
	Workflow    []*Task             `yaml:"workflow"`
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

// GetTestFromTemplate create test input file from template with given context
func GetTestFromTemplate(t *testing.T, templateFile string, ctx map[string]interface{}) string {
	// create test data yaml from the template
	template, err := pongo2.FromFile(templateFile)
	assert.NoError(t, err, "failed to read test data template")

	content, err := template.ExecuteBytes(ctx)
	assert.NoError(t, err, "failed to apply test data template")

	fileName := filepath.Base(templateFile)
	var extension = filepath.Ext(fileName)
	var prefix = fileName[0 : len(fileName)-len(extension)]
	tmpfile, err := ioutil.TempFile("", prefix)
	assert.NoError(t, err, "failed to create test data tempfile")

	_, err = tmpfile.Write(content)
	assert.NoError(t, err, "failed to write test data to tempfile")

	err = tmpfile.Close()
	assert.NoError(t, err, "failed to close tempfile")
	testFile := tmpfile.Name() + ".yml"
	err = os.Rename(tmpfile.Name(), testFile)
	assert.NoError(t, err, "failed to rename test data file to yml file")
	return testFile
}

// LoadTestScenario loads test data from file
func LoadTestScenario(testScenario *TestScenario, file string) error {
	err := common.LoadFile(file, &testScenario)
	return err
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
