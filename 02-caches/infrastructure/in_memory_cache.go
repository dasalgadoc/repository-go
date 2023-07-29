package infrastructure

import (
	domain2 "dasalgadoc.com/repository-go/02-caches/domain"
	"dasalgadoc.com/repository-go/courses/domain"
	"errors"
)

type InMemoryCache struct {
	repository domain2.CourseRepository
	courses    map[domain.CourseId]*domain.Course
}

func (m *InMemoryCache) Search(id domain.CourseId) (*domain.Course, error) {
	course, ok := m.courses[id]
	if !ok {
		return nil, errors.New("course not found")
	}
	return course, nil
}

func (m *InMemoryCache) SearchAll() ([]*domain.Course, error) {
	var courses []*domain.Course
	for _, course := range m.courses {
		courses = append(courses, course)
	}
	return courses, nil
}

func (m *InMemoryCache) Save(course domain.Course) error {
	err := m.repository.Save(course)
	if err != nil {
		return err
	}

	courseId, err := domain.NewCourseIdFromString(course.Id())
	if err != nil {
		return err
	}
	m.courses[*courseId] = &course

	return nil
}

func NewInMemoryCache(repository domain2.CourseRepository) *InMemoryCache {
	return &InMemoryCache{
		repository: repository,
		courses:    make(map[domain.CourseId]*domain.Course),
	}
}
