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

// StructuredSyslogApplicationRecordIntent
//   A struct to store attributes related to StructuredSyslogApplicationRecord
//   needed by Intent Compiler
type StructuredSyslogApplicationRecordIntent struct {
	Uuid string
}

// EvaluateStructuredSyslogApplicationRecord - evaluates the StructuredSyslogApplicationRecord
func EvaluateStructuredSyslogApplicationRecord(obj interface{}) {
	resourceObj := obj.(StructuredSyslogApplicationRecordIntent)
	log.Println("EvaluateStructuredSyslogApplicationRecord Called ", resourceObj)
}

// CreateStructuredSyslogApplicationRecord handles create request
func (service *PluginService) CreateStructuredSyslogApplicationRecord(ctx context.Context, request *services.CreateStructuredSyslogApplicationRecordRequest) (*services.CreateStructuredSyslogApplicationRecordResponse, error) {
	log.Println(" CreateStructuredSyslogApplicationRecord Entered")

	obj := request.GetStructuredSyslogApplicationRecord()

	intentObj := StructuredSyslogApplicationRecordIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("StructuredSyslogApplicationRecordIntent"); !ok {
		compilationif.ObjsCache.Store("StructuredSyslogApplicationRecordIntent", &sync.Map{})
	}
	objMap, ok := compilationif.ObjsCache.Load("StructuredSyslogApplicationRecordIntent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("CreateStructuredSyslogApplicationRecord", objMap.(*sync.Map))

	EvaluateDependencies(obj, "StructuredSyslogApplicationRecord")

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().CreateStructuredSyslogApplicationRecord(ctx, request)
}

// UpdateStructuredSyslogApplicationRecord handles update request
func (service *PluginService) UpdateStructuredSyslogApplicationRecord(ctx context.Context, request *services.UpdateStructuredSyslogApplicationRecordRequest) (*services.UpdateStructuredSyslogApplicationRecordResponse, error) {
	log.Println(" UpdateStructuredSyslogApplicationRecord ENTERED")

	obj := request.GetStructuredSyslogApplicationRecord()

	intentObj := StructuredSyslogApplicationRecordIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("StructuredSyslogApplicationRecordIntent"); !ok {
		compilationif.ObjsCache.Store("StructuredSyslogApplicationRecordIntent", &sync.Map{})
	}

	EvaluateDependencies(obj, "StructuredSyslogApplicationRecord")

	objMap, ok := compilationif.ObjsCache.Load("Intent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("UpdateStructuredSyslogApplicationRecord", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().UpdateStructuredSyslogApplicationRecord(ctx, request)
}

// DeleteStructuredSyslogApplicationRecord handles delete request
func (service *PluginService) DeleteStructuredSyslogApplicationRecord(ctx context.Context, request *services.DeleteStructuredSyslogApplicationRecordRequest) (*services.DeleteStructuredSyslogApplicationRecordResponse, error) {
	log.Println(" DeleteStructuredSyslogApplicationRecord ENTERED")

	objUUID := request.GetID()

	//intentObj := StructuredSyslogApplicationRecordIntent {
	//StructuredSyslogApplicationRecord: *obj,
	//}

	//EvaluateDependencies(intentObj, "StructuredSyslogApplicationRecord")

	objMap, ok := compilationif.ObjsCache.Load("StructuredSyslogApplicationRecordIntent")
	if ok {
		objMap.(*sync.Map).Delete(objUUID)
	}

	service.Debug("DeleteStructuredSyslogApplicationRecord", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().DeleteStructuredSyslogApplicationRecord(ctx, request)
}

// GetStructuredSyslogApplicationRecord handles get request
func (service *PluginService) GetStructuredSyslogApplicationRecord(ctx context.Context, request *services.GetStructuredSyslogApplicationRecordRequest) (*services.GetStructuredSyslogApplicationRecordResponse, error) {
	objMap, ok := compilationif.ObjsCache.Load("StructuredSyslogApplicationRecord")
	if !ok {
		return nil, errors.New("StructuredSyslogApplicationRecord get failed ")
	}

	obj, ok := objMap.(*sync.Map).Load(request.GetID())
	if !ok {
		return nil, errors.New("StructuredSyslogApplicationRecord get failed ")
	}

	response := &services.GetStructuredSyslogApplicationRecordResponse{
		StructuredSyslogApplicationRecord: obj.(*models.StructuredSyslogApplicationRecord),
	}
	return response, nil
}
