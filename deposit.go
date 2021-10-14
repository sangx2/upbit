package upbit

import (
	"bytes"
	"fmt"
	"github.com/sangx2/upbit/model"
	"github.com/sangx2/upbit/model/exchange"
	"github.com/sangx2/upbit/model/exchange/deposit"
	"io/ioutil"
	"net/url"
	"strconv"
)

// GetDeposits 입금 리스트 조회
//
// [QUERY PARAMS]
//
// currency : Currency 코드
//
// state : 입금 상태
//
// uuids : 입금 UUID의 목록
//
// txids : 입금 TXID의 목록
//
// limit : 페이지당 개수
//
// page ; 페이지 번호
//
// orderBy : 정렬 방식
//
// [HEADERS]
//
// Authorization : REQUIRED. Authorization token(JWT)
func (u *Upbit) GetDeposits(currency, state string, uuids, txids []string, limit, page, orderBy string) ([]*deposit.Deposit, *model.Remaining, error) {
	switch state {
	case exchange.DEPOSIT_STATE_SUBMITTING:
	case exchange.DEPOSIT_STATE_SUBMITTED:
	case exchange.DEPOSIT_STATE_ALMOST_ACCEPTED:
	case exchange.DEPOSIT_STATE_REJECTED:
	case exchange.DEPOSIT_STATE_ACCEPTED:
	case exchange.DEPOSIT_STATE_PROCESSING:
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

	api, e := GetApiInfo(FuncGetDeposits)
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

	deposits, e := deposit.DepositsFromJSON(resp.Body)
	if e != nil {
		return nil, nil, e
	}

	return deposits, model.RemainingFromHeader(resp.Header), nil
}

// GetDeposit 개별 입금 조회
//
// [QUERY PARAMS]
//
// uuid : 입금 UUID
//
// txid : 입금 TXID
//
// currency : Currency 코드
//
// [HEADERS]
//
// Authorization : REQUIRED. Authorization token(JWT)
func (u *Upbit) GetDeposit(uuid, txid, currency string) (*deposit.Deposit, *model.Remaining, error) {
	if (len(uuid) + len(txid) + len(currency)) == 0 {
		return nil, nil, fmt.Errorf("invalid args")
	}

	api, e := GetApiInfo(FuncGetDeposit)
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

	deposit, e := deposit.DepositFromJSON(resp.Body)
	if e != nil {
		return nil, nil, e
	}

	return deposit, model.RemainingFromHeader(resp.Header), nil
}

// GenerateDepositCoinAddress 입금 주소 생성 요청
//
// [QUERY PARAMS]
//
// currency : Currency 코드
//
// [HEADERS]
//
// Authorization : REQUIRED. Authorization token(JWT)
//
// 주소 발급 요청 시 결과로 Response1이 반환되며 주소 발급 완료 이전까지 계속 Response1이 반환됩니다.
//
// 주소가 발급된 이후부터는 새로운 주소가 발급되는 것이 아닌 이전에 발급된 주소가 Response2 형태로 반환됩니다.
func (u *Upbit) GenerateDepositCoinAddress(currency string) (*deposit.CoinAddress, *model.Remaining, error) {
	api, e := GetApiInfo(FuncGenerateDepositCoinAddress)
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

	bodyBytes, e := ioutil.ReadAll(resp.Body)
	if e != nil {
		return nil, nil, e
	}

	resp1, e := model.Response1FromJSON(bytes.NewReader(bodyBytes))
	if e != nil {
		return nil, nil, e
	}

	if resp1.Success {
		return nil, model.RemainingFromHeader(resp.Header), fmt.Errorf(resp1.Message)
	} else {
		coinAddress, e := deposit.CoinAddressFromJSON(bytes.NewReader(bodyBytes))
		if e != nil {
			return nil, nil, e
		}
		return coinAddress, model.RemainingFromHeader(resp.Header), nil
	}
}

// GetDepositCoinAddresses 전체 입금 주소 조회
//
// [HEADERS]
//
// Authorization : REQUIRED. Authorization token(JWT)
func (u *Upbit) GetDepositCoinAddresses() ([]*deposit.CoinAddress, *model.Remaining, error) {
	api, e := GetApiInfo(FuncGetDepositCoinAddresses)
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

	coinAddresses, e := deposit.CoinAddressesFromJSON(resp.Body)
	if e != nil {
		return nil, nil, e
	}

	return coinAddresses, model.RemainingFromHeader(resp.Header), nil
}

// GetDepositCoinAddress 개별 입금 주소 조회
//
// [QUERY PARAMS]
//
// currency : Currency 코드
//
// [HEADERS]
//
// Authorization : REQUIRED. Authorization token(JWT)
func (u *Upbit) GetDepositCoinAddress(currency string) (*deposit.CoinAddress, *model.Remaining, error) {
	api, e := GetApiInfo(FuncGetDepositCoinAddress)
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

	coinAddress, e := deposit.CoinAddressFromJSON(resp.Body)
	if e != nil {
		return nil, nil, e
	}

	return coinAddress, model.RemainingFromHeader(resp.Header), nil
}

// DepositKrw 원화 입금하기
//
// [BODY PARAMS]
//
// amount : 입금 원화 수량
//
// [HEADERS]
//
// Authorization : REQUIRED. Authorization token(JWT)
func (u *Upbit) DepositKrw(amount string) (*deposit.Deposit, *model.Remaining, error) {
	api, e := GetApiInfo(FuncDepositKrw)
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

	deposit, e := deposit.DepositFromJSON(resp.Body)
	if e != nil {
		return nil, nil, e
	}

	return deposit, model.RemainingFromHeader(resp.Header), nil
}
