package withdraw

import (
	"encoding/json"
	"io"

	"github.com/sangx2/upbit/model/exchange/account"
)

// Chance : 출금 가능 정보
type Chance struct {
	MemberLevel struct {
		SecurityLevel        int  `json:"security_level"`
		FeeLevel             int  `json:"fee_level"`
		EmailVerified        bool `json:"email_verified"`
		IdentityAuthVerified bool `json:"identity_auth_verified"`
		BankAccountVerified  bool `json:"bank_account_verified"`
		KakaoPayAuthVerified bool `json:"kakao_pay_auth_verified"`
		Locked               bool `json:"locked"`
		WalletLocked         bool `json:"wallet_locked"`
	} `json:"member_level"`
	Currency struct {
		Code          string   `json:"code"`
		WithdrawFee   string   `json:"withdraw_fee"`
		IsCoin        bool     `json:"is_coin"`
		WalletState   string   `json:"wallet_state"`
		WalletSupport []string `json:"wallet_support"`
	} `json:"currency"`
	Account       account.Account `json:"account"`
	WithdrawLimit struct {
		Currency          string `json:"currency"`
		Minimum           string `json:"minimum"`
		Onetime           string `json:"onetime"`
		Daily             string `json:"daily"`
		RemainingDaily    string `json:"remaining_daily"`
		RemainingDailyKrw string `json:"remaining_daily_krw"`
		Fixed             int    `json:"fixed"`
		CanWithdraw       bool   `json:"can_withdraw"`
	} `json:"withdraw_limit"`
}

func ChanceFromJSON(r io.Reader) *Chance {
	var c *Chance

	json.NewDecoder(r).Decode(&c)

	return c
}
