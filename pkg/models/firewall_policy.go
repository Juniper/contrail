package models

import (
	"github.com/Juniper/asf/pkg/models"
)

// CheckAssociatedRefsInSameScope checks scope of Firewall Policy references.
func (fp *FirewallPolicy) CheckAssociatedRefsInSameScope(fqName []string) error {
	var refs []models.Reference
	for _, ref := range fp.GetFirewallRuleRefs() {
		refs = append(refs, ref)
	}
	return checkAssociatedRefsInSameScope(fp, fqName, refs)
}
