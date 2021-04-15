package upbit

import (
	"fmt"
	"github.com/sangx2/upbit/model"
	"github.com/sangx2/upbit/model/quotation"
	"net/url"
)

// GetTradeTicks 당일 체결 내역
//
// [QUERY PARAMS]
//
// market : REQUIRED. 마켓 코드 (ex. KRW-BTC, BTC-BCC)
//
// to : 마지막 체결 시각. 형식 : [HHmmss 또는 HH:mm:ss]. 비워서 요청시 가장 최근 데이터
//
// count : 체결 개수
//
// cursor : 페이지네이션 커서 (sequentialId)
func (u *Upbit) GetTradeTicks(market, to, count, cursor string) ([]*quotation.Tick, *model.Remaining, error) {
	if len(market) == 0 {
		return nil, nil, fmt.Errorf("market length is 0")
	}

	api, e := GetApiInfo(FuncGetTradeTicks)
	if e != nil {
		return nil, nil, e
	}

	var values = url.Values{
		"market": []string{market},
		"to":     []string{to},
		"count":  []string{count},
		"cursor": []string{cursor},
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

	ticks, e := quotation.TicksFromJSON(resp.Body)
	if e != nil {
		return nil, nil, e
	}

	return ticks, model.RemainingFromHeader(resp.Header), nil
}
