package quotation

import (
	"encoding/json"
	"io"
)

// Tick : 체결 정보
type Tick struct {
	Market           string  `json:"market"`
	TradeDateUtc     string  `json:"trade_date_utc"`
	TradeTimeUtc     string  `json:"trade_time_utc"`
	Timestamp        int64   `json:"timestamp"`
	TradePrice       float64 `json:"trade_price"`
	TradeVolume      float64 `json:"trade_volume"`
	PrevClosingPrice float64 `json:"prev_closing_price"`
	ChangePrice      float64 `json:"chane_price"`
	AskBid           string  `json:"ask_bid"`
	SequentialID     int64   `json:"sequential_id"`
}

func TicksFromJSON(r io.Reader) []*Tick {
	var ticks []*Tick

	json.NewDecoder(r).Decode(&ticks)

	return ticks
}
