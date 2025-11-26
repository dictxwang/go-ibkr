package ibkr

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
)

func (s *WebsocketPublicService) UnsubscribeMarketData(
	param WebsocketPublicMarketDataParam,
) error {
	for _, contractId := range param.ContractIds {
		args := fmt.Sprintf("umd+%d+{}", contractId)
		if err := s.writeMessage(websocket.TextMessage, []byte(args)); err != nil {
			return err
		}
	}
	return nil
}

func (s *WebsocketPublicService) SubscribeMarketData(
	param WebsocketPublicMarketDataParam,
	handler func(WebsocketPublicMarketDataResponse) error,
) (func() error, error) {

	fieldsMap := map[string][]string{}
	fieldsMap["fields"] = param.Fields

	buf, err := json.Marshal(fieldsMap)
	if err != nil {
		return nil, err
	}

	for _, contractId := range param.ContractIds {
		args := fmt.Sprintf("smd+%d+%s", contractId, string(buf))

		fmt.Printf("subscribe ticker 04: %s\n", args)
		if err := s.writeMessage(websocket.TextMessage, []byte(args)); err != nil {
			return nil, err
		}
	}

	s.marketDataResponseHandler = handler

	return func() error {
		for _, contractId := range param.ContractIds {
			args := fmt.Sprintf("umd+%d+{}", contractId)
			if err := s.writeMessage(websocket.TextMessage, []byte(args)); err != nil {
				return err
			}
		}
		return nil
	}, nil
}

type WebsocketPublicMarketDataParam struct {
	ContractIds []int
	Fields      []string
}

type WebsocketPublicMarketDataResponse struct {
	Topic                  string  `json:"topic,omitempty"`
	ServerId               string  `json:"server_id,omitempty"`
	ContractIdExchange     string  `json:"conidEx,omitempty"`
	ContractId             int     `json:"conid,omitempty"`
	UpdateTime             int64   `json:"_updated,omitempty"`
	MarketDataAvailability string  `json:"6509,omitempty"`
	BidSize                float64 `json:"88,omitempty"`
	BidPrice               float64 `json:"84,omitempty"`
	AskSize                float64 `json:"85,omitempty"`
	AskPrice               float64 `json:"86,omitempty"`
}
