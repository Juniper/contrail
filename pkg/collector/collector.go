package collector

// Message represents a message to Json2Sandesh proxy
type Message struct {
	SandeshType string      `json:"sandesh_type"`
	Payload     interface{} `json:"payload"`
}

// MessageBuilder is interface to build the message
type MessageBuilder interface {
	Build() *Message
}

// Collector represent interface to Json2Sandesh proxy
type Collector interface {
	Send(m MessageBuilder)
}

type emptyMessageBuilder struct{}

// NewEmptyMessageBuilder returns nil message builder
func NewEmptyMessageBuilder() MessageBuilder { return &emptyMessageBuilder{} }

func (*emptyMessageBuilder) Build() *Message { return nil }
