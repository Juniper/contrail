package etcd

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
func (s *ObjectWriter) WriteObject(schemaID string, objUUID string, obj db.Object) error {
	return s.s.Create(schemaID, objUUID, obj)
}
