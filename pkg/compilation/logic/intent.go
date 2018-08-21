package logic

import (
	"context"

	"github.com/Juniper/contrail/pkg/services"
)

type evaluateContext struct {
	WriteService services.WriteService
}

type intent interface {
	evaluate(ctx context.Context, evaluateCtx *evaluateContext) error
}

type baseIntent struct {
}

func (b *baseIntent) evaluate(ctx context.Context, evaluateCtx *evaluateContext) error {
	return nil
}
