package infrastructure

import  (
	"dmcm/geography"
)

type ServerBasic struct {
	AgentVersion int
	Architecture string
	Budget int
	Cloud geography.CloudBasic
	Customer struct { CustomerId int }
	DataCenter struct { DataCenterId int }
	Description string
	MachineImage struct { MachineImageId int }
	Name string
	OwningGroups []struct { GroupId int }
	OwningUser struct { UserId int }
	Platform string
	PrivateIpAddresses []string
	Product struct { ProductId int }
	ProviderId string
	PublicIpAddress string
	Region struct { RegionId int }
	ServerId int
	StartDate string
	Status string
}

type Servers struct {
	Servers []ServerBasic
}
