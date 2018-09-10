package logic

import (
	"context"
	"math/rand"
	"time"

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

// LoadRoutingInstanceIntent loads RoutingInstanceIntent from cache
func LoadRoutingInstanceIntent(
	c intent.Loader,
	uuid string,
) *RoutingInstanceIntent {
	i := c.Load(models.KindRoutingInstance, intent.ByUUID(uuid))
	actual, _ := i.(*RoutingInstanceIntent)
	return actual
}

// TODO: get_autonomous_system method and int pool allocator endpoint.
const (
	defaultAutonomousSystem = 64512
	// This number should be generated from int pool allocator.
	minimumRoutingTargetNumber = 8000002
)

// CreateRoutingInstance evaluates RoutingInstance dependencies.
func (s *Service) CreateRoutingInstance(
	ctx context.Context, request *services.CreateRoutingInstanceRequest,
) (*services.CreateRoutingInstanceResponse, error) {
	i := &RoutingInstanceIntent{
		RoutingInstance: request.GetRoutingInstance(),
	}

	err := s.handleCreate(ctx, i, nil, i.RoutingInstance)
	if err != nil {
		return nil, err
	}

	return s.BaseService.CreateRoutingInstance(ctx, request)
}

// Evaluate may create default Route Target.
func (i *RoutingInstanceIntent) Evaluate(
	ctx context.Context,
	evaluateContext *intent.EvaluateContext,
) error {
	if i.GetRoutingInstanceIsDefault() {
		if i.IsIPFabric() || i.IsLinkLocal() {
			return nil
		}
		if err := i.createDefaultRouteTarget(ctx, evaluateContext); err != nil {
			return err
		}
	} else {
		// TODO: handle the situation in case if it's not default Routing Instance
		// and creating non default route targets.
	}

	return nil
}

// TODO Temporary way to generate route target number
// until allocate route target is implemented.
func generateRandomRouteTargetNumber() int {
	rand.Seed(time.Now().UTC().UnixNano())
	return minimumRoutingTargetNumber + rand.Intn(10)
}

func (i *RoutingInstanceIntent) createDefaultRouteTarget(
	ctx context.Context,
	evaluateContext *intent.EvaluateContext,
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
