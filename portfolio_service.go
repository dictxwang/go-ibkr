package ibkr

import (
	"fmt"
	"net/url"
)

type PortfolioServiceI interface {
	GetAccounts() (*[]PortfolioAccountInfo, error)
	GetSubAccounts() (*[]PortfolioAccountInfo, error)
	GetSubAccountsWithLargeAccountStructures() (*[]PortfolioAccountInfo, error)
	GetSpecificAccount(accountId string) (*PortfolioAccountInfo, error)
	GetCombinationPositions(accountId string, nocache bool) (*[]CombinationPositionItem, error)
	GetPositions(param GetPositionParam) (*[]PositionInfo, error)
	GetPositionsNew(param GetPositionParam) (*[]PositionNewInfo, error)
	GetPositionByContractId(contractId int) (*PositionInfo, error)
}

type PortfolioService struct {
	client *Client
}

func (s *PortfolioService) GetAccounts() (*[]PortfolioAccountInfo, error) {

	var res []PortfolioAccountInfo

	param := url.Values{}

	if err := s.client.getPublic("/portfolio/accounts", param, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

func (s *PortfolioService) GetSubAccounts() (*[]PortfolioAccountInfo, error) {

	var res []PortfolioAccountInfo

	param := url.Values{}

	if err := s.client.getPublic("/portfolio/subaccounts", param, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

func (s *PortfolioService) GetSubAccountsWithLargeAccountStructures() (*[]PortfolioAccountInfo, error) {

	var res []PortfolioAccountInfo

	param := url.Values{}

	if err := s.client.getPublic("/portfolio/subaccounts2", param, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

func (s *PortfolioService) GetSpecificAccount(accountId string) (*PortfolioAccountInfo, error) {

	var res PortfolioAccountInfo

	param := url.Values{}

	if err := s.client.getPublic(fmt.Sprintf("/portfolio/%s/meta", accountId), param, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

func (s *PortfolioService) GetCombinationPositions(accountId string, nocache bool) (*[]CombinationPositionItem, error) {

	var res []CombinationPositionItem

	param := url.Values{}
	param.Add("nocache", fmt.Sprintf("%t", nocache))

	if err := s.client.getPublic(fmt.Sprintf("/portfolio/%s/combo/positions", accountId), param, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

func (s *PortfolioService) GetPositions(param GetPositionParam) (*[]PositionInfo, error) {

	var res []PositionInfo

	query := url.Values{}
	if param.Model != "" {
		query.Add("model", param.Model)
	}
	if param.Sort != "" {
		query.Add("sort", param.Sort)
	}
	if param.Direction != nil {
		query.Add("direction", string(*param.Direction))
	}
	if param.Period != nil {
		query.Add("period", string(*param.Period))
	}

	if err := s.client.getPublic(fmt.Sprintf("/portfolio/%s/positions/%d", param.AccountId, param.PageId), query, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

func (s *PortfolioService) GetPositionsNew(param GetPositionParam) (*[]PositionNewInfo, error) {

	var res []PositionNewInfo

	query := url.Values{}
	if param.Model != "" {
		query.Add("model", param.Model)
	}
	if param.Sort != "" {
		query.Add("sort", param.Sort)
	}
	if param.Direction != nil {
		query.Add("direction", string(*param.Direction))
	}
	if param.Period != nil {
		query.Add("period", string(*param.Period))
	}

	if err := s.client.getPublic(fmt.Sprintf("/portfolio2/%s/positions", param.AccountId), query, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

func (s *PortfolioService) GetPositionByContractId(contractId int) (*PositionInfo, error) {

	var res PositionInfo

	query := url.Values{}

	if err := s.client.getPublic(fmt.Sprintf("/portfolio/positions/%d", contractId), query, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

type PortfolioAccountParent struct {
	MoneyManagerClientAccount []string `json:"mmc"`
	AccountId                 string   `json:"accountId"`           // Account Number for Money Manager Client
	IsMParent                 bool     `json:"isMParent,omitempty"` // Returns if this is a Multiplex Parent Account
	IsMChild                  bool     `json:"isMChild,omitempty"`  // Returns if this is a Multiplex Child Account
	IsMultiplex               bool     `json:"isMultiplex"`         // Is a Multiplex Account. These are account models with individual account being parent and managed account being child.
}
type PortfolioAccountInfo struct {
	Id                      string                 `json:"id"`
	PrepaidCryptoZ          bool                   `json:"PrepaidCrypto-Z"`
	PrepaidCryptoP          bool                   `json:"PrepaidCrypto-P"`
	BrokerageAccess         bool                   `json:"brokerageAccess"`
	AccountId               string                 `json:"accountId"`
	AccountVan              string                 `json:"accountVan"`
	AccountTitle            string                 `json:"accountTitle"`
	DisplayName             string                 `json:"displayName"`
	AccountAlias            string                 `json:"accountAlias"`
	AccountStatus           int64                  `json:"accountStatus"` // When the account was opened in unix time.
	Currency                string                 `json:"currency"`
	AccountType             string                 `json:"type"`
	TradingType             string                 `json:"tradingType"`
	BusinessType            string                 `json:"businessType"`
	IbEntity                string                 `json:"ibEntity"`
	FinancialAdvisorClient  bool                   `json:"faClient"`
	ClearingStatus          string                 `json:"clearingStatus"` // Status of the Account Potential Values: O: Open; P or N: Pending; A: Abandoned; R: Rejected; C: Closed.
	Vovestor                bool                   `json:"covestor"`
	NoClientTrading         bool                   `json:"noClientTrading"`
	TrackVirtualFXPortfolio bool                   `json:"trackVirtualFXPortfolio"`
	Parent                  PortfolioAccountParent `json:"parent"`
	Desc                    string                 `json:"desc"`
}

type CombinationPositionLeg struct {
	ContractId int     `json:"conid"`
	Ratio      float64 `json:"ratio"`
}

type GetPositionParam struct {
	AccountId string         `json:"accountId"`
	PageId    int            `json:"pageId"`
	Model     string         `json:"model"`
	Sort      string         `json:"sort"`
	Direction *SortDirection `json:"direction"`
	Period    *PeriodType    `json:"period"`
}

type PositionInfo struct {
	AccountId     string  `json:"acctId"`
	ContractId    int     `json:"conid"`
	ContractDesc  string  `json:"contractDesc"`
	Position      float64 `json:"position"`
	MarketPrice   float64 `json:"mktPrice"`
	MarketValue   float64 `json:"mktValue"`
	Currency      string  `json:"currency"`
	AverageCost   float64 `json:"avgCost"`
	AveragePrice  float64 `json:"avgPrice"`
	RealizedPnl   float64 `json:"realizedPnl"`
	UnrealizedPnl float64 `json:"unrealizedPnl"`
	AssetClass    string  `json:"assetClass"`
	PutOrCall     string  `json:"putOrCall,omitempty"`
}
type CombinationPositionItem struct {
	Name        string                   `json:"name"`
	Description string                   `json:"description"`
	Legs        []CombinationPositionLeg `json:"legs"`
	Positions   []PositionInfo           `json:"positions"`
}

type PositionNewInfo struct {
	Position      float64 `json:"position"`
	ContractId    int     `json:"conid"`
	AverageCost   float64 `json:"avgCost"`
	AveragePrice  float64 `json:"avgPrice"`
	Currency      string  `json:"currency"`
	Description   string  `json:"description"` // Returns the local symbol of the order.
	IsLastToLoq   bool    `json:"isLastToLoq"`
	MarketPrice   float64 `json:"marketPrice"`
	MarketValue   float64 `json:"marketValue"`
	RealizedPnl   float64 `json:"realizedPnl"`
	SecurityType  string  `json:"secType"`
	Timestamp     int64   `json:"timestamp"`
	UnrealizedPnl float64 `json:"unrealizedPnl"`
	AssetClass    string  `json:"assetClass"`
	Sector        string  `json:"sector"`
	Group         string  `json:"group"`
	Model         string  `json:"model"`
}
