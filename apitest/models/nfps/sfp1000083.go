package models

type FP100083I struct {
	PartId     string `json:"partId"`
	Port       string `json:"port"`
	Ip         string `json:"ip"`
	TcpStatus  string `json:"tcp_status"`
	MessageStr string `json:"messageStr"`
}

type FP100083O struct {
	Body string `json:"body"`
}

func (*FP100083I) GetServiceKey() string {
	return "FP100083"
}
