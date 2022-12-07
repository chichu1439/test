package models

type FP010002I struct {
	FPSPacs028
}

type FP010002O struct {
	FPSPacs002
	FPSAdmi002
}

func (*FP010002I) GetServiceKey() string {
	return "FP010002"
}
