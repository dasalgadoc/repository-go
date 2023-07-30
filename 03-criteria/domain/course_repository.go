package domain

import (
	"dasalgadoc.com/repository-go/courses/domain"
	shared "dasalgadoc.com/repository-go/shared/domain/criteria"
)

type CourseRepository interface {
	Save(course domain.Course) error
	Search(id domain.CourseId) (*domain.Course, error)
	SearchAll() ([]*domain.Course, error)

	// SearchBy ...
	matching(criteria shared.Criteria) ([]*domain.Course, error)
}
