package criteria

type Filters []Filter

type Filter struct {
	field    FilterField
	value    FilterValue
	operator FilterOperator
}

func (f *Filter) Field() FilterField {
	return f.field
}

func (f *Filter) Value() FilterValue {
	return f.value
}

func (f *Filter) Operator() FilterOperator {
	return f.operator
}
