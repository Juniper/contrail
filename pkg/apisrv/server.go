package apisrv

import (
	"github.com/Juniper/contrail/pkg/constants"
	"github.com/Juniper/contrail/pkg/db"
	"github.com/Juniper/contrail/pkg/db/etcd"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
	"github.com/Juniper/contrail/pkg/types"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

// TODO(dfurman): move to ASF template
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
			ReadService: dbService,
			AAAMode:     viper.GetString("aaa_mode"),
		},
		&types.ContrailTypeLogicService{
			ReadService:       dbService,
			InTransactionDoer: dbService,
			AddressManager:    dbService,
			IntPoolAllocator:  dbService,
			MetadataGetter:    dbService,
			WriteService: &services.InternalContextWriteServiceWrapper{
				WriteService: cs,
			},
		},
		services.NewQuotaCheckerService(dbService),
	}

	if viper.GetBool("server.notify_etcd") {
		// TODO(Michał): Make the codec configurable
		en, enErr := etcd.NewNotifierService(viper.GetString(constants.ETCDPathVK), models.JSONCodec)
		if enErr != nil {
			return nil, errors.Wrap(enErr, "new notifier service")
		}
		serviceChain = append(serviceChain, en)
	}

	serviceChain = append(serviceChain, extraServices...)
	serviceChain = append(serviceChain, dbService)
	services.Chain(serviceChain...)
	return cs, nil
}
