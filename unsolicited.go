package ibkr

type WebsocketServerInfo struct {
	ServerName    string `json:"serverName"`
	ServerVersion string `json:"serverVersion"`
}

type WebsocketAccountAllowFeatures struct {
	ShowGFIS               bool   `json:"showGFIS,omitempty"`
	ShowEUCostReport       bool   `json:"showEUCostReport,omitempty"`
	AllowEventContract     bool   `json:"allowEventContract,omitempty"`
	AllowFXConv            bool   `json:"allowFXConv,omitempty"`
	AllowFinancialLens     bool   `json:"allowFinancialLens,omitempty"`
	AllowMTA               bool   `json:"allowMTA,omitempty"`
	AllowTypeAhead         bool   `json:"allowTypeAhead,omitempty"`
	AllowEventTrading      bool   `json:"allowEventTrading,omitempty"`
	SnapshotRefreshTimeout int    `json:"snapshotRefreshTimeout,omitempty"`
	LiteUser               bool   `json:"liteUser,omitempty"`
	ShowWebNews            bool   `json:"showWebNews,omitempty"`
	Research               bool   `json:"research,omitempty"`
	DebugPnl               bool   `json:"debugPnl,omitempty"`
	ShowTaxOpt             bool   `json:"showTaxOpt,omitempty"`
	ShowImpactDashboard    bool   `json:"showImpactDashboard,omitempty"`
	AllowDynAccount        bool   `json:"allowDynAccount,omitempty"`
	AllowCrypto            bool   `json:"allowCrypto,omitempty"`
	AllowedAssetTypes      string `json:"allowedAssetTypes,omitempty"`
}

type WebsocketUnsolicitedAccountUpdatesArgs struct {
	Accounts                   []string                      `json:"accounts,omitempty"`
	AccountProperties          interface{}                   `json:"acctProps,omitempty"`
	Aliases                    interface{}                   `json:"aliases,omitempty"`
	AllowFeatures              WebsocketAccountAllowFeatures `json:"allowFeatures,omitempty"`
	ChartPeriods               interface{}                   `json:"chartPeriods,omitempty"`
	Groups                     interface{}                   `json:"groups,omitempty"`
	Profiles                   interface{}                   `json:"profiles,omitempty"`
	SelectedAccount            string                        `json:"selectedAccount,omitempty"`
	ServerInfo                 WebsocketServerInfo           `json:"serverInfo"`
	SessionId                  string                        `json:"sessionId,omitempty"`
	IsFractionalTradingAccount bool                          `json:"isFT,omitempty"`
	IsPaperTradingAccount      bool                          `json:"isPaper,omitempty"`
}

type WebsocketUnsolicitedAccountUpdatesResponse struct {
	Topic string                                 `json:"topic"`
	Args  WebsocketUnsolicitedAccountUpdatesArgs `json:"args,omitempty"`
}

type WebsocketUnsolicitedAuthStatusArgs struct {
	Authenticated bool   `json:"authenticated"`
	Competing     bool   `json:"competing"`
	Connected     bool   `json:"connected"`
	Message       string `json:"message,omitempty"`
	Fail          string `json:"fail,omitempty"`
	ServerName    string `json:"serverName,omitempty"`
	ServerVersion string `json:"serverVersion,omitempty"`
	UserName      string `json:"username,omitempty"`
}
type WebsocketUnsolicitedAuthStatusResponse struct {
	Topic string                             `json:"topic"`
	Args  WebsocketUnsolicitedAuthStatusArgs `json:"args,omitempty"`
}

type WebsocketUnsolicitedBulletinsArgs struct {
	Id      string `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
}
type WebsocketUnsolicitedBulletinsResponse struct {
	Topic string                            `json:"topic"`
	Args  WebsocketUnsolicitedBulletinsArgs `json:"args,omitempty"`
}

type WebsocketUnsolicitedSystemConnectionResponse struct {
	Topic                      string `json:"topic"`
	Success                    string `json:"success,omitempty"`
	IsFractionalTradingAccount bool   `json:"isFT,omitempty"`
	IsPaperTradingAccount      bool   `json:"isPaper,omitempty"`
	HB                         int64  `json:"hb,omitempty"`
}

type WebsocketUnsolicitedNotificationsArgs struct {
	Id    string `json:"id"`
	Text  string `json:"text,omitempty"`
	Title string `json:"title,omitempty"`
	Url   string `json:"url,omitempty"`
}
type WebsocketUnsolicitedNotificationsResponse struct {
	Topic string                                `json:"topic"`
	Args  WebsocketUnsolicitedNotificationsArgs `json:"args,omitempty"`
}
