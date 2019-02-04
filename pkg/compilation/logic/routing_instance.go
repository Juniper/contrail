package logic

import (
	"context"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	"github.com/Juniper/contrail/pkg/compilation/intent"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/models/basemodels"
	"github.com/Juniper/contrail/pkg/services"
)

// RoutingInstanceIntent contains Intent Compiler state for RoutingInstance.
type RoutingInstanceIntent struct {
	intent.BaseIntent
	*models.RoutingInstance
}

// GetObject returns embedded resource object
func (i *RoutingInstanceIntent) GetObject() basemodels.Object {
	return i.RoutingInstance
}

// LoadRoutingInstanceIntent loads a routing instance intent from cache.
func LoadRoutingInstanceIntent(loader intent.Loader, query intent.Query) *RoutingInstanceIntent {
	intent := loader.Load(models.KindRoutingInstance, query)
	riIntent, ok := intent.(*RoutingInstanceIntent)
	if ok == false {
		logrus.Warning("Cannot cast intent to Routing Instance Intent")
	}
	return riIntent
}

// CreateRoutingInstance evaluates RoutingInstance dependencies.
func (s *Service) CreateRoutingInstance(
	ctx context.Context, request *services.CreateRoutingInstanceRequest,
) (*services.CreateRoutingInstanceResponse, error) {
	i := &RoutingInstanceIntent{
		RoutingInstance: request.GetRoutingInstance(),
	}

	if i.GetRoutingInstanceIsDefault() {
		if i.IsIPFabric() || i.IsLinkLocal() {
			return nil, nil
		}
		if err := i.createDefaultRouteTarget(ctx, s.evaluateContext()); err != nil {
			return nil, err
		}
	} else {
		// TODO: handle the situation in case if it's not default Routing Instance
		// and creating non default route targets.
	}

	err := s.storeAndEvaluate(ctx, i)
	if err != nil {
		return nil, err
	}

	return s.BaseService.CreateRoutingInstance(ctx, request)
}

// UpdateRoutingInstance evaluates routing instance dependencies.
func (s *Service) UpdateRoutingInstance(
	ctx context.Context,
	request *services.UpdateRoutingInstanceRequest,
) (*services.UpdateRoutingInstanceResponse, error) {
	ri := request.GetRoutingInstance()
	if ri == nil {
		return nil, errors.New("failed to update Routing Instance." +
			" Routing Instance Request needs to contain resource!")
	}

	i := LoadRoutingInstanceIntent(s.cache, intent.ByUUID(ri.GetUUID()))
	if i == nil {
		return nil, errors.Errorf("cannot load intent for routing instance %v", ri.GetUUID())
	}

	i.RoutingInstance = ri
	if err := s.storeAndEvaluate(ctx, i); err != nil {
		return nil, errors.Wrapf(err, "failed to update intent for Routing Instance :%v", ri.GetUUID())
	}
	return s.BaseService.UpdateRoutingInstance(ctx, request)
}

// Evaluate may create default Route Target.
func (i *RoutingInstanceIntent) Evaluate(
	ctx context.Context, evaluateContext *intent.EvaluateContext,
) error {

	return nil
}

func (i *RoutingInstanceIntent) createDefaultRouteTarget(
	ctx context.Context, evaluateContext *intent.EvaluateContext,
) error {
	rt, err := createDefaultRouteTarget(ctx, evaluateContext)
	if err != nil {
		return err
	}

	_, err = evaluateContext.WriteService.CreateRoutingInstanceRouteTargetRef(
		ctx,
		&services.CreateRoutingInstanceRouteTargetRefRequest{
			ID: i.GetUUID(),
			RoutingInstanceRouteTargetRef: &models.RoutingInstanceRouteTargetRef{
				UUID: rt.GetUUID(),
			},
		},
	)

	return err
}
