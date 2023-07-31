# 🚧 Repositories and Caches

A cache is a component that stores data so future requests for that data can be served faster. 
The data stored in a cache might be the results of an earlier computation or a copy of data stored elsewhere.

It's possible to set a cache in front of a repository to improve performance and reduce the number of calls to the database as a complement.

## 💠 This code - Example

```go
type CourseSearcher struct {
    repository domain.CourseRepository
    cache      domain.CourseRepository
}
```

A service can have two repositories, one for the cache and another for the database and try to get data from the faster first.


```go
func (c *CourseSearcher) Invoke(id string) (any, error) {
    courseId, err := domain2.NewCourseIdFromString(id)
    if err != nil {
        return nil, err
    }
    
    course, err := c.cache.Search(*courseId)
    if err != nil {
        return nil, err
    }
    
    if course != nil {
        return course, nil
    }
    
    course, err = c.repository.Search(*courseId)
    if err != nil {
        return nil, err
    }

    return course, nil
}
```

Both implementations use the same interface, so the service doesn't know which repository is being used.

## 🖐🏻 But, stop

If the cache implements the same interface as the repository, how about the writing data methods?

There is no limitation that a cache implementation can rely on a repository implementation to write data.

```go
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

```

And some refactor in the cache based on its own repository is also valid.

```go
func (m *InMemoryCache) Search(id domain.CourseId) (*domain.Course, error) {
    course, ok := m.courses[id]
    if !ok {
    	return nil, nil
    }
    return course, nil
}
```

or

```go
func (m *InMemoryCache) Search(id domain.CourseId) (*domain.Course, error) {
    course, ok := m.courses[id]
    if !ok {
    	return m.repository.Search(id)
    }
    return course, nil
}
```

