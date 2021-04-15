package upbit

import (
	"fmt"
	"github.com/sangx2/upbit/model"
	"github.com/sangx2/upbit/model/quotation"
	"net/url"
)

// GetTickers 현재가 정보. 최대 100개의 정보를 반환
//
// [QUERY PARAMS]
//
// markets : REQUIRED. 마켓 코드 목록 (ex. KRW-BTC,BTC-BCC)
func (u *Upbit) GetTickers(markets []string) ([]*quotation.Ticker, *model.Remaining, error) {
	if len(markets) == 0 || markets == nil {
		return nil, nil, fmt.Errorf("invalid markets")
	}

	api, e := GetApiInfo(FuncGetTickers)
	if e != nil {
		return nil, nil, e
	}

	var values = url.Values{
		"markets": markets,
	}

	req, e := u.createRequest(api.Method, BaseURI+api.Url, values, api.Section)
	if e != nil {
		return nil, nil, e
	}

	resp, e := u.do(req, api.Group)
	if e != nil {
		return nil, nil, e
	}
	defer resp.Body.Close()

	tickers, e := quotation.TickersFromJSON(resp.Body)
	if e != nil {
		return nil, nil, e
	}

	return tickers, model.RemainingFromHeader(resp.Header), nil
}
