package upbit

import (
	"fmt"
	"github.com/sangx2/upbit/model"
	"github.com/sangx2/upbit/model/quotation"
	"net/url"
)

// GetMinuteCandles 분 캔들
//
// [PATH PARAMS]
//
// unit : REQUIRED. 분 단위. 가능한 값 : 1, 3, 5, 15, 10, 30, 60, 240
//
// [QUERY PARAMS]
//
// market : REQUIRED. 마켓 코드 (ex. KRW-BTC, BTC-BCC)
//
// to : 마지막 캔들 시각 (exclusive). 포맷 : yyyy-MM-dd'T'HH:mm:ssXXX. 비워서 요청시 가장 최근 캔들
//
// count : 캔들 개수(최대 200개까지 요청 가능)
func (u *Upbit) GetMinuteCandles(market, to, count, unit string) ([]*quotation.Candle, *model.Remaining, error) {
	if len(market) == 0 {
		return nil, nil, fmt.Errorf("market length is 0")
	}

	switch unit {
	case "1", "3", "5", "10", "15", "30", "60", "240":
	default:
		return nil, nil, fmt.Errorf("invalid unit. valid unit is [1, 3, 5, 10, 15, 30, 60, 240]")
	}

	api, e := GetApiInfo(FuncGetMinuteCandles)
	if e != nil {
		return nil, nil, e
	}

	var values = url.Values{
		"market": []string{market},
		"to":     []string{to},
		"count":  []string{count},
	}

	req, e := u.createRequest(api.Method, BaseURI+api.Url+unit, values, api.Section)
	if e != nil {
		return nil, nil, e
	}

	resp, e := u.do(req, api.Group)
	if e != nil {
		return nil, nil, e
	}
	defer resp.Body.Close()

	candles, e := quotation.CandlesFromJSON(resp.Body)
	if e != nil {
		return nil, nil, e
	}
	return candles, model.RemainingFromHeader(resp.Header), nil
}

// GetDayCandles 일 캔들
//
// [QUERY PARAMS]
//
// market : REQUIRED. 마켓 코드 (ex. KRW-BTC, BTC-BCC)
//
// to : 마지막 캔들 시각 (exclusive). 포맷 : yyyy-MM-dd'T'HH:mm:ssXXX. 비워서 요청시 가장 최근 캔들
//
// count : 캔들 개수
//
// convertingPriceUnit : 종가 환산 화폐 단위 (생략 가능, KRW로 명시할 시 원화 환산 가격을 반환.)
func (u *Upbit) GetDayCandles(market, to, count, convertingPriceUnit string) ([]*quotation.Candle, *model.Remaining, error) {
	if len(market) == 0 {
		return nil, nil, fmt.Errorf("market length is 0")
	}

	api, e := GetApiInfo(FuncGetDayCandles)
	if e != nil {
		return nil, nil, e
	}

	var values = url.Values{
		"market":              []string{market},
		"to":                  []string{to},
		"count":               []string{count},
		"convertingPriceUnit": []string{convertingPriceUnit},
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

	candles, e := quotation.CandlesFromJSON(resp.Body)
	if e != nil {
		return nil, nil, e
	}
	return candles, model.RemainingFromHeader(resp.Header), nil
}

// GetWeekCandles 주 캔들
//
// [QUERY PARAMS]
//
// market : REQUIRED. 마켓 코드 (ex. KRW-BTC, BTC-BCC)
//
// to : 마지막 캔들 시각 (exclusive). 포맷 : yyyy-MM-dd'T'HH:mm:ssXXX. 비워서 요청시 가장 최근 캔들
//
// count : 캔들 개수
func (u *Upbit) GetWeekCandles(market, to, count string) ([]*quotation.Candle, *model.Remaining, error) {
	if len(market) == 0 {
		return nil, nil, fmt.Errorf("market length is 0")
	}

	api, e := GetApiInfo(FuncGetWeekCandles)
	if e != nil {
		return nil, nil, e
	}

	var values = url.Values{
		"market": []string{market},
		"to":     []string{to},
		"count":  []string{count},
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

	candles, e := quotation.CandlesFromJSON(resp.Body)
	if e != nil {
		return nil, nil, e
	}
	return candles, model.RemainingFromHeader(resp.Header), nil
}

// GetMonthCandles 월 캔들
//
// [QUERY PARAMS]
//
// market : REQUIRED. 마켓 코드 (ex. KRW-BTC, BTC-BCC)
//
// to : 마지막 캔들 시각 (exclusive). 포맷 : yyyy-MM-dd'T'HH:mm:ssXXX. 비워서 요청시 가장 최근 캔들
//
// count : 캔들 개수
func (u *Upbit) GetMonthCandles(market, to, count string) ([]*quotation.Candle, *model.Remaining, error) {
	if len(market) == 0 {
		return nil, nil, fmt.Errorf("market length is 0")
	}

	api, e := GetApiInfo(FuncGetMonthCandles)
	if e != nil {
		return nil, nil, e
	}

	var values = url.Values{
		"market": []string{market},
		"to":     []string{to},
		"count":  []string{count},
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

	candles, e := quotation.CandlesFromJSON(resp.Body)
	if e != nil {
		return nil, nil, e
	}
	return candles, model.RemainingFromHeader(resp.Header), nil
}
