package ibkr

import (
	"fmt"
	"net/url"
	"strings"
)

type OrderMonitoringServiceI interface {
	GetLiveOrders(param GetLiveOrdersParam) (*GetLiveOrdersResponse, error)
	GetTrades(param GetTradesParam) (*[]TradeItem, error)
}

type OrderMonitoringService struct {
	client *Client
}

func (s *OrderMonitoringService) GetTrades(param GetTradesParam) (*[]TradeItem, error) {

	var (
		res []TradeItem
	)

	urlParam := url.Values{}
	if param.Days != nil {
		urlParam.Add("days", fmt.Sprintf("%d", *param.Days))
	}
	if err := s.client.getPublic("/iserver/account/trades", urlParam, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

func (s *OrderMonitoringService) GetLiveOrders(param GetLiveOrdersParam) (*GetLiveOrdersResponse, error) {

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
	}
	urlParam.Add("force", fmt.Sprintf("%t", param.Force))

	if err := s.client.getPublic("/iserver/account/orders", urlParam, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

type GetTradesParam struct {
	// Specify the number of days to receive executions for, up to a maximum of 7 days.
	// If unspecified, only the current day is returned.
	Days *int
}

type TradeItem struct {
	ExecutionId          string  `json:"execution_id,omitempty"`
	Symbol               string  `json:"symbol,omitempty"`
	SupportsTaxOpt       string  `json:"supports_tax_opt,omitempty"`
	Side                 string  `json:"side,omitempty"`
	OrderDescription     string  `json:"order_description,omitempty"`
	OrderRef             string  `json:"order_ref,omitempty"`
	TradeTime            string  `json:"trade_time,omitempty"`
	TradeTimeR           int64   `json:"trade_time_r,omitempty"`
	Size                 float64 `json:"size,omitempty"`
	Price                string  `json:"price,omitempty"`
	Submitter            string  `json:"submitter,omitempty"`
	Exchange             string  `json:"exchange,omitempty"`
	Commission           string  `json:"commission,omitempty"`
	NetAmount            float64 `json:"net_amount,omitempty"`
	Account              string  `json:"account,omitempty"`
	AccountCode          string  `json:"accountCode,omitempty"`
	CompanyName          string  `json:"company_name,omitempty"`
	ContractDescription1 string  `json:"contract_description_1,omitempty"`
	SecType              string  `json:"sec_type,omitempty"`
	ListingExchange      string  `json:"listing_exchange,omitempty"` // Returns the primary listing exchange of the contract.
	ContractId           int     `json:"conid,omitempty"`
	ContractIdExchange   string  `json:"conidEx,omitempty"`
	ClearingId           string  `json:"clearing_id,omitempty"`
	ClearingName         string  `json:"clearing_name,omitempty"`
	LiquidationTrade     string  `json:"liquidation_trade,omitempty"`
	IsEventTrading       string  `json:"is_event_trading,omitempty"`
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
