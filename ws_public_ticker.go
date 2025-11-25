package ibkr

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/websocket"
)

func (s *WebsocketPublicService) UnSubscribeTicker(
	param WebsocketPublicTickerParam,
) error {
	args := fmt.Sprintf("umd+%d+{}", param.ContractId)
	if err := s.writeMessage(websocket.TextMessage, []byte(args)); err != nil {
		return err
	}
	s.alreadySubscribed = false
	return nil
}

func (s *WebsocketPublicService) SubscribeTicker(
	param WebsocketPublicTickerParam,
	handler func(WebsocketPublicTickerResponse) error,
) (func() error, error) {

	s.mu.Lock()
	defer s.mu.Unlock()

	fmt.Printf("subscribe ticker 01\n")

	if s.alreadySubscribed {
		return nil, errors.New("already subscribed")
	}
	fmt.Printf("subscribe ticker 02\n")

	param.fillFields()
	fields := make([]string, 0)
	fields = append(fields, param.fieldBidSize)
	fields = append(fields, param.fieldBidPrice)
	fields = append(fields, param.fieldAskSize)
	fields = append(fields, param.fieldAskPrice)
	fieldsMap := map[string][]string{}
	fieldsMap["fields"] = fields

	buf, err := json.Marshal(fieldsMap)
	if err != nil {
		return nil, err
	}
	fmt.Printf("subscribe ticker 03\n")

	args := fmt.Sprintf("smd%d%s", param.ContractId, string(buf))

	fmt.Printf("subscribe ticker 04: %s\n", args)
	if err := s.writeMessage(websocket.TextMessage, []byte(args)); err != nil {
		return nil, err
	}

	fmt.Printf("subscribe ticker 05\n")

	s.subscribeChannel = WsPublicSubscribeChannelTicker
	s.tickerResponseHandler = handler
	s.alreadySubscribed = true

	return func() error {
		args := fmt.Sprintf("umd+%d+{}", param.ContractId)
		if err := s.writeMessage(websocket.TextMessage, []byte(args)); err != nil {
			return err
		}
		s.alreadySubscribed = false
		return nil
	}, nil
}

type WebsocketPublicTickerParam struct {
	ContractId    int
	fieldBidPrice string
	fieldBidSize  string
	fieldAskPrice string
	fieldAskSize  string
}

func (p *WebsocketPublicTickerParam) fillFields() {
	p.fieldBidPrice = "84"
	p.fieldBidSize = "88"
	p.fieldAskPrice = "86"
	p.fieldAskSize = "85"
}

type WebsocketPublicTickerResponse struct {
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
