package models

import (
	"github.com/Juniper/contrail/pkg/errutil"
)

// CheckNetworkPolicyRules validates policy rules from policy entries for Network Policy.
func (e *PolicyEntriesType) CheckNetworkPolicyRules() error {
	rules := e.GetPolicyRule()

	if err := checkPolicyEntriesRules(rules); err != nil {
		return err
	}

	for _, rule := range rules {
		if rule.ActionList == nil {
			return errutil.ErrorBadRequest("Check Policy Rules failed. Action is required.")
		}

		srcSecGroupList, dstSecGroupList := extractSecGroupLists(rule)
		if len(srcSecGroupList) != 0 || len(dstSecGroupList) != 0 {
			return errutil.ErrorBadRequest("Check Policy Rules failed." +
				"Config Error: Policy Rule refering to Security Group is not allowed")
		}
	}

	return nil
}

// CheckSecurityGroupRules validates policy rules from policy entries for Security Group.
func (e *PolicyEntriesType) CheckSecurityGroupRules() error {
	rules := e.GetPolicyRule()

	if err := checkPolicyEntriesRules(rules); err != nil {
		return err
	}

	// TODO implement

	return nil
}

func checkPolicyEntriesRules(rules []*PolicyRuleType) error {
	for i, rule := range rules {
		remainingRules := rules[i+1:]
		if isRuleInRules(rule, remainingRules) {
			return errutil.ErrorConflictf("Rule already exists: %v", rule.GetRuleUUID())
		}

		if rule.GetRuleUUID() == "" {
			// generate uuid?
		}

		if err := rule.ValidateProtocol(); err != nil {
			return err
		}
	}
	return nil
}

func isRuleInRules(rule *PolicyRuleType, rules []*PolicyRuleType) bool {
	for _, r := range rules {
		if r.EqualRule(*rule) {
			return true
		}
	}
	return false
}

func extractSecGroupLists(rule *PolicyRuleType) (srcList, dstList []string) {
	for _, addr := range rule.GetSRCAddresses() {
		if sg := addr.GetSecurityGroup(); sg != "" {
			srcList = append(srcList, sg)
		}
	}
	for _, addr := range rule.GetDSTAddresses() {
		if sg := addr.GetSecurityGroup(); sg != "" {
			dstList = append(srcList, sg)
		}
	}
	return
}
