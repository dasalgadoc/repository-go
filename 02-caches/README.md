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
    	course, err = c.repository.Search(*courseId)
    	if err != nil {
            return nil, err
    	}
    }
    
    return course, nil
}
```

Both implementations use the same interface, so the service doesn't know which repository is being used.
