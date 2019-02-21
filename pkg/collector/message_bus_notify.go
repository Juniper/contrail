package collector

import (
	"context"

	"github.com/Juniper/contrail/pkg/db/etcd"
	"github.com/Juniper/contrail/pkg/models/basemodels"
	"github.com/Juniper/contrail/pkg/services"
)

const (
	typeContrailConfigTrace   = "ContrailConfigTrace"
	typeMessageBusNotifyTrace = "MessageBusNotifyTrace"
)

type payloadMessageBusNotifyTrace struct {
	RequestID string                     `json:"request_id"`
	Operation string                     `json:"operation"`
	Body      *bodyMessageBusNotifyTrace `json:"body"`
	Error     string                     `json:"error"`
}

type bodyMessageBusNotifyTrace struct {
	Operation string      `json:"oper"`
	RequestID string      `json:"request-id"`
	Type      string      `json:"type"`
	UUID      string      `json:"uuid"`
	FQName    []string    `json:"fq_name"`
	ObjDict   interface{} `json:"obj_dict,omitempty"`
}

func (p *payloadMessageBusNotifyTrace) Build() *Message {
	return &Message{
		SandeshType: typeMessageBusNotifyTrace,
		Payload:     p,
	}
}

// MessageBusNotifyTrace sends message of MessageBusNotifyTrace type
func MessageBusNotifyTrace(operation string, obj basemodels.Object) MessageBuilder {
	/* TODO: Should be reverted as introspect service for Intent API will be introduced.
	requestID := "req-" + uuid.NewV4().String()
	var objDict interface{}
	if operation != services.OperationUpdate {
		objDict = obj
	}
	return &payloadMessageBusNotifyTrace{
		RequestID: requestID,
		Operation: operation,
		Body: &bodyMessageBusNotifyTrace{
			Operation: operation,
			RequestID: requestID,
			Type:      basemodels.KindToSchemaID(obj.Kind()),
			UUID:      obj.GetUUID(),
			FQName:    obj.GetFQName(),
			ObjDict:   objDict,
		},
		Error: "",
	}
	*/
	return NewEmptyMessageBuilder()
}

type payloadContrailConfigTrace struct {
	Table    string      `json:"table"`
	Name     string      `json:"name"`
	Elements interface{} `json:"elements"`
	Deleted  bool        `json:"deleted"`
}

func (p *payloadContrailConfigTrace) Build() *Message {
	return &Message{
		SandeshType: typeContrailConfigTrace,
		Payload:     p,
	}
}

type uveTableInfo struct {
	TableName string
	IsGlobal  bool
}

var uveTables = map[string]uveTableInfo{
	"address_group":             {TableName: "ObjectAddressGroupTable", IsGlobal: false},
	"analytics_node":            {TableName: "ObjectCollectorInfo", IsGlobal: true},
	"application_policy_set":    {TableName: "ObjectApplicationPolicySetTable", IsGlobal: false},
	"bgp_router":                {TableName: "ObjectBgpRouter", IsGlobal: true},
	"config_node":               {TableName: "ObjectConfigNode", IsGlobal: true},
	"database_node":             {TableName: "ObjectDatabaseInfo", IsGlobal: true},
	"firewall_policy":           {TableName: "ObjectFirewallPolicyTable", IsGlobal: false},
	"firewall_rule":             {TableName: "ObjectFirewallRuleTable", IsGlobal: false},
	"physical_router":           {TableName: "ObjectPRouter", IsGlobal: true},
	"project":                   {TableName: "ObjectProjectTable", IsGlobal: false},
	"service_chain":             {TableName: "ServiceChain", IsGlobal: false},
	"service_group":             {TableName: "ObjectServiceGroupTable", IsGlobal: false},
	"service_instance":          {TableName: "ObjectSITable", IsGlobal: false},
	"tag":                       {TableName: "ObjectTagTable", IsGlobal: false},
	"virtual_machine_interface": {TableName: "ObjectVMITable", IsGlobal: false},
	"virtual_machine":           {TableName: "ObjectVMTable", IsGlobal: false},
	"virtual_network":           {TableName: "ObjectVNTable", IsGlobal: false},
	"virtual_router":            {TableName: "ObjectVRouter", IsGlobal: true},
}

// ContrailConfigTrace sends message of ContrailConfigTrace type
func ContrailConfigTrace(kind, operation string, obj basemodels.Object) MessageBuilder {
	info, ok := uveTables[kind]
	if !ok {
		return NewEmptyMessageBuilder()
	}

	result := &payloadContrailConfigTrace{
		Table:    info.TableName,
		Elements: obj,
		Deleted:  operation == services.OperationDelete,
	}

	if info.IsGlobal {
		result.Name = basemodels.FQNameToName(obj.GetFQName())
	} else {
		result.Name = basemodels.FQNameToString(obj.GetFQName())
	}

	return result
}

type processor struct {
	collector Collector
}

func (p *processor) Process(ctx context.Context, event *services.Event) (*services.Event, error) {
	p.collector.Send(MessageBusNotifyTrace(event.Operation(), event.GetResource()))
	p.collector.Send(ContrailConfigTrace(event.GetKind(), event.Operation(), event.GetResource()))
	return nil, nil
}

// NewMessageBusProcessor runs etcd events watcher
func NewMessageBusProcessor(collector Collector) error {
	v := &processor{
		collector: collector,
	}
	p, err := etcd.NewEventProducer(v, "collector")
	if err != nil {
		return err
	}
	if err = p.Start(context.Background()); err != nil {
		return err
	}
	return nil
}
