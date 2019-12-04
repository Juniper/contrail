package apisrv

import (
	"encoding/json"

	"github.com/Juniper/asf/pkg/auth"
	"github.com/Juniper/asf/pkg/errutil"
	"github.com/labstack/echo"
	"golang.org/x/net/websocket"
)

func (s *Server) watchHandler(c echo.Context) error {
	ctx := c.Request().Context()
	authCtx := auth.GetIdentity(ctx)
	if !authCtx.IsAdmin() {
		return errutil.ErrorPermissionDenied
	}
	websocket.Handler(func(ws *websocket.Conn) {
		defer closeConnection(ws, c.Logger())
		watcher, err := s.Cache.AddWatcher(ctx, 0)
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
