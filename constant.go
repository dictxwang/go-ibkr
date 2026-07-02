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
	ExchangeTypeOVERNIGHT   = ExchangeType("OVERNIGHT")
)

const (
	SecurityTypeStock                 = SecurityType("STK")
	SecurityTypeOption                = SecurityType("OPT")
	SecurityTypeFuture                = SecurityType("FUT")
	SecurityTypeContractForDifference = SecurityType("CFD")
	SecurityTypeWarrant               = SecurityType("WAR")
	SecurityTypeForex                 = SecurityType("CASH")
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
	TimeInForceOVT = TimeInForce("OVT")
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

type SuppressibleMessageId string
const (
  SuppressibleMessageId_O163 = SuppressibleMessageId("o163") // The following order exceeds the price percentage limit
  SuppressibleMessageId_O354 = SuppressibleMessageId("o354") // You are submitting an order without market data. We strongly recommend against this as it may result in erroneous and unexpected trades. Are you sure you want to submit this order?
  SuppressibleMessageId_O382 = SuppressibleMessageId("o382") // The following value exceeds the tick size limit
  SuppressibleMessageId_O383 = SuppressibleMessageId("o383") // The following order "BUY 650 AAPL NASDAQ.NMS" size exceeds the Size Limit of 500. Are you sure you want to submit this order?
  SuppressibleMessageId_O403 = SuppressibleMessageId("o403") // This order will most likely trigger and fill immediately. Are you sure you want to submit this order?
  SuppressibleMessageId_O451 = SuppressibleMessageId("o451") // The following order "BUY 650 AAPL NASDAQ.NMS" value estimate of 124,995.00 USD exceeds the Total Value Limit of 100,000 USD. Are you sure you want to submit this order?
  SuppressibleMessageId_O2136 = SuppressibleMessageId("o2136") // Mixed allocation order warning
  SuppressibleMessageId_O2137 = SuppressibleMessageId("o2137") // Cross side order warning
  SuppressibleMessageId_O2165 = SuppressibleMessageId("o2165") // Warns that instrument does not support trading in fractions outside regular trading hours
  SuppressibleMessageId_O10082 = SuppressibleMessageId("o10082") // Called Bond warning
  SuppressibleMessageId_O10138 = SuppressibleMessageId("o10138") // The following order size modification exceeds the size modification limit.
  SuppressibleMessageId_O10151 = SuppressibleMessageId("o10151") // Warns about risks with Market Orders
  SuppressibleMessageId_O10152 = SuppressibleMessageId("o10152") // Warns about risks associated with stop orders once they become active
  SuppressibleMessageId_O10153 = SuppressibleMessageId("o10153") // <h4>Confirm Mandatory Cap Price</h4>To avoid trading at a price that is not consistent with a fair and orderly market, IB may set a cap (for a buy order) or sell order). THIS MAY CAUSE AN ORDER THAT WOULD OTHERWISE BE MARKETABLE TO NOT BE TRADED.
  SuppressibleMessageId_O10164 = SuppressibleMessageId("o10164") // Traders are responsible for understanding cash quantity details, which are provided on a best efforts basis only.
  SuppressibleMessageId_O10223 = SuppressibleMessageId("o10223") // <h4>Cash Quantity Order Confirmation</h4>Orders that express size using a monetary value (cash quantity) are provided on a non-guaranteed basis. The system simulates the order by cancelling it once the specified amount is spent (for buy orders) or collected (for sell orders). In addition to the monetary value, the order uses a maximum size that is calculated using the Cash Quantity Estimate Factor, which you can modify in Presets.
  SuppressibleMessageId_O10288 = SuppressibleMessageId("o10288") // Warns about risks associated with market orders for Crypto
  SuppressibleMessageId_O10331 = SuppressibleMessageId("o10331") // You are about to submit a stop order. Please be aware of the various stop order types available and the risks associated with each one. Are you sure you want to submit this order?
  SuppressibleMessageId_O10332 = SuppressibleMessageId("o10332") // OSL Digital Securities LTD Crypto Order Warning
  SuppressibleMessageId_O10333 = SuppressibleMessageId("o10333") // Option Exercise at the Money warning
  SuppressibleMessageId_O10334 = SuppressibleMessageId("o10334") // Warns that order will be placed into current omnibus account instead of currently selected global account.
  SuppressibleMessageId_O10335 = SuppressibleMessageId("o10335") // Serves internal Rapid Entry window.
  SuppressibleMessageId_O10336 = SuppressibleMessageId("o10336") // This security has limited liquidity. If you choose to trade this security, there is a heightened risk that you may not be able to close your position at the time you wish, at a price you wish, and/or without incurring a loss. Confirm that you understand the risks of trading illiquid securities. Are you sure you want to submit this order?
  SuppressibleMessageId_P6 = SuppressibleMessageId("p6") // This order will be distributed over multiple accounts. We strongly suggest you familiarize yourself with our allocation facilities before submitting orders.
  SuppressibleMessageId_P12 = SuppressibleMessageId("p12") // If your order is not immediately executable, our systems may, depending on market conditions, reject your order if its limit price is more than the allowed amount away from the reference price at that time. If this happens, you will not receive a fill. This is a control designed to ensure that we comply with our regulatory obligations to avoid submitting disruptive orders to the marketplace. Use the Price Management Algo?
)