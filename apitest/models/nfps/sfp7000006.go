// Package models will define request and response message struct
// Version: v0.0.1
package models

type FP700006I struct {
}

type FP700006O struct {
}

func (*FP700006I) GetServiceKey() string {
	return "FP700006"
}
