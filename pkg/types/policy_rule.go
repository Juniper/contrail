package types

import (
	"reflect"
	"strconv"

	"github.com/Juniper/contrail/pkg/errutil"
	"github.com/Juniper/contrail/pkg/models"
)

// CheckPolicyRules validates policy rules from policy entries for Security Group and Network Policy.
func CheckPolicyRules(entries *models.PolicyEntriesType, isNetworkPolicyRule bool) error {
	if entries == nil {
		return nil
	}

	rules := entries.GetPolicyRule()

	if err := verifyAllRulesAreDifferent(rules); err != nil {
		return err
	}

	for _, rule := range rules {

		if rule.GetRuleUUID() == "" {
			// generate uuid?
		}

		if err := verifyProtocol(rule.GetProtocol()); err != nil {
			return err
		}

		srcSecGroupList, dstSecGroupList := extractSecGroupLists(rule)

		if isNetworkPolicyRule {
			if rule.ActionList == nil {
				return errutil.ErrorBadRequest("Check Policy Rules failed. Action is required.")
			}
			if len(srcSecGroupList) != 0 || len(dstSecGroupList) != 0 {
				return errutil.ErrorBadRequest("Check Policy Rules failed." +
					"Config Error: Policy Rule refering to Security Group is not allowed")
			}
		} else {
			// TODO: Handle Security Group Policy Rules
		}
	}

	return nil
}

func verifyAllRulesAreDifferent(rules []*models.PolicyRuleType) error {
	length := len(rules)
	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			if !areRulesEqual(*rules[i], *rules[j]) {
				return errutil.ErrorConflictf("Rule already exists: %v", rules[i].GetRuleUUID())
			}
		}
	}
	return nil
}

func verifyProtocol(protocol string) error {
	if number, err := strconv.Atoi(protocol); err != nil {
		if number < 0 || number > 255 {
			return errutil.ErrorBadRequestf("Rule with invalid protocol: %v. "+
				"Protocol number range is 0 - 255.", number)
		}
		return nil
	}
	avaiableProtocols := []string{"any", "icmp", "tcp", "udp", "icmp6"}
	for _, p := range avaiableProtocols {
		if protocol == p {
			return nil
		}
	}
	return errutil.ErrorBadRequestf("Rule with invalid protocol: %v. "+
		"Protocol has to be one of these: %v", protocol, avaiableProtocols)
}

func extractSecGroupLists(rule *models.PolicyRuleType) (srcList, dstList []string) {
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

func areRulesEqual(a, b models.PolicyRuleType) bool {
	a.RuleUUID = ""
	b.RuleUUID = ""
	a.LastModified = ""
	b.LastModified = ""
	a.Created = ""
	b.Created = ""

	return reflect.DeepEqual(a, b)
}
