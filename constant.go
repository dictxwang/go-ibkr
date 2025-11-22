package ibkr

type ExchangeType string

const (
	ExchangeTypeNYSE   = ExchangeType("NYSE")
	ExchangeTypeBATS   = ExchangeType("BATS")
	ExchangeTypeINET   = ExchangeType("INET")
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
