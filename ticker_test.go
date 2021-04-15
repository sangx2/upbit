package upbit

import "testing"

func TestTicker(t *testing.T) {
	u := NewUpbit("", "")

	tickers, remaining, e := u.GetTickers([]string{marketID})
	if e != nil {
		t.Fatalf("%s's GetTickers error : %s", marketID, e.Error())
	} else {
		t.Logf("GetTickers[remaining:%+v]", *remaining)
		for _, ticker := range tickers {
			t.Logf("%+v", *ticker)
		}
	}
}
