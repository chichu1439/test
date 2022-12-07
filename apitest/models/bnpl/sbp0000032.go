//Version: v1
package models

type BP000032Request struct {
	CustomerId string `json:"customerId" description:"Customer ID" validate:"required,max=20"`
}

type BP000032Response struct {
	Birthday            string `json:"birthday" description:"Birthday - date"`
	CreateDate          string `json:"createDate" description:"Create Date"`
	CustomerId          string `json:"customerId" description:"Customer ID"`
	CustomerName        string `json:"customerName" description:"Customer Name"`
	CustomerStatus      string `json:"customerStatus" description:"Customer Status"`
	FirstName           string `json:"firstName" description:"First Name"`
	Gender              string `json:"gender" description:"Gender"`
	LastName            string `json:"lastName" description:"Last Name"`
	Marriage            string `json:"marriage" description:"Marriage"`
	NaturalStatus       string `json:"naturalStatus" description:"Natural Status"`
	ResidentFlag        string `json:"residentFlag" description:"Resident Flag"`
	SystemId            string `json:"systemId" description:"System ID"`
	Title               string `json:"title" description:"Title"`
	DefaultFlag         string `json:"defaultFlag" description:"Default Flag"`
	ExpireDate          string `json:"expireDate" description:"Expire Date"`
	IdNum               string `json:"idNum" description:"ID Number"`
	IdStatus            string `json:"idStatus" description:"ID Status"`
	IdType              string `json:"idType" description:"ID Type"`
	IssueDate           string `json:"issueDate" description:"Issue Date"`
	IdNation            string `json:"idNation" description:"ID Nation"`
	OpenCertificateFlag string `json:"openCertificateFlag" description:"Open Certificate Flag"`
	CertificateRecordId int    `json:"certificateRecordId" description:"Certificate Record ID"`
	UseType             string `json:"useType" description:"Use Type"`
}

func (*BP000032Request) GetServiceKey() string {
	return "BP000032"
}
