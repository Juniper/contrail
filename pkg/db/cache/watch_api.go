package cache

import (
	"encoding/json"

	"github.com/Juniper/asf/pkg/auth"
	"github.com/Juniper/asf/pkg/errutil"
	"github.com/Juniper/contrail/pkg/apisrv/baseapisrv"
	"github.com/labstack/echo"
	"golang.org/x/net/websocket"
)

// TODO Move this to a template in asf

// WatchPath is the path of the watch HTTP endpoint.
const WatchPath = "watch"

// RegisterHTTPAPI registers the watch HTTP endpoint.
func (cache *DB) RegisterHTTPAPI(r baseapisrv.HTTPRouter) error {
	r.GET(WatchPath, cache.watchHandler)
	return nil
}

// RegisterGRPCAPI does nothing, as there is no GRPC watch API.
func (cache *DB) RegisterGRPCAPI(r baseapisrv.GRPCRouter) error {
	return nil
}

func (cache *DB) watchHandler(c echo.Context) error {
	ctx := c.Request().Context()
	authCtx := auth.GetIdentity(ctx)
	if !authCtx.IsAdmin() {
		return errutil.ErrorPermissionDenied
	}
	websocket.Handler(func(ws *websocket.Conn) {
		defer closeConnection(ws, c.Logger())
		watcher, err := cache.AddWatcher(ctx, 0)
		if err != nil {
			errorJSON, _ := json.Marshal(map[string]interface{}{ // nolint: errcheck
				"error": err,
			})
			if sErr := websocket.Message.Send(ws, string(errorJSON)); sErr != nil {
				c.Logger().Errorf("Sending websocket error message (%v) failed: %v", err, sErr)
			}
			return
		}
		for {
			select {
			case e := <-watcher.Chan():
				update, err := json.Marshal(e)
				if err != nil {
					c.Logger().Error(err)
				}
				err = websocket.Message.Send(ws, string(update))
				if err != nil {
					c.Logger().Error(err)
				}
			case <-ctx.Done():
				return
			}
		}
	}).ServeHTTP(c.Response(), c.Request())
	return nil
}

func closeConnection(ws *websocket.Conn, l echo.Logger) { // nolint: interfacer
	if err := ws.Close(); err != nil {
		l.Errorf("Closing websocket connection failed: %v", err)
	}
}
