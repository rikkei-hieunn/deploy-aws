/*
Package model contains const.
*/
package model

var(
	//KubunInsteadOf kubun instead of @@
	KubunInsteadOf    map[string]string
)
const (
	//FirstTypeRunning running with 5 params
	FirstTypeRunning = "0"
	//SecondTypeRunning running with 2 params
	SecondTypeRunning = "1"
	//StatusFail status fail in candle management
	StatusFail = "3"
)

const (
	// TheFirstKei the first kei
	TheFirstKei = "1"
	// TheSecondKei the second kei
	TheSecondKei = "2"
)

const (
	// S3BucketKey key environment save bucket name
	S3BucketKey = "TK_SYSTEM_BUCKET_NAME"
	// S3RegionKey key environment save region
	S3RegionKey = "TK_SYSTEM_REGION"
)

const (
	//EmptyString empty string character
	EmptyString = ""
	//StrokeCharacter stroke character
	StrokeCharacter = "/"
	//UnderscoreCharacter stroke character
	UnderscoreCharacter = "_"
)
