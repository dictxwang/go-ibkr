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

type WsPrivateSubscribeChannel string

const (
	WsPrivateSubscribeChannelAccountSummary = WsPrivateSubscribeChannel("AccountSummary")
	WsPrivateSubscribeChannelAccountLedger  = WsPrivateSubscribeChannel("AccountLedger")
	WsPrivateSubscribeChannelOrder          = WsPrivateSubscribeChannel("Order")
	WsPrivateSubscribeChannelPnL            = WsPrivateSubscribeChannel("PnL")
	WsPrivateSubscribeChannelTradesData     = WsPrivateSubscribeChannel("TradesData")
)

// WebsocketPrivateServiceI :
type WebsocketPrivateServiceI interface {
	Start(context.Context, ErrHandler) error
	Run() error
	Ping() error
	Close() error

	SubscribeAccountSummary(
		WebsocketPrivateAccountSummaryParam,
		func(WebsocketPrivateAccountSummaryResponse) error,
	) (func() error, error)
	UnSubscribeAccountSummary(
		WebsocketPrivateAccountSummaryParam,
	) error
	SubscribeAccountLedger(
		WebsocketPrivateAccountLedgerParam,
		func(WebsocketPrivateAccountLedgerResponse) error,
	) (func() error, error)
	UnSubscribeAccountLedger(
		WebsocketPrivateAccountLedgerParam,
	) error

	SubscribeOrder(
		WebsocketPrivateOrderParam,
		func(WebsocketPrivateOrderResponse) error,
	) (func() error, error)
	UnSubscribeOrder(
		WebsocketPrivateOrderParam,
	) error
	SubscribePnL(
		WebsocketPrivatePnLParam,
		func(WebsocketPrivatePnLResponse) error,
	) (func() error, error)
	UnSubscribePnL(
		WebsocketPrivatePnLParam,
	) error
	SubscribeTradesData(
		WebsocketPrivateTradesDataParam,
		func(WebsocketPrivateTradesDataResponse) error,
	) (func() error, error)
	UnSubscribeTradesData(
		WebsocketPrivateTradesDataParam,
	) error
}

type WebsocketPrivateService struct {
	client            *WebSocketClient
	connection        *websocket.Conn
	alreadySubscribed bool
	writeMutex        sync.Mutex
	subscribeMutex    sync.Mutex

	subscribeChannel              WsPrivateSubscribeChannel
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

	switch s.subscribeChannel {
	case WsPrivateSubscribeChannelAccountSummary:
		var resp WebsocketPrivateAccountSummaryResponse
		if err := s.parseResponse(message, &resp); err != nil {
			return err
		}

		if err := s.accountSummaryResponseHandler(resp); err != nil {
			return err
		}
	case WsPrivateSubscribeChannelAccountLedger:
		var resp WebsocketPrivateAccountLedgerResponse
		if err := s.parseResponse(message, &resp); err != nil {
			return err
		}

		if err := s.accountLedgerResponseHandler(resp); err != nil {
			return err
		}
	case WsPrivateSubscribeChannelOrder:
		var resp WebsocketPrivateOrderResponse
		if err := s.parseResponse(message, &resp); err != nil {
			return err
		}

		if err := s.orderResponseHandler(resp); err != nil {
			return err
		}
	case WsPrivateSubscribeChannelPnL:
		var resp WebsocketPrivatePnLResponse
		if err := s.parseResponse(message, &resp); err != nil {
			return err
		}

		if err := s.pnlResponseHandler(resp); err != nil {
			return err
		}
	case WsPrivateSubscribeChannelTradesData:
		var resp WebsocketPrivateTradesDataResponse
		if err := s.parseResponse(message, &resp); err != nil {
			return err
		}

		if err := s.tradesDataResponseHandler(resp); err != nil {
			return err
		}
	}
	return nil
}

// Ping :
func (s *WebsocketPrivateService) Ping() error {
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
