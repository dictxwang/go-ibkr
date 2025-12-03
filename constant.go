package ibkr

type OrderSide string
type ExchangeType string
type SecurityType string

const (
	OrderSideBuy  = OrderSide("BUY")
	OrderSideSell = OrderSide("SELL")
)

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

type OrderType string
type TimeInForce string
type IBAlgorithm string
type TrailingType string

// https://www.interactivebrokers.com/campus/ibkr-api-page/order-types/
const (
	OrderTypeMarket               = OrderType("MKT")
	OrderTypeMarketToLimit        = OrderType("MTL")
	OrderTypeLimit                = OrderType("LMT")
	OrderTypeStop                 = OrderType("STP")
	OrderTypeRelative             = OrderType("REL")
	OrderTypeBoxTop               = OrderType("BOX TOP")
	OrderTypeLimitIfTouched       = OrderType("LIT")
	OrderTypeLimitOnClose         = OrderType("LOC")
	OrderTypeMarketIfTouch        = OrderType("MIT")
	OrderTypeMarketOnClose        = OrderType("MOC")
	OrderTypeMarketWithProtection = OrderType("MKT PRT")
	OrderTypePassiveRelative      = OrderType("PASSV REL")
	OrderTypePeggedToStock        = OrderType("PEG STK")
	OrderTypePeggedToBenchmark    = OrderType("PEG BENCH")
	OrderTypePeggedToMarket       = OrderType("PEG MKT")
	OrderTypeStopLimit            = OrderType("STP LMT")
	OrderTypeStopWithProtection   = OrderType("STP PRT")
	OrderTypeRelativeLimitCombo   = OrderType("REL + LMT")
	OrderTypeRelativeMarketCombo  = OrderType("REL + MKT")
)

const (
	TimeInForceDAY = TimeInForce("DAY")
	TimeInForceGTC = TimeInForce("GTC")
	TimeInForceIOC = TimeInForce("IOC")
	TimeInForceOPG = TimeInForce("OPG")
	TimeInForceGTD = TimeInForce("GTD")
	TimeInForceFOK = TimeInForce("FOK")
	TimeInForceDTC = TimeInForce("DTC")
	TimeInForcePAX = TimeInForce("PAX")
)

const (
	IBAlgorithmAccumulateDistribute        = IBAlgorithm("AD")
	IBAlgorithmAccumulateDistributeAlt     = IBAlgorithm("AccuDistr")
	IBAlgorithmAccumulateAdaptive          = IBAlgorithm("Adaptive")
	IBAlgorithmArrivalPrice                = IBAlgorithm("ArrivalPx")
	IBAlgorithmBalanceImpactRisk           = IBAlgorithm("BalanceImpactRisk")
	IBAlgorithmClosePrice                  = IBAlgorithm("ClosePx")
	IBAlgorithmDarkIce                     = IBAlgorithm("DarkIce")
	IBAlgorithmMiddlePrice                 = IBAlgorithm("MIDPRICE")
	IBAlgorithmMinimiseImpact              = IBAlgorithm("MinImpact")
	IBAlgorithmPercentageOfVolume          = IBAlgorithm("PctVol")
	IBAlgorithmPriceVariantPercentage      = IBAlgorithm("PctVolPx")
	IBAlgorithmPricePriceVariantPercentage = IBAlgorithm("PctVolPx")
	IBAlgorithmSizeVariantPercentage       = IBAlgorithm("PctVolSz")
	IBAlgorithmTimeVariantPercentage       = IBAlgorithm("PctVolTm")
	IBAlgorithmTWAP                        = IBAlgorithm("Twap")
	IBAlgorithmVWAP                        = IBAlgorithm("Vwap")
)

type AdaptivePriority string

const (
	AdaptivePriorityUrgent  = AdaptivePriority("Urgent")
	AdaptivePriorityNormal  = AdaptivePriority("Normal")
	AdaptivePriorityPatient = AdaptivePriority("Patient")
)

const (
	TrailingTypeAmount  = TrailingType("amt")
	TrailingTypePercent = TrailingType("%")
)

type PositionSide string
type SortDirection string
type PeriodType string

const (
	SortDirectionAscending  = SortDirection("a")
	SortDirectionDescending = SortDirection("d")

	PositionSidePut  = PositionSide("Put")
	PositionSideCall = PositionSide("Call")

	PeriodTypeOneDay    = PeriodType("1D")
	PeriodTypeSevenDays = PeriodType("7D")
	PeriodTypeOneWeek   = PeriodType("1W")
	PeriodTypeOneMonth  = PeriodType("1M")
)

type ChildOrderType string

const (
	ChildOrderTypeAttached  = ChildOrderType("A")
	ChildOrderTypeBetaHedge = ChildOrderType("B")
	ChildOrderTypeNoChild   = ChildOrderType("0")
)
