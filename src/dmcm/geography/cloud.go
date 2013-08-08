package geography

type CloudBasic struct {
	Name                      string
	ComputeDelegate           string
	WebUrl                    string
	CloudId                   int
	Status                    string
	ComputeEndpoint           string
	PrivateCloud              bool
	LogoUrl                   string
	ComputeSecretKeyLabel     string
	Provider                  string
	Endpoint                  string
	ComputeAccountNumberLabel string
	DocumentationLabel        string
	CloudProviderName         string
	DaseinComputeDelegate     string
	CloudProviderConsoleUrl   string
	CloudProviderLogoUrl      string
}

type AllClouds struct {
	Clouds []CloudBasic
}
