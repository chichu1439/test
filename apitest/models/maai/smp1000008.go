package models

type SMP1000008I struct {
	CustomerNo       string `json:"customerNo" validate:"omitempty,max=200"`
	Password         string `json:"password" validate:"required,max=1024"`
	DeviceID         string `json:"deviceId" validate:"omitempty,max=50"`
	DeivceModel      string `json:"deivceModel" validate:"omitempty,max=50"`
	DeviceSystem     string `json:"deviceSystem" validate:"omitempty,max=50"`
	DeviceVersion    string `json:"deviceVersion" validate:"omitempty,max=20"`
	DeviceMacAddress string `json:"deviceMacAddress" validate:"omitempty,max=50"`
	DeviceChannel    string `json:"deviceChannel" validate:"omitempty,max=50"`
	ReqRefNo         string `json:"reqRefNo" validate:"omitempty,max=20"`
}

type SMP1000008O struct {
	RespCode  string `json:"respCode"`
	RespDesc  string `json:"respDesc"`
	ReqRefNo  string `json:"reqRefNo"`
	RespRefNo string `json:"respRefNo"`
}

func (i *SMP1000008I) GetServiceKey() string {
	return "MP100008"
}
