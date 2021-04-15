package upbit

import (
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"github.com/sangx2/upbit/model"
	"hash"
	"net/http"
	"net/url"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

const (
	BaseURI = "https://api.upbit.com/v1"
)

// Upbit :
type Upbit struct {
	accessKey string
	secretKey string

	queryHash hash.Hash

	defaultClient      *http.Client // Group:default Min:900 Sec:30
	orderClient        *http.Client // Group:order Min:200 Sec:8
	statusWalletClient *http.Client // Group:status-wallet Min:30 Sec:1

	marketClient     *http.Client // Group:market Min:600 Sec:10
	candlesClient    *http.Client // Group:candles Min:600 Sec:10
	crixTradesClient *http.Client // Group:crix-trades Min:600 Sec:10
	tickerClient     *http.Client // Group:ticker Min:600 Sec:10
	orderbookClient  *http.Client // Group:orderbook Min:600 Sec:10
}

// NewUpbit :
func NewUpbit(accessKey, secretKey string) *Upbit {
	return &Upbit{
		accessKey: accessKey, secretKey: secretKey,
		queryHash: sha512.New(),

		defaultClient:      &http.Client{},
		orderClient:        &http.Client{},
		statusWalletClient: &http.Client{},

		marketClient:     &http.Client{},
		candlesClient:    &http.Client{},
		crixTradesClient: &http.Client{},
		tickerClient:     &http.Client{},
		orderbookClient:  &http.Client{},
	}
}

func (u *Upbit) createRequest(method, url string, values url.Values, section string) (*http.Request, error) {
	var request *http.Request

	switch method {
	case http.MethodGet, http.MethodDelete:
		req, e := http.NewRequest(method, url+"?"+values.Encode(), nil)
		if e != nil {
			return nil, e
		}
		request = req
	case http.MethodPost:
		req, e := http.NewRequest(method, url, strings.NewReader(values.Encode()))
		if e != nil {
			return nil, e
		}
		request = req

		request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	default:
		return nil, fmt.Errorf("invalid request method")
	}

	claim := jwt.MapClaims{
		"access_key": u.accessKey,
		"nonce":      uuid.New().String(),
	}

	switch section {
	case ApiSectionExchange:
		if len(values) != 0 {
			claim["query"] = values.Encode()
			u.queryHash.Reset()
			u.queryHash.Write([]byte(values.Encode()))
			claim["query_hash"] = hex.EncodeToString(u.queryHash.Sum(nil))
			claim["query_hash_alg"] = "SHA512"
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
		signedToken, e := token.SignedString([]byte(u.secretKey[:]))
		if e != nil {
			return nil, nil
		}

		request.Header.Add("Authorization", "Bearer "+signedToken)
	case ApiSectionQuotation:
	default:
		return nil, fmt.Errorf("invalid api section")
	}

	return request, nil
}

func (u *Upbit) do(request *http.Request, apiGroup string) (*http.Response, error) {
	var client *http.Client
	switch apiGroup {
	// Exchange
	case ApiGroupDefault:
		client = u.defaultClient
	case ApiGroupOrder:
		client = u.orderClient
	case ApiGroupStatusWallet:
		client = u.statusWalletClient
	// Quotation
	case ApiGroupMarket:
		client = u.marketClient
	case ApiGroupCandles:
		client = u.candlesClient
	case ApiGroupCrixTrades:
		client = u.crixTradesClient
	case ApiGroupTicker:
		client = u.tickerClient
	case ApiGroupOrderbook:
		client = u.orderbookClient
	default:
		return nil, fmt.Errorf("invalid api group")
	}

	response, e := client.Do(request)
	if e != nil {
		return nil, e
	}

	switch response.StatusCode {
	case 200: // ok
	case 201: // created
	default:
		respErr := model.ResponseErrorFromJSON(response.Body)
		if respErr == nil {
			return nil, fmt.Errorf("ResponseErrorFromJSON is nil")
		}
		return nil, fmt.Errorf("[%d:%s] %s:%s", response.StatusCode, response.Status, respErr.Detail.Name, respErr.Detail.Message)
	}

	return response, nil
}
