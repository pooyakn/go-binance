package brokerage

import (
	"context"
	"encoding/json"
)

type CreateSubAccountService struct {
	c   *Client
	tag *string
}

// Tag sets the tag parameter.
func (c *CreateSubAccountService) Tag(v string) *CreateSubAccountService {
	c.tag = &v
	return c
}

func (c *CreateSubAccountService) Do(ctx context.Context) (*CreateSubAccountResponse, error) {
	r := &request{
		method:   "POST",
		endpoint: "/sapi/v1/broker/subAccount",
		secType:  secTypeSigned,
	}

	if c.tag != nil {
		r.setParam("tag", *c.tag)
	}

	data, err := c.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}

	res := &CreateSubAccountResponse{}
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

type CreateSubAccountResponse struct {
	SubAccountID string `json:"subaccountId"`
	Email        string `json:"email"`
	Tag          string `json:"tag"`
}

type CreateApiKeyForSubAccountService struct {
	c            *Client
	subAccountID string `json:"subaccountId"`
	canTrade     bool   `json:"canTrade"`
	marginTrade  *bool  `json:"marginTrade"`
	futuresTrade *bool  `json:"futuresTrade"`
}

// SubAccountID sets the subAccountID parameter (MANDATORY).
func (c *CreateApiKeyForSubAccountService) SubAccountID(v string) *CreateApiKeyForSubAccountService {
	c.subAccountID = v
	return c
}

// CanTrade sets the canTrade parameter (MANDATORY).
func (c *CreateApiKeyForSubAccountService) CanTrade(v bool) *CreateApiKeyForSubAccountService {
	c.canTrade = v
	return c
}

// MarginTrade sets the marginTrade parameter
func (c *CreateApiKeyForSubAccountService) MarginTrade(v bool) *CreateApiKeyForSubAccountService {
	c.marginTrade = &v
	return c
}

// FuturesTrade sets the futuresTrade parameter
func (c *CreateApiKeyForSubAccountService) FuturesTrade(v bool) *CreateApiKeyForSubAccountService {
	c.futuresTrade = &v
	return c
}

func (c *CreateApiKeyForSubAccountService) Do(ctx context.Context) (*CreateApiKeyForSubAccountResponse, error) {
	r := &request{
		method:   "POST",
		endpoint: "/sapi/v1/broker/subAccountApi",
		secType:  secTypeSigned,
	}

	r.setParam("subAccountId", c.subAccountID)
	r.setParam("canTrade", c.canTrade)

	if c.marginTrade != nil {
		r.setParam("marginTrade", *c.marginTrade)
	}

	if c.futuresTrade != nil {
		r.setParam("futuresTrade", c.futuresTrade)
	}

	data, err := c.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}

	res := &CreateApiKeyForSubAccountResponse{}
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

type CreateApiKeyForSubAccountResponse struct {
	SubAccountID string `json:"subaccountId"`
	ApiKey       string `json:"apiKey"`
	SecretKey    string `json:"secretKey"`
	CanTrade     bool   `json:"canTrade"`
	MarginTrade  bool   `json:"marginTrade"`
	FuturesTrade bool   `json:"futuresTrade"`
}

type SubAccountTransferService struct {
	c            *Client
	fromID       *string `json:"fromId"`
	toID         *string `json:"toId"`
	clientTranId *string `json:"clientTranId"`
	asset        string  `json:"asset"`
	amount       string  `json:"amount"`
}

func (s *SubAccountTransferService) FromID(v string) *SubAccountTransferService {
	s.fromID = &v
	return s
}

func (s *SubAccountTransferService) ToID(v string) *SubAccountTransferService {
	s.toID = &v
	return s
}

func (s *SubAccountTransferService) ClientTranID(v string) *SubAccountTransferService {
	s.clientTranId = &v
	return s
}

func (s *SubAccountTransferService) Asset(v string) *SubAccountTransferService {
	s.asset = v
	return s
}

func (s *SubAccountTransferService) Amount(v string) *SubAccountTransferService {
	s.amount = v
	return s
}

func (s *SubAccountTransferService) Do(ctx context.Context) (*SubAccountTransferResponse, error) {
	r := &request{
		method:   "POST",
		endpoint: "/sapi/v1/broker/transfer",
		secType:  secTypeSigned,
	}

	r.setParam("asset", s.asset)
	r.setParam("amount", s.amount)

	if s.fromID != nil {
		r.setParam("fromId", *s.fromID)
	}

	if s.toID != nil {
		r.setParam("toId", *s.toID)
	}

	if s.clientTranId != nil {
		r.setParam("clientTranId", *s.clientTranId)
	}

	data, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}

	res := &SubAccountTransferResponse{}
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

type SubAccountTransferResponse struct {
	TxnID        string `json:"txnId"`
	ClientTranID string `json:"clientTranId"`
}

type GetSubAccountDepositHistoryService struct {
	c            *Client
	subAccountID *string `json:"subAccountId""`
	coin         *string `json:"coin"`
	status       *int    `json:"status"`
	startTime    *int64  `json:"startTime"`
	endTime      *int64  `json:"endTime"`
	limit        *int    `json:"limit"`
	offest       *int    `json:"offest"`
}

func (g *GetSubAccountDepositHistoryService) SubAccountID(v string) *GetSubAccountDepositHistoryService {
	g.subAccountID = &v
	return g
}

func (g *GetSubAccountDepositHistoryService) Coin(v string) *GetSubAccountDepositHistoryService {
	g.coin = &v
	return g
}

func (g *GetSubAccountDepositHistoryService) Status(v int) *GetSubAccountDepositHistoryService {
	g.status = &v
	return g
}

func (g *GetSubAccountDepositHistoryService) StartTime(v int64) *GetSubAccountDepositHistoryService {
	g.startTime = &v
	return g
}

func (g *GetSubAccountDepositHistoryService) EndTime(v int64) *GetSubAccountDepositHistoryService {
	g.endTime = &v
	return g
}

func (g *GetSubAccountDepositHistoryService) Limit(v int) *GetSubAccountDepositHistoryService {
	g.limit = &v
	return g
}

func (g *GetSubAccountDepositHistoryService) Offest(v int) *GetSubAccountDepositHistoryService {
	g.offest = &v
	return g
}

func (g *GetSubAccountDepositHistoryService) Do(ctx context.Context) ([]*GetSubAccountDepositHistoryResponse, error) {
	r := &request{
		method:   "GET",
		endpoint: "/sapi/v1/broker/subAccount/depositHist",
		secType:  secTypeSigned,
	}

	if g.subAccountID != nil {
		r.setParam("subaccountId", *g.subAccountID)
	}

	if g.coin != nil {
		r.setParam("coin", *g.coin)
	}

	if g.status != nil {
		r.setParam("status", *g.status)
	}

	if g.startTime != nil {
		r.setParam("startTime", *g.startTime)
	}

	if g.endTime != nil {
		r.setParam("endTime", *g.endTime)
	}

	if g.limit != nil {
		r.setParam("limit", *g.limit)
	}

	if g.offest != nil {
		r.setParam("offest", *g.offest)
	}

	data, err := g.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}

	res := make([]*GetSubAccountDepositHistoryResponse, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

type GetSubAccountDepositHistoryResponse struct {
	SubAccountID  string `json:"subaccountId"`
	Address       string `json:"address"`
	AddressTag    string `json:"addressTag"`
	Amount        string `json:"amount"`
	Coin          string `json:"coin"`
	InsertTime    int64  `json:"insertTime"`
	Network       string `json:"network"`
	Status        int    `json:"status"`
	TxID          string `json:"txId"`
	SourceAddress string `json:"sourceAddress"`
	ConfirmTimes  string `json:"confirmTimes"`
}
