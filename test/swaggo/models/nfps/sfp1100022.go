// Package models will define request and response message struct
// Version: v0.0.1
package models

type FP110022I struct {
	FPSPain013
	ClrRefId string
	MsgId    string
	Pain013OriMsg
}

func (*FP110022I) GetServiceKey() string {
	return "FP110022"
}

type FP110022O struct {
	FPSPacs002
}
