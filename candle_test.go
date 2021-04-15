package upbit

import (
	"testing"
)

func TestCandle(t *testing.T) {
	u := NewUpbit("", "")

	// candle
	minuteCandles, remaining, e := u.GetMinuteCandles(marketID, "", "", "1")
	if e != nil {
		t.Fatalf("%s's GetMinuteCandles error : %s", marketID, e.Error())
	} else {
		t.Logf("GetMinuteCandles[remaining:%+v]", *remaining)
		for _, minuteCandle := range minuteCandles {
			t.Logf("%+v", *minuteCandle)
		}
	}

	dayCandles, remaining, e := u.GetDayCandles(marketID, "", "", "")
	if e != nil {
		t.Fatalf("%s's GetDayCandles error : %s", marketID, e.Error())
	} else {
		t.Logf("GetDayCandles[remaining:%+v]", *remaining)
		for _, dayCandle := range dayCandles {
			t.Logf("%+v", *dayCandle)
		}
	}

	weekCandles, remaining, e := u.GetWeekCandles(marketID, "", "")
	if e != nil {
		t.Fatalf("%s's GetWeekCandles error : %s", marketID, e.Error())
	} else {
		t.Logf("GetWeekCandles[remaining:%+v]", *remaining)
		for _, weekCandle := range weekCandles {
			t.Logf("%+v", *weekCandle)
		}
	}

	monthCandles, remaining, e := u.GetMonthCandles(marketID, "", "")
	if e != nil {
		t.Fatalf("%s's GetMonthCandles error : %s", marketID, e.Error())
	} else {
		t.Logf("GetMonthCandles[remaining:%+v]", *remaining)
		for _, monthCandle := range monthCandles {
			t.Logf("%+v", *monthCandle)
		}
	}
}
