package models

import (
	"net"
	"strconv"

	"github.com/Juniper/contrail/pkg/errutil"
	"github.com/twinj/uuid"
)

// CheckNetworkPolicyRules validates policy rules from policy entries for Network Policy.
func (e *PolicyEntriesType) CheckNetworkPolicyRules() error {
	rules := e.GetPolicyRule()

	if err := checkPolicyEntriesRules(rules); err != nil {
		return err
	}

	fillRuleUUIDs(rules)

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

	fillRuleUUIDs(rules)

	for _, rule := range rules {
		if rule.GetEthertype() != "" {
			addresses := rule.GetSRCAddresses()
			addresses = append(addresses, rule.GetDSTAddresses()...)
			for _, addr := range addresses {
				if err := verifySgAddress(addr, rule.GetEthertype()); err != nil {
					return err
				}
			}
		}
		SecGroupList := append(rule.SRCSecurityGroups(), rule.DSTSecurityGroups()...)
		if !isAtLeastOneLocalSg(SecGroupList) {
			return errutil.ErrorBadRequest("At least one of source " +
				"or destination addresses must be 'local'")
		}
	}

	return nil
}

func verifySgAddress(addr *AddressType, ethertype string) error {
	if addr.GetSubnet() != nil {
		IPPrefix := addr.GetSubnet().GetIPPrefix()
		IPPrefixLen := strconv.Itoa(int(addr.GetSubnet().GetIPPrefixLen()))
		version, err := getIPVersionFromCIDR(IPPrefix, IPPrefixLen)
		if err != nil {
			return err
		}
		if ethertype != version {
			return errutil.ErrorBadRequestf("Rule subnet %v doesn't match ethertype %v",
				IPPrefix+"/"+IPPrefixLen, ethertype)
		}
	}
	return nil
}

func getIPVersionFromCIDR(IPPrefix, IPPrefixLen string) (string, error) {
	network, _, err := net.ParseCIDR(IPPrefix + "/" + IPPrefixLen)
	if err != nil {
		return "", errutil.ErrorBadRequestf("Cannot parse address %v/%v. %v.",
			IPPrefix, IPPrefixLen, err)
	}
	if network.To4() != nil {
		return "IPv4", nil
	}
	if network.To16() != nil {
		return "IPv6", nil
	}
	return "", errutil.ErrorBadRequestf("Cannot resolve ip version %v/%v.",
		IPPrefix, IPPrefixLen)
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

func fillRuleUUIDs(rules []*PolicyRuleType) {
	for i, rule := range rules {
		if rule.GetRuleUUID() == "" {
			rules[i].RuleUUID = uuid.NewV4().String()
		}
	}
}

func isRuleInRules(rule *PolicyRuleType, rules []*PolicyRuleType) bool {
	for _, r := range rules {
		if r.EqualRule(*rule) {
			return true
		}
	}
	return false
}
