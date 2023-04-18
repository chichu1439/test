// Package models will define request and response message struct
// Version: v0.0.1
package models

type FP012018I struct {
	FPSPain013
}

type FP012018O struct {
	FPSPacs002
	FPSAdmi002
}

func (*FP012018I) GetServiceKey() string {
	return "FP012018"
}
