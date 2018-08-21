package logic

import (
	"context"
	"sync"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"

	"github.com/Juniper/contrail/pkg/compilation/plugins/contrail/dependencies"
	"github.com/Juniper/contrail/pkg/compilationif"
	"github.com/Juniper/contrail/pkg/services"
)

// EvaluateDependencies evaluates the dependencies upon object change
func EvaluateDependencies(ctx context.Context, evaluateCtx *EvaluateContext,
	obj services.Resource, resourceName string) error {

	log.Printf("EvaluateDependencies called for (%s): \n", resourceName)
	d := dependencies.NewDependencyProcessor(compilationif.ObjsCache)
	d.Evaluate(obj, resourceName, "Self")
	objMap := d.GetResources()

	var err error

	objMap.Range(func(k1, v1 interface{}) bool {
		objTypeKey := k1.(string)
		objList := v1.(*sync.Map)
		log.Printf("Processing ObjType[%s] \n", objTypeKey)
		objList.Range(func(k2, v2 interface{}) bool {
			objUUID := k2.(string)
			objVal := v2
			log.Printf("Processing ObjUUID[%s] \n", objUUID)
			log.Printf("Processing Object[%v] \n", objVal)

			// Look up Intent object by objUUID
			objs, ok := compilationif.ObjsCache.Load(objTypeKey + "Intent")
			if !ok {
				return true
			}
			log.Printf("ObjMap [%v] \n", objs)

			intentObj, ok := objs.(*sync.Map).Load(objUUID)
			if !ok {
				return true
			}
			log.Printf("Intent Object[%v] \n", intentObj)

			rObj, ok := intentObj.(Intent)
			if !ok {
				err = errors.Errorf("%v is not an intent", rObj)
				return false
			}
			err = rObj.Evaluate(ctx, evaluateCtx)
			return err == nil
		})
		return err == nil
	})
	return err
}
