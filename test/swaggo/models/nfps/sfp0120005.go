// Package models will define request and response message struct
// Version: v0.0.1
package models

type FP012005I FP912002I

type FP012005O FP912002O

func (*FP012005I) GetServiceKey() string {
	return "FP012005"
}
