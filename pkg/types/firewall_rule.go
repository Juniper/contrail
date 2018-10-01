package types

import (
	"context"
	"strconv"
	"strings"

	"github.com/Juniper/contrail/pkg/common"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/models/basemodels"
	"github.com/Juniper/contrail/pkg/services"
	"github.com/gogo/protobuf/types"
)

// CreateFirewallRule performs types specific validation,
// also sets protocolID, default MatchTag, tag- and addressGroupRefs
func (sv *ContrailTypeLogicService) CreateFirewallRule(
	ctx context.Context,
	request *services.CreateFirewallRuleRequest,
) (*services.CreateFirewallRuleResponse, error) {

	var response *services.CreateFirewallRuleResponse
	firewallRule := request.GetFirewallRule()

	err := sv.InTransactionDoer.DoInTransaction(
		ctx,
		func(ctx context.Context) error {
			var err error

			if err = checkDraftModeState(ctx, firewallRule); err != nil {
				return err
			}

			if err = firewallRule.CheckAssociatedRefsInSameScope(nil); err != nil {
				return err
			}

			if err = firewallRule.CheckServiceProperties(nil); err != nil {
				return err
			}

			firewallRule.AddDefaultMatchTag(nil)

			if err = firewallRule.SetProtocolID(nil); err != nil {
				return err
			}

			if err = sv.setMatchTagTypes(ctx, firewallRule, nil); err != nil {
				return err
			}

			if err = firewallRule.CheckEndpoints(); err != nil {
				return err
			}

			if err = sv.setTagProperties(ctx, firewallRule, nil, nil); err != nil {
				return err
			}

			if err = sv.setAddressGroupRefs(ctx, firewallRule, nil, nil); err != nil {
				return err
			}

			response, err = sv.Next().CreateFirewallRule(ctx, request)
			return err
		})

	return response, err
}

// UpdateFirewallRule performs type specific validation and setup
// for updating firewallRule
func (sv *ContrailTypeLogicService) UpdateFirewallRule(
	ctx context.Context,
	request *services.UpdateFirewallRuleRequest,
) (*services.UpdateFirewallRuleResponse, error) {

	var response *services.UpdateFirewallRuleResponse
	firewallRule := request.GetFirewallRule()
	fm := request.GetFieldMask()

	err := sv.InTransactionDoer.DoInTransaction(
		ctx,
		func(ctx context.Context) error {
			var err error

			if err = checkDraftModeState(ctx, firewallRule); err != nil {
				return err
			}

			databaseFR, err := sv.getFirewallRule(ctx, firewallRule.GetUUID())
			if err != nil {
				return err
			}

			if err = firewallRule.CheckAssociatedRefsInSameScope(databaseFR); err != nil {
				return err
			}

			if err = firewallRule.CheckServiceProperties(databaseFR); err != nil {
				return err
			}

			firewallRule.AddDefaultMatchTag(&fm)

			if err = firewallRule.SetProtocolID(&fm); err != nil {
				return err
			}

			if err = sv.setMatchTagTypes(ctx, firewallRule, &fm); err != nil {
				return err
			}

			if err = firewallRule.CheckEndpoints(); err != nil {
				return err
			}

			if err = sv.setTagProperties(ctx, firewallRule, databaseFR, &fm); err != nil {
				return err
			}

			if err = sv.setAddressGroupRefs(ctx, firewallRule, databaseFR, &fm); err != nil {
				return err
			}

			response, err = sv.Next().UpdateFirewallRule(ctx, request)
			return err
		})

	return response, err
}

func checkDraftModeState(ctx context.Context, fr *models.FirewallRule) error {
	if IsInternalRequest(ctx) {
		return nil
	}

	if fr.GetDraftModeState() != "" {
		return common.ErrorBadRequest(
			"security resource property 'draft_mode_state' is only readable",
		)
	}

	return nil
}

func (sv *ContrailTypeLogicService) getFirewallRule(
	ctx context.Context, id string,
) (*models.FirewallRule, error) {

	firewallRuleResponse, err := sv.ReadService.GetFirewallRule(
		ctx,
		&services.GetFirewallRuleRequest{
			ID: id,
		},
	)

	return firewallRuleResponse.GetFirewallRule(), err
}

func (sv *ContrailTypeLogicService) setMatchTagTypes(
	ctx context.Context, fr *models.FirewallRule, fm *types.FieldMask,
) error {
	if fm != nil &&
		basemodels.FieldMaskContains(fm, models.FirewallRuleFieldMatchTags) {
		fr.MatchTagTypes = &models.FirewallRuleMatchTagsTypeIdList{
			TagType: []int64{},
		}
	}

	for _, tagType := range fr.GetMatchTags().GetTagList() {
		tagType = strings.ToLower(tagType)
		if tagType == "label" {
			return common.ErrorBadRequest("labels not allowed as match-tags")
		}

		tagTypeID, err := sv.getTagTypeID(ctx, tagType)
		if err != nil {
			return err
		}

		fr.MatchTagTypes.TagType = append(
			fr.GetMatchTagTypes().GetTagType(),
			tagTypeID,
		)
	}

	if len(fr.GetMatchTagTypes().GetTagType()) > 0 && fm != nil &&
		!basemodels.FieldMaskContains(fm, models.FirewallRuleFieldMatchTagTypes) {
		fm.Paths = append(fm.Paths, models.FirewallRuleFieldMatchTagTypes)
	}

	return nil
}

func (sv *ContrailTypeLogicService) getTagTypeID(
	ctx context.Context, tagType string,
) (int64, error) {
	tagTypeID, ok := models.TagTypeIDs[tagType]
	if ok {
		return tagTypeID, nil
	}

	m, err := sv.MetadataGetter.GetMetadata(
		ctx,
		basemodels.Metadata{
			FQName: []string{tagType},
			Type:   models.KindTagType,
		},
	)
	if err != nil {
		return -1, common.ErrorNotFoundf("cannot find tag-type %s uuid: %v", tagType, err)
	}

	tagTypeResponse, err := sv.ReadService.GetTagType(
		ctx,
		&services.GetTagTypeRequest{
			ID: m.UUID,
		})
	if err != nil {
		return -1, common.ErrorNotFoundf("cannot find tag-type %s: %v", tagType, err)
	}

	id := tagTypeResponse.GetTagType().GetTagTypeID()
	return strconv.ParseInt(id, 10, 64)
}

func (sv *ContrailTypeLogicService) setTagProperties(
	ctx context.Context,
	fr *models.FirewallRule,
	databaseFR *models.FirewallRule,
	fm *types.FieldMask,
) error {
	//TODO uncomment when tag refs definied
	// if !IsInternalRequest(ctx) && len(fr.GetTagRefs()) > 0 {
	// 	return common.ErrorBadRequestf(
	// 		"cannot directly define Tags reference from a Firewall Rule. " +
	// 			"Use 'tags' endpoints property in the Firewall Rule")
	// }

	endpoints, dbEndpoints := getEndpoints(fr, databaseFR)

	//TODO uncomment when tag refs definied
	//fr.TagRefs = []*models.FirewallRuleTagRef{}
	for i, ep := range endpoints {
		if ep == nil && dbEndpoints[i] == nil {
			continue
		}

		if ep == nil && dbEndpoints[i] != nil {
			ep = dbEndpoints[i]
		}

		if endpoints[i] != nil {
			ep.TagIds = nil
		}

		for _, tagName := range ep.GetTags() {
			tagID, err := sv.setTagRef(ctx, fr, tagName)
			if err != nil {
				return err
			}

			ep.TagIds = append(ep.GetTagIds(), tagID)
		}
	}

	//TODO uncomment when tag refs definied
	// if len(fr.GetTagRefs()) > 0 && fm != nil &&
	// 	!basemodels.FieldMaskContains(fm, models.FirewallRuleFieldTagRefs) {
	// 	fm.Paths = append(fm.Paths, models.FirewallRuleFieldTagRefs)
	// }

	return nil
}

func (sv *ContrailTypeLogicService) setTagRef(
	ctx context.Context, fr *models.FirewallRule, tagName string,
) (int64, error) {
	fqName, err := getTagFQName(fr, tagName)
	if err != nil {
		return 0, err
	}

	tag, err := sv.getTagByFQName(ctx, fqName)
	if err != nil {
		return 0, err
	}

	//TODO uncomment when tag refs definied
	// fr.TagRefs = append(
	// 	fr.GetTagRefs(),
	// 	&models.FirewallRuleTagRef{
	// 		UUID: tag.GetUUID(),
	// 		To:   tag.GetFQName(),
	// 	},
	// )

	id := strings.Replace(tag.GetTagID(), "0x", "", -1)
	return strconv.ParseInt(id, 16, 64)
}

func getTagFQName(fr *models.FirewallRule, tagName string) ([]string, error) {
	if !strings.Contains(tagName, "=") {
		return nil, common.ErrorBadRequestf("invalid tag name '%s'", tagName)
	}

	if strings.HasPrefix(tagName, "global:") {
		return []string{tagName[7:]}, nil
	}

	fqName := append([]string(nil), fr.GetFQName()...)
	if fr.GetParentType() == models.KindPolicyManagement {
		return append(fqName[:len(fqName)-2], tagName), nil
	}

	fqName = append([]string(nil), fr.GetFQName()...)
	if fr.GetParentType() == models.KindProject {
		return append(fqName[:len(fqName)-1], tagName), nil
	}

	return nil, common.ErrorBadRequestf(
		"Firewall rule %s (%s) parent type '%s' is not supported as security resource scope",
		fr.GetUUID(),
		fr.GetFQName(),
		fr.GetParentType(),
	)
}

func (sv *ContrailTypeLogicService) getTagByFQName(
	ctx context.Context, tagFQName []string,
) (*models.Tag, error) {
	m, err := sv.MetadataGetter.GetMetadata(
		ctx,
		basemodels.Metadata{
			FQName: tagFQName,
			Type:   models.KindTag,
		},
	)
	if err != nil {
		return nil, common.ErrorNotFoundf("cannot find Tag (fq_name: %s): %v", tagFQName, err)
	}

	tagResponse, err := sv.ReadService.GetTag(
		ctx,
		&services.GetTagRequest{
			ID: m.UUID,
		})
	if err != nil {
		return nil, common.ErrorNotFoundf("cannot get Tag (uuid: %s): %v", m.UUID, err)
	}

	return tagResponse.GetTag(), nil
}

func (sv *ContrailTypeLogicService) setAddressGroupRefs(
	ctx context.Context,
	fr *models.FirewallRule,
	databaseFR *models.FirewallRule,
	fm *types.FieldMask,
) error {
	if !IsInternalRequest(ctx) && len(fr.GetAddressGroupRefs()) > 0 {
		return common.ErrorBadRequestf(
			"cannot directly define Address Group reference from a Firewall Rule. " +
				"Use 'address_group' endpoints property in the Firewall Rule")
	}

	endpoints, dbEndpoints := getEndpoints(fr, databaseFR)

	fr.AddressGroupRefs = []*models.FirewallRuleAddressGroupRef{}
	for i, ep := range endpoints {
		if ep == nil && dbEndpoints[i] == nil {
			continue
		}

		if ep == nil && dbEndpoints[i] != nil {
			ep = dbEndpoints[i]
		}

		fqName := basemodels.ParseFQName(ep.GetAddressGroup())
		if fqName == nil {
			continue
		}

		m, err := sv.MetadataGetter.GetMetadata(
			ctx,
			basemodels.Metadata{
				FQName: fqName,
				Type:   models.KindAddressGroup,
			},
		)
		if err != nil {
			return common.ErrorNotFoundf(
				"no Address Group found for %s: %v",
				fqName,
				err,
			)
		}

		fr.AddressGroupRefs = append(
			fr.GetAddressGroupRefs(),
			&models.FirewallRuleAddressGroupRef{
				UUID: m.UUID,
				To:   fqName,
			},
		)
	}

	if len(fr.GetAddressGroupRefs()) > 0 && fm != nil &&
		!basemodels.FieldMaskContains(fm, models.FirewallRuleFieldAddressGroupRefs) {
		fm.Paths = append(fm.Paths, models.FirewallRuleFieldAddressGroupRefs)
	}

	return nil
}

func getEndpoints(
	fr *models.FirewallRule, databaseFR *models.FirewallRule,
) ([]*models.FirewallRuleEndpointType, []*models.FirewallRuleEndpointType) {
	endpoints := []*models.FirewallRuleEndpointType{
		fr.GetEndpoint1(),
		fr.GetEndpoint2(),
	}

	dbEndpoints := []*models.FirewallRuleEndpointType{
		databaseFR.GetEndpoint1(),
		databaseFR.GetEndpoint2(),
	}

	return endpoints, dbEndpoints
}
