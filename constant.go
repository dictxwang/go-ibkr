package ibkr

type ExchangeType string

const (
	ExchangeTypeNYSE   = ExchangeType("NYSE")
	ExchangeTypeBATS   = ExchangeType("BATS")
	ExchangeTypeISLAND = ExchangeType("ISLAND")
	ExchangeTypeNASDAQ = ExchangeType("NASDAQ")
	ExchangeTypeAEB    = ExchangeType("AEB")
	ExchangeTypeBVME   = ExchangeType("BVME")
	ExchangeTypeSWX    = ExchangeType("SWX")
	ExchangeTypeIBIS   = ExchangeType("IBIS")
	ExchangeTypeLSE    = ExchangeType("LSE")
	ExchangeTypeSBF    = ExchangeType("SBF")
	ExchangeTypeVIRTX  = ExchangeType("VIRTX")
	ExchangeTypeASX    = ExchangeType("ASX")
	ExchangeTypeOSEJPN = ExchangeType("OSE.JPN")
	ExchangeTypeSGX    = ExchangeType("SGX")
	ExchangeTypeTSEJ   = ExchangeType("TSEJ")
)

type SecurityType string

const (
	SecurityTypeStock                 = SecurityType("STK")
	SecurityTypeOption                = SecurityType("OPT")
	SecurityTypeFuture                = SecurityType("FUT")
	SecurityTypeContractForDifference = SecurityType("CFD")
	SecurityTypeWarrant               = SecurityType("WAR")
	SecurityTypeForex                 = SecurityType("SWP")
	SecurityTypeMutualFund            = SecurityType("FND")
	SecurityTypeBond                  = SecurityType("BND")
	SecurityTypeInterCommoditySpreads = SecurityType("ICS")
)

type MessageTopic string

const (
	MessageTopicSubscribeMarketData             = "smd"
	MessageTopicUnsubscribeMarketData           = "umd"
	MessageTopicSubscribeHistoricalMarketData   = "smh"
	MessageTopicUnsubscribeHistoricalMarketData = "umh"
	MessageTopicSubscribeBookTrader             = "sbd"
	MessageTopicUnsubscribeBookTrader           = "ubd"
	MessageTopicSubscribeAccountSummary         = "ssd"
	MessageTopicUnsubscribeAccountSummary       = "usd"
	MessageTopicSubscribeAccountLedger          = "sld"
	MessageTopicUnsubscribeAccountLedger        = "uld"
	MessageTopicSubscribeOrder                  = "sor"
	MessageTopicUnsubscribeOrder                = "uor"
	MessageTopicSubscribePnL                    = "spl"
	MessageTopicUnsubscribePnL                  = "upl"
	MessageTopicSubscribeTradesData             = "str"
	MessageTopicUnsubscribeTradesData           = "utr"

	UnsolicitedMessageTopicAccountUpdates   = "act"
	UnsolicitedMessageTopicAuthStatus       = "sts"
	UnsolicitedMessageTopicBulletins        = "blt"
	UnsolicitedMessageTopicSystemConnection = "system"
	UnsolicitedMessageTopicNotifications    = "ntf"
)

type OrderStatus string

const (
	OrderStatusInactive      = OrderStatus("Inactive")
	OrderStatusPendingSubmit = OrderStatus("PendingSubmit")
	OrderStatusPreSubmitted  = OrderStatus("PreSubmitted")
	OrderStatusSubmitted     = OrderStatus("Submitted")
	OrderStatusFilled        = OrderStatus("Filled")
	OrderStatusPendingCancel = OrderStatus("PendingCancel")
	OrderStatusCancelled     = OrderStatus("Cancelled")
	OrderStatusWarnState     = OrderStatus("WarnState")
)

type OrderStatusFilterValue string

const (
	OrderStatusFilterValueInactive      = OrderStatusFilterValue("inactive")
	OrderStatusFilterValuePendingSubmit = OrderStatusFilterValue("pending_submit")
	OrderStatusFilterValuePreSubmitted  = OrderStatusFilterValue("pre_submitted")
	OrderStatusFilterValueSubmitted     = OrderStatusFilterValue("submitted")
	OrderStatusFilterValueFilled        = OrderStatusFilterValue("filled")
	OrderStatusFilterValuePendingCancel = OrderStatusFilterValue("pending_cancel")
	OrderStatusFilterValueCancelled     = OrderStatusFilterValue("cancelled")
	OrderStatusFilterValueWarnState     = OrderStatusFilterValue("warnState")
)
