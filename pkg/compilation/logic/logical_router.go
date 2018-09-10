package logic

import (
	"context"

	"github.com/pkg/errors"
	"github.com/siddontang/go/log"

	"github.com/Juniper/contrail/pkg/compilation/intent"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
)

// LogicalRouterIntent contains Intent Compiler state for LogicalRouter.
type LogicalRouterIntent struct {
	intent.BaseIntent
	*models.LogicalRouter
	virtualNetworks    map[string]*models.VirtualMachineInterfaceVirtualNetworkRef
	defaultRouteTarget string
	vxlanRouting       bool
}

// NewLogicalRouterIntent creates LogicalRouterIntent from CreateLogicalRouterRequest
func NewLogicalRouterIntent(
	ctx context.Context,
	ReadService services.ReadService,
	request *services.CreateLogicalRouterRequest,
) *LogicalRouterIntent {
	lr := &LogicalRouterIntent{
		LogicalRouter: request.GetLogicalRouter(),
	}
	lr.resolveVxLan(ctx, ReadService)
	lr.virtualNetworks = make(map[string]*models.VirtualMachineInterfaceVirtualNetworkRef)

	return lr
}

// CreateLogicalRouter evaluates logical router dependencies.
func (s *Service) CreateLogicalRouter(
	ctx context.Context,
	request *services.CreateLogicalRouterRequest,
) (*services.CreateLogicalRouterResponse, error) {

	i := NewLogicalRouterIntent(ctx, s.ReadService, request)

	err := s.handleCreate(ctx, i, i.LogicalRouter)
	if err != nil {
		return nil, err
	}

	return s.BaseService.CreateLogicalRouter(ctx, request)
}

// ProcessCreate creates the default route target of a logical router.
func (i *LogicalRouterIntent) ProcessCreate(
	ctx context.Context,
	ec *intent.EvaluateContext,
) error {
	if len(i.LogicalRouter.GetRouteTargetRefs()) == 0 {
		if err := i.createDefaultRouteTarget(ctx, ec); err != nil {
			return errors.Wrap(err, "failed to create Logical Router's default Route Target")
		}
	}
	return nil
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
		ID:     i.GetParentUUID(),
		Fields: []string{"vxlan_routing"},
	})

	if err == nil {
		i.vxlanRouting = r.GetProject().VxlanRouting
	}
}

func (i *LogicalRouterIntent) updateVirtualNetworks(
	ctx context.Context,
	evaluateCtx *intent.EvaluateContext,
) error {
	vnRefs := make(map[string]*models.VirtualMachineInterfaceVirtualNetworkRef)
	for _, ref := range i.VirtualMachineInterfaceRefs {
		vmi := LoadVirtualMachineInterfaceIntent(evaluateCtx.IntentLoader, ref.UUID)
		if vmi != nil {
			if len(vmi.GetVirtualNetworkRefs()) > 0 {
				vnRefs[vmi.UUID] = vmi.GetVirtualNetworkRefs()[0]
			}
		} else {
			return errors.Errorf("failed to retrieve virtual-machine-interface "+
				"reference with uuid %s for logical-router with uuid %s", ref.UUID, i.GetUUID())
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
		return nil
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
	for _, vn := range i.getAddedNetworks(ctx, evaluateCtx, vnRefs) {
		ri := vn.GetPrimaryRoutingInstanceIntent(ctx, evaluateCtx)
		if ri != nil {
			// TODO handle all route targets
			uuid, err := i.getDefaultRouteTargetUUID()
			if err != nil {
				return err
			}
			evaluateCtx.WriteService.CreateRoutingInstanceRouteTargetRef(
				ctx, &services.CreateRoutingInstanceRouteTargetRefRequest{
					ID: ri.GetUUID(),
					RoutingInstanceRouteTargetRef: &models.RoutingInstanceRouteTargetRef{
						UUID: uuid,
					},
				},
			)
		} else {
			log.Errorf("Primary RI is None for VN: %s", vn.GetUUID())
		}
	}
	return nil
}

func (i *LogicalRouterIntent) getDeletedNetworks(
	ctx context.Context,
	evaluateCtx *intent.EvaluateContext,
	vnRefs map[string]*models.VirtualMachineInterfaceVirtualNetworkRef,
) []*VirtualNetworkIntent {
	vns := []*VirtualNetworkIntent{}
	for uuid, vnref := range i.virtualNetworks {
		if _, ok := vnRefs[uuid]; !ok {
			vn := LoadVirtualNetworkIntent(evaluateCtx.IntentLoader, vnref.UUID)
			if vn != nil {
				vns = append(vns, vn)
			} else {
				log.Errorf("Failed to load VN with uuid %s", vn.GetUUID())
			}
		}
	}
	return vns
}

func (i *LogicalRouterIntent) getAddedNetworks(
	ctx context.Context,
	evaluateCtx *intent.EvaluateContext,
	vnRefs map[string]*models.VirtualMachineInterfaceVirtualNetworkRef,
) []*VirtualNetworkIntent {
	vns := []*VirtualNetworkIntent{}
	for uuid, vnref := range vnRefs {
		if _, ok := i.virtualNetworks[uuid]; !ok {
			vn := LoadVirtualNetworkIntent(evaluateCtx.IntentLoader, vnref.UUID)
			if vn != nil {
				vns = append(vns, vn)
			} else {
				log.Errorf("failed to retrieve VN with uuid %s", vn.GetUUID())
			}
		}
	}
	return vns
}

func (i *LogicalRouterIntent) getDefaultRouteTargetUUID() (string, error) {
	if len(i.RouteTargetRefs) == 0 {
		return "", errors.Errorf("failed to get default route target for logical router with uuid %s", i.GetUUID())
	}
	return i.RouteTargetRefs[0].UUID, nil
}

func (i *LogicalRouterIntent) createDefaultRouteTarget(
	ctx context.Context,
	evaluateContext *intent.EvaluateContext,
) error {
	rt, err := createDefaultRouteTarget(ctx, evaluateContext)
	if err != nil {
		return err
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
	return err
}
