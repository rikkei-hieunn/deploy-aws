/*
Package model contain struct and constants.
*/
package model

import "chikuseki-check/configs"

const (
	// S3BucketKey key environment save bucket name
	S3BucketKey = "TK_SYSTEM_BUCKET_NAME"
	// S3RegionKey key environment save region
	S3RegionKey = "TK_SYSTEM_REGION"
)

const (
	// TheFirstKei the first kei
	TheFirstKei = "1"
	// TheSecondKei the second kei
	TheSecondKei = "2"
)

const (
	// StrokeCharacter stroke character
	StrokeCharacter = "/"
	// UnderscoreCharacter underscore character
	UnderscoreCharacter = "_"
	// EmptyString empty character
	EmptyString = ""
	// CommaCharacter comma character
	CommaCharacter = ","
	// MaxLengthQuoteCode max length of quote code
	MaxLengthQuoteCode = 42
	// DateFormatWithStroke format only date
	DateFormatWithStroke = "2006/01/02"
	// DateFormatWithoutStroke format only date without stroke
	DateFormatWithoutStroke = "20060102"
	// EnterLineCRLF enter line
	EnterLineCRLF = "\r\n"
	// SpaceString space string
	SpaceString = " "
)

var (
	// KubunsInsteadOf map[kubun]StringInsteadOfKubun
	KubunsInsteadOf map[string]string

	// TablePrefix map[DataType]Prefix
	TablePrefix map[string]string

	// QuoteCodesTheFirstKei definition quote codes in the first kei
	QuoteCodesTheFirstKei map[string]configs.QuoteCodes

	// QuoteCodesTheSecondKei definition quote codes in the second kei
	QuoteCodesTheSecondKei map[string]configs.QuoteCodes
)

const (
	// QueryCheckTableExists query string check table exists in database
	QueryCheckTableExists = `SELECT TABLE_NAME FROM information_schema.tables WHERE table_schema = '%s' AND table_name = '%s' LIMIT 1;`
	// QueryCountAllTable query string count total records in the table
	QueryCountAllTable = `SELECT COUNT(*) FROM %s;`
	// TableNameFormat format table name
	TableNameFormat = "%s_%s_%s_%s"
)

const (
	// HigaData Higa data
	HigaData = 0
	// HigaDataString Higa data string
	HigaDataString = "0"
	// TickData Tick data
	TickData = 1
	// TickDataString Tick data string
	TickDataString = "1"
	// KehaiData Kehai data
	KehaiData = 2
	// KehaiDataString Kehai data string
	KehaiDataString = "2"
	// JishouData Jishou data
	JishouData = 3
	// JishouDataString Jishou data string
	JishouDataString = "3"
	// MoneyFlowData money flow data
	MoneyFlowData = 4
	// OptionExtendedData option extended data
	OptionExtendedData = 5
	// IndexTradingData index trading data
	IndexTradingData = 6
	// StartExtendData number start extend data type
	StartExtendData = 4
	// OneMinuteData define data type for one minute table
	OneMinuteData = "-1"
)
