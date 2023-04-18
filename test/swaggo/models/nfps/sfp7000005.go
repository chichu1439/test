// Package models will define request and response message struct
// Version: v0.0.1
package models

type FP700005I struct {
}

type FP700005O struct {
}

func (*FP700005I) GetServiceKey() string {
	return "FP700005"
}
