// Package models will define request and response message struct
// Version: v0.0.1
package models

type FP012006I FP912016I

type FP012006O FP912016O

func (*FP012006I) GetServiceKey() string {
	return "FP012006"
}
