package outgress

type MC_TMPLGETI struct {
	WorkSpace    string `json:"workSpace" description:"Work space"`
	Env          string `json:"env" description:"Env"`
	OrgID        string `json:"orgId" description:"Org i d"`
	TemplateID   string `json:"templateId" description:"Template i d"`
	TemplateName string `json:"templateName" description:"Template name"`
	Status       string `json:"status" description:"Status"`
	Content      string `json:"content" description:"Content"`
	CreateUser   string `json:"createUser" description:"Create user"`
	UpdateUser   string `json:"updateUser" description:"Update user"`
	CreateTime   string `json:"createTime" description:"Create time"`
	UpdateTime   string `json:"updateTime" description:"Update time"`
}

type MC_TMPLGETO struct {
	TotalCnt  int      `json:"totalCnt" description:"Total cnt"`
	TotalPage int      `json:"totalPage" description:"Total page"`
	Result    []Result `json:"result" description:"Result"`
}

type Result struct {
	WorkSpace    string `json:"workSpace" description:"Work space"`
	Env          string `json:"env" description:"Env"`
	OrgID        string `json:"orgId" description:"Org i d"`
	TemplateID   string `json:"templateId" description:"Template i d"`
	TemplateName string `json:"templateName" description:"Template name"`
	Status       string `json:"status" description:"Status"`
	Content      string `json:"content" description:"Content"`
	CreateUser   string `json:"createUser" description:"Create user"`
	UpdateUser   string `json:"updateUser" description:"Update user"`
	CreateTime   string `json:"createTime" description:"Create time"`
	UpdateTime   string `json:"updateTime" description:"Update time"`
}

func (i *MC_TMPLGETI) GetServiceKey() string {
	return "MC-TMPLGET"
}
