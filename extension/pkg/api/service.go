package api // nolint: golint

import (
	"github.com/spf13/viper"

	"github.com/Juniper/contrail/extension/pkg/db"
	"github.com/Juniper/contrail/extension/pkg/models"
	"github.com/Juniper/contrail/extension/pkg/services"
	"github.com/Juniper/contrail/pkg/apisrv"
)

// Init initializes given server.
func Init(s *apisrv.Server) error {
	var serviceChain []services.Service
	dbService := db.NewService(s.DB(), viper.GetString("database.dialect"))
	tv, err := models.NewTypeValidatorWithFormat()
	if err != nil {
		return err
	}

	service := &services.ContrailService{
		BaseService:    services.BaseService{},
		TypeValidator:  tv,
		MetadataGetter: dbService,
	}

	serviceChain = append(serviceChain, service)
	serviceChain = append(serviceChain, dbService)
	services.Chain(serviceChain...)

	service.RegisterRESTAPI(s.Echo)
	services.RegisterContrailServiceServer(s.GRPCServer, service)
	return nil
}
