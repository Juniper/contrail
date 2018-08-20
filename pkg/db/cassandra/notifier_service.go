package cassandra

import (
	"github.com/Juniper/contrail/pkg/services"
)

type NotifierService struct {
	services.BaseService
	config Config
}

func NewNotifierService() *NotifierService {
	return &NotifierService{
		config: GetConfig(),
	}
}
