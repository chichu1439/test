//Version: v1.0.0
package models

type NT000019I struct {
	ListId []string `json:"listId" description:"List id" validate:"required"`
}

type NT000019O struct {
	ListData []ListData `json:"listData" description:"List data"`
}

type ListData struct {
	RelateMcId      string `json:"relateMcId" description:"Relate mc id" `
	TemplateContent string `json:"templateContent" description:"Template content" validate:"required"`
}

func (i *NT000019I) GetServiceKey() string {
	return "NT000019"
}
