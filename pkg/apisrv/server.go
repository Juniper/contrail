package apisrv

import (
	"github.com/Juniper/contrail/pkg/constants"
	"github.com/Juniper/contrail/pkg/db"
	"github.com/Juniper/contrail/pkg/db/etcd"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

// SetupServiceChain creates service chain using default services, given extra services and DB service.
// It puts extra services after default services but before etcd notifier service and DB service.
// It returns first service of the built service chain.
// TODO(dfurman): move to ASF template
// TODO(dfurman): return services.Service when ContrailService is split to plugins.
func SetupServiceChain(dbService *db.Service, extraServices ...services.Service) (*services.ContrailService, error) {
	tv, err := models.NewTypeValidatorWithFormat()
	if err != nil {
		return nil, errors.Wrap(err, "new type validator with format")
	}

	cs := &services.ContrailService{
		BaseService:        services.BaseService{},
		DBService:          dbService,
		TypeValidator:      tv,
		MetadataGetter:     dbService,
		InTransactionDoer:  dbService,
		IntPoolAllocator:   dbService,
		RefRelaxer:         dbService,
		UserAgentKVService: dbService,
	}

	serviceChain := []services.Service{
		cs,
		&services.RefUpdateToUpdateService{
			ReadService:       dbService,
			InTransactionDoer: dbService,
		},
		&services.SanitizerService{
			MetadataGetter: dbService,
		},
		&services.RBACService{
			ReadService:  dbService,
			AccessGetter: services.NewContrailAccessGetter(dbService),
			AAAMode:      viper.GetString("aaa_mode"),
		},
		services.NewQuotaCheckerService(dbService),
	}

	serviceChain = append(serviceChain, extraServices...)

	if viper.GetBool("server.notify_etcd") {
		// TODO(Michał): Make the codec configurable
		en, enErr := etcd.NewNotifierService(viper.GetString(constants.ETCDPathVK), models.JSONCodec)
		if enErr != nil {
			return nil, errors.Wrap(enErr, "new notifier service")
		}
		serviceChain = append(serviceChain, en)
	}

	serviceChain = append(serviceChain, dbService)
	services.Chain(serviceChain...)
	return cs, nil
}
