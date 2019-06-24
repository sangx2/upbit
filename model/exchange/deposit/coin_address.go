package deposit

import (
	"encoding/json"
	"io"
)

// CoinAddress : 입금 주소 정보
type CoinAddress struct {
	Currency         string `json:"currency"`
	DepositAddress   string `json:"deposit_address"`
	SecondaryAddress string `json:"secondary_address"`
}

func CoinAddressFromJSON(r io.Reader) *CoinAddress {
	var c *CoinAddress

	json.NewDecoder(r).Decode(&c)

	return c
}

func CoinAddressesFromJSON(r io.Reader) []*CoinAddress {
	var coinAddresses []*CoinAddress

	json.NewDecoder(r).Decode(&coinAddresses)

	return coinAddresses
}
