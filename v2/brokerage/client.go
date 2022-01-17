package brokerage

import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/pooyakn/go-binance/v2/common"
)

// SideType define side type of order
type SideType int

// OrderType define order type
type OrderType int

// TimeInForceType define time in force type of order
type TimeInForceType int

// OrderStatusType define order status type
type OrderStatusType int

// SymbolType define symbol type
type SymbolType int64

// SideEffectType define side effect type for orders
type SideEffectType string

// ListOrdersType open, history, or all
type ListOrdersType int

// ListOrdersSearchDirection searching direction: prev - in ascending order from the start order ID; next - in descending order from the start order ID
type ListOrdersSearchDirection string

// Endpoints
const (
	baseAPIMainURL = "https://api.binance.com"
)

// Global enums
const (
	ListOrdersTypeOpen    ListOrdersType = 1
	ListOrdersTypeHistory ListOrdersType = 2
	ListOrdersTypeAll     ListOrdersType = -1

	ListOrdersSearchDirectionPrev ListOrdersSearchDirection = "prev"
	ListOrdersSearchDirectionNext ListOrdersSearchDirection = "next"

	SideTypeBuy  SideType = 0
	SideTypeSell SideType = 1

	OrderTypeLimit           OrderType = 1
	OrderTypeMarket          OrderType = 2
	OrderTypeStopLoss        OrderType = 3
	OrderTypeStopLossLimit   OrderType = 4
	OrderTypeTakeProfit      OrderType = 5
	OrderTypeTakeProfitLimit OrderType = 6
	OrderTypeLimitMaker      OrderType = 7

	OrderStatusTypeNew             OrderStatusType = 0
	OrderStatusTypePartiallyFilled OrderStatusType = 1
	OrderStatusTypeFilled          OrderStatusType = 2
	OrderStatusTypeCanceled        OrderStatusType = 3
	OrderStatusTypePendingCancel   OrderStatusType = 4
	OrderStatusTypeRejected        OrderStatusType = 5
	OrderStatusTypeExpired         OrderStatusType = 6

	SymbolTypeSpot SymbolType = 1

	TimeInForceTypeGTC TimeInForceType = 0
	TimeInForceTypeIOC TimeInForceType = 1
	TimeInForceTypeFOK TimeInForceType = 2

	timestampKey  = "timestamp"
	signatureKey  = "signature"
	recvWindowKey = "recvWindow"
)

func currentTimestamp() int64 {
	return FormatTimestamp(time.Now())
}

// FormatTimestamp formats a time into Unix timestamp in milliseconds, as requested by Binance.
func FormatTimestamp(t time.Time) int64 {
	return t.UnixNano() / int64(time.Millisecond)
}

// NewClient initialize an API client instance with API key and secret key.
// You should always call this function before using this SDK.
// Services will be created by the form client.NewXXXService().
func NewClient(apiKey, secretKey string) *Client {
	return &Client{
		APIKey:     apiKey,
		SecretKey:  secretKey,
		BaseURL:    baseAPIMainURL,
		UserAgent:  "Binance/golang",
		HTTPClient: http.DefaultClient,
		Logger:     log.New(os.Stderr, "Binance-golang ", log.LstdFlags),
	}
}

type doFunc func(req *http.Request) (*http.Response, error)

// Client define API client
type Client struct {
	APIKey     string
	SecretKey  string
	BaseURL    string
	UserAgent  string
	HTTPClient *http.Client
	Debug      bool
	Logger     *log.Logger
	TimeOffset int64
	do         doFunc
}

func (c *Client) debug(format string, v ...interface{}) {
	if c.Debug {
		c.Logger.Printf(format, v...)
	}
}

func (c *Client) parseRequest(r *request, opts ...RequestOption) (err error) {
	// set request options from user
	for _, opt := range opts {
		opt(r)
	}
	err = r.validate()
	if err != nil {
		return err
	}

	fullURL := fmt.Sprintf("%s%s", c.BaseURL, r.endpoint)
	if r.recvWindow > 0 {
		r.setParam(recvWindowKey, r.recvWindow)
	}
	if r.secType == secTypeSigned {
		r.setParam(timestampKey, currentTimestamp()-c.TimeOffset)
	}
	queryString := r.query.Encode()
	body := &bytes.Buffer{}
	bodyString := r.form.Encode()
	header := http.Header{}
	if r.header != nil {
		header = r.header.Clone()
	}
	if bodyString != "" {
		header.Set("Content-Type", "application/x-www-form-urlencoded")
		body = bytes.NewBufferString(bodyString)
	}
	if r.secType == secTypeAPIKey || r.secType == secTypeSigned {
		header.Set("X-MBX-APIKEY", c.APIKey)
	}

	if r.secType == secTypeSigned {
		raw := fmt.Sprintf("%s%s", queryString, bodyString)
		mac := hmac.New(sha256.New, []byte(c.SecretKey))
		_, err = mac.Write([]byte(raw))
		if err != nil {
			return err
		}
		v := url.Values{}
		v.Set(signatureKey, fmt.Sprintf("%x", (mac.Sum(nil))))
		if queryString == "" {
			queryString = v.Encode()
		} else {
			queryString = fmt.Sprintf("%s&%s", queryString, v.Encode())
		}
	}
	if queryString != "" {
		fullURL = fmt.Sprintf("%s?%s", fullURL, queryString)
	}
	c.debug("full url: %s, body: %s", fullURL, bodyString)

	r.fullURL = fullURL
	r.header = header
	r.body = body
	return nil
}

func (c *Client) callAPI(ctx context.Context, r *request, opts ...RequestOption) (data []byte, err error) {
	err = c.parseRequest(r, opts...)
	if err != nil {
		return []byte{}, err
	}
	req, err := http.NewRequest(r.method, r.fullURL, r.body)
	if err != nil {
		return []byte{}, err
	}
	req = req.WithContext(ctx)
	req.Header = r.header
	c.debug("request: %#v", req)
	f := c.do
	if f == nil {
		f = c.HTTPClient.Do
	}
	res, err := f(req)
	if err != nil {
		return []byte{}, err
	}
	data, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return []byte{}, err
	}
	defer func() {
		cerr := res.Body.Close()
		// Only overwrite the retured error if the original error was nil and an
		// error occurred while closing the body.
		if err == nil && cerr != nil {
			err = cerr
		}
	}()
	c.debug("response: %#v", res)
	c.debug("response body: %s", string(data))
	c.debug("response status code: %d", res.StatusCode)

	if res.StatusCode >= 400 {
		apiErr := new(common.APIError)
		e := json.Unmarshal(data, apiErr)
		if e != nil {
			c.debug("failed to unmarshal json: %s", e)
		}
		return nil, apiErr
	}
	return data, nil
}

// NewCreateSubAccountService init creating order service
func (c *Client) NewCreateSubAccountService() *CreateSubAccountService {
	return &CreateSubAccountService{c: c}
}

// NewCreateApiKeyForSubAccountService init get order service
func (c *Client) NewCreateApiKeyForSubAccountService() *CreateApiKeyForSubAccountService {
	return &CreateApiKeyForSubAccountService{c: c}
}

// NewGetSubAccountDepositHistoryService init get order service
func (c *Client) NewGetSubAccountDepositHistoryService() *GetSubAccountDepositHistoryService {
	return &GetSubAccountDepositHistoryService{c: c}
}

func (c *Client) NewSubAccountTransferService() *SubAccountTransferService {
	return &SubAccountTransferService{c: c}
}

func (c *Client) NewGetSubAccountTransferHistoryService() *GetSubAccountTransferHistoryService {
	return &GetSubAccountTransferHistoryService{c: c}
}
