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

// TagTypeIntent
//   A struct to store attributes related to TagType
//   needed by Intent Compiler
type TagTypeIntent struct {
	Uuid string
}

// EvaluateTagType - evaluates the TagType
func EvaluateTagType(obj interface{}) {
	resourceObj := obj.(TagTypeIntent)
	log.Println("EvaluateTagType Called ", resourceObj)
}

// CreateTagType handles create request
func (service *PluginService) CreateTagType(ctx context.Context, request *services.CreateTagTypeRequest) (*services.CreateTagTypeResponse, error) {
	log.Println(" CreateTagType Entered")

	obj := request.GetTagType()

	intentObj := TagTypeIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("TagTypeIntent"); !ok {
		compilationif.ObjsCache.Store("TagTypeIntent", &sync.Map{})
	}
	objMap, ok := compilationif.ObjsCache.Load("TagTypeIntent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("CreateTagType", objMap.(*sync.Map))

	EvaluateDependencies(obj, "TagType")

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().CreateTagType(ctx, request)
}

// UpdateTagType handles update request
func (service *PluginService) UpdateTagType(ctx context.Context, request *services.UpdateTagTypeRequest) (*services.UpdateTagTypeResponse, error) {
	log.Println(" UpdateTagType ENTERED")

	obj := request.GetTagType()

	intentObj := TagTypeIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("TagTypeIntent"); !ok {
		compilationif.ObjsCache.Store("TagTypeIntent", &sync.Map{})
	}

	EvaluateDependencies(obj, "TagType")

	objMap, ok := compilationif.ObjsCache.Load("Intent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("UpdateTagType", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().UpdateTagType(ctx, request)
}

// DeleteTagType handles delete request
func (service *PluginService) DeleteTagType(ctx context.Context, request *services.DeleteTagTypeRequest) (*services.DeleteTagTypeResponse, error) {
	log.Println(" DeleteTagType ENTERED")

	objUUID := request.GetID()

	//intentObj := TagTypeIntent {
	//TagType: *obj,
	//}

	//EvaluateDependencies(intentObj, "TagType")

	objMap, ok := compilationif.ObjsCache.Load("TagTypeIntent")
	if ok {
		objMap.(*sync.Map).Delete(objUUID)
	}

	service.Debug("DeleteTagType", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().DeleteTagType(ctx, request)
}

// GetTagType handles get request
func (service *PluginService) GetTagType(ctx context.Context, request *services.GetTagTypeRequest) (*services.GetTagTypeResponse, error) {
	objMap, ok := compilationif.ObjsCache.Load("TagType")
	if !ok {
		return nil, errors.New("TagType get failed ")
	}

	obj, ok := objMap.(*sync.Map).Load(request.GetID())
	if !ok {
		return nil, errors.New("TagType get failed ")
	}

	response := &services.GetTagTypeResponse{
		TagType: obj.(*models.TagType),
	}
	return response, nil
}
