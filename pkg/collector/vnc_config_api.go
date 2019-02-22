package collector

import (
	"fmt"

	"github.com/labstack/echo"

	"github.com/Juniper/contrail/pkg/models/basemodels"
)

const typeVNCAPIConfigLog = "VncApiConfigLog"

type payloadVNCAPIConfigLog struct {
	UUID       string `json:"identifier_uuid"`
	ObjectType string `json:"object_type"`
	FQName     string `json:"identifier_name"`
	URL        string `json:"url"`
	Operation  string `json:"operation"`
	UserAgent  string `json:"useragent"`
	RemoteIP   string `json:"remote_ip"`
	Parameters string `json:"params"`
	Body       string `json:"body"`
	Domain     string `json:"domain"`
	Project    string `json:"project"`
	User       string `json:"user"`
	Error      string `json:"error"`
}

func (p *payloadVNCAPIConfigLog) Build() *Message {
	return &Message{
		SandeshType: typeVNCAPIConfigLog,
		Payload:     p,
	}
}

func VNCAPIConfigLog(c echo.Context, _, resBody []byte) MessageBuilder {
	metadata, ok := c.Get("VNCAPIConfigLogMetadata").(*basemodels.Metadata)
	if !ok {
		return NewEmptyMessageBuilder()
	}
	operation, ok := c.Get("VNCAPIConfigLogOperation").(string)
	if !ok {
		return NewEmptyMessageBuilder()
	}
	errorMessage, _ := c.Get("VNCAPIConfigLogError").(string)
	p := &payloadVNCAPIConfigLog{
		UUID:       metadata.UUID,
		ObjectType: metadata.Type,
		FQName:     basemodels.FQNameToString(metadata.FQName),
		URL:        c.Request().URL.String(),
		Operation:  operation,
		UserAgent:  c.Request().Header.Get("X-Contrail-Useragent"),
		RemoteIP:   c.Request().Header.Get("Host"),
		Body:       string(resBody),
		Domain:     c.Request().Header.Get("X-Domain-Name"),
		Project:    c.Request().Header.Get("X-Project-Name"),
		User:       c.Request().Header.Get("X-User-Name"),
		Error:      errorMessage,
	}
	if p.UserAgent == "" {
		p.UserAgent = c.Request().UserAgent()
	}
	fmt.Println(typeVNCAPIConfigLog, p)
	return p
}
