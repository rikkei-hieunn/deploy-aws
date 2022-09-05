/*
Package model declare all models use in application.
*/
package model

const (
	// RequestTypeAll type send command to all
	RequestTypeAll = "1"
	// RequestTypeGroupLine type send command to the same group line
	RequestTypeGroupLine = "2"
	// RequestTypeGroupID type send command follow group ID
	RequestTypeGroupID = "0"
	// RequestTypeToiawase type send command to toiawase server
	RequestTypeToiawase = "3"
)

const (
	// ConnectionProtocol protocol for socket
	ConnectionProtocol = "tcp"
	// IPSuffix suffix for get ip by machine name
	IPSuffix = "_IP"
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

	// KeyToiawaseHost os environment key for toiawase host
	KeyToiawaseHost = "TOIAWASE_HOST"
	// KeyToiawaseTheFirstPort os environment key for toiawase the first port
	KeyToiawaseTheFirstPort = "TOIAWASE_PORT1"
	// KeyToiawaseTheSecondPort os environment key for toiawase the second port
	KeyToiawaseTheSecondPort = "TOIAWASE_PORT2"
	// KeyToiawaseTheThirdPort os environment key for toiawase the third port
	KeyToiawaseTheThirdPort = "TOIAWASE_PORT3"
)

const (
	// EmptyString empty character
	EmptyString = ""
)

const (
	// S3BucketKey key environment save bucket name
	S3BucketKey = "TK_SYSTEM_BUCKET_NAME"
	// S3RegionKey key environment save region
	S3RegionKey = "TK_SYSTEM_REGION"
)
