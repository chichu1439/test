package models

type FP900032I struct {
	FieldRelationType string `json:"fieldRelationType"`
}
type FP900032O struct {
	TotalCount int             `json:"totalCount"`
	Records    []FP900032OItem `json:"records"`
}

type FP900032OItem struct {
	FieldRelationType string `json:"fieldRelationType"`
	FieldSequence     string `json:"fieldSequence"`
	Field20022        string `json:"field20022"`
	FieldLength20022  string `json:"filedLength20022"`
	FieldFormat20022  string `json:"filedFormat20022"`
	Field8583         string `json:"field8583"`
	FieldLength8583   string `json:"filedLength8583"`
	FieldFormat8583   string `json:"filedFormat8583"`
	RevSndFlag        string `json:"revSndFlag"`
}

func (*FP900032I) GetServiceKey() string {
	return "FP900032"
}
