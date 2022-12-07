//Version: v1.0.0.0
package models

type SV000015I struct {
	FrTransactionDate     string `json:"frTransactionDate" validate:"required" description:"Original transaction date"`                     //原交易日期
	FrGlobalBusinessSeqno string `json:"frGlobalBusinessSeqno" validate:"required,max=40" description:"Original transaction serial number"` //原交易全局流水号
	FrFeeFlag             string `json:"frFeeFlag" validate:"required,max=1" description:"Write off expense flag or not"`                   //是否冲销费用标志 Y-冲销费用，N-不冲销费用
	CheckType             string `json:"checkType" validate:"required,max=1" description:"Query type"`                                      //查询类型
	MediaType             string `json:"mediaType" validate:"required,max=3" description:"Media type"`                                      //介质类型
	MediaNumber           string `json:"mediaNumber" validate:"required,max=30" description:"Media number"`                                 //介质号码
	Agreement             string `json:"agreement" validate:"max=30" description:"Contract number"`                                         //合约号
	AgreementType         string `json:"agreementType" validate:"max=5" description:"Contract type"`                                        //合约类型
}

type SV000015O struct {
}

func (*SV000015I) GetServiceKey() string {
	return "ssv0000015"
}
