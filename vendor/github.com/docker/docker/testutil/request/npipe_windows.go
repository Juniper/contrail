package request

import (
	"net"
	"time"

	winio "github.com/Microsoft/go-winio"
)

func npipeDial(path string, timeout time.Duration) (net.Conn, error) {
	return winio.DialPipe(path, &timeout)
}
