package models

type CM000044I struct {
	OptType string `json:"optType" description:"Opt type" validate:"required"`
	Key     string `json:"key" description:"Key" validate:"required"`
	Value   string `json:"value" description:"Value"`
}

type CM000044O struct {
	ExtFlage string `json:"extFlage" description:"Ext flage"`
	Key      string `json:"key" description:"Key"`
	Value    string `json:"value" description:"Value"`
}

func (*CM000044I) GetServiceKey() string {
	return "CM000044"
}
