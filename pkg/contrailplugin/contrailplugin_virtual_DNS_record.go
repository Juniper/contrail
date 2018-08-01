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

// VirtualDNSRecordIntent
//   A struct to store attributes related to VirtualDNSRecord
//   needed by Intent Compiler
type VirtualDNSRecordIntent struct {
	Uuid string
}

// EvaluateVirtualDNSRecord - evaluates the VirtualDNSRecord
func EvaluateVirtualDNSRecord(obj interface{}) {
	resourceObj := obj.(VirtualDNSRecordIntent)
	log.Println("EvaluateVirtualDNSRecord Called ", resourceObj)
}

// CreateVirtualDNSRecord handles create request
func (service *PluginService) CreateVirtualDNSRecord(ctx context.Context, request *services.CreateVirtualDNSRecordRequest) (*services.CreateVirtualDNSRecordResponse, error) {
	log.Println(" CreateVirtualDNSRecord Entered")

	obj := request.GetVirtualDNSRecord()

	intentObj := VirtualDNSRecordIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("VirtualDNSRecordIntent"); !ok {
		compilationif.ObjsCache.Store("VirtualDNSRecordIntent", &sync.Map{})
	}
	objMap, ok := compilationif.ObjsCache.Load("VirtualDNSRecordIntent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("CreateVirtualDNSRecord", objMap.(*sync.Map))

	EvaluateDependencies(obj, "VirtualDNSRecord")

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().CreateVirtualDNSRecord(ctx, request)
}

// UpdateVirtualDNSRecord handles update request
func (service *PluginService) UpdateVirtualDNSRecord(ctx context.Context, request *services.UpdateVirtualDNSRecordRequest) (*services.UpdateVirtualDNSRecordResponse, error) {
	log.Println(" UpdateVirtualDNSRecord ENTERED")

	obj := request.GetVirtualDNSRecord()

	intentObj := VirtualDNSRecordIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("VirtualDNSRecordIntent"); !ok {
		compilationif.ObjsCache.Store("VirtualDNSRecordIntent", &sync.Map{})
	}

	EvaluateDependencies(obj, "VirtualDNSRecord")

	objMap, ok := compilationif.ObjsCache.Load("Intent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("UpdateVirtualDNSRecord", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().UpdateVirtualDNSRecord(ctx, request)
}

// DeleteVirtualDNSRecord handles delete request
func (service *PluginService) DeleteVirtualDNSRecord(ctx context.Context, request *services.DeleteVirtualDNSRecordRequest) (*services.DeleteVirtualDNSRecordResponse, error) {
	log.Println(" DeleteVirtualDNSRecord ENTERED")

	objUUID := request.GetID()

	//intentObj := VirtualDNSRecordIntent {
	//VirtualDNSRecord: *obj,
	//}

	//EvaluateDependencies(intentObj, "VirtualDNSRecord")

	objMap, ok := compilationif.ObjsCache.Load("VirtualDNSRecordIntent")
	if ok {
		objMap.(*sync.Map).Delete(objUUID)
	}

	service.Debug("DeleteVirtualDNSRecord", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().DeleteVirtualDNSRecord(ctx, request)
}

// GetVirtualDNSRecord handles get request
func (service *PluginService) GetVirtualDNSRecord(ctx context.Context, request *services.GetVirtualDNSRecordRequest) (*services.GetVirtualDNSRecordResponse, error) {
	objMap, ok := compilationif.ObjsCache.Load("VirtualDNSRecord")
	if !ok {
		return nil, errors.New("VirtualDNSRecord get failed ")
	}

	obj, ok := objMap.(*sync.Map).Load(request.GetID())
	if !ok {
		return nil, errors.New("VirtualDNSRecord get failed ")
	}

	response := &services.GetVirtualDNSRecordResponse{
		VirtualDNSRecord: obj.(*models.VirtualDNSRecord),
	}
	return response, nil
}
