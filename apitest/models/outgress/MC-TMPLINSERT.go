package outgress

type MC_TMPLINSERTI struct {
	Workspace    string `json:"workSpace" description:"Workspace" validate:"required"`
	Env          string `json:"env" description:"Env" validate:"required"`
	OrgID        string `json:"orgId" description:"Org i d" validate:"required"`
	TemplateID   string `json:"templateId" description:"Template i d"`
	TemplateName string `json:"templateName" description:"Template name" validate:"required"`
	Status       string `json:"status" description:"Status" validate:"required"`
	Content      string `json:"content" description:"Content" validate:"required"`
	CreateUser   string `json:"createUser" description:"Create user" validate:"required"`
}

type MC_TMPLINSERTO struct {
	Data string `json:"data" description:"Data"`
}

func (i *MC_TMPLINSERTI) GetServiceKey() string {
	return "MC-TMPLINSERT"
}
