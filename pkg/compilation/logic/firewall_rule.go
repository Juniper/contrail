package logic

import (
	"context"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"

	"github.com/Juniper/contrail/pkg/compilation/intent"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/models/basemodels"
	"github.com/Juniper/contrail/pkg/services"
)

// FirewallRuleIntent intent
type FirewallRuleIntent struct {
	intent.BaseIntent
	*models.FirewallRule
}

// GetObject returns embedded resource object
func (i *FirewallRuleIntent) GetObject() basemodels.Object {
	return i.FirewallRule
}

// NewFirewallRuleIntent returns a new firewall rule intent.
func NewFirewallRuleIntent(
	_ context.Context,
	_ services.ReadService,
	request *services.CreateFirewallRuleRequest,
) *FirewallRuleIntent {
	return &FirewallRuleIntent{
		FirewallRule: request.GetFirewallRule(),
	}
}

// LoadFirewallRuleIntent loads a firewall rule intent from cache.
func LoadFirewallRuleIntent(loader intent.Loader, query intent.Query) *FirewallRuleIntent {
	intent := loader.Load(models.KindFirewallRule, query)
	frIntent, ok := intent.(*FirewallRuleIntent)
	if ok == false {
		log.Warning("Cannot cast intent to Firewall Rule Intent")
	}
	return frIntent
}

// CreateFirewallRule evaluates FirewallRule dependencies.
func (s *Service) CreateFirewallRule(
	ctx context.Context,
	request *services.CreateFirewallRuleRequest,
) (*services.CreateFirewallRuleResponse, error) {

	i := NewFirewallRuleIntent(ctx, s.ReadService, request)

	err := s.storeAndEvaluate(ctx, i)
	if err != nil {
		return nil, err
	}

	return s.BaseService.CreateFirewallRule(ctx, request)
}

// UpdateFirewallRule evaluates UpdateFirewallRule dependencies.
func (s *Service) UpdateFirewallRule(
	ctx context.Context,
	request *services.UpdateFirewallRuleRequest,
) (*services.UpdateFirewallRuleResponse, error) {
	fr := request.GetFirewallRule()
	if fr == nil {
		return nil, errors.New("failed to update Firewall Rule." +
			" Firewall Rule Request needs to contain resource!")
	}

	// TODO: Handle update

	i := LoadFirewallRuleIntent(s.cache, intent.ByUUID(fr.GetUUID()))
	if i == nil {
		return nil, errors.Errorf("cannot load intent for Firewall Rule %v", fr.GetUUID())
	}

	i.FirewallRule = fr
	if err := s.storeAndEvaluate(ctx, i); err != nil {
		return nil, errors.Wrapf(err, "failed to update intent for Firewall Rule :%v", fr.GetUUID())
	}

	return s.BaseService.UpdateFirewallRule(ctx, request)
}
