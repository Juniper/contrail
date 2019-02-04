package logic

import (
	"context"
	"net"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	"github.com/Juniper/contrail/pkg/compilation/intent"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/models/basemodels"
	"github.com/Juniper/contrail/pkg/services"
)

const (
	ipV4 = 4
	ipV6 = 6
)

// FloatingIPIntent contains Intent Compiler state for FloatingIP.
type FloatingIPIntent struct {
	intent.BaseIntent
	*models.FloatingIP
	ipVersion int8
}

// LoadFloatingIPIntent loads a floating IP intent from cache.
func LoadFloatingIPIntent(loader intent.Loader, query intent.Query) *FloatingIPIntent {
	intent := loader.Load(models.KindFloatingIP, query)
	fipIntent, ok := intent.(*FloatingIPIntent)
	if ok == false {
		logrus.Warning("Cannot cast intent to Floating IP Intent")
	}
	return fipIntent
}

// GetObject returns embedded resource object.
func (i *FloatingIPIntent) GetObject() basemodels.Object {
	return i.FloatingIP
}

// CreateFloatingIP stores newly created floating ip resource in cache
func (s *Service) CreateFloatingIP(
	ctx context.Context,
	request *services.CreateFloatingIPRequest,
) (*services.CreateFloatingIPResponse, error) {
	i, err := newFloatingIPIntent(request.GetFloatingIP())
	if err != nil {
		return nil, err
	}

	err = s.storeAndEvaluate(ctx, i)
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
	i := LoadFloatingIPIntent(s.cache, intent.ByUUID(fip.GetUUID()))
	if i == nil {
		return nil, errors.Errorf("cannot load intent for floating ip: %v", fip.GetUUID())
	}

	if ipAddress := fip.GetFloatingIPAddress(); ipAddress != "" {
		if err := i.updateIPVersion(ipAddress); err != nil {
			return nil, err
		}
	}

	i.FloatingIP = fip
	if err := s.storeAndEvaluate(ctx, i); err != nil {
		return nil, errors.Wrapf(err, "failed to update intent for Floating IP :%v", fip.GetUUID())
	}

	return s.BaseService.UpdateFloatingIP(ctx, request)
}

// DeleteFloatingIP deletes FloatingIP from cache.
func (s *Service) DeleteFloatingIP(
	ctx context.Context,
	request *services.DeleteFloatingIPRequest,
) (*services.DeleteFloatingIPResponse, error) {

	i := LoadFloatingIPIntent(s.cache, intent.ByUUID(request.GetID()))
	if i == nil {
		return nil, errors.New("failed to process FloatingIP deletion: FloatingIPIntent not found in cache")
	}

	s.cache.Delete(models.KindFloatingIP, intent.ByUUID(i.GetUUID()))
	return s.BaseService.DeleteFloatingIP(ctx, request)
}

func newFloatingIPIntent(fip *models.FloatingIP) (*FloatingIPIntent, error) {
	i := &FloatingIPIntent{
		FloatingIP: fip,
	}

	if err := i.updateIPVersion(fip.GetFloatingIPAddress()); err != nil {
		return nil, err
	}

	return i, nil
}

func (i *FloatingIPIntent) updateIPVersion(ipAddress string) error {
	ipVersion, err := getIPVersion(ipAddress)
	if err != nil {
		return err
	}

	i.ipVersion = ipVersion
	return nil
}

func getIPVersion(ipAddress string) (int8, error) {
	ip := net.ParseIP(ipAddress)
	if ip == nil {
		return -1, errors.Errorf("Invalid ip address: %s", ipAddress)
	}

	if v := ip.To4(); v != nil {
		return ipV4, nil
	}

	return ipV6, nil
}
