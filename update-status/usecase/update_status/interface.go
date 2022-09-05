package updatestatus

import "update-status/configs"

// IUpdater method update status for database
type IUpdater interface {
	StartUpdateStatus() error
	SetNewStatus() (*configs.GroupDatabaseStatusDefinition, error)
}
