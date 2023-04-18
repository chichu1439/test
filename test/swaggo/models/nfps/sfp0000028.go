package models

//

type FP000028I struct {
	OrgId      string `json:"orgId" description:"Org id"`
	SoftwareId string `json:"softwareId" description:"Software id"`
}

type FP000028O struct {
	Certs      []CertKeys `json:"certs" description:"Certs"`
	OrgId      string     `json:"orgId" description:"Org id"`
	SoftwareId string     `json:"softwareId" description:"Software id"`
}

func (*FP000028I) GetServiceKey() string {
	return "FP000028"
}
