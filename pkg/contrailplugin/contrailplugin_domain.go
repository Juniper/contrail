// nolint
package contrailplugin

import (
	"context"
	"errors"
	"sync"

	"github.com/Juniper/contrail/pkg/compilationif"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
	log "github.com/sirupsen/logrus"
)

// DomainIntent
//   A struct to store attributes related to Domain
//   needed by Intent Compiler
type DomainIntent struct {
	Uuid string
}

// EvaluateDomain - evaluates the Domain
func EvaluateDomain(obj interface{}) {
	resourceObj := obj.(DomainIntent)
	log.Println("EvaluateDomain Called ", resourceObj)
}

// CreateDomain handles create request
func (service *PluginService) CreateDomain(ctx context.Context, request *services.CreateDomainRequest) (*services.CreateDomainResponse, error) {
	log.Println(" CreateDomain Entered")

	obj := request.GetDomain()

	intentObj := DomainIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("DomainIntent"); !ok {
		compilationif.ObjsCache.Store("DomainIntent", &sync.Map{})
	}
	objMap, ok := compilationif.ObjsCache.Load("DomainIntent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("CreateDomain", objMap.(*sync.Map))

	EvaluateDependencies(obj, "Domain")

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().CreateDomain(ctx, request)
}

// UpdateDomain handles update request
func (service *PluginService) UpdateDomain(ctx context.Context, request *services.UpdateDomainRequest) (*services.UpdateDomainResponse, error) {
	log.Println(" UpdateDomain ENTERED")

	obj := request.GetDomain()

	intentObj := DomainIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("DomainIntent"); !ok {
		compilationif.ObjsCache.Store("DomainIntent", &sync.Map{})
	}

	EvaluateDependencies(obj, "Domain")

	objMap, ok := compilationif.ObjsCache.Load("Intent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("UpdateDomain", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().UpdateDomain(ctx, request)
}

// DeleteDomain handles delete request
func (service *PluginService) DeleteDomain(ctx context.Context, request *services.DeleteDomainRequest) (*services.DeleteDomainResponse, error) {
	log.Println(" DeleteDomain ENTERED")

	objUUID := request.GetID()

	//intentObj := DomainIntent {
	//Domain: *obj,
	//}

	//EvaluateDependencies(intentObj, "Domain")

	objMap, ok := compilationif.ObjsCache.Load("DomainIntent")
	if ok {
		objMap.(*sync.Map).Delete(objUUID)
	}

	service.Debug("DeleteDomain", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().DeleteDomain(ctx, request)
}

// GetDomain handles get request
func (service *PluginService) GetDomain(ctx context.Context, request *services.GetDomainRequest) (*services.GetDomainResponse, error) {
	objMap, ok := compilationif.ObjsCache.Load("Domain")
	if !ok {
		return nil, errors.New("Domain get failed ")
	}

	obj, ok := objMap.(*sync.Map).Load(request.GetID())
	if !ok {
		return nil, errors.New("Domain get failed ")
	}

	response := &services.GetDomainResponse{
		Domain: obj.(*models.Domain),
	}
	return response, nil
}
