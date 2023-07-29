package domain

import (
	shared "dasalgadoc.com/repository-go/shared/domain"
	"time"
)

const EVENT_NAME = "dasalgadoc.com.repository.1.course.created"

type CourseCreated struct {
	fields   *shared.EventCommonFields
	CourseId string
}

func (c *CourseCreated) GetEventName() string {
	return c.fields.GetEventName()
}

func (c *CourseCreated) GetEventId() string {
	return c.fields.GetEventId()
}

func (c *CourseCreated) GetOccurredOn() time.Time {
	return c.fields.GetOccurredOn()
}

func NewCourseCreatedEvent(courseId string) *CourseCreated {
	return &CourseCreated{
		fields:   shared.NewEventCommonFields(EVENT_NAME),
		CourseId: courseId,
	}
}
