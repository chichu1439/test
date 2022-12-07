/**
 * @Author: fangwen
 * @Description:
 * @File: sfp0000050
 * @Version: 1.0.0
 * @Date: 2022/3/28 6:55 PM
 */

package models

import (
	"github.com/go-playground/validator/v10"
)

type FP000050I struct {
	PageNum         int `json:"pageNum" validate:"required"`
	PageRecordCount int `json:"pageRecordCount" validate:"required"`
}

type FP000050O struct {
	PageNum         int            `json:"pageNum"`
	PageRecordCount int            `json:"pageRecordCount"`
	TotalCount      int            `json:"totalCount"`
	ScenarioList    []ScenarioInfo `json:"scenarioList"`
}

func (*FP000050I) GetServiceKey() string {
	return "FP000050"
}

func (d *FP000050I) Validate() error {
	validate := validator.New()
	return validate.Struct(d)
}

type FP000051I struct {
	ScenarioId int64 `json:"scenarioId" validate:"required"`
}

type FP000051O struct {
	ScenarioSteps []ScenarioStep `json:"scenarioSteps"`
}

func (*FP000051I) GetServiceKey() string {
	return "FP000051"
}

func (d *FP000051I) Validate() error {
	validate := validator.New()
	return validate.Struct(d)
}

type FP000052I struct {
	ScenarioId         int64  `json:"scenarioId" validate:"required"`
	ScenarioStepId     int64  `json:"scenarioStepId" validate:"required"`
	Type               string `json:"type"`
	RandId             string `json:"randId"`
	SenderHeaderId     int64  `json:"senderHeaderId"`
	SenderRequestId    int64  `json:"senderRequestId"`
	SenderResponseId   int64  `json:"senderResponseId"`
	ReceiverHeaderId   int64  `json:"receiverHeaderId"`
	ReceiverRequestId  int64  `json:"receiverRequestId"`
	ReceiverResponseId int64  `json:"receiverResponseId"`
}

type FP000052O struct {
	//SenderHeaderFields     []ScenarioFieldValue `json:"senderHeaderFields"`
	//SenderRequestFields    []ScenarioFieldValue `json:"senderRequestFields"`
	//SenderResponseFields   []ScenarioFieldValue `json:"senderResponseFields"`
	//ReceiverHeaderFields   []ScenarioFieldValue `json:"receiverHeaderFields"`
	//ReceiverRequestFields  []ScenarioFieldValue `json:"receiverRequestFields"`
	//ReceiverResponseFields []ScenarioFieldValue `json:"receiverResponseFields"`
	HeaderFields  []ScenarioFieldValue `json:"headerFields"`
	HeaderName    string               `json:"headerName"`
	RequestFields []ScenarioFieldValue `json:"requestFields"`
	RequestName   string               `json:"requestName"`
}

func (*FP000052I) GetServiceKey() string {
	return "FP000052"
}

func (d *FP000052I) Validate() error {
	validate := validator.New()
	return validate.Struct(d)
}

type FP000053I struct {
	Keyword     map[string]interface{} `json:"keyword"`
	WorkSpaceId string                 `json:"workSpaceId"`
	SceneName   string                 `json:"scenarioName"`
	Sort        string                 `json:"sort"`
	Order       string                 `json:"order"`
	PageSize    int64                  `json:"pageSize" validate:"required"`
	Page        int64                  `json:"pageNum" validate:"required"`
}

type FP000053O struct {
	Data  []ApiDetail `json:"data"`
	Total int64       `json:"total"`
}

func (*FP000053I) GetServiceKey() string {
	return "FP000053"
}

func (d *FP000053I) Validate() error {
	validate := validator.New()
	return validate.Struct(d)
}

type FP000054I struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description"`
}

type FP000054O struct {
	Status string `json:"status"`
}

func (*FP000054I) GetServiceKey() string {
	return "FP000054"
}

func (d *FP000054I) Validate() error {
	validate := validator.New()
	return validate.Struct(d)
}

type FP000055I struct {
	//Id          int    `json:"id" validate:"required"`
	//Name        string `json:"name"`
	//Description string `json:"description"`
	Id           int64              `json:"id" validate:"required"`
	ScenarioName string             `json:"scenarioName"`
	Description  string             `json:"description"`
	ScenarioStep []ScenarioStepList `json:"scenarioStepList"`
	//ClickStep []int `json:"clickStep"`
}

type FP000055O struct {
	Status string `json:"status"`
}

func (*FP000055I) GetServiceKey() string {
	return "FP000055"
}

func (d *FP000055I) Validate() error {
	validate := validator.New()
	return validate.Struct(d)
}

type FP000056I struct {
	Id int64 `json:"id" validate:"required"`
}

type FP000056O struct {
	Status string `json:"status"`
}

func (*FP000056I) GetServiceKey() string {
	return "FP000056"
}

func (d *FP000056I) Validate() error {
	validate := validator.New()
	return validate.Struct(d)
}

type FP000057I struct {
	ScenarioName string             `json:"scenarioName" validate:"required"`
	Description  string             `json:"description"`
	ScenarioStep []ScenarioStepList `json:"scenarioStepList"`
}

type ScenarioStepList struct {
	Id                     int64                `json:"id"`
	Step                   int64                `json:"step" validate:"required"`
	Name                   string               `json:"name" validate:"required"`
	Description            string               `json:"description"`
	ApiId                  int64                `json:"apiId" validate:"required"`
	ApiName                string               `json:"apiName"`
	EventId                string               `json:"eventId"`
	Result                 bool                 `json:"result"`
	PreRequestScript       string               `json:"preRequestScript"`
	PostResponseScript     string               `json:"postResponseScript"`
	TestAssertionScript    string               `json:"testAssertionScript"`
	SenderHeaderFields     []ScenarioFieldValue `json:"senderHeaderFields"`
	SenderRequestFields    []ScenarioFieldValue `json:"senderRequestFields"`
	SenderResponseFields   []ScenarioFieldValue `json:"senderResponseFields"`
	ReceiverHeaderFields   []ScenarioFieldValue `json:"receiverHeaderFields"`
	ReceiverRequestFields  []ScenarioFieldValue `json:"receiverRequestFields"`
	ReceiverResponseFields []ScenarioFieldValue `json:"receiverResponseFields"`
}

type FP000057O struct {
	ScenarioId int64 `json:"scenarioId"`
}

func (*FP000057I) GetServiceKey() string {
	return "FP000057"
}

func (d *FP000057I) Validate() error {
	validate := validator.New()
	return validate.Struct(d)
}

type FP000058I struct {
	Id                  int64  `json:"id" validate:"required"`
	Step                int64  `json:"step"`
	Name                string `json:"name"`
	Result              bool   `json:"result"`
	Description         string `json:"description"`
	PreRequestScript    string `json:"preRequestScript"`
	PostResponseScript  string `json:"postResponseScript"`
	TestAssertionScript string `json:"testAssertionScript"`
}

type FP000058O struct {
	Status string `json:"status"`
}

func (*FP000058I) GetServiceKey() string {
	return "FP000058"
}

func (d *FP000058I) Validate() error {
	validate := validator.New()
	return validate.Struct(d)
}

type FP000059I struct {
	Id int64 `json:"id" validate:"required"`
}

type FP000059O struct {
	Status string `json:"status"`
}

func (*FP000059I) GetServiceKey() string {
	return "FP000059"
}

func (d *FP000059I) Validate() error {
	validate := validator.New()
	return validate.Struct(d)
}

type FP000060I struct {
	ScenarioStepId int64  `json:"scenarioStepId" validate:"required"`
	SchemaId       int64  `json:"schemaId" validate:"required"`
	Associated     string `json:"associated"`
	Description    string `json:"description"`
	EnumType       int64  `json:"enumType"`
	FieldName      string `json:"fieldName" validate:"required"`
	FieldType      string `json:"fieldType"`
	FieldValue     string `json:"fieldValue" validate:"required"`
}

type FP000060O struct {
	Status string `json:"status"`
}

func (*FP000060I) GetServiceKey() string {
	return "FP000060"
}

func (d *FP000060I) Validate() error {
	validate := validator.New()
	return validate.Struct(d)
}

type Fields struct {
	Id          int64  `json:"id"  validate:"required"`
	Description string `json:"description"`
	Associated  string `json:"associated"`
	EnumType    int64  `json:"enumType"`
	FieldValue  string `json:"fieldValue"`
}

type FP000061I struct {
	Fields []Fields `json:"fields"`
}

type FP000061O struct {
	Status string `json:"status"`
}

func (*FP000061I) GetServiceKey() string {
	return "FP000061"
}

func (d *FP000061I) Validate() error {
	validate := validator.New()
	return validate.Struct(d)
}

type FP000062I struct {
	Id int64 `json:"id" validate:"required"`
}

type FP000062O struct {
	Status string `json:"status"`
}

func (*FP000062I) GetServiceKey() string {
	return "FP000062"
}

func (d *FP000062I) Validate() error {
	validate := validator.New()
	return validate.Struct(d)
}

type FP000063I struct {
	ApiId string `json:"apiId" validate:"required"`
}

type FP000063O struct {
	SchemaFields map[string]SchemaFields `json:"schemaFields"`
}

func (*FP000063I) GetServiceKey() string {
	return "FP000063"
}

func (d *FP000063I) Validate() error {
	validate := validator.New()
	return validate.Struct(d)
}

type FP000064I struct {
}

type FP000064O struct {
	APIInfo  []APIInfo
	EnvId    string `json:"envId"`
	WorkSpec string `json:"wks"`
}

func (*FP000064I) GetServiceKey() string {
	return "FP000064"
}

func (d *FP000064I) Validate() error {
	validate := validator.New()
	return validate.Struct(d)
}

type FP000065I struct {
}

type FP000065O struct {
	HostTestLoad []HostTestLoad `json:"hostList"`
	ScenarioType []string       `json:"scenarioType"`
}

func (*FP000065I) GetServiceKey() string {
	return "FP000065"
}

func (d *FP000065I) Validate() error {
	validate := validator.New()
	return validate.Struct(d)
}

type FP000066I struct {
	ThreadNumber    int        `json:"threadNumber" validate:"required"`
	DurationSecond  int        `json:"durationSecond" validate:"required"`
	TransactionType string     `json:"transactionType" validate:"required"`
	CaseNumber      string     `json:"caseNumber" validate:"required"`
	HostIps         []HostInfo `json:"hostInfo"`
}

type FP000066O struct {
	Status string `json:"status"`
}

func (*FP000066I) GetServiceKey() string {
	return "FP000066"
}

func (d *FP000066I) Validate() error {
	validate := validator.New()
	return validate.Struct(d)
}

type FP000067I struct {
	HostIps []HostInfo `json:"hostInfo"`
}

type HostInfo struct {
	Ip   string `json:"ip"`
	Port string `json:"port"`
	Path string `json:"path"`
}

type FP000067O struct {
	Data map[string]interface{} `json:"data"`
}

func (*FP000067I) GetServiceKey() string {
	return "FP000067"
}

func (d *FP000067I) Validate() error {
	validate := validator.New()
	return validate.Struct(d)
}

type FP000068I struct {
}

type ScenarioInfoEnum struct {
	Id                   int                    `json:"id"`
	ScenarioName         string                 `json:"scenarioName"`
	ScenarioStepListName []ScenarioStepListName `json:"scenarioStepListName"`
}

type FP000068O struct {
	ScenarioInfoEnum []ScenarioInfoEnum `json:"scenarioInfoEnum"`
}

type ScenarioStepListName struct {
	Id               int    `json:"id"`
	ScenarioStepName string `json:"scenarioStepName"`
}

func (*FP000068I) GetServiceKey() string {
	return "FP000068"
}

func (d *FP000068I) Validate() error {
	validate := validator.New()
	return validate.Struct(d)
}

type FP000069I struct {
	TransactionType string `json:"transactionType"`
	StartTime       int64  `json:"startTime"`
	TestId          string `json:"testId"`
	EndTime         int64  `json:"endTime"`
	Page            int    `json:"page"`
	PageNum         int    `json:"pageNum"`
}

type FP000069O struct {
	TestRecord []TestResult2 `json:"testRecord"`
	TotalRows  int64         `json:"totalRows"`
}

type FP000070I struct {
}

type FP000070O struct {
	PartClearingCode []PartClearingInfo `json:"partClearingCode"`
}

type PartClearingInfo struct {
	IndirectPartClearingCode string `json:"indirectPartClearingCode"`
	PartName                 string `json:"partName"`
}

func (*FP000070I) GetServiceKey() string {
	return "FP000070"
}

func (d *FP000070I) Validate() error {
	validate := validator.New()
	return validate.Struct(d)
}

type FP000071I struct {
	PartId     string `json:"partId"`
	Port       string `json:"port"`
	Ip         string `json:"ip"`
	TcpStatus  string `json:"tcp_status"`
	MessageStr string `json:"data"`
}

type FP000071O struct {
	Body string `json:"body"`
}

func (*FP000071I) GetServiceKey() string {
	return "FP000071"
}

func (d *FP000071I) Validate() error {
	validate := validator.New()
	return validate.Struct(d)
}

type FP000072I struct {
	PartId     string `json:"partId"`
	Port       string `json:"port"`
	Ip         string `json:"ip"`
	TcpStatus  string `json:"tcp_status"`
	MessageStr string `json:"data"`
}

func (d *FP000072I) Validate() error {
	validate := validator.New()
	return validate.Struct(d)
}

type FP000072O struct {
	Body string `json:"body"`
}

func (*FP000072I) GetServiceKey() string {
	return "FP000072"
}
