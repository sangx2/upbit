package upbit

import (
	"github.com/sangx2/upbit/model"
	"github.com/sangx2/upbit/model/exchange/service"
)

// GetWalletStatus 입출금 현황
//
// [HEADERS]
//
// Authorization : REQUIRED. Authorization token(JWT)
func (u *Upbit) GetWalletStatus() ([]*service.Wallet, *model.Remaining, error) {
	api, e := GetApiInfo(FuncGetWalletStatus)
	if e != nil {
		return nil, nil, e
	}

	req, e := u.createRequest(api.Method, BaseURI+api.Url, nil, api.Section)
	if e != nil {
		return nil, nil, e
	}

	resp, e := u.do(req, api.Group)
	if e != nil {
		return nil, nil, e
	}
	defer resp.Body.Close()

	wallets, e := service.WalletsFromJSON(resp.Body)
	if e != nil {
		return nil, nil, e
	}

	return wallets, model.RemainingFromHeader(resp.Header), nil
}

// GetApiKeys API 키 리스트 조회
//
// [HEADERS]
//
// Authorization : REQUIRED. Authorization token(JWT)
func (u *Upbit) GetApiKeys() ([]*service.ApiKey, *model.Remaining, error) {
	api, e := GetApiInfo(FuncGetApiKeys)
	if e != nil {
		return nil, nil, e
	}

	req, e := u.createRequest(api.Method, BaseURI+api.Url, nil, api.Section)
	if e != nil {
		return nil, nil, e
	}

	resp, e := u.do(req, api.Group)
	if e != nil {
		return nil, nil, e
	}
	defer resp.Body.Close()

	apiKeys, e := service.ApiKeysFromJSON(resp.Body)
	if e != nil {
		return nil, nil, e
	}

	return apiKeys, model.RemainingFromHeader(resp.Header), nil
}
