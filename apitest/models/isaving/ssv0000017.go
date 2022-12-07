//Version: v1.0.0.0
package models

type SV000017I struct {
	FreezeBusinessNumber string `json:"freezeBusinessNumber" validate:"required,max=30" description:"Original frozen business number"` //原冻结业务编号
	FreezeMediumNumber   string `json:"freezeMediumNumber" validate:"required,max=30" description:"Media number"`                      //介质号码
	FreezeMediumType     string `json:"freezeMediumType" validate:"required,max=3" description:"Media type"`                           //介质类型
	CustomerId           string `json:"customerId" validate:"required,max=20" description:"Frozen customer number"`                    //冻结客户编号
	CustomerType         string `json:"customerType" validate:"required,max=2" description:"Frozen customer type"`                     //冻结客户类型
	ConDueDate           string `json:"conDueDate" validate:"required" description:"Refreeze expiration date"`                         //续冻到期日
	ConDocnm             string `json:"conDocnm" validate:"required,max=60" description:"Renew the document number"`                   //续冻文书号
	FreezeOrgnizeType    string `json:"freezeOrgnizeType" validate:"required,max=1,oneof=1" description:"Refreeze initiator"`          //续冻发起方
	FreezeOrgnizeName    string `json:"freezeOrgnizeName" validate:"required,max=60" description:"Name of refreezing initiator"`       //续冻发起方名称
	EnforcerName1        string `json:"enforcerName1" validate:"required,max=60" description:"Name of law enforcement officer 1"`      //执法人员名称1
	EnforcerId1          string `json:"enforcerId1" validate:"required,max=60" description:"Law enforcement officer ID number 1"`      //执法人员证件号码1
	EnforcerName2        string `json:"enforcerName2" validate:"max=60" description:"Name of law enforcement officer 2"`               //执法人员名称2
	EnforcerId2          string `json:"enforcerId2" validate:"max=60" description:"Law enforcement officer ID number 2"`               //执法人员证件号码2
	ConReason            string `json:"conReason" validate:"max=60" description:"Reasons for continued freezing"`                      //续冻原因
}

type SV000017O struct {
}

func (*SV000017I) GetServiceKey() string {
	return "ssv0000017"
}
