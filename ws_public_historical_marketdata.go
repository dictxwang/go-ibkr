package ibkr

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
)

func (s *WebsocketPublicService) UnsubscribeHistoricalMarketData(
	param WebsocketPublicHistoricalMarketDataParam,
) error {
	for _, contractId := range param.ContractIds {
		args := fmt.Sprintf("umh+%d+{}", contractId)
		if err := s.writeMessage(websocket.TextMessage, []byte(args)); err != nil {
			return err
		}
	}
	return nil
}

func (s *WebsocketPublicService) SubscribeHistoricalTicker(
	param WebsocketPublicHistoricalMarketDataParam,
	handler func(WebsocketPublicHistoricalMarketDataResponse) error,
) (func() error, error) {

	buf, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}

	for _, contractId := range param.ContractIds {
		args := fmt.Sprintf("smh+%d+%s", contractId, string(buf))

		if err := s.writeMessage(websocket.TextMessage, []byte(args)); err != nil {
			return nil, err
		}
	}

	s.historicalMarketDataResponseHandler = handler

	return func() error {
		for _, contractId := range param.ContractIds {
			args := fmt.Sprintf("umh+%d+{}", contractId)
			if err := s.writeMessage(websocket.TextMessage, []byte(args)); err != nil {
				return err
			}
		}
		return nil
	}, nil
}

type WebsocketPublicHistoricalMarketDataParam struct {
	ContractIds []int  `json:"-"`
	Exchange    string `json:"exchange,omitempty"`
	Period      string `json:"period,omitempty"`
	Bar         string `json:"bar,omitempty"`
	OutsideRth  bool   `json:"outsideRth"`
	Source      string `json:"source,omitempty"`
	Format      string `json:"format,omitempty"`
}

type WebsocketPublicHistoricalMarketDataResponse struct {
	Topic    string `json:"topic,omitempty"`
	ServerId string `json:"server_id,omitempty"`
	Symbol   string `json:"symbol,omitempty"`
}
