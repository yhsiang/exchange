package models

import (
	"database/sql"
)

//go:generate mockery -name=OrderModel
type OrderModel interface {
	GetOrderId() string
	CreateOrder(db *sql.DB) error
}

type Order struct {
	OrderID  string `json:"order_id"`
	Symbol   string `json:"symbol" binding:"required"`
	Side     string `json:"side"  binding:"required"`
	Price    string `json:"price" binding:"required"`
	Quantity string `json:"quantity" binding:"required"`
}

func (o *Order) GetOrderId() string {
	return o.OrderID
}

func (o *Order) CreateOrder(db *sql.DB) error {
	rs, err := db.Exec("INSERT INTO orders(order_id, symbol, side, price, quantity) VALUES (?, ?, ?, ?, ?)", o.OrderID, o.Symbol, o.Side, o.Price, o.Quantity)
	if err != nil {
		return err
	}

	_, err = rs.LastInsertId()
	if err != nil {
		return err
	}

	return nil
}
