//Version: v1.0.0.0
package models

type SV000036I struct {
	AgreementId    string `json:"agreementId" description:"Agreement Id" validate:"required"`
	AgreementType  string `json:"agreementType" description:"Agreement Type Id" validate:"required"`
	BusinessNumber string `json:"businessNumber" description:"Business Number" validate:"required"`
	DocNumber      string `json:"docNumber" description:"Doc Number"`
}

type SV000036O struct {
	DeductionBusinessNumber  string  `json:"deductionBusinessNumber" description:"Deduction Business Number"`
	DeductionDate            string  `json:"deductionDate" description:"Deduction Date"`
	DeductionTime            string  `json:"deductionTime" description:"Deduction Time"`
	DeductionMode            string  `json:"deductionMode" description:"Deduction Mode"`
	DeductionAgreement       string  `json:"deductionAgreement" description:"Deduction Agreement"`
	DeductionAgreementType   string  `json:"deductionAgreementType" description:"Deduction Agreement Type"`
	DeductionAmount          float64 `json:"deductionAmount" description:"Deduction Amount"`
	DeductionDocNumber       string  `json:"deductionDocNumber" description:"Deduction Doc Number"`
	DeductionReason          string  `json:"deductionReason" description:"Deduction Reason"`
	FrFreezeAmount           string  `json:"frFreezeAmount" description:"Fr Freeze Amount"`
	FreezeBusinessNumber     string  `json:"freezeBusinessNumber" description:"Freeze Business Number"`
	DeductionOrganize        string  `json:"deductionOrganize" description:"Deduction Organize"`
	DeductionTeller          string  `json:"deductionTeller" description:"Deduction Teller"`
	DeductionAuthorityTeller string  `json:"deductionAuthorityTeller" description:"Deduction Authority Teller"`
	DateEnforcerName1        string  `json:"dateEnforcerName1" description:"Date Enforcer Name1"`
	DateEnforcerName2        string  `json:"dateEnforcerName2" description:"Date Enforcer Name2"`
	DateEnforcerId1          string  `json:"dateEnforcerId1" description:"Date Enforcer Id1"`
	DateEnforcerId2          string  `json:"dateEnforcerId2" description:"Date Enforcer Id2"`
	ReceivAccountNumber      string  `json:"receivAccountNumber" description:"Receiving Account Number"`
	ReceivBank               string  `json:"receivBank" description:"Receiving Bank"`
}

func (*SV000036I) GetServiceKey() string {
	return "SV000036"
}
