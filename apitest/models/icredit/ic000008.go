package models
//
//import cumodels "git.sirius.io/banking/common/models/customer"
//
//type IC000008I struct {
//	CustomerId string `json:"customerId" description:"Customer id" validate:"required"` // 客户号
//}
//
//type IC000008O struct {
//	Data IC000008OData `json:"data" description:"Data"`
//}
//
//type IC000008OData struct {
//	CustomerDetail *cumodels.SCU0000003O `json:"customerDetail" description:"Customer Detail"`
//	ApproveInfo    *ApproveInfo          `json:"approveInfo" description:"Approve information"`
//}
//
//type ApproveInfo struct {
//	ApproveStatus     string                `json:"approveStatus" description:"Approve Status"` //01-approved 04-refused
//	ChangeUser        string                `json:"changeUser" description:"Change User"`
//	ApplyType         string                `json:"applyType" description:"Application type 1-add 2-modify"` //申请类型 1-新增  2-修改
//	ChangeOrgzNum     string                `json:"changeOrgzNum" description:"Change orgz number"`
//	ChangeComment     string                `json:"changeComment" description:"Change comment"`
//	ApproveUser       string                `json:"approveUser" description:"Approve user"`
//	ApproveOrgzNum    string                `json:"approveOrgzNum" description:"Approve orgz number"`
//	ApproveTime       string                `json:"approveTime" description:"Approve time"`
//	ApproveComment    string                `json:"approveComment" description:"Approve comment"`
//	ChangeContext     *cumodels.SCU0000004I `json:"changeContext" description:"Change context"`
//	ChangeContextList []ChangeContextList   `json:"changeContextList" description:"Change content"` // 变更内容
//}
//
//func (*IC000008I) GetServiceKey() string {
//	return "ic000008"
//}
