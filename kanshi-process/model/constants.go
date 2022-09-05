// Package model for store constant
package model

// number of bytes of each element
const (
	// FunctionTypeLength number of bytes
	FunctionTypeLength = 2
	// KanriIDLength number of bytes
	KanriIDLength = 6
	// UserIDLength number of bytes
	UserIDLength = 40
	// SyubetuLength number of bytes
	SyubetuLength = 4
	// QuoteCodeLength number of bytes
	QuoteCodeLength = 42
	// FromDateLength number of bytes
	FromDateLength = 8
	// ToDateLength number of bytes
	ToDateLength = 8
	// MinuteFunasiLength number of bytes
	MinuteFunasiLength = 2
	// KikanLength number of bytes
	KikanLength = 4
	// ElementsLength number of bytes
	ElementsLength = 4
	// FromTimeLength number of bytes
	FromTimeLength = 4
	// ToTimeLength number of bytes
	ToTimeLength = 4
	// YobiLength number of bytes
	YobiLength = 8
)

const (
	// ContinueFlag continue flag in 3MB, 10MB case
	ContinueFlag = "1"
	// EndFlag end flag in 3MB, 10MB case
	EndFlag = "0"
	// StrokeCharacter stroke character
	StrokeCharacter = "/"
	// UnderscoreCharacter underscore character
	UnderscoreCharacter = "_"
	// UnderscoreENCharacter underscore EN character
	UnderscoreENCharacter = "_EN"
	// EmptyString empty character
	EmptyString = ""
	// StartIndexNull null start index
	StartIndexNull = "-2"
	// StartIndexEmpty empty start index
	StartIndexEmpty = "-1"
	// TwoDotCharacter character :
	TwoDotCharacter = ":"
	// DollarCharacter chcharacter $
	DollarCharacter = "$"
	// SpaceString space string
	SpaceString = " "
	// SpaceCharacter space character
	SpaceCharacter = ' '
	// CommaCharacter comma character
	CommaCharacter = ","
	// MaxLengthQuoteCode max length of quote code
	MaxLengthQuoteCode = 42
	// DateFormatWithStroke format only date
	DateFormatWithStroke = "2006/01/02"
	// DateTimeLocal format date and time and time zone
	DateTimeLocal = "20060102 15:04 -0700"
	// DelayCharacter delay character
	DelayCharacter = "_D"
	// HighPrefix high prefix in column name
	HighPrefix = "H_"
	// EnterLineLF enter line character
	EnterLineLF = "\n"
	// EnterLineCR enter line character
	EnterLineCR = "\r"
	// TabCharacter tab character
	TabCharacter = "\t"
	// TabByte byte of tab
	TabByte = 9
	// EnterLineCRLF enter line
	EnterLineCRLF = "\r\n"
	// EmptyDate empty date
	EmptyDate = "00000000"
	// ValidDate valid date
	ValidDate = "YYYYMMDD"
	// TotalDataTypes number of data types
	TotalDataTypes = 6
	// JishouKubun Kubun of Jishou data
	JishouKubun = "J"
	// ZeroCharacter zero character
	ZeroCharacter = "0"
	// PointCharacter point character
	PointCharacter = "."
	// SpaceFourTimes four space character
	SpaceFourTimes = "    "
	// TabFourTimes four tab character
	TabFourTimes = "\t\t\t\t"
)

// Kinou types
const (
	// FirstFunctionType type 01
	FirstFunctionType = "01"
	// SecondFunctionType type 05
	SecondFunctionType = "05"
	// ErrorFunctionType type error
	ErrorFunctionType = "99"

	// TotalResponseBytesForFirstType number bytes for kinou 01
	TotalResponseBytesForFirstType = 129
	// TotalResponseBytesForSecondType number bytes for kinou 05
	TotalResponseBytesForSecondType = 132
	// TotalResponseBytesForError number bytes for error response
	TotalResponseBytesForError = 49
)

const (
	// DefaultTimeSleep time sleep after send request(second)
	DefaultTimeSleep = 60
	// DefaultTimeSleepConnect time sleep after connect to socket server fail(three times, total time sleep 60 seconds)
	DefaultTimeSleepConnect = 20
	// DefaultRetryTimesConnect times retry connect to socket server
	DefaultRetryTimesConnect = 3
	// ErrorTimeSleepConnect time sleep after connect to socket server fail fourth
	ErrorTimeSleepConnect = 60
	// ErrorRetryTimesConnect times retry connect to socket server when connect fail fourth
	ErrorRetryTimesConnect = 15
)
