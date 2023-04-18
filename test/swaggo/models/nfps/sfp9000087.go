package models

type FP900087I struct {
	TranFileShip []TranFileShip `json:"tranFileShip"`
}

type TranFileShip struct {
	InstructionId    string `json:"instruction_id"`
	EndToEndId       string `json:"end_to_end_id"`
	CreditorMemberId string `json:"creditor_member_id"` // 贷方 放款人
	DebtorMemberId   string `json:"debtor_member_id"`   // 借方 借款人
	SendingSendName  string `json:"sending_send_name"`
	SendingRcvName   string `json:"sending_rcv_name"`
	RcvRcvName       string `json:"rcv_rcv_name"`
	RcvSendName      string `json:"rcv_send_name"`
}

type FP900087O struct {
	Status string `json:"status"`
}

func (*FP900087I) GetServiceKey() string {
	return "FP900087"
}
