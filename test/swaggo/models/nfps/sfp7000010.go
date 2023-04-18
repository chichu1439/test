// Package models will define request and response message struct
// Version: v0.0.1
package models

type FP700010I struct {
}

type FP700010O struct {
}

func (*FP700010I) GetServiceKey() string {
	return "FP700010"
}
