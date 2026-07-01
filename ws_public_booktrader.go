package ibkr

import (
	"fmt"

	"github.com/gorilla/websocket"
)

type WebsocketPublicBookTraderParam struct {
	AccountId string
	ContractIds []int
	Exchange      string
}
type WebsocketPublicBookTraderDataItem struct {
	Row int `json:"row"`
	Focus int `json:"focus"` //Indicates if the value was marked as the last trade price for the contract.
	Price string `json:"price"`
	Ask string `json:"ask,omitempty"`
	Bid string `json:"bid,omitempty"`
}
type WebsocketPublicBookTraderResponse struct {
	Topic string `json:"topic"`
	Data []WebsocketPublicBookTraderDataItem `json:"data"`
}

func (s *WebsocketPublicService) SubscribeBookTrader(
	param WebsocketPublicBookTraderParam,
	handler func(WebsocketPublicBookTraderResponse) error,
) (func() error, error) {
	
	for _, contractId := range param.ContractIds {
		args := fmt.Sprintf("sbd+%s+%d", param.AccountId, contractId)
		if param.Exchange != "" {
			args = args + "+" + param.Exchange
		}

		if err := s.writeMessage(websocket.TextMessage, []byte(args)); err != nil {
			return nil, err
		}
	}

	s.bookTraderResponseHandler = handler

	return func() error {
		args := fmt.Sprintf("ubd+{%s}", param.AccountId)
		if err := s.writeMessage(websocket.TextMessage, []byte(args)); err != nil {
			return err
		}
		return nil
	}, nil
}
func (s *WebsocketPublicService) UnsubscribeBookTrader(
	param WebsocketPublicBookTraderParam,
) error {
	args := fmt.Sprintf("ubd+{%s}", param.AccountId)
	if err := s.writeMessage(websocket.TextMessage, []byte(args)); err != nil {
		return err
	}
	return nil
}
