// Package models will define request and response message struct
// Version: v0.0.1
package models

type FP700011I struct {
}

type FP700011O struct {
}

func (*FP700011I) GetServiceKey() string {
	return "FP700011"
}
