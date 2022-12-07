package models

type SMP1000001I struct {
	MobileNo         string `json:"mobileNo" validate:"required,min=11,max=12"`
	Password         string `json:"password" validate:"required,max=1024"`
	DeviceID         string `json:"deviceId" validate:"omitempty,max=50"`
	DeivceModel      string `json:"deviceModel" validate:"omitempty,max=50"`
	DeviceSystem     string `json:"deviceSystem" validate:"omitempty,max=50"`
	DeviceVersion    string `json:"deviceVersion" validate:"omitempty,max=20"`
	DeviceMacAddress string `json:"deviceMacAddress" validate:"omitempty,max=50"`
	DeviceChannel    string `json:"deviceChannel" validate:"omitempty,oneof=WEB MOBILE"`
	ReqRefNo         string `json:"reqRefNo" validate:"omitempty,max=20"`
}

type SMP1000001O struct {
	CustomerNo      string `json:"customerNo"`
	Token           string `json:"token"`
	DopaPass        bool   `json:"dopaPass"`
	FaceComparePass bool   `json:"faceComparePass"`
	LivenessPass    bool   `json:"livenessPass"`
	RespCode        string `json:"respCode"`
	RespDesc        string `json:"respDesc"`
	ReqRefNo        string `json:"reqRefNo"`
	RespRefNo       string `json:"respRefNo"`
}

func (i *SMP1000001I) GetServiceKey() string {
	return "MP100001"
}
