package order

import (
	"encoding/json"
	"io"
)

// Order : 주문 정보
type Order struct {
	UUID            string `json:"uuid"`
	Side            string `json:"side"`
	OrdType         string `json:"ord_type"`
	Price           string `json:"price"`
	AvgPrice        string `json:"avg_price"`
	State           string `json:"state"`
	Market          string `json:"market"`
	CreatedAt       string `json:"created_at"`
	Volume          string `json:"volume"`
	RemainingVolume string `json:"remainingVolume"`
	ReservedFee     string `json:"reserved_fee"`
	RemainingFee    string `json:"remaining_fee"`
	PaidFee         string `json:"paid_fee"`
	Locked          string `json:"locked"`
	ExecutedVolume  string `json:"executed_volume"`
	TradeCount      string `json:"trade_count"`
	Trades          []struct {
		Market string `json:"market"`
		UUID   string `json:"uuid"`
		Price  string `json:"price"`
		Volume string `json:"volume"`
		Funds  string `json:"funds"`
		Side   string `json:"side"`
	} `json:"trades,omitempty"`
}

func OrderFromJSON(r io.Reader) *Order {
	var o *Order

	json.NewDecoder(r).Decode(&o)

	return o
}

func OrdersFromJSON(r io.Reader) []*Order {
	var orders []*Order

	json.NewDecoder(r).Decode(&orders)

	return orders
}
