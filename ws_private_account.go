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
	// TODO
}

type WebsocketPrivateAccountLedgerResponse struct {
	// TODO
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
func (s *WebsocketPrivateService) UnSubscribeAccountSummary(
	param WebsocketPrivateAccountSummaryParam,
) error {
	args := fmt.Sprintf("usd+%s+{}", param.AccountId)
	if err := s.writeMessage(websocket.TextMessage, []byte(args)); err != nil {
		return err
	}
	return nil
}

func (s *WebsocketPrivateService) SubscribeAccountLedger(
	param, WebsocketPrivateAccountLedgerParam,
	handler func(WebsocketPrivateAccountLedgerResponse) error,
) (func() error, error) {
	// TODO
	return nil, nil
}
func (s *WebsocketPrivateService) UnSubscribeAccountLedger(
	param WebsocketPrivateAccountLedgerParam,
) error {
	// TODO
	return nil
}
