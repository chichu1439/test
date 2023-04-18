package models

type CM000041I struct {
	RandomKey string `json:"randomKey" description:"Random key" validate:"required"`
}

type CM000041O struct {
	ExtFlage string `json:"extFlage" description:"Ext flage"`
}

func (*CM000041I) GetServiceKey() string {
	return "CM000041"
}
