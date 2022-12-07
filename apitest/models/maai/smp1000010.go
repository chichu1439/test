package models

// desc /auth/login/customerNo
type SMP1000010I struct {
	CustomerNo       string `json:"customerNo" validate:"required"`       // Y 50 Customer no
	Password         string `json:"password" validate:"required"`         // Y 1024 Password hash
	DeviceId         string `json:"deviceId" validate:"required"`         // Y 50 Device unique id
	DeviceModel      string `json:"deviceModel" validate:"required"`      // Y 50 Device model
	DeviceSystem     string `json:"deviceSystem" validate:"required"`     // Y 50 Device System
	DeviceVersion    string `json:"deviceVersion" validate:"required"`    // Y 20 Device Version
	DeviceMacAddress string `json:"deviceMacAddress" validate:"required"` // Y 50 Device mac address
	DeviceChannel    string `json:"deviceChannel" validate:"required"`    // Y 50 Channel device of (WEB or MOBILE). Each channel will be available one for device only.
	ReqRefNo         string `json:"reqRefNo"`                             // Y 20 Request Reference No
}

type SMP1000010O struct {
	Token      string `json:"token"`
	CustomerNo string `json:"customerNo"`
	RespCode   string `json:"respCode"`
	RespDesc   string `json:"respDesc"`
	ReqRefNo   string `json:"reqRefNo"`
	RespRefNo  string `json:"respRefNo"`
}

type SMP1000010IOutgress struct {
	SMP1000010I
}

func (i *SMP1000010I) GetServiceKey() string {
	return "MP100010"
}

func (i *SMP1000010IOutgress) GetServiceKey() string {
	return "MP300010"
}
