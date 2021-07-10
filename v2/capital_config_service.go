package binance

import (
	"context"
	"encoding/json"
)

// CapitalConfigService get capital config
type CapitalConfigService struct {
	c *Client
}

type Network struct {
	Network                 string `json:"network"`
	Coin                    string `json:"coin"`
	WithdrawIntegerMultiple string `json:"withdrawIntegerMultiple"`
	IsDefault               bool   `json:"isDefault"`
	DepositEnable           bool   `json:"depositEnable"`
	WithdrawEnable          bool   `json:"withdrawEnable"`
	DepositDesc             string `json:"depositDesc"`
	WithdrawDesc            string `json:"withdrawDesc"`
	SpecialTips             string `json:"specialTips"`
	Name                    string `json:"name"`
	ResetAddressStatus      bool   `json:"resetAddressStatus"`
	AddressRegex            string `json:"addressRegex"`
	MemoRegex               string `json:"memoRegex"`
	WithdrawFee             string `json:"withdrawFee"`
	WithdrawMin             string `json:"withdrawMin"`
	WithdrawMax             string `json:"withdrawMax"`
	MinConfirm              int    `json:"minConfirm"`
	UnLockConfirm           int    `json:"unLockConfirm"`
}

type CapitalConfigResponse struct {
	Coin              string    `json:"coin"`
	DepositAllEnable  bool      `json:"depositAllEnable"`
	WithdrawAllEnable bool      `json:"withdrawAllEnable"`
	Name              string    `json:"name"`
	Free              string    `json:"free"`
	Locked            string    `json:"locked"`
	Freeze            string    `json:"freeze"`
	Withdrawing       string    `json:"withdrawing"`
	Ipoing            string    `json:"ipoing"`
	Ipoable           string    `json:"ipoable"`
	Storage           string    `json:"storage"`
	IsLegalMoney      bool      `json:"isLegalMoney"`
	Trading           bool      `json:"trading"`
	NetworkList       []Network `json:"networkList"`
}

// Do send request
func (s *CapitalConfigService) Do(ctx context.Context, opts ...RequestOption) (res []CapitalConfigResponse, err error) {
	r := &request{
		method:   "GET",
		endpoint: "/sapi/v1/capital/config/getall",
		secType:  secTypeSigned,
	}

	data, err := s.c.callAPI(ctx, r, opts...)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &res)

	if err != nil {
		return nil, err
	}
	return res, nil
}
