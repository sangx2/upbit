package upbit

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"net/url"
	"strconv"

	"github.com/sangx2/upbit/model"
	"github.com/sangx2/upbit/model/exchange"
	"github.com/sangx2/upbit/model/exchange/account"
	"github.com/sangx2/upbit/model/exchange/deposit"
	"github.com/sangx2/upbit/model/exchange/order"
	"github.com/sangx2/upbit/model/exchange/withdraw"
)

// GetAccounts 전체 계좌 조회
//
// [HEADERS]
//
// Authorization : REQUIRED. Authorization token (JWT)
func (u *Upbit) GetAccounts() ([]*account.Account, *model.Remaining, error) {
	resp, e := u.getResponse(METHOD_GET, URL_UPBIT_V1+"/accounts", nil, API_TYPE_EXCHANGE, API_GROUP_DEFAULT)
	if e != nil {
		return nil, nil, e
	}
	defer resp.Body.Close()

	accounts := account.AccountsFromJSON(resp.Body)
	if accounts == nil {
		return nil, nil, errors.New("account.AccountsFromJSON is nil")
	}

	return accounts, model.RemainingFromHeader(resp.Header), nil
}

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
		return nil, nil, errors.New("market length is 0")
	}

	values := url.Values{
		"market": []string{market},
	}

	resp, e := u.getResponse(METHOD_GET, URL_UPBIT_V1+"/orders/chance", values, API_TYPE_EXCHANGE, API_GROUP_DEFAULT)
	if e != nil {
		return nil, nil, e
	}
	defer resp.Body.Close()

	chance := order.ChanceFromJSON(resp.Body)
	if chance == nil {
		return nil, nil, errors.New("order.ChanceFromJSON is nil")
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
		return nil, nil, errors.New("uuid and identifier length is 0")
	} else if len(uuid) != 0 && len(identifier) != 0 {
		return nil, nil, errors.New("uuid and identifier length is not 0. you must set only one param")
	}

	values := url.Values{}
	if len(uuid) != 0 {
		values.Add("uuid", uuid)
	}
	if len(identifier) != 0 {
		values.Add("identifier", identifier)
	}

	resp, e := u.getResponse(METHOD_GET, URL_UPBIT_V1+"/order", values, API_TYPE_EXCHANGE, API_GROUP_DEFAULT)
	if e != nil {
		return nil, nil, e
	}
	defer resp.Body.Close()

	order := order.OrderFromJSON(resp.Body)
	if order == nil {
		return nil, nil, errors.New("order.OrderFromJSON is nil")
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
// orderBy : 정렬 방식
//
// [HEADERS]
//
// Authorization : Authorization token(JWT)
func (u *Upbit) GetOrders(market, state string, uuids, identifiers []string, page, orderBy string) ([]*order.Order, *model.Remaining, error) {
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

	values := url.Values{
		"market":      []string{market},
		"state":       []string{state},
		"uuids":       uuids,
		"identifiers": identifiers,
		"page":        []string{page},
		"order_by":    []string{orderBy},
	}

	resp, e := u.getResponse(METHOD_GET, URL_UPBIT_V1+"/orders", values, API_TYPE_EXCHANGE, API_GROUP_DEFAULT)
	if e != nil {
		return nil, nil, e
	}
	defer resp.Body.Close()

	orders := order.OrdersFromJSON(resp.Body)
	if orders == nil {
		return nil, nil, errors.New("order.OrdersFromJSON is nil")
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
		return nil, nil, errors.New("market length is 0")
	}

	if len(volume) == 0 {
		return nil, nil, errors.New("volume length is 0")
	}

	if len(price) == 0 {
		return nil, nil, errors.New("price length is 0")
	}

	switch orderType {
	case exchange.ORDER_TYPE_LIMIT:
	case exchange.ORDER_TYPE_PRICE:
	default:
		return nil, nil, errors.New("invalid orderType")
	}

	values := url.Values{
		"market":     []string{market},
		"side":       []string{exchange.ORDER_SIDE_BID},
		"volume":     []string{volume},
		"price":      []string{price},
		"ord_type":   []string{orderType},
		"identifier": []string{identifier},
	}

	resp, e := u.getResponse(METHOD_POST, URL_UPBIT_V1+"/orders", values, API_TYPE_EXCHANGE, API_GROUP_ORDER)
	if e != nil {
		return nil, nil, e
	}
	defer resp.Body.Close()

	order := order.OrderFromJSON(resp.Body)
	if order == nil {
		return nil, nil, errors.New("order.OrderFromJSON is nil")
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
		return nil, nil, errors.New("market length is 0")
	}

	if len(volume) == 0 {
		return nil, nil, errors.New("volume length is 0")
	}

	if len(price) == 0 {
		return nil, nil, errors.New("price length is 0")
	}

	switch orderType {
	case exchange.ORDER_TYPE_LIMIT:
	case exchange.ORDER_TYPE_MARKET:
	default:
		return nil, nil, errors.New("invalid orderType")
	}

	values := url.Values{
		"market":     []string{market},
		"side":       []string{exchange.ORDER_SIDE_ASK},
		"volume":     []string{volume},
		"price":      []string{price},
		"ord_type":   []string{orderType},
		"identifier": []string{identifier},
	}

	resp, e := u.getResponse(METHOD_POST, URL_UPBIT_V1+"/orders", values, API_TYPE_EXCHANGE, API_GROUP_ORDER)
	if e != nil {
		return nil, nil, e
	}
	defer resp.Body.Close()

	order := order.OrderFromJSON(resp.Body)
	if order == nil {
		return nil, nil, errors.New("order.OrderFromJSON is nil")
	}
	return order, model.RemainingFromHeader(resp.Header), nil
}

// CancelOrder 주문 취소 접수
//
// [QUERY PARAMS]
//
// uuid : REQUIRED. 주문 UUID
//
func (u *Upbit) CancelOrder(uuid string) (*order.Order, *model.Remaining, error) {
	if len(uuid) == 0 {
		return nil, nil, errors.New("uuid length is 0")
	}

	values := url.Values{
		"uuid": []string{uuid},
	}

	resp, e := u.getResponse(METHOD_DELETE, URL_UPBIT_V1+"/order", values, API_TYPE_EXCHANGE, API_GROUP_DEFAULT)
	if e != nil {
		return nil, nil, e
	}
	defer resp.Body.Close()

	order := order.OrderFromJSON(resp.Body)
	if order == nil {
		return nil, nil, errors.New("order.OrderFromJSON is nil")
	}
	return order, model.RemainingFromHeader(resp.Header), nil
}

// GetWithdraws 출금 리스트 조회
//
// [QUERY PARAMS]
//
// currency : Currency 코드
//
// state : 출금 상태
//
// limit : 갯수 제한. default: 100, max: 100
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
		return nil, nil, errors.New("invalid state")
	}

	l, e := strconv.Atoi(limit)
	if e != nil {
		return nil, nil, e
	}
	if l < 1 || l > 100 {
		return nil, nil, errors.New("invalid limit. 0 < limit <= 100")
	}

	switch orderBy {
	case exchange.ORDERBY_ASC:
	case exchange.ORDERBY_DESC:
	default:
		orderBy = exchange.ORDERBY_DESC
	}

	values := url.Values{
		"currency": []string{currency},
		"state":    []string{state},
		"uuids":    uuids,
		"txids":    txids,
		"limit":    []string{limit},
		"page":     []string{page},
		"order_by": []string{orderBy},
	}

	resp, e := u.getResponse(METHOD_GET, URL_UPBIT_V1+"/withdraws", values, API_TYPE_EXCHANGE, API_GROUP_DEFAULT)
	if e != nil {
		return nil, nil, e
	}
	defer resp.Body.Close()

	withdraws := withdraw.WithdrawsFromJSON(resp.Body)
	if withdraws == nil {
		return nil, nil, errors.New("withdraw.WithdrawsFromJSON is nil")
	}
	return withdraws, model.RemainingFromHeader(resp.Header), nil
}

// GetWithdraw 개별 출금 조회
//
// [QUERY PARAMS]
//
// uuid : REQUIRED. 출금 UUID
//
// [HEADERS]
//
// Authorization : REQUIRED. Authorization token(JWT)
func (u *Upbit) GetWithdraw(uuid string) (*withdraw.Withdraw, *model.Remaining, error) {
	if len(uuid) == 0 {
		return nil, nil, errors.New("uuid length is 0")
	}

	values := url.Values{
		"uuid": []string{uuid},
	}

	resp, e := u.getResponse(METHOD_GET, URL_UPBIT_V1+"/withdraw", values, API_TYPE_EXCHANGE, API_GROUP_DEFAULT)
	if e != nil {
		return nil, nil, e
	}
	defer resp.Body.Close()

	withdraw := withdraw.WithdrawFromJSON(resp.Body)
	if withdraw == nil {
		return nil, nil, errors.New("withdraw.WithdrawFromJSON is nil")
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
		return nil, nil, errors.New("currency length is 0")
	}

	values := url.Values{
		"currency": []string{currency},
	}

	resp, e := u.getResponse(METHOD_GET, URL_UPBIT_V1+"/withdraws/chance", values, API_TYPE_EXCHANGE, API_GROUP_DEFAULT)
	if e != nil {
		return nil, nil, e
	}
	defer resp.Body.Close()

	chance := withdraw.ChanceFromJSON(resp.Body)
	if chance == nil {
		return nil, nil, errors.New("withdraw.ChanceFromJSON is nil")
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
		return nil, nil, errors.New("currency length is 0")
	}

	if len(amount) == 0 {
		return nil, nil, errors.New("amount length is 0")
	}

	if len(address) == 0 {
		return nil, nil, errors.New("address length is 0")
	}

	switch transactionType {
	case exchange.WITHDRAW_TRANSACTION_DEFAULT:
	case exchange.WITHDRAW_TRANSACTION_INTERNAL:
	default:
		return nil, nil, errors.New("invalid transactionType")
	}

	values := url.Values{
		"currency":          []string{currency},
		"amount":            []string{amount},
		"address":           []string{address},
		"secondary_address": []string{secondaryAddress},
		"transaction_type":  []string{transactionType},
	}

	resp, e := u.getResponse(METHOD_POST, URL_UPBIT_V1+"/withdraws/coin", values, API_TYPE_EXCHANGE, API_GROUP_DEFAULT)
	if e != nil {
		return nil, nil, e
	}
	defer resp.Body.Close()

	withdraw := withdraw.WithdrawFromJSON(resp.Body)
	if withdraw == nil {
		return nil, nil, errors.New("withdraw.WithdrawFromJSON is nil")
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
		return nil, nil, errors.New("amount length is 0")
	}

	values := url.Values{
		"amount": []string{amount},
	}

	resp, e := u.getResponse(METHOD_POST, URL_UPBIT_V1+"/withdraws/krw", values, API_TYPE_EXCHANGE, API_GROUP_DEFAULT)
	if e != nil {
		return nil, nil, e
	}
	defer resp.Body.Close()

	withdraw := withdraw.WithdrawFromJSON(resp.Body)
	if withdraw == nil {
		return nil, nil, errors.New("withdraw.WithdrawFromJSON is nil")
	}
	return withdraw, model.RemainingFromHeader(resp.Header), nil
}

// GetDeposits 입금 리스트 조회
//
// [QUERY PARAMS]
//
// currency : Currency 코드
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
		return nil, nil, errors.New("invalid state")
	}

	l, e := strconv.Atoi(limit)
	if e != nil {
		return nil, nil, e
	}
	if l < 1 || l > 100 {
		return nil, nil, errors.New("invalid limit. 0 < limit <= 100")
	}

	switch orderBy {
	case exchange.ORDERBY_ASC:
	case exchange.ORDERBY_DESC:
	default:
		orderBy = exchange.ORDERBY_DESC
	}

	values := url.Values{
		"currency": []string{currency},
		"limit":    []string{limit},
		"page":     []string{page},
		"order_by": []string{orderBy},
	}

	resp, e := u.getResponse(METHOD_GET, URL_UPBIT_V1+"/deposits", values, API_TYPE_EXCHANGE, API_GROUP_DEFAULT)
	if e != nil {
		return nil, nil, e
	}
	defer resp.Body.Close()

	deposits := deposit.DepositsFromJSON(resp.Body)
	if deposits == nil {
		return nil, nil, errors.New("deposit.DepositsFromJSON is nil")
	}
	return deposits, model.RemainingFromHeader(resp.Header), nil
}

// GetDeposit 개별 입금 조회
//
// [QUERY PARAMS]
//
// uuid : REQUIRED. 개별 입금의 UUID
//
// [HEADERS]
//
// Authorization : REQUIRED. Authorization token(JWT)
func (u *Upbit) GetDeposit(uuid string) (*deposit.Deposit, *model.Remaining, error) {
	if len(uuid) == 0 {
		return nil, nil, errors.New("uuid length is 0")
	}

	values := url.Values{
		"uuid": []string{uuid},
	}

	resp, e := u.getResponse(METHOD_GET, URL_UPBIT_V1+"/deposit", values, API_TYPE_EXCHANGE, API_GROUP_DEFAULT)
	if e != nil {
		return nil, nil, e
	}
	defer resp.Body.Close()

	deposit := deposit.DepositFromJSON(resp.Body)
	if deposit == nil {
		return nil, nil, errors.New("deposit.DepositFromJSON is nil")
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
	values := url.Values{
		"currency": []string{currency},
	}

	resp, e := u.getResponse(METHOD_POST, URL_UPBIT_V1+"/deposits/generate_coin_address", values, API_TYPE_EXCHANGE, API_GROUP_DEFAULT)
	if e != nil {
		return nil, nil, e
	}
	defer resp.Body.Close()

	bodyBytes, e := ioutil.ReadAll(resp.Body)
	if e != nil {
		return nil, nil, e
	}

	resp1 := model.Response1FromJSON(bytes.NewReader(bodyBytes))
	if resp1 != nil {
		return nil, nil, fmt.Errorf("response1FromJSON is %+v", resp1)
	}

	coinAddress := deposit.CoinAddressFromJSON(bytes.NewReader(bodyBytes))
	if coinAddress == nil {
		return nil, nil, errors.New("deposit.CoinAddressFromJSON is nil")
	}

	return coinAddress, model.RemainingFromHeader(resp.Header), nil
}

// GetDepositCoinAddresses 전체 입금 주소 조회
//
// [HEADERS]
//
// Authorization : REQUIRED. Authorization token(JWT)
func (u *Upbit) GetDepositCoinAddresses() ([]*deposit.CoinAddress, *model.Remaining, error) {
	resp, e := u.getResponse(METHOD_GET, URL_UPBIT_V1+"/deposits/coin_addresses", nil, API_TYPE_EXCHANGE, API_GROUP_DEFAULT)
	if e != nil {
		return nil, nil, e
	}
	defer resp.Body.Close()

	coinAddresses := deposit.CoinAddressesFromJSON(resp.Body)
	if coinAddresses == nil {
		return nil, nil, errors.New("deposit.CoinAddressesFromJSON is nil")
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
	values := url.Values{
		"currency": []string{currency},
	}

	resp, e := u.getResponse(METHOD_GET, URL_UPBIT_V1+"/deposits/coin_address", values, API_TYPE_EXCHANGE, API_GROUP_DEFAULT)
	if e != nil {
		return nil, nil, e
	}
	defer resp.Body.Close()

	coinAddress := deposit.CoinAddressFromJSON(resp.Body)
	if coinAddress == nil {
		return nil, nil, errors.New("deposit.CoinAddressFromJSON is nil")
	}
	return coinAddress, model.RemainingFromHeader(resp.Header), nil
}
