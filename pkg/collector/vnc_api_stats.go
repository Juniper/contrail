package collector

import (
	"net"
	"strconv"
	"time"

	"github.com/labstack/echo"
)

const typeVNCAPIStatsLog = "VncApiStatsLog"

const (
	KeyVNCAPIStatsLogTimeStamp = "VNCAPIStatsLogTimeStamp"
	KeyVNCAPIStatsLogObjectType = "VNCAPIStatsLogObjectType"
)

type payloadVNCAPIStatsLog struct {
	OperationType      string `json:"operation_type"`
	User               string `json:"user"`
	UserAgent          string `json:"useragent"`
	RemoteIP           string `json:"remote_ip"`
	DomainName         string `json:"domain_name"`
	ProjectName        string `json:"project_name"`
	ObjectType         string `json:"object_type"`
	ResponseTimeInUSec int    `json:"response_time_in_usec"`
	ResponseSize       int    `json:"response_size"`
	ResponseCode       string `json:"resp_code"`
	RequestID          string `json:"req_id"`
}

func (p *payloadVNCAPIStatsLog) Build() *Message {
	return &Message{
		SandeshType: typeVNCAPIStatsLog,
		Payload:     p,
	}
}

// VNCAPIStatsLog sends message with type VncApiStatsLog
func VNCAPIStatsLog(c echo.Context, _, resBody []byte) MessageBuilder {
	ts, ok := c.Get(KeyVNCAPIStatsLogTimeStamp).(time.Time)
	if !ok {
		return NewEmptyMessageBuilder()
	}
	objType, ok := c.Get(KeyVNCAPIStatsLogObjectType).(string)
	if !ok {
		return NewEmptyMessageBuilder()
	}
	remoteIP, _, _ := net.SplitHostPort(c.Request().RemoteAddr) // nolint: errcheck
	p := &payloadVNCAPIStatsLog{
		OperationType:      c.Request().Method,
		User:               c.Request().Header.Get("X-User-Name"),
		UserAgent:          c.Request().Header.Get("X-Contrail-Useragent"),
		RemoteIP:           remoteIP,
		DomainName:         c.Request().Header.Get("X-Domain-Name"),
		ProjectName:        c.Request().Header.Get("X-Domain-Name"),
		ObjectType:         objType,
		ResponseTimeInUSec: int(time.Since(ts) / time.Microsecond),
		ResponseSize:       len(resBody),
		ResponseCode:       strconv.Itoa(c.Response().Status),
		RequestID:          c.Request().Header.Get("X-Request-Id"),
	}
	// set default values
	if p.UserAgent == "" {
		p.UserAgent = c.Request().UserAgent()
	}
	if p.DomainName == "" {
		p.DomainName = "default-domain"
	}
	if p.ProjectName == "" {
		p.ProjectName = "default-project"
	}
	return p
}

func VNCAPIStatsLogStamp(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Set(KeyVNCAPIStatsLogTimeStamp, time.Now())
		return next(c)
	}
}
