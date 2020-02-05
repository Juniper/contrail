package integration

import (
	"net/http"
	"net/http/httptest"
	"path"
	"testing"

	"github.com/Juniper/asf/pkg/apiserver"
	"github.com/Juniper/asf/pkg/db/basedb"
	"github.com/Juniper/asf/pkg/etcd"
	"github.com/Juniper/asf/pkg/logutil"
	"github.com/Juniper/contrail/pkg/cache"
	"github.com/Juniper/contrail/pkg/cmd/contrail"
	"github.com/Juniper/contrail/pkg/collector/analytics"
	"github.com/Juniper/contrail/pkg/db"
	"github.com/Juniper/contrail/pkg/endpoint"
	"github.com/Juniper/contrail/pkg/keystone"
	"github.com/Juniper/contrail/pkg/neutron"
	"github.com/Juniper/contrail/pkg/proxy"
	"github.com/Juniper/contrail/pkg/replication"
	"github.com/Juniper/contrail/pkg/services"
	"github.com/Juniper/contrail/pkg/testutil"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	asfkeystone "github.com/Juniper/asf/pkg/keystone"
	integrationetcd "github.com/Juniper/contrail/pkg/testutil/integration/etcd"
)

const (
	dbUser     = "root"
	dbPassword = "contrail123"
	dbName     = "contrail_test"
)

// Keystone credentials.
const (
	DefaultDomainID     = "default"
	DefaultDomainName   = "default"
	AdminProjectID      = "admin"
	AdminProjectName    = "admin"
	DemoProjectID       = "demo"
	DemoProjectName     = "demo"
	NeutronProjectID    = "aa907485e1f94a14834d8c69ed9cb3b2"
	NeutronProjectName  = "neutron"
	ServiceProjectID    = "service-uuid"
	ServiceProjectName  = "service"
	AdminRoleID         = "admin"
	AdminRoleName       = "admin"
	NeutronRoleID       = "aa907485e1f94a14834d8c69ed9cb3b2"
	NeutronRoleName     = "neutron"
	MemberRoleID        = "Member"
	MemberRoleName      = "Member"
	AdminUserID         = "alice"
	AdminUserName       = "Alice"
	AdminUserPassword   = "alice_password"
	KSAdminUserID       = "admin"
	KSAdminUserName     = "admin"
	KSAdminUserPassword = "contrail123"
	BobUserID           = "bob"
	BobUserName         = "Bob"
	BobUserPassword     = "bob_password"
	ServiceUserID       = "goapi-uuid"
	ServiceUserName     = "goapi"
	ServiceUserPassword = "goapi"
)

// APIServer is embedded API Server for testing purposes.
type APIServer struct {
	APIServer  *apiserver.Server
	testServer *httptest.Server
	dbService  *db.Service
	// TODO(Witaut): Remove this when AddKeystoneProjectAndUser is removed.
	keystone   *keystone.Keystone
	proxy      *proxy.Dynamic
	replicator *replication.Replicator
	log        *logrus.Entry
}

// APIServerConfig contains parameters for test API Server.
type APIServerConfig struct {
	CacheDB              *cache.DB
	RepoRootPath         string
	LogLevel             string
	EnableEtcdNotifier   bool
	DisableLogAPI        bool
	EnableRBAC           bool
	EnableVNCReplication bool
}

// NewRunningAPIServer creates new running test API Server for testing purposes.
// Call CloseT() or Close() method to release its resources.
func NewRunningAPIServer(t *testing.T, c *APIServerConfig) *APIServer {
	s, err := NewRunningServer(c)
	require.NoError(t, err)

	return s
}

// NewRunningServer creates new running API server with default testing configuration.
// Call CloseT() or Close() method to release its resources.
// TODO(dfurman): modify function to call contrail.StartServer() to remove duplication
// nolint: gocyclo
func NewRunningServer(c *APIServerConfig) (*APIServer, error) {
	setViperConfig(c)

	if err := logutil.Configure(c.LogLevel); err != nil {
		return nil, err
	}

	var serverHandler http.Handler
	ts := testutil.NewTestHTTPServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		serverHandler.ServeHTTP(w, r)
	}))
	viper.Set("keystone.authurl", ts.URL+keystone.LocalAuthPath)
	viper.Set("client.endpoint", ts.URL)

	analyticsCollector, err := analytics.NewCollectorFromGlobalConfig()
	if err != nil {
		return nil, err
	}
	analytics.AddLoggerHook(analyticsCollector)

	sqlDB, err := basedb.ConnectDB(analytics.WithCommitLatencyReporting(analyticsCollector))
	if err != nil {
		return nil, err
	}
	dbService := db.NewService(sqlDB)

	serviceChain, err := contrail.NewServiceChain(dbService, analyticsCollector)
	if err != nil {
		return nil, err
	}

	staticProxyPlugin, err := proxy.NewStaticByViper()
	if err != nil {
		return nil, err
	}

	es := endpoint.NewStore()
	dynamicProxy := proxy.NewDynamicFromViper(es, dbService)
	dynamicProxy.StartEndpointsSync() // TODO(dfurman): move to proxy constructor and use context for cancellation

	k, err := keystone.Init(es)
	if err != nil {
		return nil, err
	}

	plugins := []apiserver.APIPlugin{
		serviceChain,
		staticProxyPlugin,
		dynamicProxy,
		services.UploadCloudKeysPlugin{},
		analytics.BodyDumpPlugin{Collector: analyticsCollector},
		k,
		c.CacheDB,
	}

	if viper.GetBool("server.enable_vnc_neutron") {
		plugins = append(plugins, &neutron.Server{
			ReadService:       serviceChain,
			WriteService:      serviceChain,
			UserAgentKV:       serviceChain,
			IDToFQNameService: serviceChain,
			FQNameToIDService: serviceChain,
			InTransactionDoer: dbService,
			Log:               logutil.NewLogger("neutron-server"),
		})
	}

	server, err := apiserver.NewServer(plugins, contrail.NoAuthPaths())
	if err != nil {
		return nil, errors.Wrapf(err, "creating API Server failed")
	}

	// TODO(Witaut): Don't use Echo - an internal detail of Server.
	serverHandler = server.Echo

	var r *replication.Replicator
	if c.EnableVNCReplication {
		if r, err = startVNCReplicator(es); err != nil {
			return nil, errors.WithStack(err)
		}
	}

	return &APIServer{
		APIServer:  server,
		testServer: ts,
		dbService:  dbService,
		// TODO(Witaut): Remove this when AddKeystoneProjectAndUser is removed.
		keystone:   k,
		proxy:      dynamicProxy,
		replicator: r,
		log:        logutil.NewLogger("test-api-server"),
	}, nil
}

func setViperConfig(c *APIServerConfig) {
	setViper(map[string]interface{}{
		"aaa_mode":                    rbacConfig(c.EnableRBAC),
		"database.host":               "localhost",
		"database.user":               dbUser,
		"database.name":               dbName,
		"database.password":           dbPassword,
		"database.max_open_conn":      100,
		"database.connection_retries": 10,
		"database.retry_period":       3,
		"database.debug":              false,
		etcd.ETCDPathVK:               integrationetcd.Prefix,
		"keystone.local":              true,
		"keystone.assignment.type":    "static",
		"keystone.assignment.data":    keystoneAssignment(),
		"keystone.store.type":         "memory",
		"keystone.store.expire":       3600,
		"keystone.insecure":           true,
		"log_level":                   c.LogLevel,
		"server.notify_etcd":          c.EnableEtcdNotifier,
		"server.read_timeout":         10,
		"server.write_timeout":        5,
		"server.log_api":              !c.DisableLogAPI,
		"server.log_body":             !c.DisableLogAPI,
		"server.static_files.public":  path.Join(c.RepoRootPath, "public"),
		"server.enable_vnc_neutron":   true,
		"tls.enabled":                 false,
	})
}

func rbacConfig(enableRBAC bool) string {
	if enableRBAC {
		return "rbac"
	}
	return ""
}

func keystoneAssignment() *asfkeystone.StaticAssignment {
	a := asfkeystone.StaticAssignment{
		Domains: map[string]*asfkeystone.Domain{
			DefaultDomainID: {
				ID:   DefaultDomainID,
				Name: DefaultDomainName,
			},
		},
		Projects: make(map[string]*asfkeystone.Project),
		Users:    make(map[string]*asfkeystone.User),
	}
	a.Projects[AdminProjectID] = &asfkeystone.Project{
		Domain: a.Domains[DefaultDomainID],
		ID:     AdminProjectID,
		Name:   AdminProjectName,
	}
	a.Projects[DemoProjectID] = &asfkeystone.Project{
		Domain: a.Domains[DefaultDomainID],
		ID:     DemoProjectID,
		Name:   DemoProjectName,
	}
	a.Projects[NeutronProjectID] = &asfkeystone.Project{
		Domain: a.Domains[DefaultDomainID],
		ID:     NeutronProjectID,
		Name:   NeutronProjectName,
	}
	a.Projects[ServiceProjectID] = &asfkeystone.Project{
		Domain: a.Domains[DefaultDomainID],
		ID:     ServiceProjectID,
		Name:   ServiceProjectName,
	}
	a.Users[AdminUserID] = &asfkeystone.User{
		Domain:   a.Domains[DefaultDomainID],
		ID:       AdminUserID,
		Name:     AdminUserName,
		Password: AdminUserPassword,
		Roles: []*asfkeystone.Role{
			{
				ID:      AdminRoleID,
				Name:    AdminRoleName,
				Project: a.Projects[AdminProjectID],
			},
			{
				ID:      NeutronRoleID,
				Name:    NeutronRoleName,
				Project: a.Projects[NeutronProjectID],
			},
		},
	}
	a.Users[KSAdminUserID] = &asfkeystone.User{
		Domain:   a.Domains[DefaultDomainID],
		ID:       KSAdminUserID,
		Name:     KSAdminUserName,
		Password: KSAdminUserPassword,
		Roles: []*asfkeystone.Role{
			{
				ID:      AdminRoleID,
				Name:    AdminRoleName,
				Project: a.Projects[AdminProjectID],
			},
		},
	}
	a.Users[BobUserID] = &asfkeystone.User{
		Domain:   a.Domains[DefaultDomainID],
		ID:       BobUserID,
		Name:     BobUserName,
		Password: BobUserPassword,
		Roles: []*asfkeystone.Role{
			{
				ID:      MemberRoleID,
				Name:    MemberRoleName,
				Project: a.Projects[DemoProjectID],
			},
		},
	}
	a.Users[ServiceUserID] = &asfkeystone.User{
		Domain:   a.Domains[DefaultDomainID],
		ID:       ServiceUserID,
		Name:     ServiceUserName,
		Password: ServiceUserPassword,
		Roles: []*asfkeystone.Role{
			{
				ID:      AdminRoleID,
				Name:    AdminRoleName,
				Project: a.Projects[ServiceProjectID],
			},
		},
	}
	return &a
}

func setViper(config map[string]interface{}) {
	for k, v := range config {
		viper.SetDefault(k, v)
	}
}

func startVNCReplicator(es *endpoint.Store) (vncReplicator *replication.Replicator, err error) {
	vncReplicator, err = replication.New(es)
	if err != nil {
		return nil, err
	}
	err = vncReplicator.Start()
	if err != nil {
		return nil, err
	}
	return vncReplicator, nil
}

// URL returns server base URL.
func (s *APIServer) URL() string {
	return s.testServer.URL
}

// CloseT closes the server.
func (s *APIServer) CloseT(t *testing.T) {
	s.log.Debug("Closing test API server")
	err := s.Close()
	assert.NoError(t, err, "closing API Server failed")
}

// Close closes the server.
func (s *APIServer) Close() error {
	s.log.Debug("Closing test API server")
	if s.replicator != nil {
		s.replicator.Stop()
	}
	s.testServer.Close()
	s.proxy.StopEndpointsSync()
	return s.dbService.Close()
}

// ForceProxyUpdate requests an immediate update of endpoints and waits for its completion.
func (s *APIServer) ForceProxyUpdate() {
	s.proxy.ForceUpdate()
}

// AddKeystoneProjectAndUser adds Keystone project and user in Server internal state.
// TODO: Remove that, because it modifies internal state of SUT.
// TODO: Use pre-created Server's keystone assignment.
func (s *APIServer) AddKeystoneProjectAndUser(t testing.TB, testID string) func() {
	assignment, ok := s.keystone.Assignment.(*asfkeystone.StaticAssignment)
	require.True(t, ok, "s.keystone.Assignment should be a StaticAssignment")

	assignment.Projects[testID] = &asfkeystone.Project{
		Domain: assignment.Domains[DefaultDomainID],
		ID:     testID,
		Name:   testID,
	}

	assignment.Users[testID] = &asfkeystone.User{
		Domain:   assignment.Domains[DefaultDomainID],
		ID:       testID,
		Name:     testID,
		Password: testID,
		Roles: []*asfkeystone.Role{
			{
				ID:      "member",
				Name:    "Member",
				Project: assignment.Projects[testID],
			},
		},
	}

	return func() {
		delete(assignment.Projects, testID)
		delete(assignment.Users, testID)
	}
}
