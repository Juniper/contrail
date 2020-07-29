package services

import (
	"github.com/Juniper/asf/pkg/apiserver"
	"github.com/Juniper/asf/pkg/services"
)

// ContrailPlugins returns plugins specific to Contrail.
func ContrailPlugins(
	sv Service,
	itd services.InTransactionDoer,
	mg services.MetadataGetter,
) []apiserver.APIPlugin {
	return []apiserver.APIPlugin{
		&PropCollectionUpdatePlugin{
			Service:           sv,
			InTransactionDoer: itd,
			MetadataGetter:    mg,
		},
		&SetTagPlugin{
			Service:           sv,
			InTransactionDoer: itd,
			MetadataGetter:    mg,
		},
		&ChownPlugin{
			Service:           sv,
			InTransactionDoer: itd,
			MetadataGetter:    mg,
		},
	}
}
