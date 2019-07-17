package account

import (
	"encoding/json"
	"io"
)

// Account : 계좌 정보
type Account struct {
	Currency            string `json:"currency,omitempty"`
	Balance             string `json:"balance,omitempty"`
	Locked              string `json:"locked,omitempty"`
	AvgBuyPrice         string `json:"avg_buy_price,omitempty"`
	AvgBuyPriceModified bool   `json:"avg_buy_price_modified,omitempty"`
	UnitCurrency        string `json:"unit_currency,omitempty"`
}

func (a *Account) GetMarketID() string {
	return "KRW" + "-" + a.Currency
}

func AccountsFromJSON(r io.Reader) []*Account {
	var accList []*Account

	json.NewDecoder(r).Decode(&accList)

	return accList
}
