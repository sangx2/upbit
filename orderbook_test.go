package upbit

import "testing"

func TestOrderbook(t *testing.T) {
	u := NewUpbit("", "")

	// orderbook
	orderbooks, remaining, e := u.GetOrderbooks([]string{marketID})
	if e != nil {
		t.Fatalf("%s's GetOrderbooks error : %s", marketID, e.Error())
	} else {
		t.Logf("GetOrderbooks[remaining:%+v]", *remaining)
		for _, orderbook := range orderbooks {
			t.Logf("%+v", *orderbook)
		}
	}
}
