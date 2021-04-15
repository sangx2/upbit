package service

import (
	"encoding/json"
	"io"
)

// ApiKey : api 키
type ApiKey struct {
	AccessKey string `json:"access_key"`
	ExpireAt  string `json:"expire_at"`
}

func ApiKeysFromJSON(r io.Reader) ([]*ApiKey, error) {
	var apiKeys []*ApiKey

	e := json.NewDecoder(r).Decode(&apiKeys)

	return apiKeys, e
}
