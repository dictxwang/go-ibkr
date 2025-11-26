package ibkr

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/gorilla/websocket"
	"os"
	"os/signal"
	"strings"
	"sync"
	"time"
)

// WebsocketPrivateServiceI :
type WebsocketPrivateServiceI interface {
	Start(context.Context, ErrHandler) error
	Run() error
	Ping() error
	Close() error
	SetAccountUpdatesChan(channel chan *WebsocketUnsolicitedAccountUpdatesResponse)
	SetAuthStatusChan(channel chan *WebsocketUnsolicitedAuthStatusResponse)
	SetSystemChan(channel chan *WebsocketUnsolicitedSystemConnectionResponse)
	SetBulletinsChan(channel chan *WebsocketUnsolicitedBulletinsResponse)
	SetNotificationsChan(channel chan *WebsocketUnsolicitedNotificationsResponse)

	SubscribeAccountSummary(
		WebsocketPrivateAccountSummaryParam,
		func(WebsocketPrivateAccountSummaryResponse) error,
	) (func() error, error)
	UnsubscribeAccountSummary(
		WebsocketPrivateAccountSummaryParam,
	) error
	SubscribeAccountLedger(
		WebsocketPrivateAccountLedgerParam,
		func(WebsocketPrivateAccountLedgerResponse) error,
	) (func() error, error)
	UnsubscribeAccountLedger(
		WebsocketPrivateAccountLedgerParam,
	) error

	SubscribeOrder(
		WebsocketPrivateOrderParam,
		func(WebsocketPrivateOrderResponse) error,
	) (func() error, error)
	UnsubscribeOrder(
		WebsocketPrivateOrderParam,
	) error
	SubscribePnL(
		WebsocketPrivatePnLParam,
		func(WebsocketPrivatePnLResponse) error,
	) (func() error, error)
	UnsubscribePnL(
		WebsocketPrivatePnLParam,
	) error
	SubscribeTradesData(
		WebsocketPrivateTradesDataParam,
		func(WebsocketPrivateTradesDataResponse) error,
	) (func() error, error)
	UnsubscribeTradesData(
		WebsocketPrivateTradesDataParam,
	) error
}

type WebsocketPrivateService struct {
	client     *WebSocketClient
	connection *websocket.Conn
	writeMutex sync.Mutex

	accountUpdatesChan chan *WebsocketUnsolicitedAccountUpdatesResponse
	authStatusChan     chan *WebsocketUnsolicitedAuthStatusResponse
	systemChan         chan *WebsocketUnsolicitedSystemConnectionResponse
	bulletinsChan      chan *WebsocketUnsolicitedBulletinsResponse
	notificationChan   chan *WebsocketUnsolicitedNotificationsResponse

	accountSummaryResponseHandler func(WebsocketPrivateAccountSummaryResponse) error
	accountLedgerResponseHandler  func(WebsocketPrivateAccountLedgerResponse) error
	orderResponseHandler          func(WebsocketPrivateOrderResponse) error
	pnlResponseHandler            func(WebsocketPrivatePnLResponse) error
	tradesDataResponseHandler     func(WebsocketPrivateTradesDataResponse) error
}

// parseResponse :
func (s *WebsocketPrivateService) parseResponse(respBody []byte, response interface{}) error {
	if err := json.Unmarshal(respBody, &response); err != nil {
		return err
	}
	return nil
}

// parseResponseTopic :
func (s *WebsocketPrivateService) parseResponseTopic(respBody []byte) (string, error) {
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

func (s *WebsocketPrivateService) SetAccountUpdatesChan(channel chan *WebsocketUnsolicitedAccountUpdatesResponse) {
	s.accountUpdatesChan = channel
}
func (s *WebsocketPrivateService) SetAuthStatusChan(channel chan *WebsocketUnsolicitedAuthStatusResponse) {
	s.authStatusChan = channel
}
func (s *WebsocketPrivateService) SetSystemChan(channel chan *WebsocketUnsolicitedSystemConnectionResponse) {
	s.systemChan = channel
}
func (s *WebsocketPrivateService) SetBulletinsChan(channel chan *WebsocketUnsolicitedBulletinsResponse) {
	s.bulletinsChan = channel
}
func (s *WebsocketPrivateService) SetNotificationsChan(channel chan *WebsocketUnsolicitedNotificationsResponse) {
	s.notificationChan = channel
}

// Start :
func (s *WebsocketPrivateService) Start(ctx context.Context, errHandler ErrHandler) error {
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
func (s *WebsocketPrivateService) Run() error {
	_, message, err := s.connection.ReadMessage()
	if err != nil {
		return err
	}

	topic, err := s.parseResponseTopic(message)
	if err != nil {
		return err
	}

	switch topic {
	case UnsolicitedMessageTopicAccountUpdates:
		if s.accountUpdatesChan != nil {
			var resp WebsocketUnsolicitedAccountUpdatesResponse
			if err := s.parseResponse(message, &resp); err != nil {
				return err
			}
			s.accountUpdatesChan <- &resp
		}
	case UnsolicitedMessageTopicAuthStatus:
		if s.authStatusChan != nil {
			var resp WebsocketUnsolicitedAuthStatusResponse
			if err := s.parseResponse(message, &resp); err != nil {
				return err
			}
			s.authStatusChan <- &resp
		}
	case UnsolicitedMessageTopicSystemConnection:
		if s.systemChan != nil {
			var resp WebsocketUnsolicitedSystemConnectionResponse
			if err := s.parseResponse(message, &resp); err != nil {
				return err
			}
			s.systemChan <- &resp
		}
	case UnsolicitedMessageTopicBulletins:
		if s.bulletinsChan != nil {
			var resp WebsocketUnsolicitedBulletinsResponse
			if err := s.parseResponse(message, &resp); err != nil {
				return err
			}
			s.bulletinsChan <- &resp
		}
	case UnsolicitedMessageTopicNotifications:
		if s.notificationChan != nil {
			var resp WebsocketUnsolicitedNotificationsResponse
			if err := s.parseResponse(message, &resp); err != nil {
				return err
			}
			s.notificationChan <- &resp
		}
	case MessageTopicSubscribeAccountSummary:
		var resp WebsocketPrivateAccountSummaryResponse
		if err := s.parseResponse(message, &resp); err != nil {
			return err
		}
		if s.accountSummaryResponseHandler != nil {
			if err := s.accountSummaryResponseHandler(resp); err != nil {
				return err
			}
		}
	case MessageTopicSubscribeAccountLedger:
		var resp WebsocketPrivateAccountLedgerResponse
		if err := s.parseResponse(message, &resp); err != nil {
			return err
		}
		if s.accountLedgerResponseHandler != nil {
			if err := s.accountLedgerResponseHandler(resp); err != nil {
				return err
			}
		}
	case MessageTopicSubscribeOrder:
		var resp WebsocketPrivateOrderResponse
		if err := s.parseResponse(message, &resp); err != nil {
			return err
		}
		if s.orderResponseHandler != nil {
			if err := s.orderResponseHandler(resp); err != nil {
				return err
			}
		}
	case MessageTopicSubscribePnL:
		var resp WebsocketPrivatePnLResponse
		if err := s.parseResponse(message, &resp); err != nil {
			return err
		}
		if s.pnlResponseHandler != nil {
			if err := s.pnlResponseHandler(resp); err != nil {
				return err
			}
		}
	case MessageTopicSubscribeTradesData:
		var resp WebsocketPrivateTradesDataResponse
		if err := s.parseResponse(message, &resp); err != nil {
			return err
		}
		if s.tradesDataResponseHandler != nil {
			if err := s.tradesDataResponseHandler(resp); err != nil {
				return err
			}
		}
	}
	return nil
}

// Ping :
func (s *WebsocketPrivateService) Ping() error {
	if err := s.writeMessage(websocket.PingMessage, nil); err != nil {
		return err
	}
	//// below copy from bybit sdk, no need for ibkr
	//if err := s.writeMessage(websocket.TextMessage, []byte(`{"op":"ping"}`)); err != nil {
	//	return err
	//}
	return nil
}

// Close :
func (s *WebsocketPrivateService) Close() error {
	if err := s.writeMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, "")); err != nil && !errors.Is(err, websocket.ErrCloseSent) {
		return err
	}
	return nil
}

func (s *WebsocketPrivateService) writeMessage(messageType int, body []byte) error {
	s.writeMutex.Lock()
	defer s.writeMutex.Unlock()

	if err := s.connection.WriteMessage(messageType, body); err != nil {
		return err
	}
	return nil
}
