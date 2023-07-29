package domain

import "time"

type Event interface {
	GetEventName() string
	GetEventId() string
	GetOccurredOn() time.Time
}

type EventCommonFields struct {
	name       string
	id         UUIDValueObject
	occurredOn time.Time
}

func (e *EventCommonFields) GetEventName() string {
	return e.name
}

func (e *EventCommonFields) GetEventId() string {
	return e.id.Value()
}

func (e *EventCommonFields) GetOccurredOn() time.Time {
	return e.occurredOn
}

func NewEventCommonFields(name string) *EventCommonFields {
	uid, err := NewUUIDValueObject()
	if err != nil {
		return nil
	}
	return &EventCommonFields{name: name,
		id:         *uid,
		occurredOn: time.Now(),
	}
}
