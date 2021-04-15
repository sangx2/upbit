package upbit

import "testing"

func TestService(t *testing.T) {
	if len(accessKey) != 0 && len(secretKey) != 0 {
		u := NewUpbit(accessKey, secretKey)

		wallets, remaining, e := u.GetWalletStatus()
		if e != nil {
			t.Logf("GetWalletStatus error : %s", e.Error())
		} else {
			t.Logf("GetWalletStatus[remaining:%+v]", *remaining)
			for _, wallet := range wallets {
				t.Logf("%+v", *wallet)
			}
		}

		apiKeys, remaining, e := u.GetApiKeys()
		if e != nil {
			t.Logf("GetApiKeys error : %s", e.Error())
		} else {
			t.Logf("GetApiKeys[remaining:%+v]", *remaining)
			for _, apiKey := range apiKeys {
				t.Logf("%+v", *apiKey)
			}
		}
	}
}
