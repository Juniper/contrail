package sink

import "github.com/Juniper/contrail/pkg/db"

// ObjectWriter allows writing database objects to Sink.
type ObjectWriter struct {
	s Sink
}

// NewObjectWriter creates new object writer with provided Sink.
func NewObjectWriter(s Sink) *ObjectWriter {
	return &ObjectWriter{s: s}
}

// WriteObject handles database object by writing it to Sink.
func (o *ObjectWriter) WriteObject(schemaID string, objUUID string, obj db.Object) error {
	return o.s.Create(schemaID, objUUID, obj)
}
