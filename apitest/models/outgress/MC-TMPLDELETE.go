package outgress

type MC_TMPLDELETEI struct {
	TemplateID string `json:"templateID" description:"Template i d" validate:"required"`
}

type MC_TMPLDELETEO struct {
	Data bool `json:"data" description:"Data"`
}

func (i *MC_TMPLDELETEI) GetServiceKey() string {
	return "MC-TMPLDELETE"
}
