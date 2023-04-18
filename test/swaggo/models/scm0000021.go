package models

type SCM0000021I struct {
	OptionId          string   `json:"optionId" description:"Option id"`
	AvailableFlag     string   `json:"availableFlag" description:"use as default flag, D - Defalut ,null not default"`
	LangIdentify      string   `json:"langIdentify" description:"Langue identify(01 English 02 Chinese 03 Thai 04 Spanish 05 French)"`
	OptionClass       string   `json:"optionClass" description:"Option class"`
	OptionName        string   `json:"optionName" description:"Option name"`
	ParentClassId     string   `json:"parentClassId" description:"IG-integral"`
	ParentClassIdList []string `json:"parentClassIdList" description:"Parent class id list"`
	PageNum           int      `json:"pageNum" description:"Page number"`
	PageRecordCount   int      `json:"pageRecordCount" description:"Page record count"`
}

type SCM0000021O struct {
	PageNum         int          `json:"pageNum" description:"Page number"`
	PageRecordCount int          `json:"pageRecordCount" description:"Page record count"`
	TotalCount      int          `json:"totalCount" description:"Total count"`
	Records         []OptionInfo `json:"records" description:"Records"`
}

type OptionInfo struct {
	OptionId      string `json:"optionId" description:"Option id"`
	AvailableFlag string `json:"availableFlag" description:"Available flag"`
	LangIdentify  string `json:"langIdentify" description:"Langue identify"`
	OptionClass   string `json:"optionClass" description:"Option class"`
	OptionName    string `json:"optionName" description:"Option name"`
	ParentClassId string `json:"parentClassId" description:"Parent class id"`
}

func (*SCM0000021I) GetServiceKey() string {
	return "CM000021"
}
