//Version: v1
package models

type PD000125Request struct {
	ProductId string `json:"productId"`
}

type PD000125Response struct {
	RiskId                string `json:"riskId"`
	RiskName              string `json:"riskName"`
	ProductId             string `json:"productId"`
	Status                int    `json:"status"`
	AllowedOfflineDays    int    `json:"allowedOfflineDays"`
	MaxOfflinePayNum      int    `json:"maxOfflinePayNum"`
	NewPaymentAmount      int    `json:"newPaymentAmount"`
	PendingPaymentsAmount int    `json:"pendingPaymentsAmount"`
}

func (*PD000125Request) GetServiceKey() string {
	return "PD000125"
}
