package models

type Sfp9000046Tpart struct {
	PartClearingCode string `json:"part_clearing_code"`
	Echo20022Status  string `json:"echo_20022_status"` // 0 正常 1 超时 2 调用失败 3 报错  4 未知异常
	Echo8583Status   string `json:"echo_8583_status"`  // 0 正常 1 超时 2 调用失败 3 报错  4 未知异常
}

type UpdateStatus struct {
	status string
}

type FP900046I struct {
}

func (*FP900046I) GetServiceKey() string {
	return "FP900046"
}
