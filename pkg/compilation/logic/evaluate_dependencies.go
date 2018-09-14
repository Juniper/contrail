package logic

import (
	"context"
	"sync"

	log "github.com/sirupsen/logrus"

	"github.com/Juniper/contrail/pkg/compilation/intent"
	"github.com/Juniper/contrail/pkg/compilation/plugins/contrail/dependencies"
	"github.com/Juniper/contrail/pkg/services"
)

// EvaluateDependencies evaluates the dependencies upon object change
func (s *Service) EvaluateDependencies(
	ctx context.Context,
	evaluateCtx *intent.EvaluateContext,
	obj services.Resource,
) error {

	log.Printf("EvaluateDependencies called for (%s): \n", obj.TypeName())
	d := dependencies.NewDependencyProcessor(s.cache)
	d.Evaluate(obj, obj.TypeName(), "Self")
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

			intent := s.cache.Load(objTypeKey, intent.ByUUID(objUUID))
			if intent == nil {
				return false
			}
			err = intent.Evaluate(ctx, evaluateCtx)
			return err == nil
		})
		return err == nil
	})
	return err
}
