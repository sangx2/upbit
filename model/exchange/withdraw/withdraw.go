package withdraw

import (
	"encoding/json"
	"io"
)

// Withdraw : 출금 정보
type Withdraw struct {
	Type      string `json:"type"`
	UUID      string `json:"uuid"`
	Currency  string `json:"currency"`
	TxID      string `json:"txid"`
	State     string `json:"state"`
	CreatedAt string `json:"create_at"`
	DoneAt    string `json:"done_at"`
	Amount    string `json:"amount"`
	Fee       string `json:"fee"`
	KrwAmount string `json:"krw_amount"`
}

func WithdrawFromJSON(r io.Reader) *Withdraw {
	var w *Withdraw

	json.NewDecoder(r).Decode(&w)

	return w
}

func WithdrawsFromJSON(r io.Reader) []*Withdraw {
	var withdraws []*Withdraw

	json.NewDecoder(r).Decode(&withdraws)

	return withdraws
}
