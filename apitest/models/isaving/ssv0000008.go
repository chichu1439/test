//Version: v1.0.0
package models

type SSV0000008I struct {
}

type SSV0000008O struct {
	RecordNumber int                 `json:"recordNumber" description:"Record number"` //记录数量
	Result       []SSV1000008OResult `json:"result" description:"Result"`
}

type SSV1000008OResult struct {
	KeyVersion  string `json:"keyVersion" description:"Key version"`   //秘钥版本
	KeyValue    string `json:"keyValue" description:"Key value"`       //秘钥
	EffectTime  string `json:"effectTime" description:"Effect time"`   //生效时间
	DeffectTime string `json:"deffectTime" description:"Deffect time"` //失效时间
}

func (*SSV0000008I) GetServiceKey() string {
	return "ssv0000008"
}
