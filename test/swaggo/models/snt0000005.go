//Version: v1.0.0
package models

type NT000005I struct {
	TemplateId int    `json:"templateId" validate:"required" description:"Template Id"`
	RelateMcId string `json:"relateMcId" validate:"required" description:"Template Id created by MC"`
}

type NT000005O struct {
}

func (i *NT000005I) GetServiceKey() string {
	return "NT000005"
}
