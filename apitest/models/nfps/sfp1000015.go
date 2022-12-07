package models

import (
	"encoding/json"
)

type FP100015I struct {
	ApiId string `json:"apiId"`
	// SchemaId string `json:"schemaId"`
}

type FP100015O struct{
	SchemaFields map[string]SchemaFields `json:"schemaFields"`
}

type SchemaFields struct {
	Example []byte `json:"example"`
	Data  []ExtractingData `json:"data"`
	SchemaName string `json:"schemaName"`
	Total int        `json:"total"`
}

type ExtractingData struct {
	Description string `json:"description"`
	FieldName string `json:"fieldName"`
	FieldType string `json:"fieldType"`
	Required  string `json:"required"`
	FieldValue string `json:"fieldValue"`
	SchemaId int64 `json:"schemaId"`
}

func (*FP100015I) GetServiceKey() string {
	return "FP100015"
}

type SchemaFieldArry struct {
	SchemaField []SchemaField `json:"SchemaField"`
}

type SchemaField struct {
	DecimalDigits string `json:"decimalDigits"`
	DefaultValue  string `json:"defaultValue"`
	Description   string `json:"description"`
	FieldId       int64  `json:"fieldId"`
	FieldName     string `json:"fieldName"`
	FieldType     string `json:"fieldType"`
	IsArray       string `json:"isArray"`
	Length        string `json:"length"`
	MetaDataId    int64  `json:"metaDataId"`
	MetaDataName  string `json:"metaDataName"`
	Required      string `json:"required"`
	SchemaId      int64  `json:"schemaId"`
	ScopeId       string `json:"scopeId"`
	StructName    string `json:"structName"`
	Type          string `json:"type"`
	CreatorId     string `json:"creatorId"`
	CreateTime    string `json:"createTime"`
	UpdatorId     string `json:"updatorId"`
	UpdateTime    string `json:"updateTime"`
}

type QSpecs struct {
	EventId string `json:"eventId"`
	Param   Param  `json:"param"`
}

type Param struct {
	CiType      string `json:"ciType"`
	TarCiType   string `json:"tarCiType"`
	Filter      Filter `json:"filter"`
	RefDataMode string `json:"refDataMode"`
}

type Filter struct {
	Id string `json:"Id"`
}

type MuRes struct {
	ErrorCode string   `json:"errorCode"`
	ErrorMsg  string   `json:"errorMsg"`
	Response  PageData `json:"response"`
}
type PData struct {
	Total int64           `json:"Total"`
	Data  json.RawMessage `json:"Data"`
}
