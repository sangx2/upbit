package upbit

import (
	"github.com/sangx2/upbit/model"
	"github.com/sangx2/upbit/model/exchange/account"
)

// GetAccounts 전체 계좌 조회
//
// [HEADERS]
//
// Authorization : REQUIRED. Authorization token (JWT)
func (u *Upbit) GetAccounts() ([]*account.Account, *model.Remaining, error) {
	api, e := GetApiInfo(FuncGetAccounts)
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

	accounts, e := account.AccountsFromJSON(resp.Body)
	if e != nil {
		return nil, nil, e
	}

	return accounts, model.RemainingFromHeader(resp.Header), nil
}
