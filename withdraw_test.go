package upbit

import (
	"github.com/sangx2/upbit/model/exchange"
	"testing"
)

func TestWithdraw(t *testing.T) {
	var uuids []string

	if len(accessKey) != 0 && len(secretKey) != 0 {
		u := NewUpbit(accessKey, secretKey)

		// withdraw
		uuids = nil
		withdraws, remaining, e := u.GetWithdraws(currency, exchange.WITHDRAW_STATE_ACCEPTED, nil, nil, "100", "1", exchange.ORDERBY_ASC)
		if e != nil {
			t.Errorf("GetWithdraws error : %s", e.Error())
		} else {
			t.Logf("GetWithdraws[state:%s,remaining:%+v]", exchange.WITHDRAW_STATE_ACCEPTED, *remaining)
			for _, withdraw := range withdraws {
				t.Logf("%+v", *withdraw)
				uuids = append(uuids, withdraw.UUID)
			}
		}

		withdraws, remaining, e = u.GetWithdraws(currency, exchange.WITHDRAW_STATE_ALMOST_ACCEPTED, nil, nil, "100", "1", exchange.ORDERBY_ASC)
		if e != nil {
			t.Errorf("GetWithdraws error : %s", e.Error())
		} else {
			t.Logf("GetWithdraws[state:%s,remaining:%+v]", exchange.WITHDRAW_STATE_ALMOST_ACCEPTED, *remaining)
			for _, withdraw := range withdraws {
				t.Logf("%+v", *withdraw)
				uuids = append(uuids, withdraw.UUID)
			}
		}

		withdraws, remaining, e = u.GetWithdraws(currency, exchange.WITHDRAW_STATE_CANCELED, nil, nil, "100", "1", exchange.ORDERBY_ASC)
		if e != nil {
			t.Errorf("GetWithdraws error : %s", e.Error())
		} else {
			t.Logf("GetWithdraws[state:%s,remaining:%+v]", exchange.WITHDRAW_STATE_CANCELED, *remaining)
			for _, withdraw := range withdraws {
				t.Logf("%+v", *withdraw)
				uuids = append(uuids, withdraw.UUID)
			}
		}

		withdraws, remaining, e = u.GetWithdraws(currency, exchange.WITHDRAW_STATE_DONE, nil, nil, "100", "1", exchange.ORDERBY_ASC)
		if e != nil {
			t.Errorf("GetWithdraws error : %s", e.Error())
		} else {
			t.Logf("GetWithdraws[state:%s,remaining:%+v]", exchange.WITHDRAW_STATE_DONE, *remaining)
			for _, withdraw := range withdraws {
				t.Logf("%+v", *withdraw)
				uuids = append(uuids, withdraw.UUID)
			}
		}

		withdraws, remaining, e = u.GetWithdraws(currency, exchange.WITHDRAW_STATE_PROCESSING, nil, nil, "100", "1", exchange.ORDERBY_ASC)
		if e != nil {
			t.Errorf("GetWithdraws error : %s", e.Error())
		} else {
			t.Logf("GetWithdraws[state:%s,remaining:%+v]", exchange.WITHDRAW_STATE_PROCESSING, *remaining)
			for _, withdraw := range withdraws {
				t.Logf("%+v", *withdraw)
				uuids = append(uuids, withdraw.UUID)
			}
		}

		withdraws, remaining, e = u.GetWithdraws(currency, exchange.WITHDRAW_STATE_REJECTED, nil, nil, "100", "1", exchange.ORDERBY_ASC)
		if e != nil {
			t.Errorf("GetWithdraws error : %s", e.Error())
		} else {
			t.Logf("GetWithdraws[state:%s,remaining:%+v]", exchange.WITHDRAW_STATE_REJECTED, *remaining)
			for _, withdraw := range withdraws {
				t.Logf("%+v", *withdraw)
				uuids = append(uuids, withdraw.UUID)
			}
		}

		withdraws, remaining, e = u.GetWithdraws(currency, exchange.WITHDRAW_STATE_SUBMITTED, nil, nil, "100", "1", exchange.ORDERBY_ASC)
		if e != nil {
			t.Errorf("GetWithdraws error : %s", e.Error())
		} else {
			t.Logf("GetWithdraws[state:%s,remaining:%+v]", exchange.WITHDRAW_STATE_SUBMITTED, *remaining)
			for _, withdraw := range withdraws {
				t.Logf("%+v", *withdraw)
				uuids = append(uuids, withdraw.UUID)
			}
		}

		withdraws, remaining, e = u.GetWithdraws(currency, exchange.WITHDRAW_STATE_SUBMITTING, nil, nil, "100", "1", exchange.ORDERBY_ASC)
		if e != nil {
			t.Errorf("GetWithdraws error : %s", e.Error())
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
				t.Errorf("GetWithdraw error : %s", e.Error())
			} else {
				if withdraw != nil {
					t.Logf("GetWithdraw[uuid:%s, remaining:%+v]\n%+v", uuids[0], *remaining, *withdraw)
				}
			}
		}

		withdrawChance, remaining, e := u.GetWithdrawChance(currency)
		if e != nil {
			t.Errorf("GetWithdrawChance error : %s", e.Error())
		} else {
			t.Logf("GetWithdrawChance[currency:%s,remaining:%+v]\n%+v", currency, *remaining, *withdrawChance)
		}

		/*
			coin, remaining, e := u.WithdrawCoin(currency, "0.00000001", "", "", exchange.WITHDRAW_TRANSACTION_DEFAULT)
			if e != nil {
				t.Errorf("WithdrawCoin error : %s", e.Error())
			} else {
				t.Logf("WithdrawCoin[currency:%s,remaining:%+v]\n%+v", currency, *remaining, *coin)
			}
		*/

		/*
			krw, remaining, e := u.WithdrawKrw("4000")
			if e != nil {
				t.Errorf("WithdrawKrw error : %s", e.Error())
			} else {
				t.Logf("WithdrawKrw[remaining:%+v]\n%+v", *remaining, *krw)
			}
		*/
	}
}
