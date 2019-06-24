package upbit

import (
	"testing"
	"time"

	"github.com/sangx2/upbit-go/model/exchange"
)

func TestExchange(t *testing.T) {
	var uuids []string

	marketID := "KRW-BTC"
	currency := "BTC"

	// fix me
	u := NewUpbit("", "")

	// account
	accounts, remaining, e := u.GetAccounts()
	if e != nil {
		t.Errorf("GetAccounts error : %s", e)
	} else {
		t.Logf("GetAccounts[remaining:%+v]", *remaining)
		for _, account := range accounts {
			t.Logf("%+v", *account)
		}
	}

	// order
	chance, remaining, e := u.GetOrderChance(marketID)
	if e != nil {
		t.Errorf("GetOrderChance error : %s", e)
	} else {
		t.Logf("GetOrderChance[remaining:%+v]\n%+v", *remaining, *chance)
	}

	orders, remaining, e := u.GetOrders(marketID, exchange.ORDER_STATE_WAIT, nil, nil, "1", exchange.ORDERBY_ASC)
	if e != nil {
		t.Errorf("GetOrders error : %s", e)
	} else {
		t.Logf("GetOrders[state:%s,remaining:%+v]\n", exchange.ORDER_STATE_WAIT, *remaining)
		for _, order := range orders {
			t.Logf("%+v", *order)
			uuids = append(uuids, order.UUID)
		}
	}

	orders, remaining, e = u.GetOrders(marketID, exchange.ORDER_STATE_DONE, nil, nil, "1", exchange.ORDERBY_ASC)
	if e != nil {
		t.Errorf("GetOrders error : %s", e)
	} else {
		t.Logf("GetOrders[state:%s,remaining:%+v]\n", exchange.ORDER_STATE_DONE, *remaining)
		for _, order := range orders {
			t.Logf("%+v", *order)
			uuids = append(uuids, order.UUID)
		}
	}

	orders, remaining, e = u.GetOrders(marketID, exchange.ORDER_STATE_CANCEL, nil, nil, "1", exchange.ORDERBY_ASC)
	if e != nil {
		t.Errorf("GetOrders error : %s", e)
	} else {
		t.Logf("GetOrders[state:%s,remaining:%+v]\n", exchange.ORDER_STATE_CANCEL, *remaining)
		for _, order := range orders {
			t.Logf("%+v", *order)
			uuids = append(uuids, order.UUID)
		}
	}

	if len(uuids) != 0 {
		order, remaining, e := u.GetOrder(uuids[0], "")
		if e != nil {
			t.Errorf("GetOrder error : %s", e)
		} else {
			t.Logf("GetOrder[remaining:%+v]\n%+v", *remaining, *order)
		}
	}

	/*
		purchanseOrder, remaining, e := u.PurchaseOrder(market, "0.00000001", "100000000000", exchange.ORDER_TYPE_LIMIT, "")
		if e != nil {
			t.Errorf("PurchaseOrder error : %s", e)
		} else {
			t.Logf("PurchaseOrder[remaining:%+v]\n%+v", *remaining, *purchanseOrder)
		}

		sellOrder, remaining, e := u.SellOrder(market, "0.00000001", "9999999999000", exchange.ORDER_TYPE_LIMIT, "")
		if e != nil {
			t.Errorf("SellOrder error : %s", e)
		} else {
			t.Logf("SellOrder[remaining:%+v]\n%+v", *remaining, *sellOrder)
		}

		cancelOrder, remaining, e := u.CancelOrder("xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx")
		if e != nil {
			t.Errorf("CancelOrder error : %s", e)
		} else {
			t.Logf("CancelOrder[remaining:%+v]\n%+v", *remaining, *cancelOrder)
		}
	*/

	// withdraw
	uuids = nil
	withdraws, remaining, e := u.GetWithdraws(currency, exchange.WITHDRAW_STATE_ACCEPTED, "100")
	if e != nil {
		t.Errorf("GetWithdraws error : %s", e)
	} else {
		t.Logf("GetWithdraws[state:%s,remaining:%+v]", exchange.WITHDRAW_STATE_ACCEPTED, *remaining)
		for _, withdraw := range withdraws {
			t.Logf("%+v", *withdraw)
			uuids = append(uuids, withdraw.UUID)
		}
	}

	withdraws, remaining, e = u.GetWithdraws(currency, exchange.WITHDRAW_STATE_ALMOST_ACCEPTED, "100")
	if e != nil {
		t.Errorf("GetWithdraws error : %s", e)
	} else {
		t.Logf("GetWithdraws[state:%s,remaining:%+v]", exchange.WITHDRAW_STATE_ALMOST_ACCEPTED, *remaining)
		for _, withdraw := range withdraws {
			t.Logf("%+v", *withdraw)
			uuids = append(uuids, withdraw.UUID)
		}
	}

	withdraws, remaining, e = u.GetWithdraws(currency, exchange.WITHDRAW_STATE_CANCELED, "100")
	if e != nil {
		t.Errorf("GetWithdraws error : %s", e)
	} else {
		t.Logf("GetWithdraws[state:%s,remaining:%+v]", exchange.WITHDRAW_STATE_CANCELED, *remaining)
		for _, withdraw := range withdraws {
			t.Logf("%+v", *withdraw)
			uuids = append(uuids, withdraw.UUID)
		}
	}

	withdraws, remaining, e = u.GetWithdraws(currency, exchange.WITHDRAW_STATE_DONE, "100")
	if e != nil {
		t.Errorf("GetWithdraws error : %s", e)
	} else {
		t.Logf("GetWithdraws[state:%s,remaining:%+v]", exchange.WITHDRAW_STATE_DONE, *remaining)
		for _, withdraw := range withdraws {
			t.Logf("%+v", *withdraw)
			uuids = append(uuids, withdraw.UUID)
		}
	}

	withdraws, remaining, e = u.GetWithdraws(currency, exchange.WITHDRAW_STATE_PROCESSING, "100")
	if e != nil {
		t.Errorf("GetWithdraws error : %s", e)
	} else {
		t.Logf("GetWithdraws[state:%s,remaining:%+v]", exchange.WITHDRAW_STATE_PROCESSING, *remaining)
		for _, withdraw := range withdraws {
			t.Logf("%+v", *withdraw)
			uuids = append(uuids, withdraw.UUID)
		}
	}

	withdraws, remaining, e = u.GetWithdraws(currency, exchange.WITHDRAW_STATE_REJECTED, "100")
	if e != nil {
		t.Errorf("GetWithdraws error : %s", e)
	} else {
		t.Logf("GetWithdraws[state:%s,remaining:%+v]", exchange.WITHDRAW_STATE_REJECTED, *remaining)
		for _, withdraw := range withdraws {
			t.Logf("%+v", *withdraw)
			uuids = append(uuids, withdraw.UUID)
		}
	}

	withdraws, remaining, e = u.GetWithdraws(currency, exchange.WITHDRAW_STATE_SUBMITTED, "100")
	if e != nil {
		t.Errorf("GetWithdraws error : %s", e)
	} else {
		t.Logf("GetWithdraws[state:%s,remaining:%+v]", exchange.WITHDRAW_STATE_SUBMITTED, *remaining)
		for _, withdraw := range withdraws {
			t.Logf("%+v", *withdraw)
			uuids = append(uuids, withdraw.UUID)
		}
	}

	withdraws, remaining, e = u.GetWithdraws(currency, exchange.WITHDRAW_STATE_SUBMITTING, "100")
	if e != nil {
		t.Errorf("GetWithdraws error : %s", e)
	} else {
		t.Logf("GetWithdraws[state:%s,remaining:%+v]", exchange.WITHDRAW_STATE_SUBMITTING, *remaining)
		for _, withdraw := range withdraws {
			t.Logf("%+v", *withdraw)
			uuids = append(uuids, withdraw.UUID)
		}
	}

	if len(uuids) != 0 {
		withdraw, remaining, e := u.GetWithdraw(uuids[0])
		if e != nil {
			t.Errorf("GetWithdraw error : %s", e)
		} else {
			if withdraw != nil {
				t.Logf("GetWithdraw[remaining:%+v]\n%+v", *remaining, *withdraw)
			}
		}
	}

	withdrawChance, remaining, e := u.GetWithdrawChance(currency)
	if e != nil {
		t.Errorf("GetWithdrawChance error : %s", e)
	} else {
		t.Logf("GetWithdrawChance[currency:%s,remaining:%+v]\n%+v", currency, *remaining, *withdrawChance)
	}

	/*
		coin, remaining, e := u.WithdrawCoin(currency, "0.00000001", "", "", exchange.WITHDRAW_TRANSACTION_DEFAULT)
		if e != nil {
			t.Errorf("WithdrawCoin error : %s", e)
		} else {
			t.Logf("WithdrawCoin[currency:%s,remaining:%+v]\n%+v", currency, *remaining, *coin)
		}

		krw, remaining, e := u.WithdrawKrw("5000")
		if e != nil {
			t.Errorf("WithdrawKrw error : %s", e)
		} else {
			t.Logf("WithdrawKrw[remaining:%+v]\n%+v", *remaining, *krw)
		}
	*/

	// deposit
	uuids = nil
	deposits, remaining, e := u.GetDeposits(currency, "100", "1", exchange.ORDERBY_ASC)
	if e != nil {
		t.Errorf("GetDeposits error : %s", e)
	} else {
		t.Logf("GetDeposits[remaining:%+v]", *remaining)
		for _, deposit := range deposits {
			t.Logf("%+v", *deposit)
			uuids = append(uuids, deposit.UUID)
		}
	}

	if len(uuids) != 0 {
		deposit, remaining, e := u.GetDeposit(uuids[0])
		if e != nil {
			t.Errorf("GetDeposit error : %s", e)
		} else {
			if deposit != nil {
				t.Logf("GetDeposit[remaining:%+v]\n%+v", *remaining, *deposit)
			}
		}
	}

	coinAddress, remaining, e := u.GenerateDepositCoinAddress(currency)
	if e != nil {
		t.Errorf("GenerateDepositCoinAddress error : %s", e)
	} else {
		if coinAddress != nil {
			t.Logf("GenerateDepositCoinAddress[remaining:%+v]\n%+v", *remaining, *coinAddress)
		}
	}

	coinAddresses, remaining, e := u.GetDepositCoinAddresses()
	if e != nil {
		t.Errorf("GetDepositCoinAddresses error : %s", e)
	} else {
		t.Logf("GetDepositCoinAddresses[remaining:%+v]", *remaining)
		for _, coinAddress := range coinAddresses {
			t.Logf("%+v", *coinAddress)
		}
	}

	coinAddress, remaining, e = u.GetDepositCoinAddress(currency)
	if e != nil {
		t.Errorf("GetDepositCoinAddress error : %s", e)
	} else {
		if coinAddress != nil {
			t.Logf("GetDepositCoinAddress[remaining:%+v]\n%+v", *remaining, *coinAddress)
		}
	}
}

func TestQuotation(t *testing.T) {
	u := NewUpbit("", "")

	markets, remaining, e := u.GetMarkets()
	if e != nil || len(markets) == 0 {
		t.Errorf("GetMarkets error : %s", e)
	} else {
		t.Logf("GetMarkets[remaining:%+v]", *remaining)
		for _, market := range markets {
			t.Logf("%+v", *market)
		}
	}

	marketID := markets[0].Market

	// candle
	minuteCandles, remaining, e := u.GetMinuteCandles(marketID, time.Now().Add(-time.Minute).Format(time.RFC3339), "", "1")
	if e != nil {
		t.Errorf("%s's GetMinuteCandles error : %s", marketID, e)
	} else {
		t.Logf("GetMinuteCandles[remaining:%+v]", *remaining)
		for _, minuteCandle := range minuteCandles {
			t.Logf("%+v", *minuteCandle)
		}
	}

	dayCandles, remaining, e := u.GetDayCandles(marketID, time.Now().Format(time.RFC3339), "", "")
	if e != nil {
		t.Errorf("%s's GetDayCandles error : %s", marketID, e)
	} else {
		t.Logf("GetDayCandles[remaining:%+v]", *remaining)
		for _, dayCandle := range dayCandles {
			t.Logf("%+v", *dayCandle)
		}
	}

	weekCandles, remaining, e := u.GetWeekCandles(marketID, time.Now().Format(time.RFC3339), "")
	if e != nil {
		t.Errorf("%s's GetWeekCandles error : %s", marketID, e)
	} else {
		t.Logf("GetWeekCandles[remaining:%+v]", *remaining)
		for _, weekCandle := range weekCandles {
			t.Logf("%+v", *weekCandle)
		}
	}

	monthCandles, remaining, e := u.GetMonthCandles(marketID, time.Now().Format(time.RFC3339), "")
	if e != nil {
		t.Errorf("%s's GetMonthCandles error : %s", marketID, e)
	} else {
		t.Logf("GetMonthCandles[remaining:%+v]", *remaining)
		for _, monthCandle := range monthCandles {
			t.Logf("%+v", *monthCandle)
		}
	}

	// ticks
	ticks, remaining, e := u.GetTradeTicks(marketID, time.Now().Format("15:04:05"), "", "")
	if e != nil {
		t.Errorf("%s's GetTradeTicks error : %s", marketID, e)
	} else {
		t.Logf("GetTradeTicks[remaining:%+v]", *remaining)
		for _, tick := range ticks {
			t.Logf("%+v", *tick)
		}
	}

	// ticker
	tickers, remaining, e := u.GetTickers([]string{marketID})
	if e != nil {
		t.Errorf("GetTickers error : %s", e)
	} else {
		t.Logf("GetTickers[remaining:%+v]", *remaining)
		for _, ticker := range tickers {
			t.Logf("%+v", *ticker)
		}
	}

	// orderbook
	orderbooks, remaining, e := u.GetOrderbooks([]string{marketID})
	if e != nil {
		t.Errorf("GetOrderbooks error : %s", e)
	} else {
		t.Logf("GetOrderbooks[remaining:%+v]", *remaining)
		for _, orderbook := range orderbooks {
			t.Logf("%+v", *orderbook)
		}
	}
}
