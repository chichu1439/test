package models

type Sfp9000048Tpart struct {
	PartClearingCode string
	Status           string
}

type UpdatePartStatus struct {
	Status string
}
type FP900048I struct {
}

func (*FP900048I) GetServiceKey() string {
	return "FP900048"
}
