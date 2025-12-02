package ibkr

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"strings"
)

type OrdersServiceI interface {
	PlaceOrder(orders []PlaceOrderParam) (*PlaceOrderResponse, error)
	CancelOrder(param CancelOrderParam) (*CancelOrderResponse, error)
	PlaceOrderReplyConfirmation(param PlaceOrderReplyConfirmationParam) (*[]PlaceOrderReplyConfirmationResponse, error)
	RespondServerPrompt(param RespondServerPromptParam) (*RespondServerPromptResponse, error)
}

type OrdersService struct {
	client *Client
}

func (s *OrdersService) PlaceOrder(orders []PlaceOrderParam) (*PlaceOrderResponse, error) {
	if len(orders) == 0 {
		return nil, errors.New("require order params")
	}

	params := map[string][]PlaceOrderParam{}
	params["orders"] = orders
	body, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	responseBytes, err := s.client.postJSONConciseResponse(fmt.Sprintf("/iserver/account/%s/orders", orders[0].AccountId), body)
	if err != nil {
		return nil, err
	}
	var resp PlaceOrderResponse
	if strings.HasPrefix(string(responseBytes), "{") {
		var reject PlaceOrderRejectResult
		err := json.Unmarshal(responseBytes, &reject)
		if err != nil {
			return nil, err
		} else {
			resp.RejectResult = &reject
		}
	} else if strings.Contains(string(responseBytes), "\"order_id\"") {
		var normals []PlaceOrderNormalResult
		err := json.Unmarshal(responseBytes, &normals)
		if err != nil {
			return nil, err
		} else {
			resp.NormalResults = &normals
		}
	} else {
		var alternates []PlaceOrderAlternateResult
		err := json.Unmarshal(responseBytes, &alternates)
		if err != nil {
			return nil, err
		} else {
			resp.AlternateResults = &alternates
		}
	}
	return &resp, nil
}

func (s *OrdersService) CancelOrder(param CancelOrderParam) (*CancelOrderResponse, error) {

	var resp CancelOrderResponse

	queries := url.Values{}
	if err := s.client.deletePublic(fmt.Sprintf("/iserver/account/%s/order/%d", param.AccountId, param.OrderId), queries, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *OrdersService) PlaceOrderReplyConfirmation(param PlaceOrderReplyConfirmationParam) (*[]PlaceOrderReplyConfirmationResponse, error) {

	var resp []PlaceOrderReplyConfirmationResponse
	body, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}

	if err := s.client.postJSON(fmt.Sprintf("/iserver/reply/%s", param.ReplyId), body, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *OrdersService) RespondServerPrompt(param RespondServerPromptParam) (*RespondServerPromptResponse, error) {
	body, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}

	responseBytes, err := s.client.postJSONConciseResponse("/iserver/notification", body)
	if err != nil {
		return nil, err
	}
	var resp RespondServerPromptResponse
	resp.Result = string(responseBytes)
	return &resp, nil
}

type PlaceOrderParam struct {
	AccountId                  string                 `json:"acctId"`
	ContractId                 int                    `json:"conid"`
	ContractIdExchange         string                 `json:"conidex"`
	ManualIndicator            bool                   `json:"manualIndicator"`
	ExternalOperator           string                 `json:"extOperator,omitempty"`
	ContractSecurityType       string                 `json:"secType,omitempty"` // sample:265598:STK
	CustomOrderId              string                 `json:"cOID,omitempty"`
	ParentId                   string                 `json:"parentId,omitempty"`
	OrderType                  OrderType              `json:"orderType"`
	ListingExchange            string                 `json:"listingExchange,omitempty"`
	IsSingleGroup              bool                   `json:"isSingleGroup"`
	OutsideRegularTradingHours bool                   `json:"outsideRTH"`
	Price                      *float64               `json:"price,omitempty"`
	AuxPrice                   *float64               `json:"auxPrice,omitempty"`
	Side                       string                 `json:"side"`
	Ticker                     string                 `json:"ticker,omitempty"`
	TimeInForce                TimeInForce            `json:"tif"`
	TrailingAmount             *float64               `json:"trailingAmt,omitempty"`
	TrailingType               TrailingType           `json:"trailingType,omitempty"`
	AllOrNone                  bool                   `json:"allOrNone"`
	CustomerAccount            string                 `json:"customerAccount,omitempty"`
	IsProCustomer              bool                   `json:"isProCustomer"`
	Referrer                   string                 `json:"referrer,omitempty"`
	Quantity                   float64                `json:"quantity"`
	CashQty                    *float64               `json:"cashQty"`
	FxQty                      *float64               `json:"fxQty,omitempty"`
	UseAdaptive                *bool                  `json:"useAdaptive,omitempty"`
	IsCcyConversion            *bool                  `json:"isCcyConv,omitempty"'`
	AllocationMethod           string                 `json:"allocationMethod,omitempty"`
	ManualOrderTime            *int                   `json:"manualOrderTime,omitempty"`
	Deactivated                *bool                  `json:"deactivated,omitempty"`
	Strategy                   string                 `json:"strategy,omitempty"`
	StrategyParameters         map[string]interface{} `json:"strategyParameters,omitempty"`
}

type PlaceOrderNormalResult struct {
	OrderId        string      `json:"order_id"`
	OrderStatus    OrderStatus `json:"order_status"`
	EncryptMessage string      `json:"encrypt_message"`
}
type PlaceOrderAlternateResult struct {
	Id           string   `json:"id"`
	Message      []string `json:"message"`
	IsSuppressed bool     `json:"isSuppressed"`
	MessageIds   []string `json:"messageIds"`
}
type PlaceOrderRejectResult struct {
	Error string `json:"error"`
}
type PlaceOrderResponse struct {
	NormalResults    *[]PlaceOrderNormalResult
	AlternateResults *[]PlaceOrderAlternateResult
	RejectResult     *PlaceOrderRejectResult
}

type CancelOrderParam struct {
	AccountId string
	OrderId   int64
}

type CancelOrderResponse struct {
	Msg        string `json:"msg,omitempty"`
	OrderId    int64  `json:"order_id,omitempty"`
	ContractId int    `json:"conid,omitempty"`
	AccountId  string `json:"account,omitempty"`
	Error      string `json:"error,omitempty"`
}

type PlaceOrderReplyConfirmationParam struct {
	ReplyId   string `json:"-"`
	Confirmed bool   `json:"confirmed"`
}

type PlaceOrderReplyConfirmationResponse struct {
	NormalResults []PlaceOrderNormalResult
}

type RespondServerPromptParam struct {
	OrderId int64  `json:"orderId"`
	ReqId   string `json:"reqId"`
	Text    string `json:"text"`
}

type RespondServerPromptResponse struct {
	Result string
}
