package types

import (
	"database/sql"

	"github.com/Juniper/contrail/pkg/serviceif"
	"github.com/Juniper/contrail/pkg/types/ipam"
)

//ContrailTypeLogicService is a service for implementing type specific logic
type ContrailTypeLogicService struct {
	serviceif.BaseService
	DB                   *sql.DB
	VirtualNetworkIDPool ipam.IntPool
}
