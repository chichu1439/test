//Version: v1
package models

type PD000111Request struct {
	ProductIdList []string `json:"productIdList" description:"Product id list"`
	NotifyType    string   `json:"notifyType" validate:"required" description:"notification type"` //01-Loan Account reminder 02-Loan Account Late Notification 03-Account Statement 04-Loan Maturity Notice 05-Repayment Bill
}

type PD000111Response struct {
	ProdNotificationList []ProdNotification `json:"prodNotificationList" description:"notification list"`
}

type ProdNotification struct {
	Uid           int    `json:"uid"`
	ProductId     string `json:"productId"`
	NotifyType    string `json:"notifyType"`
	StrategyId    int    `json:"strategyId"`
	EffectiveDate string `json:"effectiveDate"`
	ExpireDate    string `json:"expireDate"`
	Status        string `json:"status" description:"0-inactive 1-active 2-expired 3-deleted 4-updated"`
}
