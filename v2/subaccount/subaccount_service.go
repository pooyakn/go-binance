package subaccount

import (
	"context"
	"encoding/json"
)

type GetSubAccountListService struct {
	c        *Client
	email    *string `json:"email"`
	isFreeze *string `json:"isFreeze"`
	page     *int    `json:"page"`
	limit    *int    `json:"limit"`
}

func (g *GetSubAccountListService) Email(v string) *GetSubAccountListService {
	g.email = &v
	return g
}

func (g *GetSubAccountListService) IsFreeze(v string) *GetSubAccountListService {
	g.isFreeze = &v
	return g
}

func (g *GetSubAccountListService) Page(v int) *GetSubAccountListService {
	g.page = &v
	return g
}

func (g *GetSubAccountListService) Limit(v int) *GetSubAccountListService {
	g.limit = &v
	return g
}

func (g *GetSubAccountListService) Do(ctx context.Context) (res *GetSubAccountListResponse, err error) {
	r := &request{
		method:   "GET",
		endpoint: "/sapi/v1/sub-account/list",
		secType:  secTypeSigned,
	}

	if g.email != nil {
		r.setParam("email", g.email)
	}

	if g.isFreeze != nil {
		r.setParam("isFreeze", g.isFreeze)
	}

	if g.page != nil {
		r.setParam("page", g.page)
	}

	if g.limit != nil {
		r.setParam("limit", g.limit)
	}

	data, err := g.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}

	res = new(GetSubAccountListResponse)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}

	return
}

type GetSubAccountListResponse struct {
	SubAccounts []*SubAccount `json:"subAccounts"`
}

type SubAccount struct {
	Email      string `json:"email"`
	IsFreeze   bool   `json:"isFreeze"`
	CreateTime int64  `json:"createTime"`
}
