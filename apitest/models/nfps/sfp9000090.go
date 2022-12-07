package models

type FP900090I struct {
	FileType string `json:"fileType"`
	FileName string `json:"fileName"`
}
type FP900090O struct {
	FileTrans []FileTrans `json:"file_trans"`
}

type FileTrans struct {
	InstructionId    string `json:"instruction_id"`
	EndToEndId       string `json:"end_to_end_id"`
	CreditorMemberId string `json:"creditor_member_id"`
	DebtorMemberId   string `json:"debtor_member_id"`
}

func (*FP900090I) GetServiceKey() string {
	return "FP900090"
}
