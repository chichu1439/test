package models

type FP900110I struct {
	ParamGroups []string `json:"paramGroup"`
}

type FP900110O struct {
	EnumList []EnumList `json:"enumList"`
}

type EnumList struct {
	ParamName  string `json:"paramName"`
	ParamValue string `json:"paramValue"`
}

func (*FP900110I) GetServiceKey() string {
	return "FP900110"
}
