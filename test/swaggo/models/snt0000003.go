//Version: v1.0.0
package models

type NT000003I struct {
	RelateMcId string `json:"relateMcId" validate:"required" description:"Template Id created by MC"`
}

type NT000003O struct {
	TemplateName    string `json:"templateName" description:"Template name"`
	TemplateTitle   string `json:"templateTitle" description:"Template title"`
	TemplateContent string `json:"templateContent" description:"Template content"`
}
