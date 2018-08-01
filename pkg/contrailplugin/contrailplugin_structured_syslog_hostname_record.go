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

// StructuredSyslogHostnameRecordIntent
//   A struct to store attributes related to StructuredSyslogHostnameRecord
//   needed by Intent Compiler
type StructuredSyslogHostnameRecordIntent struct {
	Uuid string
}

// EvaluateStructuredSyslogHostnameRecord - evaluates the StructuredSyslogHostnameRecord
func EvaluateStructuredSyslogHostnameRecord(obj interface{}) {
	resourceObj := obj.(StructuredSyslogHostnameRecordIntent)
	log.Println("EvaluateStructuredSyslogHostnameRecord Called ", resourceObj)
}

// CreateStructuredSyslogHostnameRecord handles create request
func (service *PluginService) CreateStructuredSyslogHostnameRecord(ctx context.Context, request *services.CreateStructuredSyslogHostnameRecordRequest) (*services.CreateStructuredSyslogHostnameRecordResponse, error) {
	log.Println(" CreateStructuredSyslogHostnameRecord Entered")

	obj := request.GetStructuredSyslogHostnameRecord()

	intentObj := StructuredSyslogHostnameRecordIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("StructuredSyslogHostnameRecordIntent"); !ok {
		compilationif.ObjsCache.Store("StructuredSyslogHostnameRecordIntent", &sync.Map{})
	}
	objMap, ok := compilationif.ObjsCache.Load("StructuredSyslogHostnameRecordIntent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("CreateStructuredSyslogHostnameRecord", objMap.(*sync.Map))

	EvaluateDependencies(obj, "StructuredSyslogHostnameRecord")

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().CreateStructuredSyslogHostnameRecord(ctx, request)
}

// UpdateStructuredSyslogHostnameRecord handles update request
func (service *PluginService) UpdateStructuredSyslogHostnameRecord(ctx context.Context, request *services.UpdateStructuredSyslogHostnameRecordRequest) (*services.UpdateStructuredSyslogHostnameRecordResponse, error) {
	log.Println(" UpdateStructuredSyslogHostnameRecord ENTERED")

	obj := request.GetStructuredSyslogHostnameRecord()

	intentObj := StructuredSyslogHostnameRecordIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("StructuredSyslogHostnameRecordIntent"); !ok {
		compilationif.ObjsCache.Store("StructuredSyslogHostnameRecordIntent", &sync.Map{})
	}

	EvaluateDependencies(obj, "StructuredSyslogHostnameRecord")

	objMap, ok := compilationif.ObjsCache.Load("Intent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("UpdateStructuredSyslogHostnameRecord", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().UpdateStructuredSyslogHostnameRecord(ctx, request)
}

// DeleteStructuredSyslogHostnameRecord handles delete request
func (service *PluginService) DeleteStructuredSyslogHostnameRecord(ctx context.Context, request *services.DeleteStructuredSyslogHostnameRecordRequest) (*services.DeleteStructuredSyslogHostnameRecordResponse, error) {
	log.Println(" DeleteStructuredSyslogHostnameRecord ENTERED")

	objUUID := request.GetID()

	//intentObj := StructuredSyslogHostnameRecordIntent {
	//StructuredSyslogHostnameRecord: *obj,
	//}

	//EvaluateDependencies(intentObj, "StructuredSyslogHostnameRecord")

	objMap, ok := compilationif.ObjsCache.Load("StructuredSyslogHostnameRecordIntent")
	if ok {
		objMap.(*sync.Map).Delete(objUUID)
	}

	service.Debug("DeleteStructuredSyslogHostnameRecord", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().DeleteStructuredSyslogHostnameRecord(ctx, request)
}

// GetStructuredSyslogHostnameRecord handles get request
func (service *PluginService) GetStructuredSyslogHostnameRecord(ctx context.Context, request *services.GetStructuredSyslogHostnameRecordRequest) (*services.GetStructuredSyslogHostnameRecordResponse, error) {
	objMap, ok := compilationif.ObjsCache.Load("StructuredSyslogHostnameRecord")
	if !ok {
		return nil, errors.New("StructuredSyslogHostnameRecord get failed ")
	}

	obj, ok := objMap.(*sync.Map).Load(request.GetID())
	if !ok {
		return nil, errors.New("StructuredSyslogHostnameRecord get failed ")
	}

	response := &services.GetStructuredSyslogHostnameRecordResponse{
		StructuredSyslogHostnameRecord: obj.(*models.StructuredSyslogHostnameRecord),
	}
	return response, nil
}
