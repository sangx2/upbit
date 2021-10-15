package upbit

import (
	"fmt"
	"github.com/sangx2/upbit/model"
	"github.com/sangx2/upbit/model/exchange"
	"github.com/sangx2/upbit/model/exchange/order"
	"net/url"
)

// GetOrderChance 주문 가능 정보
//
// [QUERY PARAMS]
//
// market : REQUIRED. 마켓 ID
//
// [HEADERS]
//
// Authorization : REQUIRED. Authorization token (JWT)
func (u *Upbit) GetOrderChance(market string) (*order.Chance, *model.Remaining, error) {
	if len(market) == 0 {
		return nil, nil, fmt.Errorf("market length is 0")
	}

	api, e := GetApiInfo(FuncGetOrderChance)
	if e != nil {
		return nil, nil, e
	}

	var values = url.Values{
		"market": []string{market},
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

	chance, e := order.ChanceFromJSON(resp.Body)
	if e != nil {
		return nil, nil, e
	}
	return chance, model.RemainingFromHeader(resp.Header), nil
}

// GetOrder 개별 주문 조회
//
// [QUERY PARAMS]
//
// uuid : 주문 UUID
//
// identifier : 조회용 사용자 지정 값
//
// * uuid 혹은 identifier 둘 중 하나의 값이 반드시 포함되어야 합니다.
//
// [HEADERS]
//
// Authorization : REQUIRED. Authorization token(JWT)
func (u *Upbit) GetOrder(uuid, identifier string) (*order.Order, *model.Remaining, error) {
	if len(uuid) == 0 && len(identifier) == 0 {
		return nil, nil, fmt.Errorf("uuid and identifier length is 0")
	} else if len(uuid) != 0 && len(identifier) != 0 {
		return nil, nil, fmt.Errorf("uuid and identifier length is not 0. you must set only one param")
	}

	api, e := GetApiInfo(FuncGetOrder)
	if e != nil {
		return nil, nil, e
	}

	var values = url.Values{}
	if len(uuid) != 0 {
		values.Add("uuid", uuid)
	}
	if len(identifier) != 0 {
		values.Add("identifier", identifier)
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

	order, e := order.OrderFromJSON(resp.Body)
	if e != nil {
		return nil, nil, e
	}

	return order, model.RemainingFromHeader(resp.Header), nil
}

// GetOrders 주문 리스트 조회
//
// market : Market ID
//
// state : 주문 상태
//
// uuids : 주문 UUID의 목록
//
// identifiers : 주문 identifier의 목록
//
// page : 페이지 수, default: 1
//
// limit : 요청 개수 (1 ~ 100)
//
// orderBy : 정렬 방식
//
// [HEADERS]
//
// Authorization : Authorization token(JWT)
func (u *Upbit) GetOrders(market, state string, states, uuids, identifiers []string, page, limit, orderBy string) ([]*order.Order, *model.Remaining, error) {
	switch state {
	case exchange.ORDER_STATE_WAIT:
	case exchange.ORDER_STATE_DONE:
	case exchange.ORDER_STATE_CANCEL:
	default:
		state = exchange.ORDER_STATE_WAIT
	}

	switch orderBy {
	case exchange.ORDERBY_ASC:
	case exchange.ORDERBY_DESC:
	default:
		orderBy = exchange.ORDERBY_DESC
	}

	api, e := GetApiInfo(FuncGetOrders)
	if e != nil {
		return nil, nil, e
	}

	var values = url.Values{
		"market":      []string{market},
		"state":       []string{state},
		"states":      states,
		"uuids":       uuids,
		"identifiers": identifiers,
		"page":        []string{page},
		"limit":       []string{limit},
		"order_by":    []string{orderBy},
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

	orders, e := order.OrdersFromJSON(resp.Body)
	if e != nil {
		return nil, nil, e
	}

	return orders, model.RemainingFromHeader(resp.Header), nil
}

// PurchaseOrder 매수하기
//
// [BODY PARAMS]
//
// market : REQUIRED. Market ID
//
// side : REQUIRED. 주문 종류
//
// volume : REQUIRED. 주문 수량. 지정가, 시장가 매도 시 필수
//
// price : REQUIRED. 유닛당 주문 가격. 지정가, 시장가 매수 시 필수
// ex) KRW-BTC 마켓에서 1BTC당 1,000 KRW로 거래할 경우, 값은 1000 이 된다.
// ex) KRW-BTC 마켓에서 1BTC당 매도 1호가가 500 KRW 인 경우, 시장가 매수 시 값을 1000으로 세팅하면 2BTC가 매수된다.
// (수수료가 존재하거나 매도 1호가의 수량에 따라 상이할 수 있음)
//
// orderType : REQUIRED. 주문 타입
//
// identifier : 조회용 사용자 지정 값
//
// [HEADERS]
//
// Authorization : REQUIRED. Authorization token(JWT)
func (u *Upbit) PurchaseOrder(market, volume, price, orderType, identifier string) (*order.Order, *model.Remaining, error) {
	if len(market) == 0 {
		return nil, nil, fmt.Errorf("market length is 0")
	}

	if len(price) == 0 {
		return nil, nil, fmt.Errorf("price length is 0")
	}

	switch orderType {
	case exchange.ORDER_TYPE_LIMIT:
	case exchange.ORDER_TYPE_PRICE:
	default:
		return nil, nil, fmt.Errorf("invalid orderType")
	}

	api, e := GetApiInfo(FuncPurchaseOrder)
	if e != nil {
		return nil, nil, e
	}

	var values = url.Values{
		"market":     []string{market},
		"side":       []string{exchange.ORDER_SIDE_BID},
		"volume":     []string{volume},
		"price":      []string{price},
		"ord_type":   []string{orderType},
		"identifier": []string{identifier},
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

	order, e := order.OrderFromJSON(resp.Body)
	if e != nil {
		return nil, nil, e
	}

	return order, model.RemainingFromHeader(resp.Header), nil
}

// SellOrder 매도하기
//
// [BODY PARAMS]
//
// market : REQUIRED. Market ID
//
// side : REQUIRED. 주문 종류
//
// volume : REQUIRED. 주문 수량. 지정가, 시장가 매도 시 필수
//
// price : REQUIRED. 유닛당 주문 가격. 지정가, 시장가 매수 시 필수
// ex) KRW-BTC 마켓에서 1BTC당 1,000 KRW로 거래할 경우, 값은 1000 이 된다.
// ex) KRW-BTC 마켓에서 1BTC당 매도 1호가가 500 KRW 인 경우, 시장가 매수 시 값을 1000으로 세팅하면 2BTC가 매수된다.
// (수수료가 존재하거나 매도 1호가의 수량에 따라 상이할 수 있음)
//
// orderType : REQUIRED. 주문 타입
//
// identifier : 조회용 사용자 지정 값
//
// [HEADERS]
//
// Authorization : REQUIRED. Authorization token(JWT)
func (u *Upbit) SellOrder(market, volume, price, orderType, identifier string) (*order.Order, *model.Remaining, error) {
	if len(market) == 0 {
		return nil, nil, fmt.Errorf("market length is 0")
	}

	if len(volume) == 0 {
		return nil, nil, fmt.Errorf("volume length is 0")
	}

	switch orderType {
	case exchange.ORDER_TYPE_LIMIT:
	case exchange.ORDER_TYPE_MARKET:
	default:
		return nil, nil, fmt.Errorf("invalid orderType")
	}

	api, e := GetApiInfo(FuncSellOrder)
	if e != nil {
		return nil, nil, e
	}

	var values = url.Values{
		"market":     []string{market},
		"side":       []string{exchange.ORDER_SIDE_ASK},
		"volume":     []string{volume},
		"price":      []string{price},
		"ord_type":   []string{orderType},
		"identifier": []string{identifier},
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

	order, e := order.OrderFromJSON(resp.Body)
	if e != nil {
		return nil, nil, e
	}

	return order, model.RemainingFromHeader(resp.Header), nil
}

// CancelOrder 주문 취소 접수
//
// [QUERY PARAMS]
//
// uuid : REQUIRED. 주문 UUID
//
func (u *Upbit) CancelOrder(uuid, identifier string) (*order.Order, *model.Remaining, error) {
	if (len(uuid) + len(identifier)) == 0 {
		return nil, nil, fmt.Errorf("invalid args")
	}

	api, e := GetApiInfo(FuncCancelOrder)
	if e != nil {
		return nil, nil, e
	}

	var values = url.Values{
		"uuid":       []string{uuid},
		"identifier": []string{identifier},
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

	order, e := order.OrderFromJSON(resp.Body)
	if e != nil {
		return nil, nil, e
	}

	return order, model.RemainingFromHeader(resp.Header), nil
}
