package ibkr

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
)

type WebsocketPrivateOrderParam struct {
	Status []string
}

type WebsocketPrivateOrder struct {
	AccountId          string  `json:"acct,omitempty"`
	ContractId         int     `json:"conid,omitempty"`
	OrderId            int64   `json:"orderId,omitempty"`
	CashCcy            string  `json:"cashCcy,omitempty"`
	SizeAndFills       string  `json:"sizeAndFills,omitempty"`
	OrderDesc          string  `json:"orderDesc,omitempty"`
	Description1       string  `json:"description1,omitempty"`
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
	LastExecutionTimeR int64   `json:"lastExecutionTime_r,omitempty"` // Returns the epoch time of the most recent execution on the order.
	OrderType          string  `json:"orderType"`
	BgColor            string  `json:"bgColor,omitempty"`
	FgColor            string  `json:"fgColor,omitempty"`
	OrderRef           string  `json:"order_ref,omitempty"` // User defined string used to identify the order. Value is set using “cOID” field while placing an order.
	TimeInForce        string  `json:"timeInForce,omitempty"`
	Side               string  `json:"side,omitempty"`
	Price              float64 `json:"price,omitempty"`
}
type WebsocketPrivateOrderResponse struct {
	Orders   []WebsocketPrivateOrder `json:"orders"`
	Snapshot bool                    `json:"snapshot,omitempty"`
}

type WebsocketPrivatePnLResponse struct {
	Topic string      `json:"topic"`
	Args  interface{} `json:"args"`
}

type WebsocketPrivateTradesDataParam struct {
	RealtimeUpdatesOnly *bool
	Days                *int
}

type WebsocketPrivateTradesData struct {
	ExecutionId          string  `json:"execution_id,omitempty"`
	Symbol               string  `json:"symbol,omitempty"`
	SupportsTaxOpt       string  `json:"supports_tax_opt,omitempty"`
	Side                 string  `json:"side,omitempty"`
	OrderDescription     string  `json:"order_description,omitempty"`
	TradeTime            string  `json:"trade_time,omitempty"`
	TradeTimeR           int64   `json:"trade_time_r,omitempty"`
	Size                 float64 `json:"size,omitempty"`
	OrderRef             string  `json:"order_ref,omitempty"`
	Price                string  `json:"price,omitempty"`
	Exchange             string  `json:"exchange,omitempty"`
	NetAmount            float64 `json:"net_amount,omitempty"`
	Account              string  `json:"account,omitempty"`
	AccountCode          string  `json:"accountCode,omitempty"`
	CompanyName          string  `json:"company_name,omitempty"`
	ContractDescription1 string  `json:"contract_description_1,omitempty"`
	ContractDescription2 string  `json:"contract_description_2,omitempty"`
	SecType              string  `json:"sec_type,omitempty"`
	ContractId           int     `json:"conid,omitempty"`
	ContractIdExchange   string  `json:"conidEx,omitempty"`
	OpenClose            string  `json:"open_close,omitempty"`
	LiquidationTrade     string  `json:"liquidation_trade,omitempty"`
	IsEventTrading       string  `json:"is_event_trading,omitempty"`
}
type WebsocketPrivateTradesDataResponse struct {
	Topic string                       `json:"topic,omitempty"`
	Args  []WebsocketPrivateTradesData `json:"args,omitempty"`
}

func (s *WebsocketPrivateService) SubscribeOrder(
	param WebsocketPrivateOrderParam,
	handler func(WebsocketPrivateOrderResponse) error,
) (func() error, error) {

	paramMap := map[string][]string{}
	paramMap["filters"] = param.Status

	buf, err := json.Marshal(paramMap)
	if err != nil {
		return nil, err
	}

	args := fmt.Sprintf("sor+%s", string(buf))

	if err := s.writeMessage(websocket.TextMessage, []byte(args)); err != nil {
		return nil, err
	}

	s.orderResponseHandler = handler

	return func() error {
		args := fmt.Sprintf("uor+{}")
		if err := s.writeMessage(websocket.TextMessage, []byte(args)); err != nil {
			return err
		}
		return nil
	}, nil
}
func (s *WebsocketPrivateService) UnsubscribeOrder(
	param WebsocketPrivateOrderParam,
) error {
	args := fmt.Sprintf("uor+{}")
	if err := s.writeMessage(websocket.TextMessage, []byte(args)); err != nil {
		return err
	}
	return nil
}

func (s *WebsocketPrivateService) SubscribeTradesData(
	param WebsocketPrivateTradesDataParam,
	handler func(WebsocketPrivateTradesDataResponse) error,
) (func() error, error) {

	paramMap := map[string]interface{}{}
	if param.RealtimeUpdatesOnly != nil {
		paramMap["realtimeUpdatesOnly"] = *(param.RealtimeUpdatesOnly)
	}
	if param.Days != nil {
		paramMap["days"] = *(param.Days)
	}

	buf, err := json.Marshal(paramMap)
	if err != nil {
		return nil, err
	}

	args := fmt.Sprintf("str+%s", string(buf))

	if err := s.writeMessage(websocket.TextMessage, []byte(args)); err != nil {
		return nil, err
	}

	s.tradesDataResponseHandler = handler

	return func() error {
		args := fmt.Sprintf("utr")
		if err := s.writeMessage(websocket.TextMessage, []byte(args)); err != nil {
			return err
		}
		return nil
	}, nil
}
func (s *WebsocketPrivateService) UnsubscribeTradesData(
	param WebsocketPrivateTradesDataParam,
) error {
	args := fmt.Sprintf("utr")
	if err := s.writeMessage(websocket.TextMessage, []byte(args)); err != nil {
		return err
	}
	return nil
}

func (s *WebsocketPrivateService) SubscribePnL(
	handler func(WebsocketPrivatePnLResponse) error,
) (func() error, error) {

	args := "spl+{}"

	if err := s.writeMessage(websocket.TextMessage, []byte(args)); err != nil {
		return nil, err
	}

	s.pnlResponseHandler = handler

	return func() error {
		args := fmt.Sprintf("upl+{}")
		if err := s.writeMessage(websocket.TextMessage, []byte(args)); err != nil {
			return err
		}
		return nil
	}, nil
}
func (s *WebsocketPrivateService) UnsubscribePnL() error {
	args := fmt.Sprintf("upl+{}")
	if err := s.writeMessage(websocket.TextMessage, []byte(args)); err != nil {
		return err
	}
	return nil
}
