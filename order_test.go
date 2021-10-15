package upbit

import (
	"github.com/sangx2/upbit/model/exchange"
	"testing"
)

func TestOrder(t *testing.T) {
	var uuids []string

	if len(accessKey) != 0 && len(secretKey) != 0 {
		u := NewUpbit(accessKey, secretKey)

		chance, remaining, e := u.GetOrderChance(marketID)
		if e != nil {
			t.Logf("GetOrderChance error : %s", e.Error())
		} else {
			t.Logf("GetOrderChance[remaining:%+v]\n%+v", *remaining, *chance)
		}

		orders, remaining, e := u.GetOrders(marketID, exchange.ORDER_STATE_WAIT, nil, nil, nil, "1", "100", exchange.ORDERBY_ASC)
		if e != nil {
			t.Logf("GetOrders error : %s", e.Error())
		} else {
			t.Logf("GetOrders[state:%s,remaining:%+v]\n", exchange.ORDER_STATE_WAIT, *remaining)
			for _, order := range orders {
				t.Logf("%+v", *order)
				uuids = append(uuids, order.UUID)
			}
		}

		orders, remaining, e = u.GetOrders(marketID, exchange.ORDER_STATE_DONE, nil, nil, nil, "1", "100", exchange.ORDERBY_ASC)
		if e != nil {
			t.Logf("GetOrders error : %s", e.Error())
		} else {
			t.Logf("GetOrders[state:%s,remaining:%+v]\n", exchange.ORDER_STATE_DONE, *remaining)
			for _, order := range orders {
				t.Logf("%+v", *order)
				uuids = append(uuids, order.UUID)
			}
		}

		orders, remaining, e = u.GetOrders(marketID, exchange.ORDER_STATE_CANCEL, nil, nil, nil, "1", "100", exchange.ORDERBY_ASC)
		if e != nil {
			t.Logf("GetOrders error : %s", e.Error())
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
				t.Logf("GetOrder error : %s", e.Error())
			} else {
				t.Logf("GetOrder[remaining:%+v]\n%+v", *remaining, *order)
			}
		}

		/*
			purchaseOrder, remaining, e := u.PurchaseOrder(marketID, "", "5000", exchange.ORDER_TYPE_PRICE, "")
			if e != nil {
				t.Errorf("PurchaseOrder error : %s", e.Error())
			} else {
				t.Logf("PurchaseOrder[remaining:%+v]\n%+v", *remaining, *purchaseOrder)
			}
		*/
		/*
			cancelOrder, remaining, e := u.CancelOrder("b2f1b30d-7bd3-4fdb-b354-b03ed3c8c57b", "")
			if e != nil {
				t.Errorf("CancelOrder error : %s", e.Error())
			} else {
				t.Logf("CancelOrder[remaining:%+v]\n%+v", *remaining, *cancelOrder)
			}
		*/
		/*
			sellOrder, remaining, e := u.SellOrder(marketID, "0.00160384", "", exchange.ORDER_TYPE_MARKET, "")
			if e != nil {
				t.Errorf("SellOrder error : %s", e.Error())
			} else {
				t.Logf("SellOrder[remaining:%+v]\n%+v", *remaining, *sellOrder)
			}
		*/
	}
}
