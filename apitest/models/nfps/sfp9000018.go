// Package models will define request and response message struct
// Version: v0.0.1
package models

type FP900018I struct {
	ParamGroups []string `json:"paramGroup" ` // if not sed query all parameters
}

type FP900018O map[string]map[string]ParamInfo

// type FP900018O struct {
// 	ParamGroups map[string]ParamGroup `json:"paramGroups"` // parameter group name -> group parameters
// }

type ParamGroup struct {
	ParamInfos map[string]ParamInfo `json:"paramInfos"` // parameter value -> parameter details
}

type ParamInfo struct {
	ParamGroup string `json:"paramGroup"`
	ParamName  string `json:"paramName"`
	ParamValue string `json:"paramValue"`
	EffectDate string `json:"effectDate"`
	ExpireDate string `json:"expireDate"`
	Remark     string `json:"remark"`
}


func (*FP900018I) GetServiceKey() string {
	return "FP900018"
}
