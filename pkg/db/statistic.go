package db

import (
	"github.com/Juniper/contrail/pkg/collector"
	"github.com/Juniper/contrail/pkg/services"
)

// StatisticService represents DBRequestTrace logger
type StatisticService struct {
	services.Service
	collector *collector.Collector
}

// NewStatisticService wrap Service into DBRequestTrace logger
func NewStatisticService(db services.Service, collector *collector.Collector) services.Service {
	return &StatisticService{
		Service:   db,
		collector: collector,
	}
}
