//Version: v1.0.0.0
package models

type AT000002I struct {
	TransactionDate string `json:"transactionDate" description:"Transaction date" validate:"required"`
	GlobalBizSeqNo  string `json:"globalBizSeqNo" description:"Global biz seq no" validate:"required"`
	SrcBizSeqNo     string `json:"srcBizSeqNo" description:"Src biz seq no" validate:"required"`
	ExtrSysDate     string `json:"extrSysDate" description:"Extr sys date"`
	ExtrSysSeq      string `json:"extrSysSeq" description:"Extr sys seq"`
	TranIniChn      string `json:"tranIniChn" description:"Tran ini chn" validate:"required"`
	BusSys          string `json:"busSys" description:"Bus sys" validate:"required"`
	ServiceId       string `json:"serviceId" description:"Service id" validate:"required"`
}

type AT000002O struct {
}

func (*AT000002I) GetServiceKey() string {
	return "AT000002"
}
