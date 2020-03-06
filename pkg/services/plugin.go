package services

import (
	"github.com/Juniper/asf/pkg/apiserver"
	"github.com/Juniper/asf/pkg/services/baseservices"
)

// ContrailPlugins returns plugins specific to Contrail.
func ContrailPlugins(
	sv Service,
	itd InTransactionDoer,
	mg baseservices.MetadataGetter,
	idToFQName IDToFQNameService,
) []apiserver.APIPlugin {
	return []apiserver.APIPlugin{
		&PropCollectionUpdatePlugin{
			Service:           sv,
			InTransactionDoer: itd,
			IDToFQNameService: idToFQName,
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
