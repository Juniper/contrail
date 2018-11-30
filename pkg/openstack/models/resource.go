package models

import "github.com/Juniper/contrail/pkg/services"

type Resource interface {
	Create(ctx RequestContext, r services.ReadService, w services.WriteService) (Response, error)
	Update(ctx RequestContext, r services.ReadService, w services.WriteService) (Response, error)
	Delete(ctx RequestContext, r services.ReadService, w services.WriteService) (Response, error)
	AddInterface(ctx RequestContext, r services.ReadService, w services.WriteService) (Response, error)
	DeleteInterface(ctx RequestContext, r services.ReadService, w services.WriteService) (Response, error)
}

type Response interface{}

type BaseResource struct{}

func (b *BaseResource) Create(
	ctx RequestContext,
	r services.ReadService,
	w services.WriteService,
) (Response, error) {
	return nil, nil
}

func (b *BaseResource) Update(
	ctx RequestContext,
	r services.ReadService,
	w services.WriteService,
) (Response, error) {
	return nil, nil
}

func (b *BaseResource) Delete(
	ctx RequestContext,
	r services.ReadService,
	w services.WriteService,
) (Response, error) {
	return nil, nil
}

func (b *BaseResource) AddInterface(
	ctx RequestContext,
	r services.ReadService,
	w services.WriteService,
) (Response, error) {
	return nil, nil
}

func (b *BaseResource) DeleteInterface(
	ctx RequestContext,
	r services.ReadService,
	w services.WriteService,
) (Response, error) {
	return nil, nil
}
