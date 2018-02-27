package models

import (
	basemodels "github.com/Juniper/contrail/pkg/generated/models"
)

//Service for custom logics
//Overwrite functions in baseservice to add custom logics.
//Don't forget call next service. see pkg/generated/services.ContrailService
type Service struct {
	basemodels.BaseService
}

//NewService makes new service for custom logics.
func NewService() basemodels.Service {
	return &Service{
		BaseService: basemodels.BaseService{},
	}
}
