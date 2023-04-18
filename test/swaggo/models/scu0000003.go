package models

import "github.com/shopspring/decimal"

type SCU0000003I struct {
	IdType     string `json:"idType" description:"id Type:101-Citizen ID 104-Passport"`
	IdNation   string `json:"idNation" description:"Id nation"`
	IdNum      string `json:"idNum" description:"Id number"`
	CustomerId string `json:"customerId" description:"Customer id"`
	SystemId   string `json:"systemId"`
}

type SCU0000003O struct {
	Birthday            string             `json:"birthday" description:"Birthday"`
	CreateDate          string             `json:"createDate" description:"Create date"`
	CustomerId          string             `json:"customerId" description:"Customer id"`
	CustomerName        string             `json:"customerName" description:"Customer name"`
	CustomerStatus      string             `json:"customerStatus" description:"Customer status"`
	FirstName           string             `json:"firstName" description:"First name"`
	Gender              string             `json:"gender" description:"gender:0-other 1-MALE 2-FEMLE"`
	LastName            string             `json:"lastName" description:"Last name"`
	Marriage            string             `json:"marriage" description:"marriage status:Single Married Divorced Widowed"`
	NaturalStatus       string             `json:"naturalStatus" description:"Natural status"`
	ResidentFlag        string             `json:"residentFlag" description:"Resident flag"`
	SystemId            string             `json:"systemId" description:"System id"`
	Title               string             `json:"title" description:"title:Mr. Mrs. Miss"`
	Age                 int                `json:"age" description:"Age"`
	DefaultFlag         string             `json:"defaultFlag" description:"Default flag"`
	ExpireDate          string             `json:"expireDate" description:"Expire date"`
	IdNum               string             `json:"idNum" description:"Id number"`
	IdStatus            string             `json:"idStatus" description:"Id status"`
	IdType              string             `json:"idType" description:"id Type:101-Citizen ID 104-Passport"`
	IssueDate           string             `json:"issueDate" description:"Issue date"`
	IdNation            string             `json:"idNation" description:"Id nation"`
	OpenCertificateFlag string             `json:"openCertificateFlag" description:"Open certificate flag"`
	CertificateRecordId int                `json:"certificateRecordId" description:"Certificate record id"`
	UseType             string             `json:"useType" description:"Use type"`
	ListAddress         []CustomerAddrInfo `json:"listAddress" description:"List address"`
	FileList            []CustomerFileInfo `json:"fileList" description:"File list"`
	ListContact         []CustomerConnInfo `json:"listContact" description:"List contact"`
	EducationList       []EducationInfo03  `json:"educationList" description:"Education list"`
	OccupationList      []OccupationInfo03 `json:"occupationList" description:"Occupation list"`
	SpouseInfo          *SpouseInfo        `json:"spouseInfo" description:"Spouse information"`
	EmergencyInfo       *EmergencyInfo     `json:"emergencyInfo" description:"Emergency information"`
	GradeInfo           *GradeInfo         `json:"gradeInfo" description:"Grade information"`
	LabelIds            string             `json:"labelIds" description:"label ids"`
	NewAddress          string             `json:"newAddress" description:"1:New Address 0:Same as Registration Address"`
}

type CustomerAddrInfo struct {
	AddressId           string `json:"addressId" description:"Address id"`
	AddressType         string `json:"addressType" description:"Address type"`
	CustomerId          string `json:"customerId" description:"Customer id"`
	Address             string `json:"address" description:"Address"`
	Address1            string `json:"address1" description:"Address1"`
	Address2            string `json:"address2" description:"Address2"`
	Address3            string `json:"address3" description:"Address3"`
	HouseResidentNumber string `json:"houseResidentNumber" description:"House resident number"`
	DefaultFlag         string `json:"defaultFlag" description:"Default flag"`
	AddressNation       string `json:"addressNation" description:"Address nation"`
	PostalCode          string `json:"postalCode" description:"Postal code"`
	EffectiveStatus     string `json:"effectiveStatus" description:"Effective status"`
	UseType             string `json:"useType" description:"Use type"`
}

type CustomerFileInfo struct {
	RecordId int64  `json:"recordId" description:"Record id"`
	FileName string `json:"fileName" description:"File name"`
	FileId   string `json:"fileId" description:"File id"`
	Status   string `json:"status" description:"Status"`
}

type CustomerConnInfo struct {
	ContactId       string `json:"contactId" description:"Contact id"`
	ContactType     string `json:"contactType" description:"Contact type"`
	ContactNum      string `json:"contactNum" description:"Contact number"`
	CustomerId      string `json:"customerId" description:"Customer id"`
	DefaultFlag     string `json:"defaultFlag" description:"Default flag"`
	EffectiveStatus string `json:"effectiveStatus" description:"Effective status"`
	Title           string `json:"title" description:"Title"`
	UseType         string `json:"useType" description:"Use type"`
}

type EducationInfo03 struct {
	EducationId         int64  `json:"educationId" description:"Education id"`
	GraduateYear        string `json:"graduateYear" description:"Graduate year"`
	Major               string `json:"major" description:"Major"`
	MaxDegree           string `json:"maxDegree" description:"Max degree"`
	SchoolName          string `json:"schoolName" description:"School name"`
	SchoolNationCountry string `json:"schoolNationCountry" description:"School nation country"`
}

type OccupationInfo03 struct {
	OccupationId      int64           `json:"occupationId" description:"Occupation id"`
	EmployerAddress   string          `json:"employerAddress" description:"Employer address"`
	EmployerName      string          `json:"employerName" description:"Employer name"`
	EmployerNation    string          `json:"employerNation" description:"Employer nation"`
	EmployerTelephone string          `json:"employerTelephone" description:"Employer telephone"`
	MonthIncome       decimal.Decimal `json:"monthIncome" description:"Month income"`
	IncomeCurrency    string          `json:"incomeCurrency" description:"Income currency"`
	JobType           string          `json:"jobType" description:"Job type"`
	Remark            string          `json:"remark" description:"Remark"`
}

func (*SCU0000003I) GetServiceKey() string {
	return "CU000003"
}
