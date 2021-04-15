package upbit

import (
	"fmt"
	"net/http"
)

const (
	FuncGetAccounts                = "GetAccounts"
	FuncGetOrderChance             = "GetOrderChance"
	FuncGetOrder                   = "GetOrder"
	FuncGetOrders                  = "GetOrders"
	FuncPurchaseOrder              = "PurchaseOrder"
	FuncSellOrder                  = "SellOrder"
	FuncCancelOrder                = "CancelOrder"
	FuncGetWithdraws               = "GetWithdraws"
	FuncGetWithdraw                = "GetWithdraw"
	FuncGetWithdrawChance          = "GetWithdrawChance"
	FuncWithdrawCoin               = "WithdrawCoin"
	FuncWithdrawKrw                = "WithdrawKrw"
	FuncGetDeposits                = "GetDeposits"
	FuncGetDeposit                 = "GetDeposit"
	FuncGenerateDepositCoinAddress = "GenerateDepositCoinAddress"
	FuncGetDepositCoinAddresses    = "GetDepositCoinAddresses"
	FuncGetDepositCoinAddress      = "GetDepositCoinAddress"
	FuncDepositKrw                 = "DepositKrw"
	FuncGetWalletStatus            = "GetWalletStatus"
	FuncGetApiKeys                 = "GetApiKeys"
	FuncGetMarkets                 = "GetMarkets"
	FuncGetMinuteCandles           = "GetMinuteCandles"
	FuncGetDayCandles              = "GetDayCandles"
	FuncGetWeekCandles             = "GetWeekCandles"
	FuncGetMonthCandles            = "GetMonthCandles"
	FuncGetTradeTicks              = "GetTradeTicks"
	FuncGetTickers                 = "GetTickers"
	FuncGetOrderbooks              = "GetOrderbooks"

	RouteAccounts                    = "/accounts"
	RouteOrderChance                 = "/orders/chance"
	RouteOrder                       = "/order"
	RouteOrders                      = "/orders"
	RouteWithdraws                   = "/withdraws"
	RouteWithdraw                    = "/withdraw"
	RouteWithdrawsChance             = "/withdraws/chance"
	RouteWithdrawsCoin               = "/withdraws/coin"
	RouteWithdrawsKrw                = "/withdraws/krw"
	RouteDeposits                    = "/deposits"
	RouteDeposit                     = "/deposit"
	RouteDepositsGenerateCoinAddress = "/deposits/generate_coin_address"
	RouteDepositsCoinAddresses       = "/deposits/coin_addresses"
	RouteDepositsCoinAddress         = "/deposits/coin_address"
	RouteDepositsKrw                 = "/deposits/krw"
	RouteStatusWallet                = "/status/wallet"
	RouteApiKeys                     = "/api_keys"
	RouteMarketAll                   = "/market/all"
	RouteCandlesMinutes              = "/candles/minutes/"
	RouteCandlesDays                 = "/candles/days"
	RouteCandlesWeeks                = "/candles/weeks"
	RouteCandlesMonths               = "/candles/months"
	RouteTradesTicks                 = "/trades/ticks"
	RouteTicker                      = "/ticker"
	RouteOrderbook                   = "/orderbook"

	ApiSectionExchange  = "exchange"
	ApiSectionQuotation = "quotation"

	ApiGroupDefault      = "default"
	ApiGroupOrder        = "order"
	ApiGroupStatusWallet = "status-wallet"
	ApiGroupMarket       = "market"
	ApiGroupCandles      = "candles"
	ApiGroupCrixTrades   = "crix-trades"
	ApiGroupTicker       = "ticker"
	ApiGroupOrderbook    = "orderbook"
)

type ApiInfo struct {
	Method, Url, Section, Group string
}

func GetApiInfo(funcName string) (*ApiInfo, error) {
	switch funcName {
	case FuncGetAccounts: //전체 계좌 조회
		return &ApiInfo{
			Method: http.MethodGet, Url: RouteAccounts,
			Section: ApiSectionExchange, Group: ApiGroupDefault,
		}, nil
	case FuncGetOrderChance: //주문 가능 정보
		return &ApiInfo{
			Method: http.MethodGet, Url: RouteOrderChance,
			Section: ApiSectionExchange, Group: ApiGroupDefault,
		}, nil
	case FuncGetOrder: //개별 주문 조회
		return &ApiInfo{
			Method: http.MethodGet, Url: RouteOrder,
			Section: ApiSectionExchange, Group: ApiGroupDefault,
		}, nil
	case FuncGetOrders: //주문 리스트 조회
		return &ApiInfo{
			Method: http.MethodGet, Url: RouteOrders,
			Section: ApiSectionExchange, Group: ApiGroupDefault,
		}, nil
	case FuncPurchaseOrder: //매수하기
		return &ApiInfo{
			Method: http.MethodPost, Url: RouteOrders,
			Section: ApiSectionExchange, Group: ApiGroupOrder,
		}, nil
	case FuncSellOrder: //매도하기
		return &ApiInfo{
			Method: http.MethodPost, Url: RouteOrders,
			Section: ApiSectionExchange, Group: ApiGroupOrder,
		}, nil
	case FuncCancelOrder: //주문 취소 접수
		return &ApiInfo{
			Method: http.MethodDelete, Url: RouteOrder,
			Section: ApiSectionExchange, Group: ApiGroupDefault,
		}, nil
	case FuncGetWithdraws: //출금 리스트 조회
		return &ApiInfo{
			Method: http.MethodGet, Url: RouteWithdraws,
			Section: ApiSectionExchange, Group: ApiGroupDefault,
		}, nil
	case FuncGetWithdraw: //개별 출금 조회
		return &ApiInfo{
			Method: http.MethodGet, Url: RouteWithdraw,
			Section: ApiSectionExchange, Group: ApiGroupDefault,
		}, nil
	case FuncGetWithdrawChance: //출금 가능 정보
		return &ApiInfo{
			Method: http.MethodGet, Url: RouteWithdrawsChance,
			Section: ApiSectionExchange, Group: ApiGroupDefault,
		}, nil
	case FuncWithdrawCoin: //코인 출금하기
		return &ApiInfo{
			Method: http.MethodPost, Url: RouteWithdrawsCoin,
			Section: ApiSectionExchange, Group: ApiGroupDefault,
		}, nil
	case FuncWithdrawKrw: //원화 출금하기
		return &ApiInfo{
			Method: http.MethodPost, Url: RouteWithdrawsKrw,
			Section: ApiSectionExchange, Group: ApiGroupDefault,
		}, nil
	case FuncGetDeposits: //입금 리스트 조회
		return &ApiInfo{
			Method: http.MethodGet, Url: RouteDeposits,
			Section: ApiSectionExchange, Group: ApiGroupDefault,
		}, nil
	case FuncGetDeposit: //개별 입금 조회
		return &ApiInfo{
			Method: http.MethodGet, Url: RouteDeposit,
			Section: ApiSectionExchange, Group: ApiGroupDefault,
		}, nil
	case FuncGenerateDepositCoinAddress: //입금 주소 생성 요청
		return &ApiInfo{
			Method: http.MethodPost, Url: RouteDepositsGenerateCoinAddress,
			Section: ApiSectionExchange, Group: ApiGroupDefault,
		}, nil
	case FuncGetDepositCoinAddresses: //전체 입금 주소 조회
		return &ApiInfo{
			Method: http.MethodGet, Url: RouteDepositsCoinAddresses,
			Section: ApiSectionExchange, Group: ApiGroupDefault,
		}, nil
	case FuncGetDepositCoinAddress: //개별 입금 주소 조회
		return &ApiInfo{
			Method: http.MethodGet, Url: RouteDepositsCoinAddress,
			Section: ApiSectionExchange, Group: ApiGroupDefault,
		}, nil
	case FuncDepositKrw: //원화 입금하기
		return &ApiInfo{
			Method: http.MethodPost, Url: RouteDepositsKrw,
			Section: ApiSectionExchange, Group: ApiGroupDefault,
		}, nil
	case FuncGetWalletStatus: //입출금 현황
		return &ApiInfo{
			Method: http.MethodGet, Url: RouteStatusWallet,
			Section: ApiSectionExchange, Group: ApiGroupStatusWallet,
		}, nil
	case FuncGetApiKeys: //API 키 리스트 조회
		return &ApiInfo{
			Method: http.MethodGet, Url: RouteApiKeys,
			Section: ApiSectionExchange, Group: ApiGroupDefault,
		}, nil
	case FuncGetMarkets: //마켓 코드 조회
		return &ApiInfo{
			Method: http.MethodGet, Url: RouteMarketAll,
			Section: ApiSectionQuotation, Group: ApiGroupMarket,
		}, nil
	case FuncGetMinuteCandles: //분 캔들
		return &ApiInfo{
			Method: http.MethodGet, Url: RouteCandlesMinutes,
			Section: ApiSectionQuotation, Group: ApiGroupCandles,
		}, nil
	case FuncGetDayCandles: //일 캔들
		return &ApiInfo{
			Method: http.MethodGet, Url: RouteCandlesDays,
			Section: ApiSectionQuotation, Group: ApiGroupCandles,
		}, nil
	case FuncGetWeekCandles: //주 캔들
		return &ApiInfo{
			Method: http.MethodGet, Url: RouteCandlesWeeks,
			Section: ApiSectionQuotation, Group: ApiGroupCandles,
		}, nil
	case FuncGetMonthCandles: //월 캔들
		return &ApiInfo{
			Method: http.MethodGet, Url: RouteCandlesMonths,
			Section: ApiSectionQuotation, Group: ApiGroupCandles,
		}, nil
	case FuncGetTradeTicks: //당일 체결 내역
		return &ApiInfo{
			Method: http.MethodGet, Url: RouteTradesTicks,
			Section: ApiSectionQuotation, Group: ApiGroupCrixTrades,
		}, nil
	case FuncGetTickers: //현재가 정보
		return &ApiInfo{
			Method: http.MethodGet, Url: RouteTicker,
			Section: ApiSectionQuotation, Group: ApiGroupTicker,
		}, nil
	case FuncGetOrderbooks: //호가 정보 조회
		return &ApiInfo{
			Method: http.MethodGet, Url: RouteOrderbook,
			Section: ApiSectionQuotation, Group: ApiGroupOrderbook,
		}, nil
	default:
		return nil, fmt.Errorf("function is not defined")
	}
}
