package collector

import (
	"context"
	"strings"

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
	"virtual_network":           uveTableInfo{TableName: "ObjectVNTable", IsGlobal: false},
	"virtual_machine":           uveTableInfo{TableName: "ObjectVMTable", IsGlobal: false},
	"virtual_machine_interface": uveTableInfo{TableName: "ObjectVMITable", IsGlobal: false},
	"service_instance":          uveTableInfo{TableName: "ObjectSITable", IsGlobal: false},
	"virtual_router":            uveTableInfo{TableName: "ObjectVRouter", IsGlobal: true},
	"analytics_node":            uveTableInfo{TableName: "ObjectCollectorInfo", IsGlobal: true},
	"database_node":             uveTableInfo{TableName: "ObjectDatabaseInfo", IsGlobal: true},
	"config_node":               uveTableInfo{TableName: "ObjectConfigNode", IsGlobal: true},
	"service_chain":             uveTableInfo{TableName: "ServiceChain", IsGlobal: false},
	"physical_router":           uveTableInfo{TableName: "ObjectPRouter", IsGlobal: true},
	"bgp_router":                uveTableInfo{TableName: "ObjectBgpRouter", IsGlobal: true},
	"tag":                       uveTableInfo{TableName: "ObjectTagTable", IsGlobal: false},
	"project":                   uveTableInfo{TableName: "ObjectProjectTable", IsGlobal: false},
	"firewall_policy":           uveTableInfo{TableName: "ObjectFirewallPolicyTable", IsGlobal: false},
	"firewall_rule":             uveTableInfo{TableName: "ObjectFirewallRuleTable", IsGlobal: false},
	"address_group":             uveTableInfo{TableName: "ObjectAddressGroupTable", IsGlobal: false},
	"service_group":             uveTableInfo{TableName: "ObjectServiceGroupTable", IsGlobal: false},
	"application_policy_set":    uveTableInfo{TableName: "ObjectApplicationPolicySetTable", IsGlobal: false},
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

	fqName := obj.GetFQName()
	if info.IsGlobal && len(fqName) > 0 {
		result.Name = fqName[len(fqName)-1]
	} else {
		result.Name = strings.Join(fqName, ":")
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
