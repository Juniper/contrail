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

// TagIntent
//   A struct to store attributes related to Tag
//   needed by Intent Compiler
type TagIntent struct {
	Uuid string
}

// EvaluateTag - evaluates the Tag
func EvaluateTag(obj interface{}) {
	resourceObj := obj.(TagIntent)
	log.Println("EvaluateTag Called ", resourceObj)
}

// CreateTag handles create request
func (service *PluginService) CreateTag(ctx context.Context, request *services.CreateTagRequest) (*services.CreateTagResponse, error) {
	log.Println(" CreateTag Entered")

	obj := request.GetTag()

	intentObj := TagIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("TagIntent"); !ok {
		compilationif.ObjsCache.Store("TagIntent", &sync.Map{})
	}
	objMap, ok := compilationif.ObjsCache.Load("TagIntent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("CreateTag", objMap.(*sync.Map))

	EvaluateDependencies(obj, "Tag")

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().CreateTag(ctx, request)
}

// UpdateTag handles update request
func (service *PluginService) UpdateTag(ctx context.Context, request *services.UpdateTagRequest) (*services.UpdateTagResponse, error) {
	log.Println(" UpdateTag ENTERED")

	obj := request.GetTag()

	intentObj := TagIntent{
		Uuid: obj.GetUUID(),
	}

	if _, ok := compilationif.ObjsCache.Load("TagIntent"); !ok {
		compilationif.ObjsCache.Store("TagIntent", &sync.Map{})
	}

	EvaluateDependencies(obj, "Tag")

	objMap, ok := compilationif.ObjsCache.Load("Intent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intentObj)
	}

	service.Debug("UpdateTag", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().UpdateTag(ctx, request)
}

// DeleteTag handles delete request
func (service *PluginService) DeleteTag(ctx context.Context, request *services.DeleteTagRequest) (*services.DeleteTagResponse, error) {
	log.Println(" DeleteTag ENTERED")

	objUUID := request.GetID()

	//intentObj := TagIntent {
	//Tag: *obj,
	//}

	//EvaluateDependencies(intentObj, "Tag")

	objMap, ok := compilationif.ObjsCache.Load("TagIntent")
	if ok {
		objMap.(*sync.Map).Delete(objUUID)
	}

	service.Debug("DeleteTag", objMap.(*sync.Map))

	if service.Next() == nil {
		return nil, nil
	}
	return service.Next().DeleteTag(ctx, request)
}

// GetTag handles get request
func (service *PluginService) GetTag(ctx context.Context, request *services.GetTagRequest) (*services.GetTagResponse, error) {
	objMap, ok := compilationif.ObjsCache.Load("Tag")
	if !ok {
		return nil, errors.New("Tag get failed ")
	}

	obj, ok := objMap.(*sync.Map).Load(request.GetID())
	if !ok {
		return nil, errors.New("Tag get failed ")
	}

	response := &services.GetTagResponse{
		Tag: obj.(*models.Tag),
	}
	return response, nil
}
