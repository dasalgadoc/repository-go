package application

import (
	"dasalgadoc.com/repository-go/01-infrastructure_leaks/domain"
	shared "dasalgadoc.com/repository-go/shared/domain"
)

type CourseCreator struct {
	repository domain.CourseRepository
	bus        shared.EventBus
}

func NewCourseCreator(
	repository domain.CourseRepository, bus shared.EventBus) *CourseCreator {
	return &CourseCreator{
		repository: repository,
		bus:        bus,
	}
}

func (c *CourseCreator) Invoke(name string) error {
	course, err := domain.CreateCourse(name)
	if err != nil {
		return err
	}

	err = c.repository.Save(*course)
	if err != nil {
		return err
	}

	err = c.repository.Flush(*course)
	if err != nil {
		return err
	}

	return c.bus.PublishAll(course.PullEvents())
}
