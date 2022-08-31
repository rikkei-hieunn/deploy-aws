// Package model store constants
package model

const (
	//StrokeCharacter stroke character
	StrokeCharacter = "/"
	//UnderscoreCharacter underscore character
	UnderscoreCharacter = "_"
	//CommaCharacter comma character
	CommaCharacter = ","
	//EmptyString empty string
	EmptyString = ""
)

const (
	//BP03 process BP03
	BP03 = "BP03"
	//BP06 process BP06
	BP06 = "BP06"
	//BP05 process BP05
	BP05 = "BP05"
	//BP07 process BP07
	BP07 = "BP07"

	//BP03FirstRunningType define first kind to run service BP-03
	BP03FirstRunningType = "0"
	//BP03SecondRunningType define second kind to run service BP-03
	BP03SecondRunningType = "1"

	//BP07FirstRunningType define first kind to run service BP-07
	BP07FirstRunningType = "0"
	//BP07SecondRunningType define second kind to run service BP-07
	BP07SecondRunningType = "1"
	//BP07ThirdRunningType define third kind to run service BP-07
	BP07ThirdRunningType = "2"

	//BP05DemegetRunningType define DEMEGET Operation Type for service BP-05
	BP05DemegetRunningType = "DEMEGET"
	//BP05DemegetEvRunningType define DEMEGET-EV Operation Type for service BP-05
	BP05DemegetEvRunningType = "DEMEGET-EV"
	//BP05DemegetPtsRunningType define DEMEGET-PTS Operation Type for service BP-05
	BP05DemegetPtsRunningType = "DEMEGET-PTS"
	//BP05DemegetPtsEvRunningType define DEMEGET-PTSEV Operation Type for service BP-05
	BP05DemegetPtsEvRunningType = "DEMEGET-PTSEV"

	//BP05DownloadAllRunningType define DOWNLOADALL Operation Type for service BP-05
	BP05DownloadAllRunningType = "DOWNLOADALL"
	//BP05DownloadAllEvoRunningType define DOWNLOADALL-EVO Operation Type for service BP-05
	BP05DownloadAllEvoRunningType = "DOWNLOADALL-EVO"
	//BP05DownloadAllPtsRunningType define DOWNLOADALL-PTS Operation Type for service BP-05
	BP05DownloadAllPtsRunningType = "DOWNLOADALL-PTS"
	//BP05DownloadAllEvPtsRunningType define DOWNLOADALL-EVPTS Operation Type for service BP-05
	BP05DownloadAllEvPtsRunningType = "DOWNLOADALL-EVPTS"

	//BP05RecreateRunningType define RECREATE Operation Type for service BP-05
	BP05RecreateRunningType = "RECREATE"
	//BP05RecreatePtsRunningType define RECREATE-PTS Operation Type for service BP-05
	BP05RecreatePtsRunningType = "RECREATE-PTS"

	//BP05ReDownloadRunningType define REDOWNLOAD Operation Type for service BP-05
	BP05ReDownloadRunningType = "REDOWNLOAD"
	//BP05ReDownloadEvRunningType define REDOWNLOAD-EV Operation Type for service BP-05
	BP05ReDownloadEvRunningType = "REDOWNLOAD-EV"

	//TaskStatusRunning status of ecs
	TaskStatusRunning = "RUNNING"
	//ClusterName a param of ecs
	ClusterName = "CLUSTER_NAME"
	//TaskDefinition a  param of ecs
	TaskDefinition = "TASK_DEFINITION"
	//Subnets a  param of ecs
	Subnets = "SUBNETS"
	//SecurityGroups a  param of ecs
	SecurityGroups = "SECURITY_GROUPS"
	//ContainerName a  param of ecs
	ContainerName = "CONTAINER_NAME"
)
