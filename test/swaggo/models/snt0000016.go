//Version: v1.0.0
package models

type NT000016I struct {
	TemplateName    string `json:"templateName" description:"Template name" validate:"required"`
	TemplateContent string `json:"templateContent" description:"Template content" validate:"required"`
}

type NT000016O struct {
	RelateMcId string `json:"relateMcId" description:"Relate mc id"`
}

func (i *NT000016I) GetServiceKey() string {
	return "NT000016"
}
