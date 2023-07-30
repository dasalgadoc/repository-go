package criteria

type OrderBy struct {
	value string
}

func (o *OrderBy) Value() string {
	return o.value
}

func NewOrderBy(value string) *OrderBy {
	return &OrderBy{value: value}
}
