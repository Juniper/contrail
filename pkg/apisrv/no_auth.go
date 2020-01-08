package apisrv

import (
	"context"

	"github.com/Juniper/contrail/pkg/apisrv/baseapisrv"
	"github.com/Juniper/contrail/pkg/auth"
	"github.com/labstack/echo"
	"google.golang.org/grpc"
)

type noAuthPlugin struct{}

func (p noAuthPlugin) RegisterHTTPAPI(r baseapisrv.HTTPRouter) error {
	r.Use(p.middleware)
	return nil
}

func (noAuthPlugin) middleware(next baseapisrv.HandlerFunc) baseapisrv.HandlerFunc {
	return func(c echo.Context) error {
		r := c.Request()
		ctx := auth.NoAuth(r.Context())
		newRequest := r.WithContext(ctx)
		c.SetRequest(newRequest)
		return next(c)
	}
}

func (p noAuthPlugin) RegisterGRPCAPI(r baseapisrv.GRPCRouter) error {
	r.AddServerOptions(grpc.UnaryInterceptor(p.interceptor))
	return nil
}

func (p noAuthPlugin) interceptor(
	ctx context.Context, req interface{},
	info *grpc.UnaryServerInfo, handler grpc.UnaryHandler,
) (interface{}, error) {
	newCtx := auth.NoAuth(ctx)
	return handler(newCtx, req)
}
