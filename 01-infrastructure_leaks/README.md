# 🎢 Infrastructure Leaks

A infrastructure leak occurs when an interface discloses implementation details to its consumers. This is a concern since services shouldn't be aware of the specifics of the implementation or how a repository operates.

## 💠 This code - Example

Some [ORMs](https://gorm.io/docs/index.html) use the [unit of work pattern](https://martinfowler.com/eaaCatalog/unitOfWork.html) to handle database transactions. 
The unit of work pattern maintains a list of objects affected by a business transaction and coordinates the writing out of changes and the resolution of concurrency problems.  Flush methods are used to commit the changes in the database.

The methods in a repository interface could reveal details about specific implementation like ORMs, for example:

```go
type CourseRepository interface {
    Save(course Course) error
    Search(id string) (*Course, error)
    Flush(course Course) error // Infrastructure leak
}
```

If we want to be sure that data was persisted in the database, a use case must be invoked and utilize the `Flush` function to commit the unit of work.

```go
package application

func (c *CourseCreator) Invoke(name string) error {
    course, err := domain.CreateCourse(name)
    if err != nil {
    	return err
    }
    
    err = c.repository.Save(*course)
    if err != nil {
    	return err
    }
    
    // Infrastructure leak
    err = c.repository.Flush(*course)
    if err != nil {
    	return err
    }
    
    return c.bus.PublishAll(course.PullEvents())
}
```

This is call structural coupling by a [header interface](https://martinfowler.com/bliki/HeaderInterface.html), because the service layer should not know how the repository works and not specific details of the implementation.

Other forms of structural coupling are:

- naming conventions, for example `Add` method for redis implementation.
- parameter names, for example `_id` for ids in Mongo.
