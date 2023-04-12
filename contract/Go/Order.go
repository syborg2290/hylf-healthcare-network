package healthcare

type Order struct {
	//	OrderDate string `json:"OrderDate"`
	OrderID     string  `json:""`
	OrderTime   string  `json:"OrderTime"`
	OrderAmount float64 `json:"OrderAmount"`
	//	OrderDeliveryTime string  `json:"OrderETD"`
	OrderMedicines []Medicine `json:"OrderMedicines"`
	OrderCompleted bool       `json:"OrderCompleted"`
}

/*func(order *Order) getOrderDate() string{
	return order.OrderDate
}*/
func (order *Order) getOrderTime() string {
	return order.OrderTime
}
func (order *Order) getOrderAmount() float64 {
	return order.OrderAmount
}

/*func (order *Order) getOrderDeliveryTime() string {
	return order.OrderDeliveryTime
}*/
func (order *Order) getOrderMedicines() []Medicine {
	return order.OrderMedicines
}
func (order *Order) getOrderCompleted() bool {
	return order.OrderCompleted
}

/*func(order *Order) setOrderDate(date string){
	order.OrderDate=date
}*/
func (order *Order) setOrderCompleted(status bool) {
	order.OrderCompleted = status
}
func (order *Order) setOrderTime(t string) {
	order.OrderTime = t
}
func (order *Order) setOrderAmount(amount float64) {
	order.OrderAmount = amount
}

/*func (order *Order) setOrderDeliveryTime(t string) {
	order.OrderDeliveryTime = t
}*/
func (order *Order) setOrderMedicines(medicines []Medicine) {
	order.OrderMedicines = medicines
}
