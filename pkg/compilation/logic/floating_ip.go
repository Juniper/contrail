package logic

import (
	"context"

	"github.com/Juniper/contrail/pkg/compilation/intent"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/models/basemodels"
	"github.com/Juniper/contrail/pkg/services"
	"github.com/pkg/errors"
)

// FloatingIPIntent contains Intent Compiler state for FloatingIP
type FloatingIPIntent struct {
	intent.BaseIntent
	*models.FloatingIP
}

// GetObject returns embedded resource object
func (i *FloatingIPIntent) GetObject() basemodels.Object {
	return i.FloatingIP
}

// CreateFloatingIP evaluates FloatingIP dependencies.
func (s *Service) CreateFloatingIP(
	ctx context.Context,
	request *services.CreateFloatingIPRequest,
) (*services.CreateFloatingIPResponse, error) {

	i := &FloatingIPIntent{
		FloatingIP: request.GetFloatingIP(),
	}

	err := s.storeAndEvaluate(ctx, i)
	if err != nil {
		return nil, err
	}

	return s.BaseService.CreateFloatingIP(ctx, request)
}

// UpdateFloatingIP evaluates FloatingIP dependencies.
func (s *Service) UpdateFloatingIP(
	ctx context.Context,
	request *services.UpdateFloatingIPRequest,
) (*services.UpdateFloatingIPResponse, error) {
	fip := request.GetFloatingIP()
	i := LoadFloatingIPIntent(s.cache, fip.GetUUID())
	if i == nil {
		return nil, errors.Errorf("cannot load intent for floating ip: %v", fip.GetUUID())
	}

	i.FloatingIP = fip
	if err := s.storeAndEvaluate(ctx, i); err != nil {
		return nil, err
	}

	//TODO floating-ip version???
	return s.BaseService.UpdateFloatingIP(ctx, request)
}

// LoadFloatingIPIntent loads a floating ip intent from cache.
func LoadFloatingIPIntent(
	loader intent.Loader,
	uuid string,
) *FloatingIPIntent {
	i := loader.Load(models.KindFloatingIP, intent.ByUUID(uuid))
	actual, _ := i.(*FloatingIPIntent) //nolint: errcheck
	return actual
}
