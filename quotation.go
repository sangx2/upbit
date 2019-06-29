package upbit

import (
	"errors"
	"net/url"

	"github.com/sangx2/upbit/model"
	"github.com/sangx2/upbit/model/quotation"
)

// GetMarkets 마켓 코드 조회
func (u *Upbit) GetMarkets() ([]*quotation.Market, *model.Remaining, error) {
	resp, e := u.getResponse(METHOD_GET, URL_UPBIT_V1+"/market/all", nil, API_TYPE_QUOTATION, API_GROUP_MARKET)
	if e != nil {
		return nil, nil, e
	}
	defer resp.Body.Close()

	markets := quotation.MarketsFromJSON(resp.Body)
	if markets == nil {
		return nil, nil, errors.New("quotation.MarketsFromJSON is nil")
	}
	return markets, model.RemainingFromHeader(resp.Header), nil
}

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
		return nil, nil, errors.New("market length is 0")
	}

	switch unit {
	case "1", "3", "5", "10", "15", "30", "60", "240":
	default:
		return nil, nil, errors.New("invalid unit. valid unit is [1, 3, 5, 10, 15, 30, 60, 240]")
	}

	values := url.Values{
		"market": []string{market},
		"to":     []string{to},
		"count":  []string{count},
	}

	resp, e := u.getResponse(METHOD_GET, URL_UPBIT_V1+"/candles/minutes/"+unit, values, API_TYPE_QUOTATION, API_GROUP_CANDLES)
	if e != nil {
		return nil, nil, e
	}
	defer resp.Body.Close()

	candles := quotation.CandlesFromJSON(resp.Body)
	if candles == nil {
		return nil, nil, errors.New("quotation.CandlesFromJSON is nil")
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
		return nil, nil, errors.New("market length is 0")
	}

	values := url.Values{
		"market":              []string{market},
		"to":                  []string{to},
		"count":               []string{count},
		"convertingPriceUnit": []string{convertingPriceUnit},
	}

	resp, e := u.getResponse(METHOD_GET, URL_UPBIT_V1+"/candles/days", values, API_TYPE_QUOTATION, API_GROUP_CANDLES)
	if e != nil {
		return nil, nil, e
	}
	defer resp.Body.Close()

	candles := quotation.CandlesFromJSON(resp.Body)
	if candles == nil {
		return nil, nil, errors.New("quotation.CandlesFromJSON is nil")
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
		return nil, nil, errors.New("market length is 0")
	}

	values := url.Values{
		"market": []string{market},
		"to":     []string{to},
		"count":  []string{count},
	}

	resp, e := u.getResponse(METHOD_GET, URL_UPBIT_V1+"/candles/weeks", values, API_TYPE_QUOTATION, API_GROUP_CANDLES)
	if e != nil {
		return nil, nil, e
	}
	defer resp.Body.Close()

	candles := quotation.CandlesFromJSON(resp.Body)
	if candles == nil {
		return nil, nil, errors.New("quotation.CandlesFromJSON is nil")
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
		return nil, nil, errors.New("market length is 0")
	}

	values := url.Values{
		"market": []string{market},
		"to":     []string{to},
		"count":  []string{count},
	}

	resp, e := u.getResponse(METHOD_GET, URL_UPBIT_V1+"/candles/months", values, API_TYPE_QUOTATION, API_GROUP_CANDLES)
	if e != nil {
		return nil, nil, e
	}
	defer resp.Body.Close()

	candles := quotation.CandlesFromJSON(resp.Body)
	if candles == nil {
		return nil, nil, errors.New("quotation.CandlesFromJSON is nil")
	}
	return candles, model.RemainingFromHeader(resp.Header), nil
}

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
		return nil, nil, errors.New("market length is 0")
	}

	values := url.Values{
		"market": []string{market},
		"to":     []string{to},
		"count":  []string{count},
		"cursor": []string{cursor},
	}

	resp, e := u.getResponse(METHOD_GET, URL_UPBIT_V1+"/trades/ticks", values, API_TYPE_QUOTATION, API_GROUP_CRIX_TRADES)
	if e != nil {
		return nil, nil, e
	}
	defer resp.Body.Close()

	ticks := quotation.TicksFromJSON(resp.Body)
	if ticks == nil {
		return nil, nil, errors.New("quotation.TicksFromJSON is nil")
	}

	return ticks, model.RemainingFromHeader(resp.Header), nil
}

// GetTickers 현재가 정보. 최대 100개의 정보를 반환
//
// [QUERY PARAMS]
//
// markets : REQUIRED. 마켓 코드 목록 (ex. KRW-BTC,BTC-BCC)
func (u *Upbit) GetTickers(markets []string) ([]*quotation.Ticker, *model.Remaining, error) {
	if len(markets) == 0 || markets == nil {
		return nil, nil, errors.New("invalid markets")
	}

	values := url.Values{
		"markets": markets,
	}

	resp, e := u.getResponse(METHOD_GET, URL_UPBIT_V1+"/ticker", values, API_TYPE_QUOTATION, API_GROUP_TICKER)
	if e != nil {
		return nil, nil, e
	}
	defer resp.Body.Close()

	tickers := quotation.TickersFromJSON(resp.Body)
	if tickers == nil {
		return nil, nil, errors.New("quotation.TickersFromJSON is nil")
	}

	return tickers, model.RemainingFromHeader(resp.Header), nil
}

// GetOrderbooks 호가 정보 조회. 최대 100개의 정보를 반환
//
// [QUERY PARAMS]
//
// markets : REQUIRED. 마켓 코드 목록 (ex. KRW-BTC,KRW-ADA)
func (u *Upbit) GetOrderbooks(markets []string) ([]*quotation.Orderbook, *model.Remaining, error) {
	if len(markets) == 0 || markets == nil {
		return nil, nil, errors.New("invalid marketss")
	}

	values := url.Values{
		"markets": markets,
	}

	resp, e := u.getResponse(METHOD_GET, URL_UPBIT_V1+"/orderbook", values, API_TYPE_QUOTATION, API_GROUP_ORDERBOOK)
	if e != nil {
		return nil, nil, e
	}
	defer resp.Body.Close()

	orderbooks := quotation.OrderbooksFromJSON(resp.Body)
	if orderbooks == nil {
		return nil, nil, errors.New("quotation.OrderbooksFromJSON is nil")
	}

	return orderbooks, model.RemainingFromHeader(resp.Header), nil
}
