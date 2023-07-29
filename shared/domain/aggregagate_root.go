package domain

type AggregateRoot struct {
	events []Event
}

func NewAggregateRoot() *AggregateRoot {
	return &AggregateRoot{events: make([]Event, 0)}
}

func (a *AggregateRoot) PullEvents() []Event {
	events := a.events
	a.events = make([]Event, 0)

	return events
}

func (a *AggregateRoot) Record(event Event) {
	a.events = append(a.events, event)
}
