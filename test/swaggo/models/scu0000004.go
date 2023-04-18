package models

import "github.com/shopspring/decimal"

type SCU0000004I struct {
	CustomerId          string                  `json:"customerId" validate:"required,max=20" description:"Customer id"` //客户号 Y
	CustomerName        string                  `json:"customerName" description:"Customer name" validate:"max=320"`
	LastName            string                  `json:"lastName" description:"Last name" validate:"max=160"`
	FirstName           string                  `json:"firstName" description:"First name" validate:"max=160"`
	Title               string                  `json:"title" validate:"max=10" description:"Title:Mr. Mrs. Miss"`
	Gender              string                  `json:"gender" validate:"max=1" description:"Gender:0-other 1-MALE 2-FEMLE"` //0-other 1-MALE 2-FEMLE
	Birthday            string                  `json:"birthday" description:"Birthday" validate:"max=10"`
	Marriage            string                  `json:"marriage" validate:"max=10" description:"Marriage:Single Married Divorced Widowed"`
	CustomerStatus      string                  `json:"customerStatus" description:"Customer status"`
	NaturalStatus       string                  `json:"naturalStatus" description:"Natural status"`
	ResidentFlag        string                  `json:"residentFlag" description:"Resident flag"`
	SystemId            string                  `json:"systemId" description:"System id" `
	EducationList       []EducationInfo02       `json:"educationList" description:"Education list"  `
	OccupationList      []OccupationInfo02      `json:"occupationList" description:"Occupation list"  `
	ListContact         []ListContact03         `json:"listContact" description:"List contact"  `
	ListAddress         []ListAddress03         `json:"listAddress" description:"List address"  `
	IdNation            string                  `json:"idNation" description:"Id nation" validate:"max=2"`
	IdType              string                  `json:"idType" validate:"required,max=4" description:"Id Type:101-Citizen ID 104-Passport"`
	IdNum               string                  `json:"idNum" description:"Id number" validate:"required,max=20"`
	IssueDate           string                  `json:"issueDate" description:"Issue date" validate:"max=10"`
	ExpireDate          string                  `json:"expireDate" description:"Expire date" validate:"max=10"`
	IdStatus            string                  `json:"idStatus" description:"Id status"`
	UseType             string                  `json:"useType" description:"Use type"`
	CertificateRecordId int                     `json:"certificateRecordId" description:"Certificate record id"`
	FileList            []CustomerIdentAppend02 `json:"fileList" description:"File list" `
	SpouseInfo          *SpouseInfo             `json:"spouseInfo" description:"Spouse information"`
	EmergencyInfo       *EmergencyInfo          `json:"emergencyInfo" description:"Emergency information"`
	GradeInfo           *GradeInfo              `json:"gradeInfo" description:"Grade information"`
	LabelIds            string                  `json:"labelIds" description:"label ids"`
}

type CustomerIdentAppend02 struct {
	RecordId int64  `json:"recordId" description:"Record id"`
	FileName string `json:"fileName" description:"File name"`
	FileId   string `json:"fileId" description:"File id"`
	Status   string `json:"status" description:"Status"`
}

type EducationInfo02 struct {
	EducationId         int64  `json:"educationId" description:"Education id" validate:"required"`
	CustomerId          string `json:"customerId" description:"Customer id"`
	GraduateYear        string `json:"graduateYear" description:"Graduate year"`
	Major               string `json:"major" description:"Major"`
	MaxDegree           string `json:"maxDegree" description:"Max degree"`
	SchoolName          string `json:"schoolName" description:"School name"`
	SchoolNationCountry string `json:"schoolNationCountry" description:"School nation country"`
	FinalModifyDate     string `json:"finalModifyDate" description:"Final modify date"`
	FinalModifyOrgz     string `json:"finalModifyOrgz" description:"Final modify organization"`
	FinalModifyTime     string `json:"finalModifyTime" description:"Final modify time"`
	FinalModifyUser     string `json:"finalModifyUser" description:"Final modify user"`
}

type OccupationInfo02 struct {
	OccupationId      int64           `json:"occupationId" description:"Occupation id" validate:"required"`
	EmployerAddress   string          `json:"employerAddress" description:"Employer address"`
	EmployerName      string          `json:"employerName" description:"Employer name"`
	EmployerNation    string          `json:"employerNation" description:"Employer nation"`
	EmployerTelephone string          `json:"employerTelephone" description:"Employer telephone"`
	MonthIncome       decimal.Decimal `json:"monthIncome" description:"Month income"`
	IncomeCurrency    string          `json:"incomeCurrency" description:"Income currency"`
	JobType           string          `json:"jobType" description:"Job type"`
	Remark            string          `json:"remark" description:"Remark"`
}

type CertificateInfo struct {
	IdNation     string                  `json:"idNation" description:"Id nation" validate:"max=2"`
	IdType       string                  `json:"idType" description:"Id type" validate:"required,max=4"`
	IdNum        string                  `json:"idNum" description:"Id number" validate:"required,max=20"`
	IssueDate    string                  `json:"issueDate" description:"Issue date" validate:"max=10"`
	ExpireDate   string                  `json:"expireDate" description:"Expire date" validate:"max=10"`
	IdStatus     string                  `json:"idStatus" description:"Id status"`
	UseType      string                  `json:"useType" description:"Use type"`
	RecordId     int                     `json:"recordId" description:"Record id" validate:"required"`
	FileInfoList []CustomerIdentAppend02 `json:"fileInfoList" description:"File information list" `
}

type ListContact03 struct {
	ContactId       string `json:"contactId" description:"Contact id" validate:"required"`
	ContactType     string `json:"contactType" description:"Contact type"`
	ContactNum      string `json:"contactNum" description:"Contact number"`
	Title           string `json:"title" description:"Title"`
	EffectiveStatus string `json:"effectiveStatus" description:"Effective status"`
	UseType         string `json:"useType" description:"Use type"`
	DefaultFlag     string `json:"defaultFlag" description:"Default flag"`
}

type ListAddress03 struct {
	AddressId           string `json:"addressId" description:"Address id" validate:"required"`
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

type SCU0000004O struct {
	CustomerId string `json:"customerId" description:"Customer id"`
}

func (*SCU0000004I) GetServiceKey() string {
	return "CU000004"
}
