//Version: v1.0.0
package models

type NT000004I struct {
	TemplateId      int    `json:"templateId" validate:"required" description:"Template Id"`
	TemplateName    string `json:"templateName" validate:"required" description:"Template Name"`
	TemplateTitle   string `json:"templateTitle" validate:"required" description:"Template Title"`
	TemplateContent string `json:"templateContent" validate:"required" description:"Template Content"`
	RelateMcId      string `json:"relateMcId" validate:"required" description:"Template Id created by MC"`
}

type NT000004O struct {
}

func (i *NT000004I) GetServiceKey() string {
	return "NT000004"
}
