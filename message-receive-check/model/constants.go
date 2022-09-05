package model

const (
	// LogFileExtension extension for log file
	LogFileExtension = ".txt"

	// EmptyString empty character
	EmptyString = ""
	// CommaCharacter comma character
	CommaCharacter = ","
	// KehaiPrefix prefix Kehai file
	KehaiPrefix = "K"
)

const (
	// BestQuoteData Tick data
	BestQuoteData = "1"
	// MultipleQuoteData Kehai data
	MultipleQuoteData = "2"
	// MasterQuoteData Jishou data
	MasterQuoteData = "3"
	// MoneyFlowData money flow data
	MoneyFlowData = "4"
	// OptionExtendedData option extended data
	OptionExtendedData = "5"
	// IndexTradingData Index trading
	IndexTradingData = "6"
)

const (
	// S3BucketKey key environment save bucket name
	S3BucketKey = "TK_SYSTEM_BUCKET_NAME"
	// S3RegionKey key environment save region
	S3RegionKey = "TK_SYSTEM_REGION"
)
