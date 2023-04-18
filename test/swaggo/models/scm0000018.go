package models

type SCM0000018I struct {
	SerialCode     string   `json:"serialCode" description:"Serial code"`
	SerialCodeList []string `json:"serialCodeList" description:"Serial code"`
}

type SCM0000018O struct {
	SerialNum     string   `json:"serialNum" description:"Serial number"`
	SerialNumList []string `json:"serialNumList" description:"Serial number"`
}

func (*SCM0000018I) GetServiceKey() string {
	return "CM000018"
}
