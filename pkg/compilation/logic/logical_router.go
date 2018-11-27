package logic

import (
	"context"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"

	"github.com/Juniper/contrail/pkg/compilation/intent"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/models/basemodels"
	"github.com/Juniper/contrail/pkg/services"
	"github.com/Juniper/contrail/pkg/services/baseservices"
)

// LogicalRouterIntent contains Intent Compiler state for LogicalRouter.
type LogicalRouterIntent struct {
	intent.BaseIntent
	*models.LogicalRouter
	virtualNetworks        map[string]*models.VirtualMachineInterfaceVirtualNetworkRef
	defaultRouteTargetUUID string
	vxlanRouting           bool
}

// GetObject returns embedded resource object
func (i *LogicalRouterIntent) GetObject() basemodels.Object {
	return i.LogicalRouter
}

// NewLogicalRouterIntent creates LogicalRouterIntent from CreateLogicalRouterRequest
func NewLogicalRouterIntent(
	ctx context.Context,
	ReadService services.ReadService,
	request *services.CreateLogicalRouterRequest,
) *LogicalRouterIntent {
	lr := &LogicalRouterIntent{
		LogicalRouter:   request.GetLogicalRouter(),
		virtualNetworks: make(map[string]*models.VirtualMachineInterfaceVirtualNetworkRef),
	}
	lr.resolveVxLan(ctx, ReadService)

	return lr
}

// LoadLogicalRouterIntent loads a logical router intent from cache.
func LoadLogicalRouterIntent(loader intent.Loader, query intent.Query) *LogicalRouterIntent {
	intent := loader.Load(models.KindLogicalRouter, query)
	lrIntent, ok := intent.(*LogicalRouterIntent)
	if ok == false {
		log.Warning("Cannot cast intent to Logical Router Intent")
	}
	return lrIntent
}

// CreateLogicalRouter evaluates logical router dependencies.
func (s *Service) CreateLogicalRouter(
	ctx context.Context,
	request *services.CreateLogicalRouterRequest,
) (*services.CreateLogicalRouterResponse, error) {
	i := NewLogicalRouterIntent(ctx, s.ReadService, request)

	ec := s.evaluateContext()

	if len(i.LogicalRouter.GetRouteTargetRefs()) == 0 {
		if rt, err := i.createDefaultRouteTarget(ctx, ec); err == nil {
			i.defaultRouteTargetUUID = rt.GetUUID()
		} else {
			return nil, errors.Wrap(err, "failed to create Logical Router's default Route Target")
		}
	} else {
		i.defaultRouteTargetUUID = i.LogicalRouter.GetRouteTargetRefs()[0].GetUUID()
	}

	if err := s.storeAndEvaluate(ctx, i); err != nil {
		return nil, err
	}

	return s.BaseService.CreateLogicalRouter(ctx, request)
}

// UpdateLogicalRouter evaluates logical router dependencies.
func (s *Service) UpdateLogicalRouter(
	ctx context.Context,
	request *services.UpdateLogicalRouterRequest,
) (*services.UpdateLogicalRouterResponse, error) {
	lr := request.GetLogicalRouter()
	if lr == nil {
		return nil, errors.New("failed to update Logical Router." +
			" Logical Router Request needs to contain resource!")
	}

	i := LoadLogicalRouterIntent(s.cache, intent.ByUUID(lr.GetUUID()))
	if i == nil {
		return nil, errors.Errorf("cannot load intent for logical router %v", lr.GetUUID())
	}

	i.LogicalRouter = lr
	if err := s.storeAndEvaluate(ctx, i); err != nil {
		return nil, errors.Wrapf(err, "failed to update intent for Logical Router :%v", lr.GetUUID())
	}

	return s.BaseService.UpdateLogicalRouter(ctx, request)
}

func (i *LogicalRouterIntent) checkVnDiff(
	vns map[string]*models.VirtualMachineInterfaceVirtualNetworkRef,
) bool {
	for k := range vns {
		_, present := i.virtualNetworks[k]
		if !present {
			return false
		}
	}
	return true
}

// Evaluate updates references from default routing instances of virtual networks
// to the logical router's route target.
func (i *LogicalRouterIntent) Evaluate(
	ctx context.Context,
	evaluateCtx *intent.EvaluateContext,
) error {
	return i.updateVirtualNetworks(ctx, evaluateCtx)
}

func (i *LogicalRouterIntent) resolveVxLan(
	ctx context.Context,
	ReadService services.ReadService,
) {
	r, err := ReadService.GetProject(ctx, &services.GetProjectRequest{
		ID: i.GetParentUUID(),
		Spec: &baseservices.GetSpec{
			Fields: []string{"vxlan_routing"},
		},
	})

	if err != nil {
		log.Errorf("failed to retrieve project for LogicalRouter with uuid %s", i.GetUUID())
	}

	if r.GetProject() != nil {
		i.vxlanRouting = r.GetProject().VxlanRouting
	}
}

func (i *LogicalRouterIntent) updateVirtualNetworks(
	ctx context.Context,
	evaluateCtx *intent.EvaluateContext,
) error {
	vnRefs := make(map[string]*models.VirtualMachineInterfaceVirtualNetworkRef)
	for _, ref := range i.VirtualMachineInterfaceRefs {
		vmi := LoadVirtualMachineInterfaceIntent(evaluateCtx.IntentLoader, intent.ByUUID(ref.UUID))
		if vmi == nil {
			return errors.Errorf("failed to retrieve virtual-machine-interface "+
				"reference with uuid %s for logical-router with uuid %s", ref.UUID, i.GetUUID())
		}
		if len(vmi.GetVirtualNetworkRefs()) > 0 {
			vnRefs[vmi.UUID] = vmi.GetVirtualNetworkRefs()[0]
		}
	}

	if i.checkVnDiff(vnRefs) {
		return nil
	}

	return i.setVirtualNetworks(ctx, evaluateCtx, vnRefs)
}

func (i *LogicalRouterIntent) setVirtualNetworks(
	ctx context.Context,
	evaluateCtx *intent.EvaluateContext,
	vnRefs map[string]*models.VirtualMachineInterfaceVirtualNetworkRef,
) error {
	if i.vxlanRouting {
		i.virtualNetworks = vnRefs
		return nil
	}
	err := i.handleDeletedNetworks(ctx, evaluateCtx, vnRefs)
	if err != nil {
		return err
	}
	return i.handleAddedNetworks(ctx, evaluateCtx, vnRefs)

}

func (i *LogicalRouterIntent) handleDeletedNetworks(
	ctx context.Context,
	evaluateCtx *intent.EvaluateContext,
	vnRefs map[string]*models.VirtualMachineInterfaceVirtualNetworkRef,
) error {
	for _, vn := range i.getDeletedNetworks(ctx, evaluateCtx, vnRefs) {
		ri := vn.GetPrimaryRoutingInstanceIntent(ctx, evaluateCtx)
		if ri == nil {
			log.Errorf("Primary RI is None for VN: %s", vn.GetUUID())
			continue
		}
		// TODO handle delete
	}
	return nil
}

func (i *LogicalRouterIntent) handleAddedNetworks(
	ctx context.Context,
	evaluateCtx *intent.EvaluateContext,
	vnRefs map[string]*models.VirtualMachineInterfaceVirtualNetworkRef,
) error {
	if i.defaultRouteTargetUUID == "" {
		return errors.Errorf("missing default route target of logical router with uuid: %s", i.GetUUID())
	}
	for _, vn := range i.getAddedNetworks(ctx, evaluateCtx, vnRefs) {
		ri := vn.GetPrimaryRoutingInstanceIntent(ctx, evaluateCtx)
		if ri == nil {
			log.Errorf("Primary RI is None for VN: %s", vn.GetUUID())
			continue
		}
		// TODO handle all route targets
		_, err := evaluateCtx.WriteService.CreateRoutingInstanceRouteTargetRef(
			ctx, &services.CreateRoutingInstanceRouteTargetRefRequest{
				ID: ri.GetUUID(),
				RoutingInstanceRouteTargetRef: &models.RoutingInstanceRouteTargetRef{
					UUID: i.defaultRouteTargetUUID,
				},
			},
		)
		if err != nil {
			return err
		}
	}
	return nil
}

func (i *LogicalRouterIntent) getDeletedNetworks(
	_ context.Context,
	evaluateCtx *intent.EvaluateContext,
	vnRefs map[string]*models.VirtualMachineInterfaceVirtualNetworkRef,
) []*VirtualNetworkIntent {
	vns := []*VirtualNetworkIntent{}
	for uuid, vnref := range i.virtualNetworks {
		if _, ok := vnRefs[uuid]; !ok {
			vn := LoadVirtualNetworkIntent(evaluateCtx.IntentLoader, intent.ByUUID(vnref.UUID))
			if vn == nil {
				log.Errorf("Failed to load VN with uuid %s", vnref.UUID)
				continue
			}
			vns = append(vns, vn)
		}
	}
	return vns
}

func (i *LogicalRouterIntent) getAddedNetworks(
	_ context.Context,
	evaluateCtx *intent.EvaluateContext,
	vnRefs map[string]*models.VirtualMachineInterfaceVirtualNetworkRef,
) []*VirtualNetworkIntent {
	vns := []*VirtualNetworkIntent{}
	for uuid, vnRef := range vnRefs {
		if _, ok := i.virtualNetworks[uuid]; !ok {
			vn := LoadVirtualNetworkIntent(evaluateCtx.IntentLoader, intent.ByUUID(vnRef.UUID))
			if vn == nil {
				log.Errorf("failed to retrieve VN with uuid %s", vnRef.UUID)
				continue
			}
			vns = append(vns, vn)
		}
	}
	return vns
}

func (i *LogicalRouterIntent) createDefaultRouteTarget(
	ctx context.Context,
	evaluateContext *intent.EvaluateContext,
) (*models.RouteTarget, error) {
	rt, err := createDefaultRouteTarget(ctx, evaluateContext)
	if err != nil {
		return nil, err
	}

	_, err = evaluateContext.WriteService.CreateLogicalRouterRouteTargetRef(
		ctx,
		&services.CreateLogicalRouterRouteTargetRefRequest{
			ID: i.GetUUID(),
			LogicalRouterRouteTargetRef: &models.LogicalRouterRouteTargetRef{
				UUID: rt.GetUUID(),
			},
		},
	)

	return rt, err
}
