// Package models will define request and response message struct
// Version: v0.0.1
package models

type FP013018I struct {
	//Message message.Pain013BizData
}

type FP013018O struct {
	//Message message.Pain014BizData
}

func (*FP013018I) GetServiceKey() string {
	return "FP013018"
}
