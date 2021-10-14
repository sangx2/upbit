package withdraw

import (
	"encoding/json"
	"io"
)

// Withdraw : 출금 정보
type Withdraw struct {
	Type            string `json:"type"`
	UUID            string `json:"uuid"`
	Currency        string `json:"currency"`
	TxID            string `json:"txid"`
	State           string `json:"state"`
	CreatedAt       string `json:"create_at"`
	DoneAt          string `json:"done_at"`
	Amount          string `json:"amount"`
	Fee             string `json:"fee"`
	KrwAmount       string `json:"krw_amount"`
	TransactionType string `json:"transaction_type"`
}

func WithdrawFromJSON(r io.Reader) (*Withdraw, error) {
	var withdraw *Withdraw

	e := json.NewDecoder(r).Decode(&withdraw)

	return withdraw, e
}

func WithdrawsFromJSON(r io.Reader) ([]*Withdraw, error) {
	var withdraws []*Withdraw

	e := json.NewDecoder(r).Decode(&withdraws)

	return withdraws, e
}
