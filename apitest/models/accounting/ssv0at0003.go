//Version: v1.0.0.0
package models

type SV0AT003I struct {
	Item []SV0AT003IItem `json:"item" description:"Item"`
}

type SV0AT003IItem struct {
	TransationDate      string  `json:"transationDate" description:"Transation date"`
	TransationTime      string  `json:"transationTime" description:"Transation time"`
	ActDate             string  `json:"actDate" description:"Act date"`
	GlobalBizSeqNo      string  `json:"globalBizSeqNo" description:"Global biz seq no"`
	SrcBizSeqNo         string  `json:"srcBizSeqNo" description:"Src biz seq no"`
	Sequence            int     `json:"sequence" description:"Sequence"`
	BusSys              string  `json:"busSys" description:"Bus sys"`
	ExtrSysDate         string  `json:"extrSysDate" description:"Extr sys date"`
	ExtrSysSeq          string  `json:"extrSysSeq" description:"Extr sys seq"`
	RefSysDate          string  `json:"refSysDate" description:"Ref sys date"`
	RefSysSeq           string  `json:"refSysSeq" description:"Ref sys seq"`
	DcnId               string  `json:"dcnId" description:"Dcn id"`
	TranBranch          string  `json:"tranBranch" description:"Tran branch"`
	AntDept             string  `json:"antDept" description:"Ant dept"`
	ActDept             string  `json:"actDept" description:"Act dept"`
	Currency            string  `json:"currency" description:"Currency"`
	DebCrtFlag          string  `json:"debCrtFlag" description:"Deb crt flag"`
	TranAmount          float64 `json:"tranAmount" description:"Tran amount"`
	AccTitle            string  `json:"accTitle" description:"Acc title"`
	AccTitleSub         string  `json:"accTitleSub" description:"Acc title sub"`
	Account             string  `json:"account" description:"Account"`
	VchNm               string  `json:"vchNm" description:"Vch nm"`
	NedGenFlag          string  `json:"nedGenFlag" description:"Ned gen flag"`
	SndGrFlag           string  `json:"sndGrFlag" description:"Snd gr flag"`
	AccDistinction      string  `json:"accDistinction" description:"Acc distinction"`
	ReverseTradeFlag    string  `json:"reverseTradeFlag" description:"Reverse trade flag"`
	CashTranFlag        string  `json:"cashTranFlag" description:"Cash tran flag"`
	ReverseCnlFlag      string  `json:"reverseCnlFlag" description:"Reverse cnl flag"`
	ReverseDate         string  `json:"reverseDate" description:"Reverse date"`
	ReverseGlobalBiz    string  `json:"reverseGlobalBiz" description:"Reverse global biz"`
	CorrectionDate      string  `json:"correctionDate" description:"Correction date"`
	CorrectionGlobalBiz string  `json:"correctionGlobalBiz" description:"Correction global biz"`
	TransationEm        string  `json:"transationEm" description:"Transation em"`
	ReviewEm            string  `json:"reviewEm" description:"Review em"`
	SmmryCode           string  `json:"smmryCode" description:"Smmry code"`
	Summary             string  `json:"summary" description:"Summary"`
	Total               int     `json:"total" description:"Total"`
	LastUpDatetime      string  `json:"lastUpDatetime" description:"Last up datetime"`
	LastUpBn            string  `json:"lastUpBn" description:"Last up bn"`
	LastUpEm            string  `json:"lastUpEm" description:"Last up em"`
}

type SV0AT003O struct {
}

func (*SV0AT003I) GetServiceKey() string {
	return "ssv0at0003"
}
