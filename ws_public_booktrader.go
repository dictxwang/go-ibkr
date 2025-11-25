package ibkr

type WebsocketPublicBookTraderParam struct {
	// TODO
}
type WebsocketPublicBookTraderResponse struct {
	// TODO
}

func (s *WebsocketPublicService) SubscribeBookTrader(
	param WebsocketPublicBookTraderParam,
	handler func(WebsocketPublicBookTraderResponse) error,
) (func() error, error) {
	// TODO
	return nil, nil
}
func (s *WebsocketPublicService) UnSubscribeBookTrader(
	param WebsocketPublicBookTraderParam,
) error {
	// TODO
	return nil
}
