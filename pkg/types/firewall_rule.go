package types

import (
	"context"
	"strconv"
	"strings"

	"github.com/gogo/protobuf/types"

	"github.com/Juniper/asf/pkg/models/basemodels"
	"github.com/Juniper/asf/pkg/services/baseservices"
	"github.com/Juniper/contrail/pkg/errutil"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
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

			if err = sv.validateFirewallRule(ctx, firewallRule, nil, nil); err != nil {
				return err
			}

			firewallRule.AddDefaultMatchTag(nil)

			if err = SetProtocolID(firewallRule, nil); err != nil {
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

			response, err = sv.BaseService.CreateFirewallRule(ctx, request)
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
			databaseFR, err := sv.getFirewallRule(ctx, firewallRule.GetUUID())
			if err != nil {
				return err
			}

			if err = sv.validateFirewallRule(ctx, firewallRule, databaseFR, &fm); err != nil {
				return err
			}

			firewallRule.AddDefaultMatchTag(&fm)

			if err = SetProtocolID(firewallRule, &fm); err != nil {
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

			response, err = sv.BaseService.UpdateFirewallRule(
				ctx,
				&services.UpdateFirewallRuleRequest{
					FirewallRule: firewallRule,
					FieldMask:    fm,
				},
			)
			return err
		})

	return response, err
}

func (sv *ContrailTypeLogicService) validateFirewallRule(
	ctx context.Context,
	firewallRule *models.FirewallRule,
	databaseFR *models.FirewallRule,
	fm *types.FieldMask,
) error {

	fqName := firewallRule.GetFQName()
	if len(fqName) == 0 {
		fqName = databaseFR.GetFQName()
	}

	if err := checkDraftModeState(ctx, firewallRule); err != nil {
		return err
	}

	if err := sv.complementRefs(ctx, firewallRule); err != nil {
		return err
	}

	if err := firewallRule.CheckAssociatedRefsInSameScope(fqName); err != nil {
		return errutil.ErrorBadRequest(err.Error())
	}

	return CheckServiceProperties(firewallRule, databaseFR, fm)
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

// CheckServiceProperties checks for existence of service and serviceGroupRefs property
func CheckServiceProperties(
	fr *models.FirewallRule, databaseFR *models.FirewallRule, fm *types.FieldMask,
) error {
	serviceGroupRefs := fr.GetServiceGroupRefs()
	if fm != nil && !basemodels.FieldMaskContains(fm, models.FirewallRuleFieldServiceGroupRefs) {
		serviceGroupRefs = databaseFR.GetServiceGroupRefs()
	}

	service := fr.GetService()
	if fm != nil && !basemodels.FieldMaskContains(fm, models.FirewallRuleFieldService) {
		service = databaseFR.GetService()
	}

	if service == nil || service.GetProtocol() == "" {
		if len(serviceGroupRefs) == 0 {
			return errutil.ErrorBadRequest("firewall Rule requires at least 'service' property or Service Group references(s)")
		}
	} else {
		if len(serviceGroupRefs) > 0 {
			return errutil.ErrorBadRequest(
				"firewall Rule cannot have both defined 'service' property and Service Group reference(s)",
			)
		}
	}

	return nil
}

// SetProtocolID sets protocolID based on protocol property
func SetProtocolID(fr *models.FirewallRule, fm *types.FieldMask) error {
	if fr.GetService() == nil || (fm != nil &&
		!basemodels.FieldMaskContains(
			fm,
			models.FirewallRuleFieldService,
			models.FirewallServiceTypeFieldProtocol,
		)) {
		return nil
	}

	protocolID, err := fr.GetProtocolID()
	fr.Service.ProtocolID = protocolID

	basemodels.FieldMaskAppend(fm, models.FirewallRuleFieldService, models.FirewallServiceTypeFieldProtocolID)
	return err
}

func (sv *ContrailTypeLogicService) setMatchTagTypes(
	ctx context.Context, fr *models.FirewallRule, fm *types.FieldMask,
) error {

	fr.MatchTagTypes = &models.FirewallRuleMatchTagsTypeIdList{
		TagType: []int64{},
	}

	for _, tagType := range fr.GetMatchTags().GetTagList() {
		tagType = strings.ToLower(tagType)
		if tagType == "label" {
			return errutil.ErrorBadRequest("labels not allowed as match-tags")
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

	if len(fr.GetMatchTagTypes().GetTagType()) > 0 {
		basemodels.FieldMaskAppend(
			fm,
			models.FirewallRuleFieldMatchTagTypes,
			models.FirewallRuleMatchTagsTypeIdListFieldTagType,
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
		return -1, errutil.ErrorNotFoundf("cannot find tag-type %s uuid: %v", tagType, err)
	}

	tagTypeResponse, err := sv.ReadService.GetTagType(
		ctx,
		&services.GetTagTypeRequest{
			ID:     m.UUID,
			Fields: []string{models.TagTypeFieldTagTypeID},
		},
	)
	if err != nil {
		return -1, errutil.ErrorNotFoundf("cannot find tag-type %s: %v", tagType, err)
	}

	id := tagTypeResponse.GetTagType().GetTagTypeID()
	if id == "" {
		return 0, nil
	}
	return strconv.ParseInt(id, 0, 64)
}

func (sv *ContrailTypeLogicService) setTagProperties(
	ctx context.Context,
	fr *models.FirewallRule,
	databaseFR *models.FirewallRule,
	fm *types.FieldMask,
) error {
	if !baseservices.IsInternalRequest(ctx) && len(fr.GetTagRefs()) > 0 {
		return errutil.ErrorBadRequestf(
			"cannot directly define Tags reference from a Firewall Rule. " +
				"Use 'tags' endpoints property in the Firewall Rule")
	}

	fr.TagRefs = []*models.FirewallRuleTagRef{}

	basemodels.FieldMaskAppend(fm, models.FirewallRuleFieldTagRefs)
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
	endpointIndices := []int{models.Endpoint1Index, models.Endpoint2Index}

	for i, ep := range endpoints {
		if ep == nil && dbEndpoints[i] == nil {
			continue
		}

		if ep == nil && !basemodels.FieldMaskContains(fm, fields[i]) {
			ep = dbEndpoints[i]
		}

		if ep != nil && basemodels.FieldMaskContains(fm, fields[i]) {
			ep.TagIds = nil
			basemodels.FieldMaskAppend(fm, fields[i], models.FirewallRuleEndpointTypeFieldTagIds)
		}

		if err := sv.processTags(ctx, fr, databaseFR, ep); err != nil {
			return err
		}

		if err := fr.SetEndpoint(endpoints[i], endpointIndices[i]); err != nil {
			return err
		}
	}

	return nil
}

func (sv *ContrailTypeLogicService) processTags(
	ctx context.Context,
	fr *models.FirewallRule,
	databaseFR *models.FirewallRule,
	ep *models.FirewallRuleEndpointType,
) error {
	for _, tagName := range ep.GetTags() {
		tagID, err := sv.setTagRef(ctx, fr, databaseFR, tagName)
		if err != nil {
			return err
		}

		ep.TagIds = append(ep.GetTagIds(), tagID)
	}
	return nil
}

func (sv *ContrailTypeLogicService) setTagRef(
	ctx context.Context,
	fr *models.FirewallRule,
	databaseFR *models.FirewallRule,
	tagName string,
) (int64, error) {
	parentType := fr.GetParentType()
	if parentType == "" {
		parentType = databaseFR.GetParentType()
	}

	frFQName := fr.GetFQName()
	if len(frFQName) == 0 {
		frFQName = databaseFR.GetFQName()
	}

	fqName, err := fr.GetTagFQName(tagName, parentType, frFQName)
	if err != nil {
		return 0, err
	}

	tag, err := sv.getTagByFQName(ctx, fqName)
	if err != nil {
		return 0, err
	}

	fr.TagRefs = append(
		fr.GetTagRefs(),
		&models.FirewallRuleTagRef{
			UUID: tag.GetUUID(),
			To:   tag.GetFQName(),
		},
	)

	if tag.GetTagID() == "" {
		return 0, nil
	}
	return strconv.ParseInt(tag.GetTagID(), 0, 64)
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
		return nil, errutil.ErrorNotFoundf("cannot find Tag (fq_name: %s): %v", tagFQName, err)
	}

	tagResponse, err := sv.ReadService.GetTag(
		ctx,
		&services.GetTagRequest{
			ID: m.UUID,
		})
	if err != nil {
		return nil, errutil.ErrorNotFoundf("cannot get Tag (uuid: %s): %v", m.UUID, err)
	}

	return tagResponse.GetTag(), nil
}

func (sv *ContrailTypeLogicService) setAddressGroupRefs(
	ctx context.Context,
	fr *models.FirewallRule,
	databaseFR *models.FirewallRule,
	fm *types.FieldMask,
) error {
	if !baseservices.IsInternalRequest(ctx) && len(fr.GetAddressGroupRefs()) > 0 {
		return errutil.ErrorBadRequestf(
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

		if ep == nil && !basemodels.FieldMaskContains(fm, fields[i]) {
			ep = dbEndpoints[i]
		}

		if err := sv.addAddressGroupRef(ctx, fr, ep.GetAddressGroup()); err != nil {
			return err
		}
	}

	handleAddressGroupRefsFieldMask(fr, fm)
	return nil
}

func handleAddressGroupRefsFieldMask(fr *models.FirewallRule, fm *types.FieldMask) {
	if len(fr.GetAddressGroupRefs()) > 0 {
		basemodels.FieldMaskAppend(fm, models.FirewallRuleFieldAddressGroupRefs)
	}
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
		return errutil.ErrorNotFoundf(
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
