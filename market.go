package upbit

import (
	"github.com/sangx2/upbit/model"
	"github.com/sangx2/upbit/model/quotation"
)

// GetMarkets 마켓 코드 조회
func (u *Upbit) GetMarkets() ([]*quotation.Market, *model.Remaining, error) {
	api, e := GetApiInfo(FuncGetMarkets)
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

	markets, e := quotation.MarketsFromJSON(resp.Body)
	if e != nil {
		return nil, nil, e
	}
	return markets, model.RemainingFromHeader(resp.Header), nil
}
