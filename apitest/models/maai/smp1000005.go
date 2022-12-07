package models

type SMP1000005I struct {
	IDPhoto  string `json:"idPhoto" validate:"required"`
	Selfie   string `json:"selfie" validate:"required"`
	ReqRefNo string `json:"reqRefNo" validate:"omitempty,max=20"`
}

type SMP1000005O struct {
	FacePass     bool   `json:"facePass"`
	LivenessPass bool   `json:"livenessPass"`
	RespCode     string `json:"respCode"`
	RespDesc     string `json:"respDesc"`
	ReqRefNo     string `json:"reqRefNo"`
	RespRefNo    string `json:"respRefNo"`
}

func (i *SMP1000005I) GetServiceKey() string {
	return "MP100005"
}
