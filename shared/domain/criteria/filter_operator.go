package criteria

import "errors"

type Operator string

const (
	EQUAL    Operator = "="
	GT       Operator = ">"
	GTE      Operator = ">="
	LT       Operator = "<"
	LTE      Operator = "<="
	CONTAINS Operator = "CONTAINS"
)

type FilterOperator struct {
	value Operator
}

func (f *FilterOperator) Value() string {
	return string(f.value)
}

func isValidOperator(operator Operator) bool {
	switch operator {
	case EQUAL, GT, GTE, LT, LTE, CONTAINS:
		return true
	}
	return false
}

func NewFilterOperator(value string) (*FilterOperator, error) {
	operator := Operator(value)
	if !isValidOperator(operator) {
		return nil, errors.New("invalid filter operator")
	}
	return &FilterOperator{value: operator}, nil
}
