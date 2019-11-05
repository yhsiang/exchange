package models

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestGetOrderID(t *testing.T) {
	order := &Order{
		OrderID:  uuid.New().String(),
		Symbol:   "SBI",
		Side:     "sell",
		Price:    "2275.5",
		Quantity: "100",
	}
	assert.Equal(t, order.OrderID, order.GetOrderId())
}

func TestCreateOrder(t *testing.T) {
	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	order := &Order{
		OrderID:  uuid.New().String(),
		Symbol:   "SBI",
		Side:     "sell",
		Price:    "2275.5",
		Quantity: "100",
	}

	mock.ExpectExec("^INSERT INTO orders").WithArgs(order.OrderID, order.Symbol, order.Side, order.Price, order.Quantity).WillReturnResult(sqlmock.NewResult(1, 1))

	// now we execute our method
	if err = order.CreateOrder(db); err != nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
