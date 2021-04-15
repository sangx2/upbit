package upbit

import (
	"github.com/sangx2/upbit/model/exchange"
	"testing"
)

func TestDeposit(t *testing.T) {
	var uuids []string

	if len(accessKey) != 0 && len(secretKey) != 0 {
		u := NewUpbit(accessKey, secretKey)
		if u == nil {
			t.Fatalf("NewUpbit is nil")
		}

		deposits, remaining, e := u.GetDeposits(currency, exchange.DEPOSIT_STATE_ACCEPTED, nil, nil, "100", "1", exchange.ORDERBY_ASC)
		if e != nil {
			t.Errorf("GetDeposits error : %s", e.Error())
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
				t.Logf("GetDeposit error : %s", e.Error())
			} else {
				if deposit != nil {
					t.Logf("GetDeposit[remaining:%+v]\n%+v", *remaining, *deposit)
				}
			}
		}

		coinAddress, remaining, e := u.GenerateDepositCoinAddress(currency)
		if e != nil {
			t.Logf("GenerateDepositCoinAddress error : %s", e.Error())
		} else {
			if coinAddress != nil {
				t.Logf("GenerateDepositCoinAddress[remaining:%+v]\n%+v", *remaining, *coinAddress)
			}
		}

		coinAddresses, remaining, e := u.GetDepositCoinAddresses()
		if e != nil {
			t.Logf("GetDepositCoinAddresses error : %s", e.Error())
		} else {
			t.Logf("GetDepositCoinAddresses[remaining:%+v]", *remaining)
			for _, coinAddress := range coinAddresses {
				t.Logf("%+v", *coinAddress)
			}
		}

		coinAddress, remaining, e = u.GetDepositCoinAddress(currency)
		if e != nil {
			t.Fatalf("GetDepositCoinAddress error : %s", e.Error())
		} else {
			if coinAddress != nil {
				t.Logf("GetDepositCoinAddress[remaining:%+v]\n%+v", *remaining, *coinAddress)
			}
		}

		/*
			deposit, remaining, e := u.DepositKrw("5000")
			if e != nil {
				t.Fatalf("DepositKrw error : %s", e.Error())
			} else {
				if deposit != nil {
					t.Logf("DepositKrw[remaining:%+v]\n%+v", *remaining, *deposit)
				}
			}
		*/
	}
}
