package criteria

import "errors"

type Type string

const (
	ASC  Type = "ASC"
	DESC Type = "DESC"
)

type OrderType struct {
	value Type
}

func (o *OrderType) Value() string {
	return string(o.value)
}

func isValidOrderType(value OrderType) bool {
	switch value {
	case ASC, DESC:
		return true
	}
	return false
}

func NewOrderType(value string) (*OrderType, error) {
	if value == "" {
		value = string(ASC)
	}

	orderType := Type(value)
	if !isValidOrderType(orderType) {
		return nil, errors.New("invalid order type")
	}

	return &OrderType{value: orderType}, nil
}
