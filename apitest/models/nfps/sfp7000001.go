// Package models will define request and response message struct
// Version: v0.0.1
package models

type FP700001I struct {
}

type FP700001O struct {
	PartClearingCode string
	PartType         string
	PartName         string
}
