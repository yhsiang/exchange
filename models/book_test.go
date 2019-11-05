package models

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestQuery(t *testing.T) {
	db, mock, err := sqlmock.New()

	mockSellRows := sqlmock.NewRows([]string{"price", "quantity"}).
		AddRow("2275", "20")

	mock.ExpectQuery("SELECT price, quantity FROM orders WHERE side = 'sell'").
		WithArgs("SBI").
		WillReturnRows(mockSellRows)

	mockBuyRows := sqlmock.NewRows([]string{"price", "quantity"}).
		AddRow("2272.5", "100").
		AddRow("2271", "80").
		AddRow("2271", "20")

	mock.ExpectQuery("SELECT price, quantity FROM orders WHERE side = 'buy'").
		WithArgs("SBI").
		WillReturnRows(mockBuyRows)

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	book := &OrderBook{
		Symbol: "SBI",
	}

	// now we execute our method
	if err = book.Query(db); err != nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
