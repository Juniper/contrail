package models

import (
	"github.com/Juniper/contrail/pkg/errutil"
)

func (m *QosIdForwardingClassPair) CheckValueForDscp() error {
	if m != nil && (m.GetKey() < 0 || m.GetKey() > 63) {
		return errutil.ErrorBadRequestf("Invalid DSCP value %v.", m.GetKey())
	}
	return nil
}

func (m *QosIdForwardingClassPair) CheckValueForVlanPriority() error {
	if m != nil && (m.GetKey() < 0 || m.GetKey() > 7) {
		return errutil.ErrorBadRequestf("Invalid 802.1p value %v.", m.GetKey())
	}
	return nil
}

func (m *QosIdForwardingClassPair) CheckValueForMplsExp() error {
	if m != nil && (m.GetKey() < 0 || m.GetKey() > 7) {
		return errutil.ErrorBadRequestf("Invalid MPLS EXP value %v.", m.GetKey())
	}
	return nil
}
