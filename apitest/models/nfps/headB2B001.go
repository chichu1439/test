package models

type HeadB2B001 struct {
	FromMemberId   string `xml:"Fr>FIId>FinInstnId>ClrSysMmbId>MmbId,omitempty" validate:"required,max=35"`
	FromBranchId   string `xml:"Fr>BrnchId>Id,omitempty" validate:"max=35"` //itmx add
	ToMemberId     string `xml:"To>FIId>FinInstnId>ClrSysMmbId>MmbId,omitempty" validate:"required,max=35"`
	ToBranchId     string `xml:"To>BrnchId>Id,omitempty" validate:"max=35"` //itmx add
	BizMsgId       string `xml:"BizMsgIdr,omitempty" validate:"required,max=35"`
	MsgDefId       string `xml:"MsgDefIdr,omitempty" validate:"required,max=35"`
	BizSvc         string `xml:"BizSvc,omitempty" validate:"required,max=35"`
	CreateDatetime string `xml:"CreDt,omitempty" validate:"required,max=35"`
	CopyDup        string `xml:"CpyDplct,omitempty" validate:"max=35"` //itmx add
}
