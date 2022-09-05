/*
Package model define constants
*/
package model

const (
	// FormatDate format date
	FormatDate = "20060102"
	// UnderscoreCharacter underscore character
	UnderscoreCharacter = "_"
	// StrokeCharacter slash character
	StrokeCharacter = "/"
	// QKBNAll kubun `all`
	QKBNAll = "ALL"
	// ConfigFilename Config Filename
	ConfigFilename = "environment_variables.json"
	//TypeTick data tick
	TypeTick = "1"
	//TypeKehai data kehai
	TypeKehai = "2"
	//FirstKei index first ke DB
	FirstKei = "1"
	//SecondKei index seconds ke DB
	SecondKei = "2"
	//PercentSign use for sql like select
	PercentSign = "%"
	//EmptyString empty string character
	EmptyString = ""
)
const (
	// S3BucketKey key environment save bucket name
	S3BucketKey = "TK_SYSTEM_BUCKET_NAME"
	// S3RegionKey key environment save region
	S3RegionKey = "TK_SYSTEM_REGION"
)

var (
	//InsteadOfKubun instead of @@
	InsteadOfKubun map[string]string
)
