package service

import (
	"encoding/json"
	"io"
)

type Wallet struct {
	Currency    string `json:"currency"`
	WalletState string `json:"wallet_state"`

	BlockState     string `json:"block_state"`
	BlockHeight    int64  `json:"block_height"`
	BlockUpdatedAt string `json:"block_updated_at"`
}

func WalletsFromJSON(r io.Reader) ([]*Wallet, error) {
	var wallets []*Wallet

	e := json.NewDecoder(r).Decode(&wallets)

	return wallets, e
}
