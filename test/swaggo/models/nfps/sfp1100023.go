// Package models will define request and response message struct
// Version: v0.0.1
package models

type FP110023I struct {
	FPSPain013
	Pain013OriMsg
}

func (*FP110023I) GetServiceKey() string {
	return "FP110023"
}

type FP110023O struct {
	FPSPacs002
}
type FP110023Outgress struct {
	Body []byte `json:"body"`
}

func (*FP110023Outgress) GetServiceKey() string {
	return "FP210023"
}

type FP1100223OutgressResponse struct {
	Body []byte `json:"body"`
}
