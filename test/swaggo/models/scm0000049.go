package models

type CM000049I struct {
	TaskId string `json:"taskId" description:"Task id" validate:"required"`
}

type CM000049O struct {
	ContactPersonType string  `json:"contactPersonType" description:"Contact person type"`
	ContactPersonName string  `json:"contactPersonName" description:"Contact person name"`
	ContactTime       string  `json:"contactTime" description:"Contact time"`
	ContactMethod     string  `json:"contactMethod" description:"Contact method"`
	ContactStatus     string  `json:"contactStatus" description:"Contact status"`
	PromisePayDate    string  `json:"promisePayDate" description:"Promise pay date"`
	PromisePayAmount  float64 `json:"promisePayAmount" description:"Promise pay amount"`
	PromiseInfo       string  `json:"promiseInfo" description:"Promise information"`
	FamilyReason      string  `json:"familyReason" description:"Family reason"`
	BusinessReason    string  `json:"businessReason" description:"Business reason"`
	JobReason         string  `json:"jobReason" description:"Job reason"`
	Remark            string  `json:"remark" description:"Remark"`
	CurrentAddress    string  `json:"currentAddress" description:"Current address"`
	OrgzNum           string  `json:"orgzNum" description:"Orgz number"`
	UserId            string  `json:"userId" description:"User id"`
}

func (*CM000049I) GetServiceKey() string {
	return "CM000049"
}
