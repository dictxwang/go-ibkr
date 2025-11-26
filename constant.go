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
	MessageTopicSubscribeTicker           = "smd"
	MessageTopicUnsubscribeTicker         = "umd"
	MessageTopicSubscribeBookTrader       = "sbd"
	MessageTopicUnsubscribeBookTrader     = "ubd"
	MessageTopicSubscribeAccountSummary   = "ssd"
	MessageTopicUnsubscribeAccountSummary = "usd"
	MessageTopicSubscribeAccountLedger    = "sld"
	MessageTopicUnsubscribeAccountLedger  = "uld"
	MessageTopicSubscribeOrder            = "sor"
	MessageTopicUnsubscribeOrder          = "uor"
	MessageTopicSubscribePnL              = "spl"
	MessageTopicUnsubscribePnL            = "upl"
	MessageTopicSubscribeTradesData       = "str"
	MessageTopicUnsubscribeTradesData     = "utr"
)
