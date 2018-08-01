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

// CredentialIntent
//   A struct to store attributes related to Credential
//   needed by Intent Compiler
type CredentialIntent struct {
	Uuid string
}

// EvaluateCredential - evaluates the Credential
func EvaluateCredential(obj interface{}) {
	resourceObj := obj.(CredentialIntent)
	log.Println("EvaluateCredential Called ", resourceObj)
}

// CreateCredential handles create request
func (service *PluginService) CreateCredential(ctx context.Context, request *services.CreateCredentialRequest) (*services.CreateCredentialResponse, error) {
	log.Println(" CreateCredential Entered")

	obj := request.GetCredential()

	intentObj := CredentialIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("CredentialIntent"); !ok {
		compilationif.ObjsCache.Store("CredentialIntent", &sync.Map{})
	}
	objMap, ok := compilationif.ObjsCache.Load("CredentialIntent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("CreateCredential", objMap.(*sync.Map))

	EvaluateDependencies(obj, "Credential")

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().CreateCredential(ctx, request)
}

// UpdateCredential handles update request
func (service *PluginService) UpdateCredential(ctx context.Context, request *services.UpdateCredentialRequest) (*services.UpdateCredentialResponse, error) {
	log.Println(" UpdateCredential ENTERED")

	obj := request.GetCredential()

	intentObj := CredentialIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("CredentialIntent"); !ok {
		compilationif.ObjsCache.Store("CredentialIntent", &sync.Map{})
	}

	EvaluateDependencies(obj, "Credential")

	objMap, ok := compilationif.ObjsCache.Load("Intent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("UpdateCredential", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().UpdateCredential(ctx, request)
}

// DeleteCredential handles delete request
func (service *PluginService) DeleteCredential(ctx context.Context, request *services.DeleteCredentialRequest) (*services.DeleteCredentialResponse, error) {
	log.Println(" DeleteCredential ENTERED")

	objUUID := request.GetID()

	//intentObj := CredentialIntent {
	//Credential: *obj,
	//}

	//EvaluateDependencies(intentObj, "Credential")

	objMap, ok := compilationif.ObjsCache.Load("CredentialIntent")
	if ok {
		objMap.(*sync.Map).Delete(objUUID)
	}

	service.Debug("DeleteCredential", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().DeleteCredential(ctx, request)
}

// GetCredential handles get request
func (service *PluginService) GetCredential(ctx context.Context, request *services.GetCredentialRequest) (*services.GetCredentialResponse, error) {
	objMap, ok := compilationif.ObjsCache.Load("Credential")
	if !ok {
		return nil, errors.New("Credential get failed ")
	}

	obj, ok := objMap.(*sync.Map).Load(request.GetID())
	if !ok {
		return nil, errors.New("Credential get failed ")
	}

	response := &services.GetCredentialResponse{
		Credential: obj.(*models.Credential),
	}
	return response, nil
}
