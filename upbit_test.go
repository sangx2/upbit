package upbit

import (
	"os"
	"testing"
)

// fix me
var (
	accessKey = "8nSmG42bkK7CmgN8L23zARByPLy5c538PkhkgsOf"
	secretKey = "Z8aX43xRHcNXg2e2Aw8ijHzLW9n6WlNy5K5XENIq"

	marketID = "KRW-ETH"
	currency = "ETH"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}
