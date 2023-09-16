package utilities

type CVEInformation struct {
	Id           string      `json:"id"`
	Description  string      `json:"description"`
	CVSS3xScore  CVSS3xScore `json:"cVSS3xScore"`
	CVSS20Score  CVSS20Score `json:"cVSS20Score"`
	CVEID        string      `json:"cveID"`
	CreationDate string      `json:"creationDate"`
}

type CVSS3xScore struct {
	Id                 string                   `json:"id"`
	BaseScore          CVSS3xBaseScore          `json:"baseScore"`
	TemporalScore      TemporalScore            `json:"temporalScore"`
	Vector             string                   `json:"vector"`
	ImpactScore        string                   `json:"impactScore"`
	Source             string                   `json:"source"`
	EnvironmentalScore CVSS3xEnvironmentalScore `json:"environmentalScore"`
}

type CVSS20Score struct {
	Id                 string                   `json:"id"`
	Vector             string                   `json:"vector"`
	ImpactScore        string                   `json:"impactScore"`
	Source             string                   `json:"source"`
	EnvironmentalScore CVSS20EnvironmentalScore `json:"EnvironmentalScore"`
	TemporalScore      TemporalScore            `json:"TemporalScore"`
	CVSS20BaseScore    CVSS20BaseScore          `json:"baseScore"`
}

type TemporalScore struct {
	Id                  string `json:"id"`
	Score               string `json:"Score"`
	Severity            string `json:"severity"`
	ExploitCodeMaturity string `json:"exploitCodeMaturity"`
	RemeditationLevel   string `json:"remeditationLevel"`
	ReportConfidence    string `json:"reportConfidence"`
}

type CVSS20EnvironmentalScore struct {
	Id                            string `json:"id"`
	Score                         string `json:"score"`
	Severity                      string `json:"severity"`
	CollateralDamagePotentialType string `json:"collateralDamagePotentialType"`
	TargetDistributionType        string `json:"targetDistributionType"`
	ConfidentialityRequirement    string `json:"confidentialityRequirement"`
	AvailabilityRequirement       string `json:"availabilityRequirement"`
	IntegrityRequirement          string `json:"integrityRequirement"`
	PrivilegesRequired            string `json:"privilegesRequired"`
	UserInteraction               string `json:"userInteraction"`
	Scope                         string `json:"scope"`
	Confidentiality               string `json:"confidentiality"`
	Integrity                     string `json:"integrity"`
	Availability                  string `json:"availability"`
}

type CVSS3xBaseScore struct {
	Id                 string `json:"id"`
	Score              string `json:"score"`
	Severity           string `json:"severity"`
	AttackVector       string `json:"attackVector"`
	AttackComplexity   string `json:"attackComplexity"`
	PrivilegesRequired string `json:"privilegesRequired"`
	UserInteraction    string `json:"userInteraction"`
	Scope              string `json:"scope"`
	Confidentiality    string `json:"confidentiality"`
	Integrity          string `json:"integrity"`
	Availability       string `json:"availability"`
}
type CVSS20BaseScore struct {
	Id               string `json:"id"`
	Score            string `json:"score"`
	Severity         string `json:"severity"`
	AttackVector     string `json:"attackVector"`
	Authentication   string `json:"authentication"`
	AccessComplexity string `json:"accessComplexity"`
	Confidentiality  string `json:"confidentiality"`
	Integrity        string `json:"integrity"`
	Availability     string `json:"availability"`
}
type CVSS3xEnvironmentalScore struct {
	Id                         string `json:"id"`
	Score                      string `json:"score"`
	Severity                   string `json:"severity"`
	AttackVector               string `json:"attackVector"`
	AttackComplexity           string `json:"attackComplexity"`
	PrivilegesRequired         string `json:"privilegesRequired"`
	UserInteraction            string `json:"userInteraction"`
	Scope                      string `json:"scope"`
	Confidentiality            string `json:"confidentiality"`
	Integrity                  string `json:"integrity"`
	Availability               string `json:"availability"`
	ConfidentialityRequirement string `json:"confidentialityRequirement"`
	AvailabilityRequirement    string `json:"availabilityRequirement"`
	IntegrityRequirement       string `json:"integrityRequirement"`
}

type CommonAppInformation struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Identifier  string `json:"identifier"`
	Description string `json:"description"`
	Version     string `json:"version"`
	BuildNumber string `json:"build_number"`
}

type CommonHealthCheckResponse struct {
	ServiceName string `json:"serviceName"`
	Status      string `json:"status"`
}

type APIDoc struct {
	Link string `json:"link"`
}
