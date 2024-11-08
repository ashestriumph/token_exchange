package main

import (
	"fmt"
	"time"
)

type Order struct {
	Size      float64
	Bid       bool
	Limit     *Limit
	Timestamp int64
}

func NewOrder(bid bool, size float64) *Order {
	return &Order{
		Size:      size,
		Bid:       bid,
		Timestamp: time.Now().Unix(),
	}
}

func (o *Order) String() string {
	return fmt.Sprintf("Order{Size: %.2f, Bid: %t, Timestamp: %d}", o.Size, o.Bid, o.Timestamp)
}

type Limit struct {
	Price       float64 // Price of the limit
	Orders      []*Order
	TotalVolume float64
}

func NewLimit(price float64) *Limit {
	return &Limit{
		Price:  price,
		Orders: []*Order{},
	}
}

func (l *Limit) AddOrder(o *Order) {
	o.Limit = l
	l.Orders = append(l.Orders, o)
	l.TotalVolume += o.Size
}

func (l *Limit) DeleteOrder(o *Order) {
	for i := 0; i < len(l.Orders); i++ {
		if l.Orders[i] == o {
			l.Orders[i] = l.Orders[len(l.Orders)-1]
			l.Orders = l.Orders[:len(l.Orders)-1]
			return
		}
	}

	o.Limit = nil
	l.TotalVolume -= o.Size
}

type OrderBook struct {
	Ask []*Limit
	Bid []*Limit
}
