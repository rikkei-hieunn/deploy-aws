/*
Package model declare all models use in application.
*/
package model

import "update-status/configs"

const (
	// UpdateStatusTypeQuoteCode update status follow kubun and hassin
	UpdateStatusTypeQuoteCode = "1"
	// UpdateStatusTypeDBName update status follow db name
	UpdateStatusTypeDBName = "2"
	// UpdateStatusTypeGroupID update status follow group ID
	UpdateStatusTypeGroupID = "3"
)

var (
	// TickDatabaseStatuses list of status database for two kei
	TickDatabaseStatuses configs.ArrayDatabaseStatus
	// KehaiDatabaseStatuses list of status database for two kei
	KehaiDatabaseStatuses configs.ArrayDatabaseStatus
	// QuoteCodeDefinition quote code definition follow group ID
	QuoteCodeDefinition map[string][]configs.QuoteCodes
)

const (
	// TheFirstKei the first kei
	TheFirstKei = "1"
	// TheSecondKei the second kei
	TheSecondKei = "2"

	// TickData tick data type
	TickData = "1"
	// KehaiData kehai data type
	KehaiData = "2"
)

const (
	// EmptyString empty character
	EmptyString = ""
	// TabString tab character
	TabString = "\t"
)

const (
	// S3BucketKey key environment save bucket name
	S3BucketKey = "TK_SYSTEM_BUCKET_NAME"
	// S3RegionKey key environment save region
	S3RegionKey = "TK_SYSTEM_REGION"
)
