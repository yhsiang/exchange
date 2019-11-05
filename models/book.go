package models

import (
	"database/sql"
	"sort"
	"strconv"
)

//go:generate mockery -name=OrderBookModel
type OrderBookModel interface {
	Query(db *sql.DB) error
}

type Depth struct {
	Depth    int    `json:"depth"`
	Price    string `json:"price"`
	Quantity string `json:"quantity"`
}

type OrderBookSymbol struct {
	Symbol string `json:"symbol"`
}

type OrderBook struct {
	Asks   []Depth `json:"sell"`
	Bids   []Depth `json:"buy"`
	Symbol string  `json:"-"`
}

type ByPrice []Depth

func (a ByPrice) Len() int      { return len(a) }
func (a ByPrice) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByPrice) Less(i, j int) bool {
	ip, _ := strconv.ParseInt(a[i].Price, 10, 64)
	jp, _ := strconv.ParseInt(a[j].Price, 10, 64)
	return ip < jp
}

func aggregate(depths []Depth) []Depth {
	dict := make(map[string]int64)

	for _, depth := range depths {
		quantity, _ := strconv.ParseInt(depth.Quantity, 10, 64)
		if val, ok := dict[depth.Price]; ok {
			dict[depth.Price] = val + quantity
		} else {
			dict[depth.Price] = quantity
		}
	}

	var data []Depth
	for k, v := range dict {
		data = append(data, Depth{
			Price:    k,
			Quantity: strconv.FormatInt(v, 10),
		})
	}

	sort.Sort(ByPrice(data))

	var result []Depth
	for index, val := range data {
		result = append(result, Depth{
			Depth:    index,
			Price:    val.Price,
			Quantity: val.Quantity,
		})
	}
	return result
}

func (b *OrderBook) Query(db *sql.DB) error {
	rows, err := db.Query("SELECT price, quantity FROM orders WHERE side = 'sell' and symbol=?", b.Symbol)

	if err != nil {
		return err
	}

	asks := make([]Depth, 0)
	for rows.Next() {
		var depth Depth
		err = rows.Scan(&depth.Price, &depth.Quantity)
		if err != nil {
			return err
		}
		asks = append(asks, depth)
	}

	rows, err = db.Query("SELECT price, quantity FROM orders WHERE side = 'buy' and symbol=?", b.Symbol)

	if err != nil {
		return err
	}

	bids := make([]Depth, 0)
	for rows.Next() {
		var depth Depth
		err = rows.Scan(&depth.Price, &depth.Quantity)
		if err != nil {
			return err
		}
		bids = append(bids, depth)
	}

	b.Asks = aggregate(asks)
	b.Bids = aggregate(bids)

	if err = rows.Err(); err != nil {
		return err
	}

	return nil
}
