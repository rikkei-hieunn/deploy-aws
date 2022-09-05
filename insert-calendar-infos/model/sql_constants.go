package model

const (
	// SQLTruncateTable sql syntax truncate
	SQLTruncateTable = "DELETE FROM "
	// FormatInsert sql syntax insert
	FormatInsert = "INSERT INTO "
	// FormatInsertValue sql syntax values
	FormatInsertValue = " VALUES "
	// FormatFieldInsert sql syntax insert field value
	FormatFieldInsert = " (CALKBN, JAPANESE_DATE, DAY_WEEK, RTYPE_CODE, SPECIAL_CODE, CLOSED_REASON)"
	// FormatInsertOneField sql syntax insert batch one value
	FormatInsertOneField = "(?,?,?,?,?,?)"
	// NullString NULL string
	NullString = "NULL"
)
