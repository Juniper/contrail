package apisrv

import (
	"encoding/json"
	"time"

	"github.com/Juniper/contrail/pkg/common"
	"github.com/labstack/echo"
	"golang.org/x/net/websocket"
)

const (
	watchTimeout = 10 * time.Second
)

func (s *Server) watchHandler(c echo.Context) error {
	ctx := c.Request().Context()
	auth := common.GetAuthCTX(ctx)
	if !auth.IsAdmin() {
		return common.ErrorPermissionDenied
	}
	websocket.Handler(func(ws *websocket.Conn) {
		defer ws.Close()
		watcher := s.Cache.AddWatcher(ctx, 0, watchTimeout)
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
