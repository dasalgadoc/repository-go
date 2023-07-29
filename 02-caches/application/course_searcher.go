package application

import (
	"dasalgadoc.com/repository-go/02-caches/domain"
	domain2 "dasalgadoc.com/repository-go/courses/domain"
)

type CourseSearcher struct {
	repository domain.CourseRepository
	cache      domain.CourseRepository
}

func NewCourseSearcher(repository domain.CourseRepository,
	cache domain.CourseRepository) *CourseSearcher {
	return &CourseSearcher{
		repository: repository,
		cache:      cache,
	}
}

func (c *CourseSearcher) Invoke(id string) (any, error) {
	courseId, err := domain2.NewCourseIdFromString(id)
	if err != nil {
		return nil, err
	}

	course, err := c.cache.Search(*courseId)
	if err != nil {
		course, err = c.repository.Search(*courseId)
		if err != nil {
			return nil, err
		}
	}

	return course, nil
}
