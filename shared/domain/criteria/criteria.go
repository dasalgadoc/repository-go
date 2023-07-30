package criteria

type Criteria struct {
	offset  int
	limit   int
	order   Order
	filters Filters
}

func (c *Criteria) Offset() int {
	return c.offset
}

func (c *Criteria) Limit() int {
	return c.limit
}

func (c *Criteria) Order() Order {
	return c.order
}

func (c *Criteria) Filters() Filters {
	return c.filters
}

func NewCriteria(offset int, limit int, order Order, filters Filters) *Criteria {
	return &Criteria{offset: offset, limit: limit, order: order, filters: filters}
}
