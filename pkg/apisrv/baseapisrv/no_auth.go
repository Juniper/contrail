package baseapisrv

import (
	"context"

	"github.com/Juniper/asf/pkg/auth"
	"github.com/labstack/echo"
	"google.golang.org/grpc"
)

type noAuthPlugin struct{}

func (p noAuthPlugin) RegisterHTTPAPI(r HTTPRouter) {
	r.Use(p.middleware)
}

func (noAuthPlugin) middleware(next HandlerFunc) HandlerFunc {
	return func(c echo.Context) error {
		r := c.Request()
		ctx := auth.NoAuth(r.Context())
		newRequest := r.WithContext(ctx)
		c.SetRequest(newRequest)
		return next(c)
	}
}

func (p noAuthPlugin) RegisterGRPCAPI(r GRPCRouter) {
	r.AddServerOptions(grpc.UnaryInterceptor(p.interceptor))
}

func (p noAuthPlugin) interceptor(
	ctx context.Context, req interface{},
	info *grpc.UnaryServerInfo, handler grpc.UnaryHandler,
) (interface{}, error) {
	newCtx := auth.NoAuth(ctx)
	return handler(newCtx, req)
}
