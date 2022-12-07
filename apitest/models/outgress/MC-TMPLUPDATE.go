package outgress

type MC_TMPLUPDATEI struct {
	WorkSpace    string `json:"workSpace" description:"Work space" validate:"required"`
	Env          string `json:"env" description:"Env" validate:"required"`
	OrgID        string `json:"orgID" description:"Org i d" validate:"required"`
	TemplateID   string `json:"templateID" description:"Template i d" validate:"required"`
	TemplateName string `json:"templateName" description:"Template name" validate:"required"`
	Status       string `json:"status" description:"Status" validate:"required"`
	Content      string `json:"content" description:"Content" validate:"required"`
	CreateUser   string `json:"createUser" description:"Create user" validate:"required"`
	UpdateUser   string `json:"updateUser" description:"Update user" validate:"required"`
}

type MC_TMPLUPDATEO struct {
	Data bool `json:"data" description:"Data"`
}

func (i *MC_TMPLUPDATEI) GetServiceKey() string {
	return "MC-TMPLUPDATE"
}
