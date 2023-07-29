package domain

type EventBus interface {
	Publish(event Event) error
	PublishAll(events []Event) error
}
