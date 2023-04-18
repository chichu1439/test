package models

type FP900072I struct {
	ScenarioStepId int64  `json:"scenarioStepId" validate:"required"`
	Associated     string `json:"associated"`
	SchemaId       int64  `json:"schemaId" validate:"required"`
	Description    string `json:"description"`
	EnumType       int64  `json:"enumType"`
	FieldName      string `json:"fieldName" validate:"required"`
	FieldType      string `json:"fieldType"`
	FieldValue     string `json:"fieldValue" validate:"required"`
}

type FP900072O struct {
	Status string `json:"status"`
}

func (*FP900072I) GetServiceKey() string {
	return "FP900072"
}
