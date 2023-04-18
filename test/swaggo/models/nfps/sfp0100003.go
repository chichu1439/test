package models

type FP010003I struct {
	FPSCamt060
}

type FP010003O struct {
	// FPSAdmi002 fail
	// FPSCamt052 success
}

func (*FP010003I) GetServiceKey() string {
	return "FP010003"
}
