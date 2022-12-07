package models

type FP900069I struct {
	ScenarioName string             `json:"scenarioName"`
	Description  string             `json:"description"`
	ScenarioStepListDao []ScenarioStepListDao `json:"scenarioStepListDao"`
}

type ScenarioStepListDao struct {
	Id                     int64 `json:"id"`
	Step                   int64                `json:"step" validate:"required"`
	Name                   string               `json:"name" validate:"required"`
	Description            string               `json:"description"`
	Result                 bool                  `json:"result"`
	ApiId                  int64                `json:"apiId" validate:"required"`
	ApiName                string               `json:"apiName"`
	EventId                string               `json:"eventId"`
	SenderHeaderId         string                `json:"senderHeaderId"`
	SenderRequestId        string                `json:"senderRequestId"`
	SenderResponseId       string                `json:"senderResponseId"`
	ReceiverHeaderId       string                `json:"receiverHeaderId"`
	ReceiverRequestId      string                `json:"receiverRequestId"`
	ReceiverResponseId     string                `json:"receiverResponseId"`
	PreRequestScript       string               `json:"preRequestScript"`
	PostResponseScript     string               `json:"postResponseScript"`
	TestAssertionScript    string               `json:"testAssertionScript"`
	RequestTemplate        string               `json:"requestTemplate"`
	ResponseTemplate       string               `json:"responseTemplate"`
	SenderHeaderFields     []ScenarioFieldValue `json:"senderHeaderFields"`
	SenderRequestFields    []ScenarioFieldValue `json:"senderRequestFields"`
	SenderResponseFields   []ScenarioFieldValue `json:"senderResponseFields"`
	ReceiverHeaderFields   []ScenarioFieldValue `json:"receiverHeaderFields"`
	ReceiverRequestFields  []ScenarioFieldValue `json:"receiverRequestFields"`
	ReceiverResponseFields []ScenarioFieldValue `json:"receiverResponseFields"`
}


type ScenarioFieldValue struct {
	Id          int64    `json:"id"`
	Description string `json:"description"  xorm:"'description'"`
	Required    string `json:"required" xorm:"'required'"`
	FieldName   string `json:"fieldName"  xorm:"'field_name'"`
	FieldType   string `json:"fieldType"  xorm:"'field_type'"`
	FieldValue  string `json:"fieldValue"  xorm:"'field_value'"`
	EnumType    int64  `json:"enumType"  xorm:"'enum_type'"`
	SchemaId    int64  `json:"schemaId"  xorm:"'schema_id'"`
}
type FP900069O struct {
	ScenarioId int64 `json:"scenarioId"`
}

func (*FP900069I) GetServiceKey() string {
	return "FP900069"
}
