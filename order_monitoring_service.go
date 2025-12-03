package ibkr

import (
	"fmt"
	"net/url"
	"strings"
)

type OrderMonitoringServiceI interface {
	GetLiveOrders(param GetLiveOrdersParam) (*GetLiveOrdersResponse, error)
}

type OrderMonitoringService struct {
	client *Client
}

func (s OrderMonitoringService) GetLiveOrders(param GetLiveOrdersParam) (*GetLiveOrdersResponse, error) {

	var (
		res GetLiveOrdersResponse
	)

	urlParam := url.Values{}
	if param.StatusValueFilters != nil && len(param.StatusValueFilters) > 0 {
		filters := make([]string, 0)
		for _, value := range param.StatusValueFilters {
			filters = append(filters, string(value))
		}
		urlParam.Add("filters", strings.Join(filters, ","))
		urlParam.Add("force", fmt.Sprintf("%t", param.Force))
	}
	if err := s.client.getPublic("/iserver/account/orders", urlParam, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

type GetLiveOrdersParam struct {
	StatusValueFilters []OrderStatusFilterValue
	/* Please be aware that filtering orders using the /iserver/account/orders endpoint will prevent order details from coming through over the websocket “sor” topic. To resolve this issue, developers should set “force=true” in a follow-up /iserver/account/orders call to clear any cached behavior surrounding the endpoint prior to calling for the websocket request */
	Force bool
}

type LiveOrderItem struct {
	AccountId          string  `json:"acct"`
	ContractIdExchange string  `json:"conidEx,omitempty"`
	ContractId         int     `json:"conid,omitempty"`
	OrderId            int64   `json:"orderId,omitempty"`
	CashCcy            string  `json:"cashCcy,omitempty"`
	SizeAndFills       string  `json:"sizeAndFills,omitempty"`
	OrderDesc          string  `json:"orderDesc,omitempty"`
	Description1       string  `json:"description1,omitempty"` // Returns the local symbol of the order
	Ticker             string  `json:"ticker,omitempty"`
	SecurityType       string  `json:"secType,omitempty"`
	ListingExchange    string  `json:"listingExchange,omitempty"`
	RemainingQuantity  float64 `json:"remainingQuantity,omitempty"`
	FilledQuantity     float64 `json:"filledQuantity,omitempty"`
	CompanyName        string  `json:"companyName,omitempty"`
	Status             string  `json:"status,omitempty"`
	OrderCcpStatus     string  `json:"order_ccp_status,omitempty"`
	OrigOrderType      string  `json:"origOrderType,omitempty"`
	SupportsTaxOpt     string  `json:"supportsTaxOpt,omitempty"`
	LastExecutionTime  string  `json:"lastExecutionTime"`
	OrderType          string  `json:"orderType"`
	BgColor            string  `json:"bgColor,omitempty"`
	FgColor            string  `json:"fgColor,omitempty"`
	OrderRef           string  `json:"order_ref,omitempty"` // User defined string used to identify the order. Value is set using “cOID” field while placing an order.
	TimeInForce        string  `json:"timeInForce,omitempty"`
	Side               string  `json:"side,omitempty"`
	AveragePrice       float64 `json:"avgPrice,omitempty"`
}

type GetLiveOrdersResponse struct {
	Orders   []LiveOrderItem `json:"orders,omitempty"`
	Snapshot bool            `json:"snapshot"`
}
