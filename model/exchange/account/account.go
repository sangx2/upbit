package account

import (
	"encoding/json"
	"io"
)

// Account : 계좌 정보
type Account struct {
	Currency       string `json:"currency"`
	Balance        string `json:"balance"`
	Locked         string `json:"locked"`
	AvgKrwBuyPrice string `json:"avg_krw_buy_price"`
	Modified       bool   `json:"modified"`
}

func AccountsFromJSON(r io.Reader) []*Account {
	var accList []*Account

	json.NewDecoder(r).Decode(&accList)

	return accList
}
