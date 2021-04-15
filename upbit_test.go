package upbit

import (
	"os"
	"testing"
)

// fix me
var (
	accessKey = ""
	secretKey = ""

	marketID = "KRW-ETH"
	currency = "ETH"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}
