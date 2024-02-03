package entity

import (
	"fmt"
	"time"
)

type Status = string

const (
	ORDER_STATUS_CREATED Status = "CREATED"
)

type Order struct {
	ID         string
	TotalPrice int32
	DeliveryID string
	Status     Status
	DeletedAt  time.Time
	OrderItems []OrderItem
	Delivery
}

type OrderItem struct {
	ID           string
	ProductID    string
	Quantity     int32
	ProductPrice int32
	OrderID      string
	DeletedAt    time.Time
}

type Delivery struct {
	ID            string
	CEP           string
	Address       string
	Number        string
	Country       string
	District      string
	City          string
	DeliveryPrice int32
	DeletedAt     time.Time
}

func (o *Order) Calculate() (err error) {
	var priceProduct int32

	for i := range o.OrderItems {
		if o.OrderItems[i].ProductPrice == 0 {
			err = fmt.Errorf("Price product not valid, by product ID %s", o.OrderItems[i].ProductID)
			return
		}
		priceProduct += o.OrderItems[i].ProductPrice
	}

	o.TotalPrice = (o.DeliveryPrice + priceProduct)
	return
}
