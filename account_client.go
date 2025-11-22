package ibkr

import (
	"net/url"
)

type AccountServiceI interface {
	GetProfitAndLoss() (*AccountProfitAndLossResponse, error)
}

// AccountService :
type AccountService struct {
	client *Client
}

type AccountProfitAndLossInfo struct {
	RowType         int     `json:"rowType"` // Returns the positional value of the returned account. Always returns 1 for individual accounts.
	DailyPnL        float64 `json:"dpl"`     // Daily PnL for the specified account profile.
	NetLiquidity    float64 `json:"nl"`      // Net Liquidity for the specified account profile.
	UnPnL           float64 `json:"upl"`     // Unrealized PnL for the specified account profile.
	ExcessLiquidity float64 `json:"el"`      // Excess Liquidity for the specified account profile.
	MarginValue     float64 `json:"mv"`      // Margin value for the specified account profile.
}

type AccountProfitAndLossResponse struct {
	UserPnL map[string]AccountProfitAndLossInfo `json:"upnl"`
}

func (r *AccountProfitAndLossResponse) GetUserPnL(accountId string) (*AccountProfitAndLossInfo, bool) {
	origin, has := r.UserPnL[accountId+".Core"]
	if has {
		return &origin, has
	} else {
		return nil, false
	}
}

func (s *AccountService) GetProfitAndLoss() (*AccountProfitAndLossResponse, error) {

	var (
		res AccountProfitAndLossResponse
	)

	param := url.Values{}

	if err := s.client.getPublic("/iserver/account/pnl/partitioned", param, &res); err != nil {
		return nil, err
	}

	return &res, nil
}
