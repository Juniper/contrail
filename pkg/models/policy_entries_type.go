package models

import (
	"github.com/Juniper/contrail/pkg/errutil"
	"github.com/twinj/uuid"
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

		srcSecGroupList, dstSecGroupList := rule.SRCSecurityGroups(), rule.DSTSecurityGroups()
		if len(srcSecGroupList) != 0 || len(dstSecGroupList) != 0 {
			return errutil.ErrorBadRequest("Config Error: Policy Rule refering to Security Group is not allowed")
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

	for _, rule := range rules {
		if err := rule.ValidateSubnetsWithEthertype(); err != nil {
			return err
		}
		if !rule.IsAnySecurityGroupAddrLocal() {
			return errutil.ErrorBadRequest("At least one of source " +
				"or destination addresses must be 'local'")
		}
	}

	return nil
}

// FillRuleUUIDs adds UUID to every PolicyRule within PolicyEntriesType
// which doesn't have one.
func (e *PolicyEntriesType) FillRuleUUIDs() {
	for i, rule := range e.PolicyRule {
		if rule.GetRuleUUID() == "" {
			e.PolicyRule[i].RuleUUID = uuid.NewV4().String()
		}
	}
}

func isAtLeastOneLocalSg(secGroups []string) bool {
	for _, sg := range secGroups {
		if sg == "local" {
			return true
		}
	}
	return false
}

func checkPolicyEntriesRules(rules []*PolicyRuleType) error {
	for i, rule := range rules {
		remainingRules := rules[i+1:]
		if isRuleInRules(rule, remainingRules) {
			return errutil.ErrorConflictf("Rule already exists: %v", rule.GetRuleUUID())
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
