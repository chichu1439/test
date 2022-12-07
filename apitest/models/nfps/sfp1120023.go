// Package models will define request and response message struct
// Version: v0.0.1
package models

type FP112023I struct {
	FPSPain013
	Pain013OriMsg
	MessageContent []byte `json:"messageContent"`
}

func (*FP112023I) GetServiceKey() string {
	return "FP112023"
}

type FP112023O struct {
	FPSPacs002
}
type FP112023Outgress struct {
	Body []byte `json:"body"`
}

func (*FP112023Outgress) GetServiceKey() string {
	return "FP212023"
}

type FP112023IOutgressResponse struct {
	Body []byte `json:"body"`
}
