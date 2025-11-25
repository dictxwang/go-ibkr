package ibkr

type WebsocketPrivateOrderParam struct {
	// TODO
}

type WebsocketPrivateOrderResponse struct {
	// TODO
}

type WebsocketPrivatePnLParam struct {
	// TODO
}

type WebsocketPrivatePnLResponse struct {
	// TODO
}

type WebsocketPrivateTradesDataParam struct {
	// TODO
}

type WebsocketPrivateTradesDataResponse struct {
	// TODO
}

func (s *WebsocketPrivateService) SubscribeOrder(
	param WebsocketPrivateOrderParam,
	handler func(WebsocketPrivateOrderResponse) error,
) (func() error, error) {
	// TODO
	return nil, nil
}
func (s *WebsocketPrivateService) UnSubscribeOrder(
	param WebsocketPrivatePnLParam,
) error {
	// TODO
	return nil
}

func (s *WebsocketPrivateService) SubscribeTradesData(
	param WebsocketPrivateTradesDataParam,
	handler func(WebsocketPrivateTradesDataResponse) error,
) (func() error, error) {
	// TODO
	return nil, nil
}
func (s *WebsocketPrivateService) UnSubscribeTradesData(
	param WebsocketPrivateTradesDataParam,
) error {
	// TODO
	return nil
}

func (s *WebsocketPrivateService) SubscribePnL(
	param WebsocketPrivatePnLParam,
	handler func(WebsocketPrivatePnLResponse) error,
) (func() error, error) {
	// TODO
	return nil, nil
}
func (s *WebsocketPrivateService) UnSubscribePnL(
	param WebsocketPrivatePnLParam,
) error {
	// TODO
	return nil
}
