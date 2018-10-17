package models

import (
	"github.com/Juniper/contrail/pkg/models/basemodels"
)

// CheckAssociatedRefsInSameScope checks scope of Firewall Policy references.
func (fp *FirewallPolicy) CheckAssociatedRefsInSameScope(fqName []string) error {
	var refs []basemodels.Reference
	for _, ref := range fp.GetFirewallRuleRefs() {
		refs = append(refs, ref)
	}
	return CheckAssociatedRefsInSameScope(fp, fqName, refs)
}
