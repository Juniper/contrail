package event

//Consumer consumes events
type Consumer interface {
	Consume(e Event) (err error)
}

//Producer produces event and send it to consumer.
type Producer interface {
	SetConsumer(ec Consumer)
}

//Consumers is list of consumers.
type Consumers []Consumer

//Consume dispatch event for each consumers.
func (cs Consumers) Consume(e Event) (err error) {
	for _, c := range cs {
		err = c.Consume(e)
		if err != nil {
			return err
		}
	}
	return nil
}
