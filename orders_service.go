package ibkr

type OrdersServiceI interface {
	PlaceOrder(params []PlaceOrderParam) (*[]PlaceOrderResponse, error)
	CancelOrder(param CancelOrderParam) (*CancelOrderResponse, error)
	PlaceOrderReplyConfirmation(param PlaceOrderReplyConfirmationParam) (*[]PlaceOrderReplyConfirmationResponse, error)
	RespondServerPrompt(param RespondServerPromptParam) (*RespondServerPromptResponse, error)
}

type PlaceOrderParam struct {
	AccountId          string `json:"acctId"`
	ContractId         string `json:"conid"`
	ContractIdExchange string `json:"conidex"`
	ManualIndicator    bool   `json:"manualIndicator"`
	ExternalOperator   string `json:"extOperator,omitempty"`
	ConSecType         string `json:"secType,omitempty"` // sample:265598:STK
	CustomOrderId      string `json:"cOID,omitempty"`
	ParentId           string `json:"parentId,omitempty"`
	OrderType          string `json:"orderType"`
}

type PlaceOrderResponse struct {
	// TODO
}

type CancelOrderParam struct {
	// TODO
}

type CancelOrderResponse struct {
	// TODO
}

type PlaceOrderReplyConfirmationParam struct {
	// TODO
}

type PlaceOrderReplyConfirmationResponse struct {
	// TODO
}

type RespondServerPromptParam struct {
	// TODO
}

type RespondServerPromptResponse struct {
	// TODO
}
