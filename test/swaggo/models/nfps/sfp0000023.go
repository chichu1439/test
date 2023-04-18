package models

// upload cert and key

type FP000023I struct {
	CertKeys []CertKeys `json:"certKeys" description:"Cert keys"`
	Org      Org        `json:"org" description:"Org"`
	Software Software   `json:"software" description:"Software"`
}

type CertKeys struct {
	Certificate string `json:"certificate" description:"Certificate"`
	PriKey      string `json:"priKey" description:"Pri key"`
	Usage       string `json:"usage" description:"Usage"`
}

type Org struct {
	OrgId   string `json:"orgId" description:"Org id"`
	OrgName string `json:"orgName" description:"Org name"`
}

type Software struct {
	SoftwareClientName string `json:"softwareClientName" description:"Software client name"`
	SoftwareId         string `json:"softwareId" description:"Software id"`
}

type FP000023O struct {
}

func (*FP000023I) GetServiceKey() string {
	return "FP000023"
}
