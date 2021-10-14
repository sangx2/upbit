package upbit

import (
	"fmt"
	"github.com/sangx2/upbit/model"
	"github.com/sangx2/upbit/model/exchange"
	"github.com/sangx2/upbit/model/exchange/withdraw"
	"net/url"
	"strconv"
)

// GetWithdraws 출금 리스트 조회
//
// [QUERY PARAMS]
//
// currency : Currency 코드
//
// state : 출금 상태
//
// uuids : 출금 UUID의 목록
//
// txids : 출금 TXID의 목록
//
// limit : 갯수 제한. default: 100, max: 100
//
// page  : 페이지 수, default: 1
//
// orderBy : 정렬
//
// [HEADERS]
//
// Authorization : REQUIRED. Authorization token(JWT)
func (u *Upbit) GetWithdraws(currency, state string, uuids, txids []string, limit, page, orderBy string) ([]*withdraw.Withdraw, *model.Remaining, error) {
	switch state {
	case exchange.WITHDRAW_STATE_SUBMITTING:
	case exchange.WITHDRAW_STATE_SUBMITTED:
	case exchange.WITHDRAW_STATE_ALMOST_ACCEPTED:
	case exchange.WITHDRAW_STATE_REJECTED:
	case exchange.WITHDRAW_STATE_ACCEPTED:
	case exchange.WITHDRAW_STATE_PROCESSING:
	case exchange.WITHDRAW_STATE_DONE:
	case exchange.WITHDRAW_STATE_CANCELED:
	default:
		return nil, nil, fmt.Errorf("invalid state")
	}

	l, e := strconv.Atoi(limit)
	if e != nil {
		return nil, nil, e
	}
	if l < 1 || l > 100 {
		return nil, nil, fmt.Errorf("invalid limit. 0 < limit <= 100")
	}

	switch orderBy {
	case exchange.ORDERBY_ASC:
	case exchange.ORDERBY_DESC:
	default:
		orderBy = exchange.ORDERBY_DESC
	}

	api, e := GetApiInfo(FuncGetWithdraws)
	if e != nil {
		return nil, nil, e
	}

	var values = url.Values{
		"currency": []string{currency},
		"state":    []string{state},
		"uuids":    uuids,
		"txids":    txids,
		"limit":    []string{limit},
		"page":     []string{page},
		"order_by": []string{orderBy},
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

	withdraws, e := withdraw.WithdrawsFromJSON(resp.Body)
	if e != nil {
		return nil, nil, e
	}

	return withdraws, model.RemainingFromHeader(resp.Header), nil
}

// GetWithdraw 개별 출금 조회
//
// [QUERY PARAMS]
//
// uuid : 출금 UUID
//
// txid : 출금 TXID
//
// currency : Currency 코드
//
// [HEADERS]
//
// Authorization : REQUIRED. Authorization token(JWT)
func (u *Upbit) GetWithdraw(uuid, txid, currency string) (*withdraw.Withdraw, *model.Remaining, error) {
	if (len(uuid) + len(txid) + len(currency)) == 0 {
		return nil, nil, fmt.Errorf("invalid args")
	}

	api, e := GetApiInfo(FuncGetWithdraw)
	if e != nil {
		return nil, nil, e
	}

	var values = url.Values{
		"uuid":     []string{uuid},
		"txid":     []string{txid},
		"currency": []string{currency},
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

	withdraw, e := withdraw.WithdrawFromJSON(resp.Body)
	if e != nil {
		return nil, nil, e
	}

	return withdraw, model.RemainingFromHeader(resp.Header), nil
}

// GetWithdrawChance 출금 가능 정보
//
// [QUERY PARAMS]
//
// currency : REQUIRED. Currency symbol
//
// [HEADERS]
//
// Authorization : REQUIRED. Authorization token (JWT)
func (u *Upbit) GetWithdrawChance(currency string) (*withdraw.Chance, *model.Remaining, error) {
	if len(currency) == 0 {
		return nil, nil, fmt.Errorf("currency length is 0")
	}

	api, e := GetApiInfo(FuncGetWithdrawChance)
	if e != nil {
		return nil, nil, e
	}

	var values = url.Values{
		"currency": []string{currency},
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

	chance, e := withdraw.ChanceFromJSON(resp.Body)
	if e != nil {
		return nil, nil, e
	}

	return chance, model.RemainingFromHeader(resp.Header), nil
}

// WithdrawCoin 코인 출금하기
//
// [QUERY PARAMS]
//
// currency : REQUIRED. Currency symbol
//
// amount : REQUIRED. 출금 코인 수량
//
// address : REQUIRED. 출금 지갑 주소
//
// secondaryAddress : 2차 출금주소 (필요한 코인에 한해서)
//
// transactionType : 출금 유형
//
// [HEADERS]
//
// Authorization : REQUIRED. Authorization token(JWT)
func (u *Upbit) WithdrawCoin(currency, amount, address, secondaryAddress, transactionType string) (*withdraw.Withdraw, *model.Remaining, error) {
	if len(currency) == 0 {
		return nil, nil, fmt.Errorf("currency length is 0")
	}

	if len(amount) == 0 {
		return nil, nil, fmt.Errorf("amount length is 0")
	}

	if len(address) == 0 {
		return nil, nil, fmt.Errorf("address length is 0")
	}

	switch transactionType {
	case exchange.WITHDRAW_TRANSACTION_DEFAULT:
	case exchange.WITHDRAW_TRANSACTION_INTERNAL:
	default:
		return nil, nil, fmt.Errorf("invalid transactionType")
	}

	api, e := GetApiInfo(FuncWithdrawCoin)
	if e != nil {
		return nil, nil, e
	}

	var values = url.Values{
		"currency":          []string{currency},
		"amount":            []string{amount},
		"address":           []string{address},
		"secondary_address": []string{secondaryAddress},
		"transaction_type":  []string{transactionType},
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

	withdraw, e := withdraw.WithdrawFromJSON(resp.Body)
	if e != nil {
		return nil, nil, e
	}

	return withdraw, model.RemainingFromHeader(resp.Header), nil
}

// WithdrawKrw 원화 출금하기
//
// [QUERY PARAMS]
//
// amount : REQUIRED. 출금 코인 수량
//
// [HEADERS]
//
// Authorization : REQUIRED. Authorization token(JWT)
func (u *Upbit) WithdrawKrw(amount string) (*withdraw.Withdraw, *model.Remaining, error) {
	if len(amount) == 0 {
		return nil, nil, fmt.Errorf("amount length is 0")
	}

	api, e := GetApiInfo(FuncWithdrawKrw)
	if e != nil {
		return nil, nil, e
	}

	var values = url.Values{
		"amount": []string{amount},
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

	withdraw, e := withdraw.WithdrawFromJSON(resp.Body)
	if e != nil {
		return nil, nil, e
	}

	return withdraw, model.RemainingFromHeader(resp.Header), nil
}
