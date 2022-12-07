// Package models will define request and response message struct
// Version: v0.0.1
package models

type FP730001I struct {
}

type FP730001O struct {
}

func (*FP730001I) GetServiceKey() string {
	return "FP730001"
}
