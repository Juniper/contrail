package logic

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/pkg/errors"

	"github.com/Juniper/contrail/pkg/compilationif"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
)

// RoutingInstanceIntent contains Intent Compiler state for RoutingInstance
type RoutingInstanceIntent struct {
	BaseIntent
	*models.RoutingInstance
}

// TODO: get_autonomous_system method and int pool allocator endpoint
const (
	defaultAutonomousSystem = 64512
	// This number should be generated from int pool allocator.
	minimumRoutingTargetNumber = 8000002
)

// CreateRoutingInstance evaluates RoutingInstance dependencies
func (s *Service) CreateRoutingInstance(
	ctx context.Context, request *services.CreateRoutingInstanceRequest,
) (*services.CreateRoutingInstanceResponse, error) {

	obj := request.GetRoutingInstance()

	intent := &RoutingInstanceIntent{
		RoutingInstance: obj,
	}

	if _, ok := compilationif.ObjsCache.Load("RoutingInstanceIntent"); !ok {
		compilationif.ObjsCache.Store("RoutingInstanceIntent", &sync.Map{})
	}

	objMap, ok := compilationif.ObjsCache.Load("RoutingInstanceIntent")
	if ok {
		objMap.(*sync.Map).Store(obj.GetUUID(), intent)
	}

	ec := &EvaluateContext{
		WriteService: s.WriteService,
	}
	err := EvaluateDependencies(ctx, ec, obj, "RoutingInstance")
	if err != nil {
		return nil, errors.Wrap(err, "failed to evaluate Routing Instance dependencies")
	}

	return s.BaseService.CreateRoutingInstance(ctx, request)
}

// Evaluate may create default Route Target.
func (s *RoutingInstanceIntent) Evaluate(ctx context.Context, evaluateContext *EvaluateContext) error {
	if s.GetRoutingInstanceIsDefault() {
		if s.IsIPFabric() || s.IsLinkLocal() {
			return nil
		}
		if err := s.createDefaultRouteTarget(ctx, evaluateContext); err != nil {
			return err
		}
	} else {
		// TODO: handle the situation in case if it's not default Routing Instance
		// and creating non default route targets
	}

	return nil
}

// TODO Temporary way to generate route target number
// until allocate route target is implemented
func generateRandomRouteTargetNumber() int {
	rand.Seed(time.Now().UTC().UnixNano())
	return minimumRoutingTargetNumber + rand.Intn(10)
}

func (s *RoutingInstanceIntent) createDefaultRouteTarget(ctx context.Context, evaluateContext *EvaluateContext) error {
	rtKey := fmt.Sprintf("target:%v:%v", defaultAutonomousSystem, generateRandomRouteTargetNumber())

	rtResponse, err := evaluateContext.WriteService.CreateRouteTarget(
		ctx,
		&services.CreateRouteTargetRequest{
			RouteTarget: &models.RouteTarget{
				FQName:      []string{rtKey},
				DisplayName: rtKey,
			},
		},
	)

	if err != nil {
		return err
	}

	_, err = evaluateContext.WriteService.CreateRoutingInstanceRouteTargetRef(
		ctx,
		&services.CreateRoutingInstanceRouteTargetRefRequest{
			ID: s.GetUUID(),
			RoutingInstanceRouteTargetRef: &models.RoutingInstanceRouteTargetRef{
				UUID: rtResponse.RouteTarget.GetUUID(),
			},
		},
	)

	return err
}
