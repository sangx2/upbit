package upbit

import "testing"

func TestAsset(t *testing.T) {
	if len(accessKey) != 0 && len(secretKey) != 0 {
		u := NewUpbit(accessKey, secretKey)

		// account
		accounts, remaining, e := u.GetAccounts()
		if e != nil {
			t.Logf("GetAccounts error : %s", e.Error())
		} else {
			t.Logf("GetAccounts[remaining:%+v]", *remaining)
			for _, account := range accounts {
				t.Logf("%+v", *account)
			}
		}
	}
}
