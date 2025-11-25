package ibkr

type WebsocketPrivateAccountSummaryParam struct {
	// TODO
}

type WebsocketPrivateAccountSummaryResponse struct {
	// TODO
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
	// TODO
	return nil, nil
}
func (s *WebsocketPrivateService) UnSubscribeAccountSummary(
	param WebsocketPrivateAccountSummaryParam,
) error {
	// TODO
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
