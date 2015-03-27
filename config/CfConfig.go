package config

type CfConfig struct {
	Target              string
	ConfigVersion       int
	ApiVersion          string
	LoggregatorEndPoint string
	AccessToken         string
	OrganizationFields  OrganizationFields
	SpaceFields         SpaceFields
	SSLDisabled         bool
}

type OrganizationFields struct {
	Guid string
	Name string
}

type SpaceFields struct {
	Guid string
	Name string
}
