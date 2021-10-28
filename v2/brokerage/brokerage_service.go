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
