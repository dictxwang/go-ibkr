package ibkr

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/websocket"
	"os"
	"os/signal"
	"strings"
	"sync"
	"time"
)

// market data fields: https://www.interactivebrokers.com/campus/ibkr-api-page/cpapi-v1/#market-data-fields

// WebsocketPublicServiceI :
type WebsocketPublicServiceI interface {
	Start(context.Context, ErrHandler) error
	Run() error
	Ping() error
	Close() error

	SubscribeMarketData(
		WebsocketPublicMarketDataParam,
		func(WebsocketPublicMarketDataResponse) error,
	) (func() error, error)
	UnsubscribeMarketData(
		WebsocketPublicMarketDataParam,
	) error
	SubscribeHistoricalMarketData(
		WebsocketPublicMarketDataParam,
		func(WebsocketPublicMarketDataResponse) error,
	) (func() error, error)
	UnsubscribeHistoricalMarketData(
		WebsocketPublicMarketDataParam,
	) error
	SubscribeBookTrader(
		WebsocketPublicBookTraderParam,
		func(WebsocketPublicBookTraderResponse) error,
	) (func() error, error)
	UnsubscribeBookTrader(
		WebsocketPublicBookTraderParam,
	) error
}

type WebsocketPublicService struct {
	client     *WebSocketClient
	connection *websocket.Conn
	writeMutex sync.Mutex

	marketDataResponseHandler           func(WebsocketPublicMarketDataResponse) error
	historicalMarketDataResponseHandler func(WebsocketPublicHistoricalMarketDataResponse) error
	bookTraderResponseHandler           func(WebsocketPublicBookTraderResponse) error
}

// parseResponse :
func (s *WebsocketPublicService) parseResponse(respBody []byte, response interface{}) error {
	if err := json.Unmarshal(respBody, &response); err != nil {
		return err
	}
	return nil
}

// parseResponseTopic :
func (s *WebsocketPublicService) parseResponseTopic(respBody []byte) (string, error) {
	if len(respBody) == 0 {
		return "", nil
	}
	resp := map[string]interface{}{}
	if err := json.Unmarshal(respBody, &resp); err != nil {
		return "", err
	}
	if topic, has := resp["topic"]; has {
		topicParts := strings.Split(topic.(string), "-")
		return topicParts[0], nil
	} else {
		return "", nil
	}
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
	fmt.Printf("before read message\n")
	_, message, err := s.connection.ReadMessage()
	fmt.Printf("after read message: %s\n", message)
	if err != nil {
		return err
	}

	topic, err := s.parseResponseTopic(message)
	if err != nil {
		return err
	}

	switch topic {
	case MessageTopicSubscribeMarketData:
		var resp WebsocketPublicMarketDataResponse
		if err := s.parseResponse(message, &resp); err != nil {
			return err
		}

		if err := s.marketDataResponseHandler(resp); err != nil {
			return err
		}
	case MessageTopicSubscribeHistoricalMarketData:
		var resp WebsocketPublicHistoricalMarketDataResponse
		if err := s.parseResponse(message, &resp); err != nil {
			return err
		}

		if err := s.historicalMarketDataResponseHandler(resp); err != nil {
			return err
		}
	case MessageTopicSubscribeBookTrader:
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
	s.writeMutex.Lock()
	defer s.writeMutex.Unlock()

	if err := s.connection.WriteMessage(messageType, body); err != nil {
		return err
	}
	return nil
}
