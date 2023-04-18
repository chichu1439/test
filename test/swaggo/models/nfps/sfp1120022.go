// Package models will define request and response message struct
// Version: v0.0.1
package models

type FP112022I struct {
	FPSPain013
	ClrRefId string
	MsgId    string
	Pain013OriMsg
	RemInf         string
	MessageContent []byte `json:"messageContent"`
}

func (*FP112022I) GetServiceKey() string {
	return "FP112022"
}

type FP112022O struct {
	FPSPacs002
}
