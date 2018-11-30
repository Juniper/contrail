package models

import "github.com/Juniper/contrail/pkg/services"

type Resource interface {
	Create(ctx RequestContext, r services.ReadService, w services.WriteService) Response
	Update(ctx RequestContext, r services.ReadService, w services.WriteService) Response
	Delete(ctx RequestContext, r services.ReadService, w services.WriteService) Response
	AddInterface(ctx RequestContext, r services.ReadService, w services.WriteService) Response
	RemoveInterface(ctx RequestContext, r services.ReadService, w services.WriteService) Response
}

type Response interface{}

type BaseResource struct{}

func (b *BaseResource) Create(
	ctx RequestContext,
	r services.ReadService,
	w services.WriteService,
) Response {
	return nil
}

func (b *BaseResource) Update(
	ctx RequestContext,
	r services.ReadService,
	w services.WriteService,
) Response {
	return nil
}

func (b *BaseResource) Delete(
	ctx RequestContext,
	r services.ReadService,
	w services.WriteService,
) Response {
	return nil
}

func (b *BaseResource) AddInterface(
	ctx RequestContext,
	r services.ReadService,
	w services.WriteService,
) Response {
	return nil
}

func (b *BaseResource) RemoveInterface(
	ctx RequestContext,
	r services.ReadService,
	w services.WriteService,
) Response {
	return nil
}
