package loaddata

// ILoader interface load database status
type ILoader interface {
	LoadDatabaseStatus() error
}
