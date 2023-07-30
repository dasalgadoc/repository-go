package criteria

type FilterField struct {
	value string
}

func (f *FilterField) Value() string {
	return f.value
}

func NewFilterField(value string) *FilterField {
	return &FilterField{value: value}
}
