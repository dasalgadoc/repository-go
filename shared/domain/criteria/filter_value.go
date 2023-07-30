package criteria

type FilterValue struct {
	value string
}

func (f *FilterValue) Value() string {
	return f.value
}

func NewFilterValue(value string) *FilterValue {
	return &FilterValue{value: value}
}
