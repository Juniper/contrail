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
// also sets protocolID, default MatchTag, tagRefs and addressGroupRefs
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

			if err = firewallRule.Validate(); err != nil {
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
	err := sv.InTransactionDoer.DoInTransaction(
		ctx,
		func(ctx context.Context) error {
			var err error

			//TODO update validation
			response, err = sv.BaseService.UpdateFirewallRule(ctx, request)
			return err
		})

	return response, err
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
			ID:     m.UUID,
			Fields: []string{models.TagTypeFieldTagTypeID},
		},
	)
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
	if !IsInternalRequest(ctx) && len(fr.GetTagRefs()) > 0 {
		return common.ErrorBadRequestf(
			"cannot directly define Tags reference from a Firewall Rule. " +
				"Use 'tags' endpoints property in the Firewall Rule")
	}

	//TODO initialize tagRefs
	return sv.setTagRefs(ctx, fr, databaseFR, fm)
}

func (sv *ContrailTypeLogicService) setTagRefs(
	ctx context.Context,
	fr *models.FirewallRule,
	databaseFR *models.FirewallRule,
	fm *types.FieldMask,
) error {
	endpoints, dbEndpoints := fr.GetEndpoints(databaseFR)
	fields := []string{models.FirewallRuleFieldEndpoint1, models.FirewallRuleFieldEndpoint2}

	for i, ep := range endpoints {
		if ep == nil && dbEndpoints[i] == nil {
			continue
		}

		if ep == nil && (fm == nil || !basemodels.FieldMaskContains(fm, fields[i])) {
			ep = dbEndpoints[i]
		}

		if ep != nil {
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

	return nil
}

func (sv *ContrailTypeLogicService) setTagRef(
	ctx context.Context, fr *models.FirewallRule, tagName string,
) (int64, error) {
	fqName, err := fr.GetTagFQName(tagName)
	if err != nil {
		return 0, err
	}

	tag, err := sv.getTagByFQName(ctx, fqName)
	if err != nil {
		return 0, err
	}

	//TODO append to tagRefs tag with given tagName

	id := strings.Replace(tag.GetTagID(), "0x", "", -1)
	return strconv.ParseInt(id, 16, 64)
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

	endpoints, dbEndpoints := fr.GetEndpoints(databaseFR)
	fields := []string{models.FirewallRuleFieldEndpoint1, models.FirewallRuleFieldEndpoint2}

	fr.AddressGroupRefs = []*models.FirewallRuleAddressGroupRef{}
	for i, ep := range endpoints {
		if ep == nil && dbEndpoints[i] == nil {
			continue
		}

		if ep == nil && (fm == nil || !basemodels.FieldMaskContains(fm, fields[i])) {
			ep = dbEndpoints[i]
		}

		if err := sv.addAddressGroupRef(ctx, fr, ep.GetAddressGroup()); err != nil {
			return err
		}
	}

	return nil
}

func (sv *ContrailTypeLogicService) addAddressGroupRef(
	ctx context.Context,
	fr *models.FirewallRule,
	fqname string,
) error {
	fqName := basemodels.ParseFQName(fqname)
	if fqName == nil {
		return nil
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

	return nil
}
