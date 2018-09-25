package types

import (
	"context"
	"strconv"
	"strings"

	"github.com/Juniper/contrail/pkg/common"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/models/basemodels"
	"github.com/Juniper/contrail/pkg/services"
)

// CreateFirewallRule TODO
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

			if err = firewallRule.CheckAssociatedRefsInSameScope(); err != nil {
				return err
			}

			if err = firewallRule.CheckServiceProperties(); err != nil {
				return err
			}

			firewallRule.AddDefaultMatchTag()

			if err = firewallRule.SetProtocolID(); err != nil {
				return err
			}

			if err = sv.setMatchTagTypes(ctx, firewallRule); err != nil {
				return err
			}

			if err = firewallRule.CheckEndpoints(); err != nil {
				return err
			}

			if err = sv.setTagRefs(ctx, firewallRule, nil); err != nil {
				return err
			}

			if err = sv.setAddressGroupRefs(ctx, firewallRule, nil); err != nil {
				return err
			}

			response, err = sv.Next().CreateFirewallRule(ctx, request)
			return err
		})

	return response, err
}

// UpdateFirewallRule TODO
func (sv *ContrailTypeLogicService) UpdateFirewallRule(
	ctx context.Context,
	request *services.UpdateFirewallRuleRequest,
) (*services.UpdateFirewallRuleResponse, error) {

	var response *services.UpdateFirewallRuleResponse
	//firewallRule := request.GetFirewallRule()

	err := sv.InTransactionDoer.DoInTransaction(
		ctx,
		func(ctx context.Context) error {
			var err error

			response, err = sv.BaseService.UpdateFirewallRule(ctx, request)
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

func (sv *ContrailTypeLogicService) setMatchTagTypes(
	ctx context.Context, fr *models.FirewallRule,
) error {
	fr.MatchTagTypes = &models.FirewallRuleMatchTagsTypeIdList{
		TagType: []int64{},
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

func (sv *ContrailTypeLogicService) setTagRefs(
	ctx context.Context,
	fr *models.FirewallRule,
	databaseFR *models.FirewallRule,
) error {
	//TODO delete tag_refs from here , not existing in schema
	if !IsInternalRequest(ctx) && len(fr.GetTagRefs()) > 0 {
		return common.ErrorBadRequestf(
			"cannot directly define Tags reference from a Firewall Rule. " +
				"Use 'tags' endpoints property in the Firewall Rule")
	}

	endpoints := []*models.FirewallRuleEndpointType{
		fr.GetEndpoint1(),
		fr.GetEndpoint2(),
	}

	dbEndpoints := []*models.FirewallRuleEndpointType{
		databaseFR.GetEndpoint1(),
		databaseFR.GetEndpoint2(),
	}

	fr.TagRefs = []*models.FirewallRuleTagRef{}
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
			fqName, err := getTagFQName(fr, tagName)
			if err != nil {
				return err
			}

			tag, err := sv.getTagByFQName(ctx, fqName)
			if err != nil {
				return err
			}

			tagID, err := strconv.ParseInt(tag.GetTagID(), 10, 64)
			if err != nil {
				return err
			}

			if endpoints[i] != nil {
				ep.TagIds = append(ep.GetTagIds(), tagID)
			}

			fr.TagRefs = append(
				fr.GetTagRefs(),
				&models.FirewallRuleTagRef{
					UUID: tag.GetUUID(),
					To:   tag.GetFQName(),
				},
			)
		}
	}

	return nil
}

func getTagFQName(fr *models.FirewallRule, tagName string) ([]string, error) {
	if !strings.Contains(tagName, "=") {
		return nil, common.ErrorNotFoundf("invalid tag name '%s'", tagName)
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
) error {
	if !IsInternalRequest(ctx) && len(fr.GetAddressGroupRefs()) > 0 {
		return common.ErrorBadRequestf(
			"cannot directly define Address Group reference from a Firewall Rule. " +
				"Use 'address_group' endpoints property in the Firewall Rule")
	}

	endpoints := []*models.FirewallRuleEndpointType{
		fr.GetEndpoint1(),
		fr.GetEndpoint2(),
	}

	dbEndpoints := []*models.FirewallRuleEndpointType{
		databaseFR.GetEndpoint1(),
		databaseFR.GetEndpoint2(),
	}

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

	return nil
}
