// Package models will define request and response message struct
// Version: v0.0.1
package models

import "net/http"

type FP110025I ISO85830200RequestToPayData

type FP110025O ISO85830200RequestToPayData

func (*FP110025I) GetServiceKey() string {
	return "FP110025"
}

type FP110025Outgress struct {
	Headers http.Header `json:"headers"`
	Body    []byte      `json:"body"`
}

func (*FP110025Outgress) GetServiceKey() string {
	return "FP210025"
}

type FP110025OutgressResponse struct {
	Headers http.Header `json:"headers"`
	Body    []byte      `json:"body"`
}
