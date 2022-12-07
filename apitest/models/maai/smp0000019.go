package models

// desc /auth/login/customerNo
type SMP0000019I struct {
	CustomerNo string `json:"customerNo" validate:"required,max=50"` // Y 50 Customer no
	Password   string `json:"password" validate:"required,max=1024"` // Y 1024 Password hash
	//DeviceId         string `json:"deviceId" validate:"required,max=50"`         // Y 50 Device unique id
	//DeviceModel      string `json:"deviceModel" validate:"required,max=50"`      // Y 50 Device model
	//DeviceSystem     string `json:"deviceSystem" validate:"required,max=50"`     // Y 50 Device System
	//DeviceVersion    string `json:"deviceVersion" validate:"required,max=20"`    // Y 20 Device Version
	//DeviceMacAddress string `json:"deviceMacAddress" validate:"required,max=50"` // Y 50 Device mac address
	//DeviceChannel    string `json:"deviceChannel" validate:"required,max=50"`    // Y 50 Channel device of (WEB or MOBILE). Each channel will be available one for device only.
}

type SMP0000019O struct {
	Token      string `json:"token"`
	CustomerNo string `json:"customerNo"`
}

func (i *SMP0000019I) GetServiceKey() string {
	return "MP000019"
}
