package model

// UpdateTypeDBName type of request for update status with db name
type UpdateTypeDBName struct {
	DBName    string
	NewStatus bool
}

// UpdateTypeGroupID type of request for update status with group ID
type UpdateTypeGroupID struct {
	GroupID   string
	NewStatus bool
}

// UpdateTypeQuoteCode type of request for update status with kubun and hassin
type UpdateTypeQuoteCode struct {
	Kubun     string
	Hassin    string
	NewStatus bool
}
