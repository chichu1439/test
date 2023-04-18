package models

import "github.com/shopspring/decimal"

type SCU0000002I struct {
	NeedApprove    string           `json:"needApprove"  description:"Need Approve flag Y-Need approve N-No need"`
	CustomerName   string           `json:"customerName" description:"Customer name" validate:"max=320"`
	IdNation       string           `json:"idNation" description:"Id nation" validate:"max=2"`
	IdType         string           `json:"idType" validate:"required,max=4" description:"Id Type 101-Citizen ID 104-Passport"`
	IdNum          string           `json:"idNum" description:"Id number" validate:"required,max=20"`
	FileList       []FileInfo       `json:"fileList" description:"File list"`
	LastName       string           `json:"lastName" description:"Last name" validate:"required,max=160"`
	FirstName      string           `json:"firstName" description:"First name" validate:"required,max=160"`
	Title          string           `json:"title" description:"Title Mr. Mrs. Miss"`
	IssueDate      string           `json:"issueDate" description:"Issue date" validate:"required,max=10"`
	ExpireDate     string           `json:"expireDate" description:"Expire date" validate:"required,max=10"`
	Gender         string           `json:"gender" validate:"required,max=1" description:"gender,0-other 1-MALE 2-FEMLE"` //0-other 1-MALE 2-FEMLE
	Birthday       string           `json:"birthday" description:"Birthday" validate:"required,max=10"`
	Marriage       string           `json:"marriage"  description:"marriage flag:Single Married Divorced Widowed"`
	SystemId       string           `json:"systemId" description:"System id" `
	ResidentFlag   string           `json:"residentFlag" description:"Resident flag"`
	NaturalStatus  string           `json:"naturalStatus" description:"Natural status"`
	EducationList  []EducationInfo  `json:"educationList" description:"Education list"`
	OccupationList []OccupationInfo `json:"occupationList" description:"Occupation list"`
	ListContact    []ListContact02  `json:"listContact" description:"contact List "`
	ListAddress    []ListAddress02  `json:"listAddress" description:"address List "`
	ChangeComment  string           `json:"changeComment" description:"Change comment"`
	MakerComment   string           `json:"makerComment" description:"Maker comment"`
	SpouseInfo     SpouseInfo       `json:"spouseInfo" description:"Spouse information"`
	EmergencyInfo  EmergencyInfo    `json:"emergencyInfo" description:"Emergency information"`
	GradeInfo      GradeInfo        `json:"gradeInfo" description:"Grade information"`
	LabelIds       string           `json:"labelIds" description:"label ids"`
}

type FileInfo struct {
	FileName string `json:"fileName" description:"File name"`
	FileId   string `json:"fileId" description:"File id"`
}

type EducationInfo struct {
	GraduateYear        string `json:"graduateYear" description:"Graduate year"`
	Major               string `json:"major" description:"Major"`
	MaxDegree           string `json:"maxDegree" description:"Max degree"`
	SchoolName          string `json:"schoolName" description:"School name"`
	SchoolNationCountry string `json:"schoolNationCountry" description:"School nation country"`
}

type OccupationInfo struct {
	EmployerAddress   string          `json:"employerAddress" description:"Employer address"`
	EmployerName      string          `json:"employerName" description:"Employer name"`
	EmployerNation    string          `json:"employerNation" description:"Employer nation"`
	EmployerTelephone string          `json:"employerTelephone" description:"Employer telephone"`
	MonthIncome       decimal.Decimal `json:"monthIncome" description:"Month income"`
	IncomeCurrency    string          `json:"incomeCurrency" description:"Income currency"`
	JobType           string          `json:"jobType" description:"Job type"`
	Remark            string          `json:"remark" description:"Remark"`
}

type EmergencyInfo struct {
	ContactName           string `json:"contactName" description:"Contact name"`
	ContactRelation       string `json:"contactRelation" description:"Contact relation"`
	ContactTelephoneNum   string `json:"contactTelephoneNum" description:"Contact telephone number"`
	ContactAddressCountry string `json:"contactAddressCountry" description:"Contact address country"`
	ContactAddress1       string `json:"contactAddress1" description:"Contact address1"`
	ContactAddress2       string `json:"contactAddress2" description:"Contact address2"`
	ContactAddress3       string `json:"contactAddress3" description:"Contact address3"`
	ContactPostCode       string `json:"contactPostCode" description:"Contact post code"`
}

type SpouseInfo struct {
	SpouseName            string          `json:"spouseName" description:"Spouse name"`
	SpouseFirstName       string          `json:"spouseFirstName" description:"Spouse first name"`
	SpouseLastName        string          `json:"spouseLastName" description:"Spouse last name"`
	SpouseIdNation        string          `json:"spouseIdNation" description:"Spouse id nation"`
	SpouseIdType          string          `json:"spouseIdType" description:"Spouse id type"`
	SpouseIdNum           string          `json:"spouseIdNum" description:"Spouse id number"`
	SpouseTelephoneNumber string          `json:"spouseTelephoneNumber" description:"Spouse telephone number"`
	SpouseEmployerName    string          `json:"spouseEmployerName" description:"Spouse employer name"`
	SpouseMonthlyIncome   decimal.Decimal `json:"spouseMonthlyIncome" description:"Spouse monthly income"`
	IncomeCurrency        string          `json:"incomeCurrency" description:"Income currency"`
	NumberOfChildren      int             `json:"numberOfChildren" description:"Number of children"`
}

type GradeInfo struct {
	Source string `json:"source" description:"Source come of the grade"`
	Grade  string `json:"grade" description:"grade"`
}
type SCU0000002O struct {
	CustomerId string `json:"customerId" description:"Customer id"`
}

func (*SCU0000002I) GetServiceKey() string {
	return "CU000002"
}
