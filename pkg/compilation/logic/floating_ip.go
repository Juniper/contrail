package logic

import (
	"context"
	"net"

	"github.com/Juniper/contrail/pkg/compilation/intent"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/models/basemodels"
	"github.com/Juniper/contrail/pkg/services"
	"github.com/pkg/errors"
)

const (
	ipV4 = 4
	ipV6 = 6
)

// FloatingIPIntent contains Intent Compiler state for FloatingIP.
type FloatingIPIntent struct {
	intent.BaseIntent
	*models.FloatingIP
	ipVersion int64
}

// GetObject returns embedded resource object.
func (i *FloatingIPIntent) GetObject() basemodels.Object {
	return i.FloatingIP
}

// CreateFloatingIP evaluates FloatingIP dependencies.
func (s *Service) CreateFloatingIP(
	ctx context.Context,
	request *services.CreateFloatingIPRequest,
) (*services.CreateFloatingIPResponse, error) {
	fip := request.GetFloatingIP()
	ipVersion, err := getIPVersion(fip.GetFloatingIPAddress())
	if err != nil {
		return nil, err
	}

	i := &FloatingIPIntent{
		FloatingIP: fip,
		ipVersion:  ipVersion,
	}

	err = s.StoreAndEvaluate(ctx, i)
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

	if ipAddress := fip.GetFloatingIPAddress(); ipAddress != "" {
		ipVersion, err := getIPVersion(ipAddress)
		if err != nil {
			return nil, err
		}

		i.ipVersion = ipVersion
	}

	i.FloatingIP = fip
	if err := s.StoreAndEvaluate(ctx, i); err != nil {
		return nil, err
	}

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

// DeleteFloatingIP deletes FloatingIP from cache.
func (s *Service) DeleteFloatingIP(
	ctx context.Context,
	request *services.DeleteFloatingIPRequest,
) (*services.DeleteFloatingIPResponse, error) {

	i := LoadFloatingIPIntent(s.cache, request.GetID())
	if i == nil {
		return nil, errors.New("failed to process FloatingIP deletion: FloatingIPIntent not found in cache")
	}

	s.cache.Delete(models.KindFloatingIP, intent.ByUUID(i.GetUUID()))
	return s.BaseService.DeleteFloatingIP(ctx, request)
}

func getIPVersion(ipAddress string) (int64, error) {
	ip := net.ParseIP(ipAddress)
	if ip == nil {
		return -1, errors.Errorf("Invalid ip address: %s", ipAddress)
	}

	if v := ip.To4(); v != nil {
		return ipV4, nil
	}

	return ipV6, nil
}
