//Version: v1.0.0
package models

type NT000020I struct {
	Body             interface{} `json:"body" validate:"required"`
	AttachFileIDList []string    `json:"attachFileIDList" `
	BizDate          string      `json:"bizDate" validate:"required"`
	BizSeqNo         string      `json:"bizSeqNo" validate:"required"`
	//ListChannelTarget []ListChannelTarget`json:"listChannelTarget" `
}

type NT000020O struct {
	//ListChannel []ListChannelResponse `json:"listChannel" description:"List channel"`
}

func (i *NT000020I) GetServiceKey() string {
	return "NT000020"
}
