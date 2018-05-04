package sink

import "github.com/Juniper/contrail/pkg/db"

// Sink represents service that handler transfers data to.
type Sink interface {
	Create(resourceName string, pk string, obj db.Object) error
	Update(resourceName string, pk string, obj db.Object) error
	Delete(resourceName string, pk string) error
}
