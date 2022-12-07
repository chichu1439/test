//Version: v1.0.0.0
package models

type AT000001I struct {
	TransactionDate   string          `json:"transactionDate" description:"Transaction date" validate:"required"`
	ActDate           string          `json:"actDate" description:"Act date" validate:"required"`
	GlobalBizSeqNo    string          `json:"globalBizSeqNo" description:"Global biz seq no" validate:"required"`
	SrcBizSeqNo       string          `json:"srcBizSeqNo" description:"Src biz seq no" validate:"required"`
	ExtrSysDate       string          `json:"extrSysDate" description:"Extr sys date"`
	ExtrSysSeq        string          `json:"extrSysSeq" description:"Extr sys seq"`
	TranBranch        string          `json:"tranBranch" description:"Tran branch"`
	TransactionEm     string          `json:"transactionEm" description:"Transaction em"`
	ReviewEm          string          `json:"reviewEm" description:"Review em"`
	TranIniChn        string          `json:"tranIniChn" description:"Tran ini chn" validate:"required"`
	BusSys            string          `json:"busSys" description:"Bus sys" validate:"required"`
	ServiceId         string          `json:"serviceId" description:"Service id" validate:"required"`
	TranType          string          `json:"tranType" description:"Tran type"`
	OrgTranDate       string          `json:"orgTranDate" description:"Org tran date"`
	OrgGlobalBizSeqNo string          `json:"orgGlobalBizSeqNo" description:"Org global biz seq no"`
	Item              []AT000001IItem `json:"item" description:"Item"`
}

type AT000001IItem struct {
	TransferFlag   string  `json:"transferFlag" description:"Transfer flag"`
	PrcType        string  `json:"prcType" description:"Prc type"`
	PrcTypeSq      int     `json:"prcTypeSq" description:"Prc type sq"`
	AccDistinction string  `json:"accDistinction" description:"Acc distinction"`
	AntBranch      string  `json:"antBranch" description:"Ant branch"`
	Account        string  `json:"account" description:"Account"`
	AgreementId    string  `json:"agreementId" description:"Agreement id"`
	PrductId       string  `json:"prductId" description:"Prduct id"`
	PrductNm       int     `json:"prductNm" description:"Prduct nm"`
	CusDif1        string  `json:"cusDif1" description:"Cus dif1"`
	CusDif2        string  `json:"cusDif2" description:"Cus dif2"`
	CusDif3        string  `json:"cusDif3" description:"Cus dif3"`
	AmtType        string  `json:"amtType" description:"Amt type"`
	TermUnit       string  `json:"termUnit" description:"Term unit"`
	Term           int     `json:"term" description:"Term"`
	CusDft1        string  `json:"cusDft1" description:"Cus dft1"`
	CusDft2        string  `json:"cusDft2" description:"Cus dft2"`
	CusDft3        string  `json:"cusDft3" description:"Cus dft3"`
	CusDft4        string  `json:"cusDft4" description:"Cus dft4"`
	EvnFlag        string  `json:"evnFlag" description:"Evn flag"`
	EvnSub         string  `json:"evnSub" description:"Evn sub"`
	EvnDft1        string  `json:"evnDft1" description:"Evn dft1"`
	EvnDft2        string  `json:"evnDft2" description:"Evn dft2"`
	AccTitle       string  `json:"accTitle" description:"Acc title"`
	AccTitleSub    string  `json:"accTitleSub" description:"Acc title sub"`
	Currency       string  `json:"currency" description:"Currency"`
	ExgRate        float64 `json:"exgRate" description:"Exg rate"`
	TranAmount     float64 `json:"tranAmount" description:"Tran amount"`
	ExgAmount      float64 `json:"exgAmount" description:"Exg amount"`
	PairNm         int     `json:"pairNm" description:"Pair nm"`
}

type AT000001O struct {
}

func (*AT000001I) GetServiceKey() string {
	return "AT000001"
}
