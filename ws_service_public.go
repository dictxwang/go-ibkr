package ibkr

import (
	"context"
	"github.com/gorilla/websocket"
	"sync"
)

// market data fields: https://www.interactivebrokers.com/campus/ibkr-api-page/cpapi-v1/#market-data-fields

type WsPublicSubscribeChannel string

const (
	WsPublicSubscribeChannelTicker     = WsPublicSubscribeChannel("Ticker")
	WsPublicSubscribeChannelBookTrader = WsPublicSubscribeChannel("BookTrader")
)

// WebsocketPublicServiceI :
type WebsocketPublicServiceI interface {
	Start(context.Context, ErrHandler) error
	Run() error
	Ping() error
	Close() error
	SubscribeTicker(
		WebsocketPublicTickerParam,
		func(WebsocketPublicTickerResponse) error,
	) (func() error, error)
}

type WebsocketPublicService struct {
	client            *WebSocketClient
	connection        *websocket.Conn
	alreadySubscribed bool
	mu                sync.Mutex

	subscribeChannel      WsPublicSubscribeChannel
	fillResponseFunc      func(data map[string]interface{}, response interface{}) error
	tickerResponseHandler func(WebsocketPublicTickerResponse) error
}

func (s *WebsocketPublicService) writeMessage(messageType int, body []byte) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if err := s.connection.WriteMessage(messageType, body); err != nil {
		return err
	}
	return nil
}
