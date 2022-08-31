/*
Package model contain struct and constants.
*/
package model

import "process-get-data/configs"

const (
	// GetDataTypeKei1 get data for kei 1
	GetDataTypeKei1 = "1"
	// GetDataTypeKei2 get data for kei 2
	GetDataTypeKei2 = "2"
	// GetDataTypeBoth get data for both
	GetDataTypeBoth = "3"

	// RequestType type for request cron tab
	RequestType = 1
)

const (
	// WorkingFlag flag business day
	WorkingFlag = "1"
	// NotWorkingFlag flag not business day
	NotWorkingFlag = "0"

	// DefaultFolderPath default value for folder path column
	DefaultFolderPath = "-"
	// DefaultStartIndex default value for start index column
	DefaultStartIndex = "-1"

	// CreatingStatusNotCreated status for candle management not created
	CreatingStatusNotCreated = "0"

	// QueryStringInsertCandleManagement base query insert data to candle management table
	QueryStringInsertCandleManagement = `INSERT INTO %s (TKQKBN,SNDC,OPERATION_TYPE,START_INDEX,CREATE_TIME,CREATE_DAY,
	FOLDER_PATH,START_TIME,END_TIME,TABLE_NAME,CREATE_STATUS,
	END_INDEX,QUOTE_CODE,TOOL_ID,BACK_UP_FLG,REQUEST_DATE,UPDATE_DATE) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`
)

var (
	// OneMinuteConfigs slice one minute config
	OneMinuteConfigs []configs.OneMinuteConfig

	// QuoteCodesDefinitionTheFirstKei map[kubun/hassin]configs.QuoteCodes
	QuoteCodesDefinitionTheFirstKei map[string]configs.QuoteCodes

	// QuoteCodesDefinitionTheSecondKei map[kubun/hassin]configs.QuoteCodes
	QuoteCodesDefinitionTheSecondKei map[string]configs.QuoteCodes

	// CandleManagementPrefix prefix for candle management table name
	CandleManagementPrefix string

	// KubunsInsteadOf map[kubun]StringInsteadOfKubun
	KubunsInsteadOf map[string]string

	// TablePrefix map[DataType]Prefix
	TablePrefix map[string]string
)

const (
	// StrokeCharacter stroke character
	StrokeCharacter = "/"
	// UnderscoreCharacter underscore character
	UnderscoreCharacter = "_"
	// EnterLine enter line string
	EnterLine = "\n"
	// EmptyString empty character
	EmptyString = ""
	// NullString null character
	NullString = "NULL"
	// DateFormatWithDash format only date
	DateFormatWithDash = "2006-01-02"
	// DateFormatWithoutStroke format only date without stroke
	DateFormatWithoutStroke = "20060102"
)

const (
	// S3BucketKey key environment save bucket name
	S3BucketKey = "TK_SYSTEM_BUCKET_NAME"
	// S3RegionKey key environment save region
	S3RegionKey = "TK_SYSTEM_REGION"
)
