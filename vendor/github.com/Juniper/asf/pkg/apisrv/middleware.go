package apisrv

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"io"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
	"time"

	"github.com/labstack/echo"
	"google.golang.org/grpc"
	// TODO(dfurman): Decouple from below packages
	//"github.com/Juniper/asf/pkg/auth"
)

func removePathPrefixMiddleware(prefix string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			req := c.Request()
			req.URL.Path = strings.TrimPrefix(req.URL.Path, prefix)
			return next(c)
		}
	}
}

func proxyMiddleware(target *url.URL, insecure bool) func(next echo.HandlerFunc) echo.HandlerFunc {
	server := httputil.NewSingleHostReverseProxy(target)
	if target.Scheme == "https" {
		server.Transport = &http.Transport{
			Dial: (&net.Dialer{
				Timeout:   30 * time.Second,
				KeepAlive: 30 * time.Second,
			}).Dial,
			TLSClientConfig:     &tls.Config{InsecureSkipVerify: insecure},
			TLSHandshakeTimeout: 10 * time.Second,
		}
	}
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			r := c.Request()
			w := c.Response()
			server.ServeHTTP(w, r)
			return nil
		}
	}
}

func noAuthMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			r := c.Request()
			ctx := auth.NoAuth(r.Context())
			newRequest := r.WithContext(ctx)
			c.SetRequest(newRequest)
			return next(c)
		}
	}
}

func noAuthInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{},
		info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		newCtx := auth.NoAuth(ctx)
		return handler(newCtx, req)
	}
}

func gRPCMiddleware(grpcServer http.Handler) func(next echo.HandlerFunc) echo.HandlerFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			r := c.Request()
			w := c.Response()
			if r.ProtoMajor == 2 && strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
				grpcServer.ServeHTTP(w, r)
				return nil
			}
			if err := next(c); err != nil {
				c.Error(err)
			}
			return nil
		}
	}
}

type customBinder struct{}

func (*customBinder) Bind(i interface{}, c echo.Context) (err error) {
	rq := c.Request()
	ct := rq.Header.Get(echo.HeaderContentType)
	err = echo.ErrUnsupportedMediaType
	if !strings.HasPrefix(ct, echo.MIMEApplicationJSON) {
		db := new(echo.DefaultBinder)
		return db.Bind(i, c)
	}

	dec := json.NewDecoder(rq.Body)
	dec.UseNumber()
	err = dec.Decode(i)
	if err == io.EOF {
		return nil
	}
	return err
}
