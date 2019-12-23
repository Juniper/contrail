package contrail

import (
	"github.com/Juniper/asf/pkg/logutil"
	"github.com/Juniper/contrail/pkg/apisrv"
	"github.com/Juniper/contrail/pkg/neutron"
	"github.com/spf13/viper"
)

func NewServer() (*apisrv.Server, error) {
	return apisrv.NewServer(
		neutronAPIPlugin{},
	)
}

type neutronAPIPlugin struct{}

func (p neutronAPIPlugin) RegisterAPI(s *apisrv.Server) {
	if viper.GetBool("server.enable_vnc_neutron") {
		setupNeutronService(s)
	}
}

func setupNeutronService(s *apisrv.Server) {
	n := &neutron.Server{
		ReadService:       s.Service,
		WriteService:      s.Service,
		UserAgentKV:       s.UserAgentKVServer,
		IDToFQNameService: s.IDToFQNameServer,
		FQNameToIDService: s.FQNameToIDServer,
		InTransactionDoer: s.DBService,
		Log:               logutil.NewLogger("neutron-server"),
	}
	n.RegisterNeutronAPI(s.Echo)
}
