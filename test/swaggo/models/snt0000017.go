//Version: v1.0.0
package models

type NT000017I struct {
	RelateMcId      string `json:"relateMcId" description:"Relate mc id" validate:"required"`
	TemplateName    string `json:"templateName" description:"Template name" validate:"required"`
	TemplateContent string `json:"templateContent" description:"Template content" validate:"required"`
}

type NT000017O struct {
}

func (i *NT000017I) GetServiceKey() string {
	return "NT000017"
}
