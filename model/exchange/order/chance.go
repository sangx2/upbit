package order

import (
	"encoding/json"
	"io"

	"github.com/sangx2/upbit/model/exchange/account"
)

// Chance : 주문 가능 정보
type Chance struct {
	BidFee string `json:"bid_fee"`
	AskFee string `json:"ask_fee"`

	Market struct {
		ID         string   `json:"id"`
		Name       string   `json:"name"`
		OrderTypes []string `json:"order_types"`
		OrderSides []string `json:"order_sides"`

		Bid struct {
			Currency  string `json:"currency"`
			PriceUnit string `json:"price_unit"`
			MinTotal  int64  `json:"min_total"`
		} `json:"bid"`

		Ask struct {
			Currency  string `json:"currency"`
			PriceUnit string `json:"price_unit"`
			MinTotal  int64  `json:"min_total"`
		} `json:"ask"`

		MaxTotal string `json:"max_total"`
		State    string `json:"state"`
	} `json:"market"`

	BidAccount account.Account `json:"bid_account"`
	AskAccount account.Account `json:"ask_account"`
}

func ChanceFromJSON(r io.Reader) *Chance {
	var c *Chance

	json.NewDecoder(r).Decode(&c)

	return c
}
