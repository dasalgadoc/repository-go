# 📎 Criteria Pattern
[_aka Specification Pattern_](https://en.wikipedia.org/wiki/Specification_pattern)

Criteria (or specification) pattern allows to build any kind of queries without a method explosion, via an object that specifies limits, offsets, filters, and order in a collection of query results is produced.

If a criteria is used with a repository, the repository will be able to build a query based on the criteria and return the results.
Concrete implementations must translate the criteria into a query that the repository can understand.

## 💠 This code - Example

Generate a method that inputs a criteria and outputs a list of courses (that meet the criteria).

```go
type CourseRepository interface {
    Save(course domain.Course) error
    Search(id domain.CourseId) (*domain.Course, error)
    SearchAll() ([]*domain.Course, error)

    // SearchBy ...
    matching(criteria shared.Criteria) ([]*domain.Course, error)
}
```

Criteria is a object that custom a query (limits, offsets, order by and filters) 

```go
type Criteria struct {
    offset  int
    limit   int
    order   Order
    filters Filters
}
```

This struct a its components are in the shared package, so it can be used by any package and never applies concrete business logic.

A concrete implementation must convert this object to a valid query.
