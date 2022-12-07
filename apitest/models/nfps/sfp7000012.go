// Package models will define request and response message struct
// Version: v0.0.1
package models

type FP700012I struct {
}

type FP700012O struct {
}

func (*FP700012I) GetServiceKey() string {
	return "FP700012"
}
