package services

//Chain setup chain of services.
func Chain(services []Service) {
	if len(services) < 2 {
		return
	}
	previous := services[0]
	for i := 1; i < len(services); i++ {
		current := services[i]
		previous.SetNext(current)
		previous = current
	}
}

type BaseService struct {
	next Service
}

func (service *BaseService) Next() Service {
	return service.next
}

func (service *BaseService) SetNext(next Service) {
	service.next = next
}
