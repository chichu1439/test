package models

type XX100001I struct {
	Date string
}

type XX100001O struct {
}

func (*XX100001I) GetServiceKey() string {
	return "XX100001"
}
