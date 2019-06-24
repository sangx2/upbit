package quotation

import (
	"encoding/json"
	"io"
)

// Orderbook 시세 호가 정보
type Orderbook struct {
	Market         string  `json:"market"`
	Timestamp      int64   `json:"timestamp"`
	TotalAskSize   float64 `json:"total_ask_size"`
	TotalBidSize   float64 `json:"total_bid_size"`
	OrderbookUnits []struct {
		AskPrice float32 `json:"ask_price"`
		BidPrice float64 `json:"bid_price"`
		AskSize  float64 `json:"ask_size"`
		BidSize  float64 `json:"bid_size"`
	} `json:"orderbook_units"`
}

func OrderbooksFromJSON(r io.Reader) []*Orderbook {
	var orderbooks []*Orderbook

	json.NewDecoder(r).Decode(&orderbooks)

	return orderbooks
}
