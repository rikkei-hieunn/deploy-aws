/*
Package model contain struct and constants.
*/
package model

import (
	"create-table/configs"
)

const (
	// S3BucketKey key environment save bucket name
	S3BucketKey = "TK_SYSTEM_BUCKET_NAME"
	// S3RegionKey key environment save region
	S3RegionKey = "TK_SYSTEM_REGION"
)

const (
	// CreateTypeKei1 create table for kei 1
	CreateTypeKei1 = "1"
	// CreateTypeKei2 create table for kei 2
	CreateTypeKei2 = "2"
	// CreateTypeBoth create table for both
	CreateTypeBoth = "3"
)

var (
	// TargetCreateTableTheFirstKei []TargetCreateTable
	TargetCreateTableTheFirstKei []configs.TargetCreateTable

	// TargetCreateTableTheSecondKei []TargetCreateTable
	TargetCreateTableTheSecondKei []configs.TargetCreateTable

	// Tables map[QKBN]map[DATA_TYPE]TableDefinition
	Tables map[string]map[string]configs.TableDefinition

	// KubunsInsteadOf map[kubun]StringInsteadOfKubun
	KubunsInsteadOf map[string]string

	// TablePrefix map[DataType]Prefix
	TablePrefix map[string]string

	// OneMinuteTablePrefix one minute table prefix
	OneMinuteTablePrefix string
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
)

const (
	// QueryCheckTableExists query string check table exists in database
	QueryCheckTableExists = `SELECT TABLE_NAME FROM information_schema.tables WHERE table_schema = '%s' AND table_name = '%s' LIMIT 1;`
	// TableNameFormat format table name
	TableNameFormat = "%s_%s_%s_%s"
	// QueryCreateTableForMainData query create table for main data
	QueryCreateTableForMainData = `
create table %s
(
    QCD            varchar(42) not null,
    TIME           varchar(5)  not null,
    TKZXD          date        not null,
    TKTIM          varchar(15) not null,
    TKSERIALNUMBER int         not null,
    TKQKBN         varchar(2)  not null,
    SNDC           varchar(5)  not null,
    %s
    primary key (QCD, TIME, TKZXD, TKTIM, TKSERIALNUMBER)
);
`
	// QueryCreateTableForExtendData query create table for extend data
	QueryCreateTableForExtendData = `
create table %s
(
    QCD            varchar(47) not null,
    TIME           varchar(5)  not null,
    HTKZXD         date        not null,
    HTKTIM         varchar(15) not null,
    TKSERIALNUMBER int         not null,
    TKQKBN         varchar(2)  not null,
    SNDC           varchar(5)  not null,
    %s
    primary key (QCD, TIME, HTKZXD, HTKTIM, TKSERIALNUMBER)
);
`
	// QueryCreateTableForOneMinuteData query create table for one minute data
	QueryCreateTableForOneMinuteData = `
create table %s
(
    QCD    varchar(42) not null,
    TIME   varchar(5)  not null,
    TKQKBN varchar(2)  not null,
    SNDC   varchar(5)  not null,
    %s
    primary key (QCD, TIME)
);
`
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
