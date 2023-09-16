package utilities

const (
	APIBase         = "/api"
	APIVersion      = "/v1"
	CVEInfo         = "/cveinfo"
	HealthInfo      = "/health"
	Info            = "/info"
	CVEEndpoint     = APIBase + APIVersion + CVEInfo
	CVEEndpointById = CVEEndpoint + "/cveId"
	HealthEndpoint  = APIBase + APIVersion + HealthInfo
	InfoEndpoint    = APIBase + APIVersion + Info
	ApiDoc          = "/api-docs"
	ApiDocEndpoint  = APIBase + APIVersion + ApiDoc
	DocURL          = "https://stoplight.dell.com/docs/cveinformation-1/zhn1e72hisfsn-cve-info"
)

var CveData []CVEInformation
