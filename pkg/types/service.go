package types

import (
	"github.com/Juniper/contrail/pkg/serviceif"
	"github.com/Juniper/contrail/pkg/types/ipam"
)

const (
	//VirtualNetworkIDPoolKey is a key for id pool for virutal network id.
	VirtualNetworkIDPoolKey = "virtual_network_id"
)

//ContrailTypeLogicService is a service for implementing type specific logic
type ContrailTypeLogicService struct {
	serviceif.BaseService
	DB   *db.Service
}
