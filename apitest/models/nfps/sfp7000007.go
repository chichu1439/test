// Package models will define request and response message struct
// Version: v0.0.1
package models

type FP700007I struct {
}

type FP700007O struct {
}

func (*FP700007I) GetServiceKey() string {
	return "FP700007"
}
