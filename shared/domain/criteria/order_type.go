package criteria

import "errors"

type AscDesType string

const (
	ASC  AscDesType = "ASC"
	DESC AscDesType = "DESC"
)

type OrderType struct {
	value AscDesType
}

func (o *OrderType) Value() string {
	return string(o.value)
}

func isValidOrderType(value AscDesType) bool {
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

	orderType := AscDesType(value)
	if !isValidOrderType(orderType) {
		return nil, errors.New("invalid order type")
	}

	return &OrderType{value: orderType}, nil
}
