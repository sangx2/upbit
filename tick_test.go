package upbit

import (
	"testing"
	"time"
)

func TestTicks(t *testing.T) {
	u := NewUpbit("", "")

	ticks, remaining, e := u.GetTradeTicks(marketID, time.Now().Format("15:04:05"), "", "")
	if e != nil {
		t.Fatalf("%s's GetTradeTicks error : %s", marketID, e.Error())
	} else {
		t.Logf("GetTradeTicks[remaining:%+v]", *remaining)
		for _, tick := range ticks {
			t.Logf("%+v", *tick)
		}
	}
}
