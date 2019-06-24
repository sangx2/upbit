package deposit

import (
	"encoding/json"
	"io"
)

// Deposit : 입금 정보
type Deposit struct {
	Type            string `json:"type"`
	UUID            string `json:"uuid"`
	Currency        string `json:"currency"`
	TxID            string `json:"txid"`
	State           string `json:"state"`
	CreatedAt       string `json:"create_at"`
	DoneAt          string `json:"done_at"`
	Amount          string `json:"amount"`
	Fee             string `json:"fee"`
	TransactionType string `json:"transaction_type"`
}

func DepositFromJSON(r io.Reader) *Deposit {
	var d *Deposit

	json.NewDecoder(r).Decode(d)

	return d
}

func DepositsFromJSON(r io.Reader) []*Deposit {
	var deposits []*Deposit

	json.NewDecoder(r).Decode(&deposits)

	return deposits
}
