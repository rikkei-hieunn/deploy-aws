package loadconfig

import "update-status/configs"

// IConfigurationLoader interface load config service
type IConfigurationLoader interface {
	LoadConfig() error
	LoadDatabaseStatus() (*configs.GroupDatabaseStatusDefinition, error)
	LoadQuoteCodeData() (map[string][]configs.QuoteCodes, error)
}
