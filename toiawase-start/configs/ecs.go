package configs

// ECS structure config ECS
type ECS struct {
	Region    string
	Cluster   string
	Service   string
	TaskCount int32
}
