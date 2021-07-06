package tr

import (
	"context"
	"encoding/json"
)

// CreateOrderService create order
type CreateOrderService struct {
	c               *Client
	symbol          string
	side            SideType
	orderType       OrderType
	quantity        *string
	quoteOrderQty   *string
	price           *string
	clientId        *string
	stopPrice       *string
	icebergQuantity *string
}

type CreateOrderResponse struct {
	Code    int    `json:"code"`
	Message string `json:"msg"`
	Data    *struct {
		OrderID    string `json:"orderId"`
		CreateTime int64  `json:"createTime"`
	} `json:"data"`
	Timestamp int64 `json:"timestamp"`
}

// Symbol set symbol
func (s *CreateOrderService) Symbol(symbol string) *CreateOrderService {
	s.symbol = symbol
	return s
}

// Side set side
func (s *CreateOrderService) Side(side SideType) *CreateOrderService {
	s.side = side
	return s
}

// Type set type
func (s *CreateOrderService) Type(orderType OrderType) *CreateOrderService {
	s.orderType = orderType
	return s
}

// Quantity set quantity
func (s *CreateOrderService) Quantity(quantity string) *CreateOrderService {
	s.quantity = &quantity
	return s
}

// QuoteOrderQty set clientId
func (s *CreateOrderService) QuoteOrderQty(quoteOrderQty string) *CreateOrderService {
	s.clientId = &quoteOrderQty
	return s
}

// Price set price
func (s *CreateOrderService) Price(price string) *CreateOrderService {
	s.price = &price
	return s
}

// ClientId set clientId
func (s *CreateOrderService) ClientId(clientId string) *CreateOrderService {
	s.clientId = &clientId
	return s
}

// StopPrice set stopPrice
func (s *CreateOrderService) StopPrice(stopPrice string) *CreateOrderService {
	s.stopPrice = &stopPrice
	return s
}

// IcebergQuantity set icebergQuantity
func (s *CreateOrderService) IcebergQuantity(icebergQuantity string) *CreateOrderService {
	s.icebergQuantity = &icebergQuantity
	return s
}

func (s *CreateOrderService) createOrder(ctx context.Context, endpoint string, opts ...RequestOption) (data []byte, err error) {
	r := &request{
		method:   "POST",
		endpoint: endpoint,
		secType:  secTypeSigned,
	}
	m := params{
		"symbol": s.symbol,
		"side":   s.side,
		"type":   s.orderType,
	}

	if s.quantity != nil {
		m["quantity"] = *s.quantity
	}

	if s.quoteOrderQty != nil {
		m["quoteOrderQty"] = *s.quoteOrderQty
	}

	if s.price != nil {
		m["price"] = *s.price
	}

	if s.clientId != nil {
		m["clientId"] = *s.clientId
	}

	if s.stopPrice != nil {
		m["stopPrice"] = *s.stopPrice
	}

	if s.icebergQuantity != nil {
		m["icebergQty"] = *s.icebergQuantity
	}

	r.setFormParams(m)

	data, err = s.c.callAPI(ctx, r, opts...)

	if err != nil {
		return []byte{}, err
	}

	return data, nil
}

// Do send request
func (s *CreateOrderService) Do(ctx context.Context, opts ...RequestOption) (res *CreateOrderResponse, err error) {
	data, err := s.createOrder(ctx, "/open/v1/orders", opts...)

	if err != nil {
		return nil, err
	}

	res = new(CreateOrderResponse)
	err = json.Unmarshal(data, res)

	if err != nil {
		return nil, err
	}

	return res, nil
}

// GetOrderService get an order
type GetOrderService struct {
	c       *Client
	orderID *int64
}

type GetOrderDetailResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    *struct {
		OrderID          int             `json:"orderId"`
		OrderListID      int             `json:"orderListId"`
		ClientID         string          `json:"clientId"`
		Symbol           string          `json:"symbol"`
		Side             SideType        `json:"side"`
		Type             OrderType       `json:"type"`
		Price            float64         `json:"price"`
		Status           OrderStatusType `json:"status"`
		OrigQty          float64         `json:"origQty"`
		OrigQuoteQty     float64         `json:"origQuoteQty"`
		ExecutedQty      float64         `json:"executedQty"`
		ExecutedPrice    float64         `json:"executedPrice"`
		ExecutedQuoteQty float64         `json:"executedQuoteQty"`
		CreateTime       int64           `json:"createTime"`
	} `json:"data"`
	Timestamp int64 `json:"timestamp"`
}

// OrderID set orderID
func (s *GetOrderService) OrderID(orderID int64) *GetOrderService {
	s.orderID = &orderID
	return s
}

// Do send request
func (s *GetOrderService) Do(ctx context.Context, opts ...RequestOption) (res *GetOrderDetailResponse, err error) {
	r := &request{
		method:   "GET",
		endpoint: "/open/v1/orders/detail",
		secType:  secTypeSigned,
	}

	if s.orderID != nil {
		r.setParam("orderId", *s.orderID)
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(GetOrderDetailResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// CancelOrderService cancel an order
type CancelOrderService struct {
	c       *Client
	orderID int64
}

type CancelOrderResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    *struct {
		OrderID          int             `json:"orderId"`
		OrderListID      int             `json:"orderListId"`
		ClientID         string          `json:"clientId"`
		Symbol           string          `json:"symbol"`
		Side             SideType        `json:"side"`
		Type             OrderType       `json:"type"`
		Price            float64         `json:"price"`
		Status           OrderStatusType `json:"status"`
		OrigQty          float64         `json:"origQty"`
		OrigQuoteQty     float64         `json:"origQuoteQty"`
		ExecutedQty      float64         `json:"executedQty"`
		ExecutedPrice    float64         `json:"executedPrice"`
		ExecutedQuoteQty float64         `json:"executedQuoteQty"`
		CreateTime       int64           `json:"createTime"`
	} `json:"data"`
	Timestamp int64 `json:"timestamp"`
}

// OrderID set orderID
func (s *CancelOrderService) OrderID(orderID int64) *CancelOrderService {
	s.orderID = orderID
	return s
}

// Do send request
func (s *CancelOrderService) Do(ctx context.Context, opts ...RequestOption) (res *CancelOrderResponse, err error) {
	r := &request{
		method:   "POST",
		endpoint: "/open/v1/orders/cancel",
		secType:  secTypeSigned,
	}

	r.setParam("orderId", s.orderID)

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(CancelOrderResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// ListOrdersService all account orders; active, canceled, or filled
type ListOrdersService struct {
	c              *Client
	symbol         string
	listOrdersType *ListOrdersType
	side           *SideType
	startTime      *int64
	endTime        *int64
	fromId         *string
	direct         *ListOrdersSearchDirection
	limit          *int
}

type ListOrdersResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data *struct {
		List []struct {
			OrderID          string          `json:"orderId"`
			ClientID         string          `json:"clientId"`
			Symbol           string          `json:"symbol"`
			SymbolType       SymbolType      `json:"symbolType"`
			Side             SideType        `json:"side"`
			Type             OrderType       `json:"type"`
			Price            string          `json:"price"`
			OrigQty          string          `json:"origQty"`
			OrigQuoteQty     string          `json:"origQuoteQty"`
			ExecutedQty      string          `json:"executedQty"`
			ExecutedPrice    string          `json:"executedPrice"`
			ExecutedQuoteQty string          `json:"executedQuoteQty"`
			TimeInForce      TimeInForceType `json:"timeInForce"`
			StopPrice        string          `json:"stopPrice"`
			IcebergQty       string          `json:"icebergQty"`
			Status           int             `json:"status"`
			IsWorking        int             `json:"isWorking"`
			CreateTime       int64           `json:"createTime"`
		} `json:"list"`
	} `json:"data"`
	Timestamp int64 `json:"timestamp"`
}

// Symbol set symbol
func (s *ListOrdersService) Symbol(symbol string) *ListOrdersService {
	s.symbol = symbol
	return s
}

// ListOrdersType set listOrdersType
func (s *ListOrdersService) ListOrdersType(listOrdersType ListOrdersType) *ListOrdersService {
	s.listOrdersType = &listOrdersType
	return s
}

// Side set side
func (s *ListOrdersService) Side(side SideType) *ListOrdersService {
	s.side = &side
	return s
}

// StartTime set starttime
func (s *ListOrdersService) StartTime(startTime int64) *ListOrdersService {
	s.startTime = &startTime
	return s
}

// EndTime set endtime
func (s *ListOrdersService) EndTime(endTime int64) *ListOrdersService {
	s.endTime = &endTime
	return s
}

// FromId set fromId
func (s *ListOrdersService) FromId(fromId string) *ListOrdersService {
	s.fromId = &fromId
	return s
}

// Direction set direction
func (s *ListOrdersService) Direction(direction ListOrdersSearchDirection) *ListOrdersService {
	s.direct = &direction
	return s
}

// Limit set limit
func (s *ListOrdersService) Limit(limit int) *ListOrdersService {
	s.limit = &limit
	return s
}

// Do send request
func (s *ListOrdersService) Do(ctx context.Context, opts ...RequestOption) (res []*ListOrdersResponse, err error) {
	r := &request{
		method:   "GET",
		endpoint: "/open/v1/orders",
		secType:  secTypeSigned,
	}

	r.setParam("symbol", s.symbol)

	if s.listOrdersType != nil {
		r.setParam("listOrdersType", s.listOrdersType)
	}

	if s.side != nil {
		r.setParam("side", s.side)
	}

	if s.startTime != nil {
		r.setParam("startTime", *s.startTime)
	}

	if s.endTime != nil {
		r.setParam("endTime", *s.endTime)
	}

	if s.fromId != nil {
		r.setParam("fromId", s.fromId)
	}

	if s.direct != nil {
		r.setParam("direct", s.direct)
	}

	if s.limit != nil {
		r.setParam("limit", *s.limit)
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []*ListOrdersResponse{}, err
	}
	res = make([]*ListOrdersResponse, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return []*ListOrdersResponse{}, err
	}
	return res, nil
}
