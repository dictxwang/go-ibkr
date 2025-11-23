package ibkr

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/gorilla/websocket"
	"os"
	"os/signal"
	"sync"
	"time"
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
	UnSubscribeTicker(
		WebsocketPublicTickerParam,
	) error
}

type WebsocketPublicService struct {
	client            *WebSocketClient
	connection        *websocket.Conn
	alreadySubscribed bool
	mu                sync.Mutex

	subscribeChannel      WsPublicSubscribeChannel
	tickerResponseHandler func(WebsocketPublicTickerResponse) error
}

// parseResponse :
func (s *WebsocketPublicService) parseResponse(respBody []byte, response interface{}) error {
	if err := json.Unmarshal(respBody, &response); err != nil {
		return err
	}
	return nil
}

// Start :
func (s *WebsocketPublicService) Start(ctx context.Context, errHandler ErrHandler) error {
	done := make(chan struct{})

	go func() {
		defer close(done)
		defer s.connection.Close()

		_ = s.connection.SetReadDeadline(time.Now().Add(60 * time.Second))
		s.connection.SetPongHandler(func(string) error {
			_ = s.connection.SetReadDeadline(time.Now().Add(60 * time.Second))
			return nil
		})

		for {
			if err := s.Run(); err != nil {
				if errHandler == nil {
					return
				}
				errHandler(IsErrWebsocketClosed(err), err)
				return
			}
		}
	}()

	ticker := time.NewTicker(20 * time.Second)
	defer ticker.Stop()

	ctx, stop := signal.NotifyContext(ctx, os.Interrupt)
	defer stop()

	for {
		select {
		case <-done:
			return nil
		case <-ticker.C:
			if err := s.Ping(); err != nil {
				return err
			}
		case <-ctx.Done():
			s.client.debugf("caught websocket public service interrupt signal")

			if err := s.Close(); err != nil {
				return err
			}
			select {
			case <-done:
			case <-time.After(time.Second):
			}
			return nil
		}
	}
}

// Run :
func (s *WebsocketPublicService) Run() error {
	_, message, err := s.connection.ReadMessage()
	if err != nil {
		return err
	}

	switch s.subscribeChannel {
	case WsPublicSubscribeChannelTicker:
		var resp WebsocketPublicTickerResponse
		if err := s.parseResponse(message, &resp); err != nil {
			return err
		}

		if err := s.tickerResponseHandler(resp); err != nil {
			return err
		}
	case WsPublicSubscribeChannelBookTrader:
		// TODO
		return nil
	}
	return nil
}

// Ping :
func (s *WebsocketPublicService) Ping() error {
	// NOTE: It appears that two messages need to be sent.
	// REF: https://github.com/hirokisan/bybit/pull/127#issuecomment-1537479346
	if err := s.writeMessage(websocket.PingMessage, nil); err != nil {
		return err
	}
	if err := s.writeMessage(websocket.TextMessage, []byte(`{"op":"ping"}`)); err != nil {
		return err
	}
	return nil
}

// Close :
func (s *WebsocketPublicService) Close() error {
	if err := s.writeMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, "")); err != nil && !errors.Is(err, websocket.ErrCloseSent) {
		return err
	}
	return nil
}

func (s *WebsocketPublicService) writeMessage(messageType int, body []byte) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if err := s.connection.WriteMessage(messageType, body); err != nil {
		return err
	}
	return nil
}
