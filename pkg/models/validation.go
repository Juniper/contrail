package models

import (
	"net"
	"regexp"
	"strconv"
	"strings"

	"github.com/pkg/errors"

	"github.com/Juniper/contrail/pkg/models/basemodels"
)

//NewTypeValidatorWithFormat creates new TypeValidator with format validators
func NewTypeValidatorWithFormat() (*TypeValidator, error) {
	base, err := basemodels.NewBaseValidatorWithFormat()
	if err != nil {
		return nil, err
	}
	tv := &TypeValidator{
		SchemaValidator: SchemaValidator{
			BaseValidator: base,
		},
	}

	// Create regex used while validating CommunityAttributes
	tv.communityAttributeRegexStr = "^[0-9]+:[0-9]+$"
	r, err := regexp.Compile(tv.communityAttributeRegexStr)
	if err != nil {
		return nil, err
	}
	tv.communityAttributeRegex = r
	return tv, nil
}

//SchemaValidator implementing basic checks based on information in schema
type SchemaValidator struct {
	*basemodels.BaseValidator
}

//TypeValidator embedding SchemaValidator validator. It enables defining custom validation for each type
type TypeValidator struct {
	SchemaValidator
	communityAttributeRegex    *regexp.Regexp
	communityAttributeRegexStr string
}

//ValidateAllowedAddressPair custom validation for AllowedAddressPair
func (tv *TypeValidator) ValidateAllowedAddressPair(obj *AllowedAddressPair) error {
	err := tv.SchemaValidator.ValidateAllowedAddressPair(obj)
	if err != nil {
		return err
	}

	if obj.AddressMode != "active-standby" {
		return nil
	}

	ip := net.ParseIP(obj.IP.IPPrefix)

	if ip.To4() == nil {
		if obj.IP.IPPrefixLen < 120 {
			return errors.Errorf("IPv6 Prefix length lesser than 120 is not acceptable")
		}
	} else {
		if obj.IP.IPPrefixLen < 24 {
			return errors.Errorf("IPv4 Prefix length lesser than 24 is not acceptable")
		}
	}
	return nil
}

//ValidateCommunityAttributes custom validation for AllowedAddressPair
func (tv *TypeValidator) ValidateCommunityAttributes(obj *CommunityAttributes) error {

	restrictions := map[string]struct{}{
		"no-export":           {},
		"accept-own":          {},
		"no-advertise":        {},
		"no-export-subconfed": {},
		"no-reoriginate":      {},
	}

	for _, value := range obj.CommunityAttribute {
		if _, ok := restrictions[value]; ok {
			continue
		}

		if !tv.communityAttributeRegex.MatchString(value) {
			return errors.Errorf("CommunityAttribute value (%s) must be either one of "+
				"[%s] or match \"%s\"", value, strings.Join(mapKeys(restrictions), ", "), tv.communityAttributeRegexStr)
		}
		asn := strings.Split(value, ":")

		asn0, err := strconv.Atoi(asn[0])
		if err != nil {
			return errors.Wrapf(err, "error while parsing CommunityAttribute.")
		}

		if asn0 > 65535 {
			return errors.Errorf("Out of range ASN value %v. ASN values cannot exceed 65535.", asn0)

		}
	}

	return nil
}

// Returns array of map keys
func mapKeys(m map[string]struct{}) (keys []string) {
	for s := range m {
		keys = append(keys, s)
	}
	return keys
}
