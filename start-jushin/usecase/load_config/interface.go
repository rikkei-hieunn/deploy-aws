/*
Package loadconfig implements logics repository.
*/
package loadconfig

//ILoadConfig provide all services for load config
type ILoadConfig interface {
	LoadConfig(startType, keiType, groupID, dataType, groupLine string) error
}
