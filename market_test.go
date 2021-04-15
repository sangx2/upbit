package upbit

import "testing"

func TestMarket(t *testing.T) {
	u := NewUpbit("", "")

	markets, remaining, e := u.GetMarkets()
	if e != nil || len(markets) == 0 {
		t.Fatalf("GetMarkets error : %s", e.Error())
	} else {
		t.Logf("GetMarkets[remaining:%+v]", *remaining)
		for _, market := range markets {
			t.Logf("%+v", *market)
		}
	}
}
