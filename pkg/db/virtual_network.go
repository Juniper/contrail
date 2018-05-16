package db

import (
	"context"

	"github.com/Juniper/contrail/pkg/models"
	"github.com/pkg/errors"
)

const (
	refLogicalRouterBGPVPNTable                   = "ref_logical_router_bgpvpn"
	refLogicalRouterVirtualMachineInterfaceTable  = "ref_logical_router_virtual_machine_interface"
	refVirtualMachineInterfaceVirtualNetowrkTable = "ref_virtual_machine_interface_virtual_network"
)

//GetLinkedBGPVPNviaRouter find bgp vpn linked to a virtual network via router.
func (db *Service) GetLinkedBGPVPNviaRouter(
	ctx context.Context, virtualNetwork *models.VirtualNetwork) (bgpVPNIDs []string, err error) {
	d := db.Dialect
	tx := GetTransaction(ctx)
	query := "select " + d.quote("to") + " from " + d.quote(refLogicalRouterBGPVPNTable) +
		"left join " + d.quote(refLogicalRouterVirtualMachineInterfaceTable) + " on " +
		d.quote(refLogicalRouterVirtualMachineInterfaceTable, "from") + " = " + d.quote(refLogicalRouterBGPVPNTable, "from") +
		"left join " + d.quote(refVirtualMachineInterfaceVirtualNetowrkTable) + " on " +
		d.quote(refVirtualMachineInterfaceVirtualNetowrkTable+"to") + " = " + d.quote(refLogicalRouterVirtualMachineInterfaceTable+"to") +
		"where " + d.quote(refVirtualMachineInterfaceVirtualNetowrkTable, "to") + " = " + d.placeholder(1)
	rows, err := tx.QueryContext(ctx, query, virtualNetwork.UUID)
	if err != nil {
		err = handleError(err)
		return nil, errors.Wrap(err, "select query failed")
	}
	defer rows.Close()
	if err := rows.Err(); err != nil {
		err = handleError(err)
		return nil, errors.Wrap(err, "row error")
	}
	for rows.Next() {
		var bgpVPNID string
		if err := rows.Scan(&bgpVPNID); err != nil {
			return nil, errors.Wrap(err, "scan failed")
		}
		bgpVPNIDs = append(bgpVPNIDs, bgpVPNID)
	}
	return bgpVPNIDs, nil
}
