package models

type FP100082I struct {
	PartId     string `json:"partId"`
	Port       string `json:"port"`
	Ip         string `json:"ip"`
	TcpStatus  string `json:"tcp_status"`
	MessageStr string `json:"messageStr"`
}

type FP100082O struct {
	Body string `json:"body"`
}

func (*FP100082I) GetServiceKey() string {
	return "FP100082"
}
