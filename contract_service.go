package ibkr

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

type ContractServiceI interface {
	GetAllContractIds(exchange ExchangeType) (*[]ContractIdItem, error)
	SearchSecurityDefinitionByContactId(contractIds []int) (*SearchSecurityDefinitionResponse, error)
	GetContractInfoByContractId(contractId int) (*GetContractInfoResponse, error)
	GetCurrencyPairs(currency string) (*map[string][]CurrencyPairItem, error)
	GetCurrencyExchangeRate(source, target string) (*GetCurrencyExchangeRateResponse, error)
	GetContractFullAllInfoAndRules(contractId int, isBuy *bool) (*GetContractAllInfoAndRulesResponse, error)
	SearchContractBySymbol(query SearchContractBySymbolQuery) (*[]SearchContractBySymbolItem, error)
	SearchContractRules(query SearchContractRulesQuery) (*ContractRules, error)
	GetSecurityFuturesBySymbol(symbols []string) (*map[string][]FuturesInfoItem, error)
	GetSecurityStocksBySymbol(symbols []string) (*map[string][]StockInfoItem, error)
	GetTradingScheduleBySymbol(query TradingScheduleQuery) (*[]TradingScheduleItem, error)
}

type TradingScheduleQuery struct {
	AssetClass     string  `json:"assetClass"`
	ContractId     int     `json:"conid"`
	Symbol         string  `json:"symbol"`
	Exchange       *string `json:"exchange,omitempty"`
	ExchangeFilter *string `json:"exchangeFilter,omitempty"`
}

type ScheduleSession struct {
	OpeningTime string `json:"openingTime"`
	ClosingTime string `json:"closingTime"`
	Prop        string `json:"prop"`
}

type ScheduleTradingTime struct {
	OpeningTime     string `json:"openingTime"`
	ClosingTime     string `json:"closingTime"`
	CancelDayOrders string `json:"cancelDayOrders"`
}

type TradingScheduleItem struct {
	ClearingCycleEndTime string                `json:"clearingCycleEndTime"`
	TradingScheduleDate  string                `json:"tradingScheduleDate"`
	Sessions             []ScheduleSession     `json:"sessions"`
	TradingTimes         []ScheduleTradingTime `json:"tradingtimes"`
}

type StockContractItem struct {
	ContractId int    `json:"conid"`
	Exchange   string `json:"exchange"`
	IsUS       bool   `json:"isUS"`
}

type StockInfoItem struct {
	Name        string              `json:"name"`
	ChineseName string              `json:"chineseName"`
	AssetClass  string              `json:"assetClass"`
	Contracts   []StockContractItem `json:"contracts"`
}

type FuturesInfoItem struct {
	Symbol               string `json:"symbol"`
	ContractId           int    `json:"conid"`
	UnderlyingContractId int    `json:"underlyingConid"`
	ExpirationDate       int    `json:"expirationDate"`
	LastTradeDate        int    `json:"ltd"`
	ShortFuturesCutOff   int    `json:"shortFuturesCutOff"`
	LongFuturesCutOff    int    `json:"longFuturesCutOff"`
}

type SearchContractRulesQuery struct {
	ContractId  int     `json:"conid"`
	Exchange    *string `json:"exchange,omitempty"`
	IsBuy       *bool   `json:"isBuy,omitempty"`
	ModifyOrder *bool   `json:"modifyOrder,omitempty"`
	OrderId     *int    `json:"orderId,omitempty"`
}

type SearchContractBySymbolQuery struct {
	Symbol       string
	Name         *bool
	SecurityType *SecurityType
}

type SearchContractBySymbolItem struct {
	ContractId    int     `json:"conid"`
	CompanyHeader string  `json:"companyHeader"`
	CompanyName   string  `json:"companyName"`
	Symbol        string  `json:"symbol"`
	Description   *string `json:"description,omitempty"`
	Restricted    *bool   `json:"restricted"`
	SecType       string  `json:"secType"`
}

type ContractRuleOrderDefaultItem struct {
	Lp string `json:"LP,omitempty"`
}

type ContractRules struct {
	AlgoEligible      bool                                    `json:"algoEligible"`
	OvernightEligible bool                                    `json:"overnightEligible"`
	CostReport        bool                                    `json:"costReport"`
	CanTradeAcctIds   []string                                `json:"canTradeAcctIds"`
	Error             *string                                 `json:"error,omitempty"`
	OrderTypes        []string                                `json:"orderTypes"`
	IbAlgoTypes       []string                                `json:"ibAlgoTypes"`
	FraqTypes         []string                                `json:"fraqTypes"`
	ForceOrderPreview bool                                    `json:"forceOrderPreview"`
	CqtTypes          []string                                `json:"cqtTypes"`
	OrderDefaults     map[string]ContractRuleOrderDefaultItem `json:"orderDefaults,omitempty"`
	OrderTypesOutside []string                                `json:"orderTypesOutside"`
	DefaultSize       float64                                 `json:"defaultSize"`
	CashSize          float64                                 `json:"cashSize"`
	SizeIncrement     float64                                 `json:"sizeIncrement"`
	TifTypes          []string                                `json:"tifTypes"`
	TifDefaults       map[string]interface{}                  `json:"tifDefaults"`
	LimitPrice        float64                                 `json:"limitPrice"`
	StopPrice         float64                                 `json:"stopprice"`
	OrderOrigination  *string                                 `json:"orderOrigination"`
	Preview           bool                                    `json:"preview"`
	DisplaySize       *string                                 `json:"displaySize,omitempty"`
	FraqInt           int                                     `json:"fraqInt"`
	CashCcy           string                                  `json:"cashCcy"`
	CashQtyIncr       float64                                 `json:"cashQtyIncr"`
	PriceMagnifier    *string                                 `json:"priceMagnifier"`
	NegativeCapable   bool                                    `json:"negativeCapable"`
	IncrementType     int                                     `json:"incrementType"`
	IncrementRules    []IncrementRule                         `json:"incrementRules"`
	HasSecondary      bool                                    `json:"hasSecondary"`
	Increment         bool                                    `json:"increment"`
	IncrementDigits   int                                     `json:"incrementDigits"`
}

type GetContractAllInfoAndRulesResponse struct {
	CfiCode                   string        `json:"cfi_code"`
	Symbol                    string        `json:"symbol"`
	Cusip                     *string       `json:"cusip,omitempty"`
	ExpiryFull                *string       `json:"expiry_full,omitempty"`
	ContractId                int           `json:"con_id"`
	MaturityDate              *string       `json:"maturity_date,omitempty"`
	Industry                  string        `json:"industry"`
	InstrumentType            string        `json:"instrument_type"'`
	TradingClass              string        `json:"trading_class"`
	ValidExchanges            string        `json:"valid_exchanges"`
	AllowSellLong             bool          `json:"allow_sell_long"`
	IsZeroCommissionSecurity  bool          `json:"is_zero_commission_security"`
	LocalSymbol               string        `json:"local_symbol"`
	ContractClarificationType *string       `json:"contract_clarification_type,omitempty"`
	Classifier                *string       `json:"classifier,omitempty"`
	Currency                  string        `json:"currency"`
	Text                      *string       `json:"text,omitempty"`
	UnderlyingContractId      int           `json:"underlying_con_id"`
	RegularTradingHour        bool          `json:"r_t_h"`
	Multiplier                *string       `json:"multiplier,omitempty"`
	UnderlyingIssuer          *string       `json:"underlying_issuer,omitempty"`
	ContractMonth             *string       `json:"contract_month,omitempty"`
	CompanyName               string        `json:"company_name"`
	SmartAvailable            bool          `json:"smart_available"`
	Exchange                  string        `json:"exchange"`
	Category                  string        `json:"category"`
	Rules                     ContractRules `json:"rules"`
}

type GetCurrencyExchangeRateResponse struct {
	Rate float64 `json:"rate"`
}

type CurrencyPairItem struct {
	Symbol     string `json:"symbol"`
	ContractId int    `json:"conid"`
	CcyPair    string `json:"ccyPair"`
}

type GetContractInfoResponse struct {
	CfiCode                   string  `json:"cfi_code"`
	Symbol                    string  `json:"symbol"`
	Cusip                     *string `json:"cusip,omitempty"`
	ExpiryFull                *string `json:"expiry_full,omitempty"`
	ContractId                int     `json:"con_id"`
	MaturityDate              *string `json:"maturity_date,omitempty"`
	Industry                  string  `json:"industry"`
	InstrumentType            string  `json:"instrument_type"'`
	TradingClass              string  `json:"trading_class"`
	ValidExchanges            string  `json:"valid_exchanges"`
	AllowSellLong             bool    `json:"allow_sell_long"`
	IsZeroCommissionSecurity  bool    `json:"is_zero_commission_security"`
	LocalSymbol               string  `json:"local_symbol"`
	ContractClarificationType *string `json:"contract_clarification_type,omitempty"`
	Classifier                *string `json:"classifier,omitempty"`
	Currency                  string  `json:"currency"`
	Text                      *string `json:"text,omitempty"`
	UnderlyingContractId      int     `json:"underlying_con_id"`
	RegularTradingHour        bool    `json:"r_t_h"`
	Multiplier                *string `json:"multiplier,omitempty"`
	UnderlyingIssuer          *string `json:"underlying_issuer,omitempty"`
	ContractMonth             *string `json:"contract_month,omitempty"`
	CompanyName               string  `json:"company_name"`
	SmartAvailable            bool    `json:"smart_available"`
	Exchange                  string  `json:"exchange"`
	Category                  string  `json:"category"`
}

type IncrementRule struct {
	LowerEdge float64 `json:"lowerEdge"`
	Increment float64 `json:"increment"`
}

type DisplayRuleStep struct {
	DecimalDigits int     `json:"decimalDigits"`
	LowerEdge     float64 `json:"lowerEdge"`
	WholeDigits   int     `json:"wholeDigits"`
}

type DisplayRule struct {
	Magnification   int             `json:"magnification"`
	DisplayRuleStep DisplayRuleStep `json:"displayRuleStep"`
}

type SearchSecurityDefinitionItem struct {
	Conid           int             `json:"conid"`
	Currency        string          `json:"currency"`
	Time            int             `json:"time"`
	ChineseName     string          `json:"chineseName"`
	AllExchanges    string          `json:"allExchanges"`
	ListingExchange string          `json:"listingExchange"`
	CountryCode     string          `json:"countryCode"`
	Name            string          `json:"name"`
	AssetClass      string          `json:"assetClass"`
	Expiry          *string         `json:"expiry,omitempty"`
	LastTradingDay  *string         `json:"lastTradingDay,omitempty"`
	Group           string          `json:"group"`
	PutOrCall       *string         `json:"putOrCall,omitempty"`
	Sector          string          `json:"sector"`
	SectorGroup     string          `json:"sectorGroup"`
	Strike          string          `json:"strike"`
	Ticker          string          `json:"ticker"`
	UndConid        int             `json:"undConid"`
	Multiplier      float64         `json:"multiplier"`
	Type            string          `json:"type"`
	HasOptions      bool            `json:"hasOptions"`
	FullName        string          `json:"fullName"`
	IsUS            bool            `json:"isUS"`
	IncrementRules  []IncrementRule `json:"incrementRules"`
	DisplaceRule    DisplayRule     `json:"displayRule"`
	IsEventContract bool            `json:"isEventContract"`
	PageSize        int             `json:"pageSize"`
}

type SearchSecurityDefinitionResponse struct {
	SecurityDefinitions []SearchSecurityDefinitionItem `json:"SearchSecurityDefinitionItem"`
}

type ContractIdItem struct {
	Ticker     string `json:"ticker"`
	ContractId int    `json:"conid"`
	Exchange   string `json:"exchange"`
}

type ContractService struct {
	client *Client
}

func (s *ContractService) GetAllContractIds(exchange ExchangeType) (*[]ContractIdItem, error) {

	var (
		res []ContractIdItem
	)

	param := url.Values{}
	param.Add("exchange", string(exchange))

	if err := s.client.getPublic("/trsrv/all-conids", param, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

func (s *ContractService) SearchSecurityDefinitionByContactId(contractIds []int) (*SearchSecurityDefinitionResponse, error) {
	var res SearchSecurityDefinitionResponse

	conIds := make([]string, 0)
	for _, contractId := range contractIds {
		conIds = append(conIds, string(contractId))
	}
	param := url.Values{}
	param.Add("conids", strings.Join(conIds, ","))

	if err := s.client.getPublic("/trsrv/secdef", param, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

func (s *ContractService) GetContractInfoByContractId(contractId int) (*GetContractInfoResponse, error) {

	var res GetContractInfoResponse

	param := url.Values{}
	path := fmt.Sprintf("/iserver/contract/%d/info", contractId)

	if err := s.client.getPublic(path, param, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

func (s *ContractService) GetCurrencyPairs(currency string) (*map[string][]CurrencyPairItem, error) {

	var (
		res map[string][]CurrencyPairItem
	)

	param := url.Values{}
	param.Add("currency", currency)

	if err := s.client.getPublic("/iserver/currency/pairs", param, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

func (s *ContractService) GetCurrencyExchangeRate(source, target string) (*GetCurrencyExchangeRateResponse, error) {

	var (
		res GetCurrencyExchangeRateResponse
	)

	param := url.Values{}
	param.Add("source", source)
	param.Add("target", target)

	if err := s.client.getPublic("/iserver/exchangerate", param, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

func (s *ContractService) GetContractFullAllInfoAndRules(contractId int, isBuy *bool) (*GetContractAllInfoAndRulesResponse, error) {

	var res GetContractAllInfoAndRulesResponse

	param := url.Values{}
	param.Add("isBuy", strconv.FormatBool(*isBuy))

	path := fmt.Sprintf("/iserver/contract/%d/info-and-rules", contractId)

	if err := s.client.getPublic(path, param, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

func (s *ContractService) SearchContractBySymbol(query SearchContractBySymbolQuery) (*[]SearchContractBySymbolItem, error) {

	var res []SearchContractBySymbolItem

	param := url.Values{}
	param.Add("symbol", query.Symbol)
	if query.Name != nil {
		param.Add("name", strconv.FormatBool(*(query.Name)))
	}
	if query.SecurityType != nil {
		param.Add("secType", string(*query.SecurityType))
	}

	if err := s.client.getPublic("/iserver/secdef/search", param, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

func (s *ContractService) SearchContractRules(query SearchContractRulesQuery) (*ContractRules, error) {

	var res ContractRules

	body, err := json.Marshal(query)
	if err != nil {
		return nil, err
	}

	if err := s.client.postJSON("/iserver/contract/rules", body, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

func (s *ContractService) GetSecurityFuturesBySymbol(symbols []string) (*map[string][]FuturesInfoItem, error) {

	var res map[string][]FuturesInfoItem

	param := url.Values{}
	param.Add("symbols", strings.Join(symbols, ","))

	if err := s.client.getPublic("/trsrv/futures", param, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

func (s *ContractService) GetSecurityStocksBySymbol(symbols []string) (*map[string][]StockInfoItem, error) {

	var res map[string][]StockInfoItem

	param := url.Values{}
	param.Add("symbols", strings.Join(symbols, ","))

	if err := s.client.getPublic("/trsrv/stocks", param, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

func (s *ContractService) GetTradingScheduleBySymbol(query TradingScheduleQuery) (*[]TradingScheduleItem, error) {

	var res []TradingScheduleItem

	param := url.Values{}
	param.Add("assetClass", query.AssetClass)
	param.Add("conid", fmt.Sprintf("%d", query.ContractId))
	param.Add("symbol", query.Symbol)
	if query.Exchange != nil {
		param.Add("exchange", *query.Exchange)
	}
	if query.ExchangeFilter != nil {
		param.Add("exchangeFilter", *query.ExchangeFilter)
	}
	if err := s.client.getPublic("/trsrv/secdef/schedule", param, &res); err != nil {
		return nil, err
	}

	return &res, nil
}
