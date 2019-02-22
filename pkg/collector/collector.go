package collector

const (
	// KeyVNCAPIConfigLogMetadata is metadata for VncApiConfigLog message
	KeyVNCAPIConfigLogMetadata = "VNCAPIConfigLogMetadata"
	// KeyVNCAPIConfigLogOperation is operation for VncApiConfigLog message
	KeyVNCAPIConfigLogOperation = "VNCAPIConfigLogOperation"
	// KeyVNCAPIConfigLogError is error string for VncApiConfigLog message
	KeyVNCAPIConfigLogError = "VNCAPIConfigLogError"
)

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
