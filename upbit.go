package upbit

import (
	"crypto/sha512"
	"encoding/hex"
	"errors"
	"fmt"
	"hash"
	"net/http"
	"net/url"
	"strings"

	"github.com/sangx2/upbit/model"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

const (
	URL_UPBIT_V1 = "https://api.upbit.com/v1"

	METHOD_GET    = "GET"
	METHOD_POST   = "POST"
	METHOD_DELETE = "DELETE"

	API_TYPE_EXCHANGE  = "exchange"
	API_TYPE_QUOTATION = "quotation"

	API_GROUP_DEFAULT     = "default"
	API_GROUP_ORDER       = "order"
	API_GROUP_MARKET      = "market"
	API_GROUP_CANDLES     = "candles"
	API_GROUP_CRIX_TRADES = "crix-trades"
	API_GROUP_TICKER      = "ticker"
	API_GROUP_ORDERBOOK   = "orderbook"
)

// Upbit :
type Upbit struct {
	accessKey string
	secretKey string

	queryHash hash.Hash

	defaultClient *http.Client // Group:default Min:1800 Sec:30
	orderClient   *http.Client //

	marketClient     *http.Client // Group:market Min:600 Sec:10
	candlesClient    *http.Client // Group:candles Min:600 Sec:10
	crixTradesClient *http.Client // Group:crix-trades Min:600 Sec:10
	tickerClient     *http.Client // Group:ticker Min:600 Sec:10
	orderbookClient  *http.Client // Group:orderbook Min:600 Sec:10

	ResponseError *model.ResponseError // for debug
}

// NewUpbit :
func NewUpbit(accessKey, secretKey string) *Upbit {
	return &Upbit{accessKey: accessKey, secretKey: secretKey,
		queryHash:     sha512.New(),
		defaultClient: &http.Client{}, orderClient: &http.Client{},
		marketClient: &http.Client{}, candlesClient: &http.Client{}, crixTradesClient: &http.Client{}, tickerClient: &http.Client{}, orderbookClient: &http.Client{},
	}
}

func (u *Upbit) getResponse(requestMethod, requestURL string, values url.Values, apiType, apiGroup string) (*http.Response, error) {
	var request *http.Request

	claim := jwt.MapClaims{
		"access_key": u.accessKey,
		"nonce":      uuid.New().String(),
	}

	switch requestMethod {
	case METHOD_GET, METHOD_DELETE:
		req, e := http.NewRequest(requestMethod, requestURL+"?"+values.Encode(), nil)
		if e != nil {
			return nil, e
		}
		request = req
	case METHOD_POST:
		req, e := http.NewRequest(requestMethod, requestURL, strings.NewReader(values.Encode()))
		if e != nil {
			return nil, e
		}
		request = req

		request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	default:
		return nil, errors.New("invalid request method")
	}

	switch apiType {
	case API_TYPE_EXCHANGE:
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
	case API_TYPE_QUOTATION:
	default:
		return nil, errors.New("invalid api type")
	}

	var client *http.Client
	switch apiGroup {
	case API_GROUP_DEFAULT:
		client = u.defaultClient
	case API_GROUP_ORDER:
		client = u.orderClient

	case API_GROUP_MARKET:
		client = u.marketClient
	case API_GROUP_CANDLES:
		client = u.candlesClient
	case API_GROUP_CRIX_TRADES:
		client = u.crixTradesClient
	case API_GROUP_TICKER:
		client = u.tickerClient
	case API_GROUP_ORDERBOOK:
		client = u.orderbookClient
	}

	response, e := client.Do(request)
	if e != nil {
		return nil, e
	}

	switch response.StatusCode {
	case 200: // ok
	case 201: // created
	default:
		u.ResponseError = model.ResponseErrorFromJSON(response.Body)
		if u.ResponseError == nil {
			return nil, errors.New("model.ResponseErrorFromJSON is nil")
		}
		return nil, fmt.Errorf("[%d:%s]%s:%s", response.StatusCode, response.Status, u.ResponseError.Detail.Name, u.ResponseError.Detail.Message)
	}

	return response, nil
}
