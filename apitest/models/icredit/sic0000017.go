package models
//
//import cumodels "git.sirius.io/banking/common/models/customer"
//
//type IC000017I struct {
//	CustomerId        string                `json:"customerId" description:"Customer id" validate:"required"`
//	ChangeComment     string                `json:"changeComment" description:"Change comment"`
//	ApproveStatus     string                `json:"approveStatus" description:"Approval status:1-to be approved 2-approved 3-rejected"` // 审批状态 1-待审批 2-审批通过 3-审批拒绝
//	ApproveComment    string                `json:"approveComment" description:"Approval description"`                                  // 审批说明
//	ChangeContext     *cumodels.SCU0000004I `json:"changeContext" description:"Change context"`
//	ChangeContextList []ChangeContextList   `json:"changeContextList" description:"Change context list"`
//}
//
//type ChangeContextList struct {
//	ChangeField       string `json:"changeField" description:"Change field"`
//	ChangeValueBefore string `json:"changeValueBefore" description:"Change value before"`
//	ChangeValueAfter  string `json:"changeValueAfter" description:"Change value after"`
//}
//
//type IC000017O struct {
//	CustomerId string `json:"customerId" description:"Customer id"`
//}
//
//func (*IC000017I) GetServiceKey() string {
//	return "ic000017"
//}
