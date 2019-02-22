package collector

import (
	"context"
	"fmt"

	"github.com/sirupsen/logrus"
)

const (
	// KeyVNCAPIConfigLogMetadata is metadata for VncApiConfigLog message
	KeyVNCAPIConfigLogMetadata = "VNCAPIConfigLogMetadata"
	// KeyVNCAPIConfigLogOperation is operation for VncApiConfigLog message
	KeyVNCAPIConfigLogOperation = "VNCAPIConfigLogOperation"
	// KeyVNCAPIConfigLogError is error string for VncApiConfigLog message
	KeyVNCAPIConfigLogError = "VNCAPIConfigLogError"
)

var collectorKey interface{} = "collectorKey"

// WithContext returns new context with stored collector
func WithContext(ctx context.Context, c Collector) context.Context {
	return context.WithValue(ctx, collectorKey, c)
}

// FromContext returns stored collector in context
func FromContext(ctx context.Context) Collector {
	value := ctx.Value(collectorKey)
	if value == nil {
		return nil
	}

	c, ok := value.(Collector)
	if ok != true {
		logrus.Error(
			fmt.Errorf("there is object of type '%T' in context for key '%s'",
				value, collectorKey),
		)
		return nil
	}

	return c
}
