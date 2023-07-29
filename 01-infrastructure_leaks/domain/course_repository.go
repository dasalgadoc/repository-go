package domain

type CourseRepository interface {
	Save(course Course) error
	Search(id string) (*Course, error)
	Flush(course Course) error
}
