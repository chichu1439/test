package models

import "github.com/go-playground/validator/v10"

type SF011002I struct {
	Fileaction     string `json:"fileaction"`
	FileSize       int64  `json:"fileSize"`
	FileName       string `json:"fileName" validate:"required"`
	FileDate       int64  `json:"fileDate"`
	SendBackNumber int    `json:"send_back_number"`
}

type SF011002O struct {
	Data string `json:"data"`
}

/*
	OriginalMsgId:          t.request.XMLHead.BizMsgId,
	OriginalMsgNameId:      t.request.XMLHead.MsgDefId,
	OriginalCreateDatetime: t.request.Document.CreateDatetime, //todo
	OriginalNumOfTrans:     "1",
	OriginalInstructionId:  t.request.Document.OrgInstructionId,
	OriginalTransId:        t.request.Document.TransId,
	OriginalEndToEndId:     t.request.Document.EndToEndId,
	TransStatus:            status,
	// TransRejectReason:     code,
	TransRejectReason2:    code,
	InstructingAgent:      t.request.Document.InstructingMemberId,
	InstructedAgent:       t.request.Document.InstructedMemberId,
	ProcessCode:           t.request.Document.ProcessCode,
	MessageType:           t.request.Document.MessageType,
	MemberId:              t.request.Document.MemberId,
	MerchantType:          t.request.Document.MerchantType,
	TermType:              t.request.Document.TermType,
	RecvType:              t.request.Document.RecvType,
	SendType:              t.request.Document.SendType,
	LocalAmount:           t.request.Document.LocalAmount,
*/
type RestFieldAggr struct {
	ToMemberId             string   `json:"to_member_id"`
	OriginalMsgId          string   `json:"original_msg_id"`
	OriginalMsgNameId      string   `json:"original_msg_name_id"`
	OriginalCreateDatetime string   `json:"original_create_datetime"`
	OriginalNumOfTrans     string   `json:"original_num_of_trans"`
	OriginalInstructionId  string   `json:"original_instruction_id"`
	OriginalTransId        string   `json:"original_trans_id"`
	OriginalEndToEndId     string   `json:"original_end_to_end_id"`
	TransStatus            string   `json:"trans_status"`
	TransRejectReason2     string   `json:"trans_reject_reason_2"`
	InstructingAgent       string   `json:"instructing_agent"`
	InstructedAgent        string   `json:"instructed_agent"`
	ProcessCode            string   `json:"process_code"`
	MessageType            string   `json:"message_type"`
	MemberId               string   `json:"member_id"`
	MerchantType           string   `json:"merchant_type"`
	TermType               string   `json:"term_type"`
	RecvType               string   `json:"recv_type"`
	SendType               string   `json:"send_type"`
	LocalAmount            AmtField `json:"local_amount"`
}

func (*SF011002I) GetServiceKey() string {
	return "FP011002"
}

func (d *SF011002I) Validate() error {
	validate := validator.New()
	return validate.Struct(d)
}
