package domain

import "dasalgadoc.com/repository-go/courses/domain"

type CourseRepository interface {
	Save(course domain.Course) error
	Search(id domain.CourseId) (*domain.Course, error)
	SearchAll() ([]*domain.Course, error)
}
