//Version: v1
package models

type PD000110Request struct {
}

type PD000110Response struct {
	ListTemplate []PD000110List `json:"listTemplate"`
}
type PD000110List struct {
	TemplateName         string `json:"templateName"`
	Version              int    `json:"version"`
	TemplateType         string `json:"templateType"`
	TemplateGenerateTime string `json:"templateGenerateTime"`
	TemplateGenerateDate string `json:"templateGenerateDate"`
	TemplateFileType     string `json:"templateFileType"`
	TemplateFileKey      string `json:"templateFileKey"`
	TemplateContent      string `json:"templateContent"`
}
