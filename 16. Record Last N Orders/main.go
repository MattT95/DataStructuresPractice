package main

import "fmt"

// Record last N order ids in a log
// Implement a data structure to accomplish this, with following API:
// record(order_id):
// 	adds the order id to the log
// get_last(i):
// 	gets the ith last element from the log.
// 	i is guaranteed to be smaller than or equal to N
// You should be as efficient as possible with time and space

func main() {

	orders := Orders{50, nil, -1}

	for i := 0; i <= 72; i++ {
		orders.Record(i)
		fmt.Println(orders.orderHistory)
	}

	fmt.Println(orders.GetLast(2))
}

type Orders struct {
	OrderLimit        int
	orderHistory      []int
	lastOrderPosition int
}

func (orders *Orders) Record(orderId int) {
	if orders.orderHistory == nil {
		orders.orderHistory = make([]int, orders.OrderLimit)
		orders.lastOrderPosition = -1
	}

	orders.lastOrderPosition = (orders.lastOrderPosition + 1) % orders.OrderLimit
	orders.orderHistory[orders.lastOrderPosition] = orderId
}

func (orders *Orders) GetLast(nthLastElement int) int {

	zeroBasedLastElement := nthLastElement - 1
	position := (orders.lastOrderPosition - zeroBasedLastElement) % orders.OrderLimit

	return orders.orderHistory[position]
}
