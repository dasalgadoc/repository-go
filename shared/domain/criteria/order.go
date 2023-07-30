package criteria

type Order struct {
	orderBy   OrderBy
	orderType OrderType
}

func (o *Order) OrderBy() OrderBy {
	return o.orderBy
}

func (o *Order) OrderType() OrderType {
	return o.orderType
}

func NewOrder(orderBy OrderBy, orderType OrderType) *Order {
	return &Order{orderBy: orderBy, orderType: orderType}
}
