package ibkr

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
)

type WebsocketPrivateAccountSummaryParam struct {
	AccountId string
	Keys      []string
	Fields    []string
}

type WebsocketPrivateAccountSummaryResponse struct {
	Result map[string]interface{} `json:"result"`
}

type WebsocketPrivateAccountLedgerParam struct {
	AccountId string
	Keys      []string
	Fields    []string
}

type WebsocketPrivateAccountLedgerItem struct {
	Key                       string  `json:"key"`
	SecondKey                 string  `json:"secondKey,omitempty"`
	Timestamp                 int64   `json:"timestamp,omitempty"`
	Dividends                 float64 `json:"dividends,omitempty"`
	ExchangeRate              float64 `json:"exchangeRate,omitempty"`
	Funds                     float64 `json:"funds,omitempty"`
	AccountCode               string  `json:"acctCode,omitempty"`
	CashBalance               float64 `json:"cashbalance,omitempty"`
	CashBalanceFXSegment      float64 `json:"cashBalanceFXSegment,omitempty"`
	CommodityMarketValue      float64 `json:"commodityMarketValue,omitempty"`
	CorporateBondsMarketValue float64 `json:"corporateBondsMarketValue,omitempty"`
	MarketValue               float64 `json:"marketValue,omitempty"`
	OptionMarketValue         float64 `json:"optionMarketValue,omitempty"`
	Interest                  float64 `json:"interest,omitempty"`
	IssueOptionsMarketValue   float64 `json:"issueOptionsMarketValue,omitempty"`
	MoneyFunds                float64 `json:"moneyFunds,omitempty"`
	NetLiquidationValue       float64 `json:"netLiquidationValue,omitempty"`
	RealizedPnl               float64 `json:"realizedPnl,omitempty"`
	UnrealizedPnl             float64 `json:"unrealizedPnl,omitempty"`
	SettledCash               float64 `json:"settledCash,omitempty"`
	Severity                  float64 `json:"severity,omitempty"`
	StockMarketValue          float64 `json:"stockMarketValue,omitempty"`
	TBillsMarketValue         float64 `json:"tBillsMarketValue,omitempty"`
	TBondsMarketValue         float64 `json:"tBondsMarketValue,omitempty"`
	WarrantsMarketValue       float64 `json:"warrantsMarketValue,omitempty"`
}

type WebsocketPrivateAccountLedgerResponse struct {
	Result []WebsocketPrivateAccountLedgerItem `json:"result"`
}

func (s *WebsocketPrivateService) SubscribeAccountSummary(
	param WebsocketPrivateAccountSummaryParam,
	handler func(WebsocketPrivateAccountSummaryResponse) error,
) (func() error, error) {

	paramMap := map[string][]string{}
	paramMap["keys"] = param.Keys
	paramMap["fields"] = param.Fields

	buf, err := json.Marshal(paramMap)
	if err != nil {
		return nil, err
	}

	args := fmt.Sprintf("ssd+%s+%s", param.AccountId, string(buf))

	if err := s.writeMessage(websocket.TextMessage, []byte(args)); err != nil {
		return nil, err
	}

	s.accountSummaryResponseHandler = handler

	return func() error {
		args := fmt.Sprintf("usd+%s+{}", param.AccountId)
		if err := s.writeMessage(websocket.TextMessage, []byte(args)); err != nil {
			return err
		}
		return nil
	}, nil
}
func (s *WebsocketPrivateService) UnsubscribeAccountSummary(
	param WebsocketPrivateAccountSummaryParam,
) error {
	args := fmt.Sprintf("usd+%s+{}", param.AccountId)
	if err := s.writeMessage(websocket.TextMessage, []byte(args)); err != nil {
		return err
	}
	return nil
}

func (s *WebsocketPrivateService) SubscribeAccountLedger(
	param WebsocketPrivateAccountLedgerParam,
	handler func(WebsocketPrivateAccountLedgerResponse) error,
) (func() error, error) {

	paramMap := map[string][]string{}
	paramMap["keys"] = param.Keys
	paramMap["fields"] = param.Fields

	buf, err := json.Marshal(paramMap)
	if err != nil {
		return nil, err
	}

	args := fmt.Sprintf("sld+%s+%s", param.AccountId, string(buf))

	if err := s.writeMessage(websocket.TextMessage, []byte(args)); err != nil {
		return nil, err
	}

	s.accountLedgerResponseHandler = handler

	return func() error {
		args := fmt.Sprintf("uld+%s+{}", param.AccountId)
		if err := s.writeMessage(websocket.TextMessage, []byte(args)); err != nil {
			return err
		}
		return nil
	}, nil
}

func (s *WebsocketPrivateService) UnsubscribeAccountLedger(
	param WebsocketPrivateAccountLedgerParam,
) error {
	args := fmt.Sprintf("uld+%s+{}", param.AccountId)
	if err := s.writeMessage(websocket.TextMessage, []byte(args)); err != nil {
		return err
	}
	return nil
}
