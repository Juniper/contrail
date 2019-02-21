package analytics

import (
	"context"

	"github.com/Juniper/contrail/pkg/collector"
	"github.com/Juniper/contrail/pkg/db/etcd"
	"github.com/Juniper/contrail/pkg/models"
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

func (p *payloadMessageBusNotifyTrace) Build() *collector.Message {
	return &collector.Message{
		SandeshType: typeMessageBusNotifyTrace,
		Payload:     p,
	}
}

// MessageBusNotifyTrace sends message of MessageBusNotifyTrace type
func MessageBusNotifyTrace(ctx context.Context, operation string, obj basemodels.Object) collector.MessageBuilder {
	/* TODO: Should be reverted as introspect service for Intent API will be introduced.
	requestID := services.GetRequestID(ctx)
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
	return collector.NewEmptyMessageBuilder()
}

type payloadContrailConfigTrace struct {
	Table    string      `json:"table"`
	Name     string      `json:"name"`
	Elements interface{} `json:"elements"`
	Deleted  bool        `json:"deleted"`
}

func (p *payloadContrailConfigTrace) Build() *collector.Message {
	return &collector.Message{
		SandeshType: typeContrailConfigTrace,
		Payload:     p,
	}
}

type uveTableInfo struct {
	TableName string
	IsGlobal  bool
}

var uveTables = map[string]uveTableInfo{
	models.KindAddressGroup:            {TableName: "ObjectAddressGroupTable", IsGlobal: false},
	models.KindAnalyticsNode:           {TableName: "ObjectCollectorInfo", IsGlobal: true},
	models.KindApplicationPolicySet:    {TableName: "ObjectApplicationPolicySetTable", IsGlobal: false},
	models.KindBGPRouter:               {TableName: "ObjectBgpRouter", IsGlobal: true},
	models.KindConfigNode:              {TableName: "ObjectConfigNode", IsGlobal: true},
	models.KindDatabaseNode:            {TableName: "ObjectDatabaseInfo", IsGlobal: true},
	models.KindFirewallPolicy:          {TableName: "ObjectFirewallPolicyTable", IsGlobal: false},
	models.KindFirewallRule:            {TableName: "ObjectFirewallRuleTable", IsGlobal: false},
	models.KindPhysicalRouter:          {TableName: "ObjectPRouter", IsGlobal: true},
	models.KindProject:                 {TableName: "ObjectProjectTable", IsGlobal: false},
	models.KindServiceGroup:            {TableName: "ObjectServiceGroupTable", IsGlobal: false},
	models.KindServiceInstance:         {TableName: "ObjectSITable", IsGlobal: false},
	models.KindTag:                     {TableName: "ObjectTagTable", IsGlobal: false},
	models.KindVirtualMachineInterface: {TableName: "ObjectVMITable", IsGlobal: false},
	models.KindVirtualMachine:          {TableName: "ObjectVMTable", IsGlobal: false},
	models.KindVirtualNetwork:          {TableName: "ObjectVNTable", IsGlobal: false},
	models.KindVirtualRouter:           {TableName: "ObjectVRouter", IsGlobal: true},
	// TODO: Introduce KindServiceChain
	"service-chain": {TableName: "ServiceChain", IsGlobal: false},
}

// ContrailConfigTrace sends message of ContrailConfigTrace type
func ContrailConfigTrace(operation string, obj basemodels.Object) collector.MessageBuilder {
	info, ok := uveTables[obj.Kind()]
	if !ok {
		return collector.NewEmptyMessageBuilder()
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
	collector collector.Collector
}

func (p *processor) Process(ctx context.Context, event *services.Event) (*services.Event, error) {
	p.collector.Send(MessageBusNotifyTrace(ctx, event.Operation(), event.GetResource()))
	p.collector.Send(ContrailConfigTrace(event.Operation(), event.GetResource()))
	return nil, nil
}

// NewMessageBusProcessor runs etcd events watcher
func NewMessageBusProcessor(collector collector.Collector) error {
	v := &processor{
		collector: collector,
	}
	p, err := etcd.NewEventProducer(v, "collector")
	if err != nil {
		return err
	}
	return p.Start(context.Background())
}
