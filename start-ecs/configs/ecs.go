package configs
//ECS ecs struct definition
type ECS struct {
	Region         string
	ClusterName    string
	TaskDefinition string
	Subnets        []string
	SecurityGroups []string
	ContainerName  string
}
