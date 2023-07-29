package domain

import (
	shared "dasalgadoc.com/repository-go/shared/domain"
)

type Course struct {
	aggregate shared.AggregateRoot
	id        CourseId
	name      CourseName
}

func (c *Course) Id() string {
	return c.id.Value()
}

func (c *Course) Name() string {
	return c.name.Value()
}

func (c *Course) PullEvents() []shared.Event {
	return c.aggregate.PullEvents()
}

func CreateCourse(name string) (*Course, error) {
	id, err := NewCourseId()
	if err != nil {
		return nil, err
	}
	studentName, err := NewCourseName(name)
	if err != nil {
		return nil, err
	}

	course := &Course{
		aggregate: *shared.NewAggregateRoot(),
		id:        *id,
		name:      *studentName,
	}

	course.aggregate.Record(NewCourseCreatedEvent(course.id.Value()))

	return course, nil
}

func NewCourseWithId(id string, name string) (*Course, error) {
	courseId, err := NewCourseIdFromString(id)
	if err != nil {
		return nil, err
	}
	courseName, err := NewCourseName(name)
	if err != nil {
		return nil, err
	}

	course := &Course{
		aggregate: *shared.NewAggregateRoot(),
		id:        *courseId,
		name:      *courseName,
	}

	return course, nil
}
