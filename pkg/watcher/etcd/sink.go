package etcd

// Sink represents service that handler transfers data to.
type Sink interface {
	Create(resourceName string, pk string, obj interface{}) error // TODO change to proto Message
	Update(resourceName string, pk string, obj interface{}) error // TODO proto message
	Delete(resourceName string, pk string) error
}
