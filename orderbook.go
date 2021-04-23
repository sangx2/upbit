package upbit

import (
	"fmt"
	"github.com/sangx2/upbit/model"
	"github.com/sangx2/upbit/model/quotation"
	"net/url"
)

// GetOrderbooks 호가 정보 조회. 최대 100개의 정보를 반환
//
// [QUERY PARAMS]
//
// markets : REQUIRED. 마켓 코드 목록 (ex. KRW-BTC,KRW-ADA)
func (u *Upbit) GetOrderbooks(markets []string) ([]*quotation.Orderbook, *model.Remaining, error) {
	if len(markets) == 0 || markets == nil {
		return nil, nil, fmt.Errorf("invalid markets")
	}

	api, e := GetApiInfo(FuncGetOrderbooks)
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

	orderbooks, e := quotation.OrderbooksFromJSON(resp.Body)
	if e != nil {
		return nil, nil, e
	}

	return orderbooks, model.RemainingFromHeader(resp.Header), nil
}
