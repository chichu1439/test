// Package models will define request and response message struct
// Version: v0.0.1
package models

type FP010018I struct {
	FPSPain013
}

type FP010018O struct {
	FPSPacs002
	FPSAdmi002
}

func (*FP010018I) GetServiceKey() string {
	return "FP010018"
}
