package ibkr

import "net/url"

type ContractServiceI interface {
	GetAllCoinIds(exchange ExchangeType) (*[]CoinIdItem, error)
}

type CoinIdItem struct {
	Ticker   string `json:"ticker"`
	CoinId   int    `json:"conid"`
	Exchange string `json:"exchange"`
}

type ContractService struct {
	client *Client
}

func (s *ContractService) GetAllCoinIds(exchange ExchangeType) (*[]CoinIdItem, error) {

	var (
		res []CoinIdItem
	)

	param := url.Values{}
	param.Add("exchange", string(exchange))

	if err := s.client.getPublic("/trsrv/all-conids", param, &res); err != nil {
		return nil, err
	}

	return &res, nil
}
