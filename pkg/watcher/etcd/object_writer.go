package etcd

import "github.com/Juniper/contrail/pkg/db"

type ObjectWriter struct {
	s Sink
}

func NewObjectWriter(s Sink) *ObjectWriter {
	return &ObjectWriter{s: s}
}

func (s *ObjectWriter) WriteObject(schemaID string, objUUID string, obj db.Object) error {
	return s.s.Create(schemaID, objUUID, obj)
}
