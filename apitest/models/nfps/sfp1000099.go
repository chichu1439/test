/**
 * @Author: fangwen
 * @Description:
 * @File: sfp1000099
 * @Version: 1.0.0
 * @Date: 2022/3/27 1:19 PM
 */

package models

type FP100099I struct {
	TestScenarioId int64  `json:"testScenarioId"`
	ScenarioStepId int64  `json:"scenarioStepId"`
	Type           string `json:"type"`
	TcpStatus      string `json:"tcpStatus"`
	//ApiId int64 `json:"apiId"`
	Fields []FieldUpdates `json:"fields"`
}

type FieldUpdates struct {
	Id          int64  `json:"id"`
	Description string `json:"description"`
	FieldName   string `json:"fieldName"`
	FieldType   string `json:"fieldType"`
	FieldValue  string `json:"fieldValue"`
	EnumType    int    `json:"enumType"`
	SchemaId    int64  `json:"schemaId"`
}

type FP100099O struct {
	HeaderFields []ScenarioFieldValue `json:"senderHeaderFields"`
	RequestData  []byte               `json:"sendRequestData"`
	//ReceiverHeaderFields   []ScenarioFieldValue `json:"receiverHeaderFields"`
	//ReceiverRequestData  []byte `json:"receiverRequestData"`
	ApiInfo    RequestApiInfo `json:"sedTemplate"`
	TraceId    string         `json:"traceId"`
	Types      string         `json:"type"`
	Simulation bool           `json:"simulation"`
}

type RequestApiInfo struct {
	Name         string `json:"name"`
	Category     string `json:"category"`
	Version      string `json:"version"`
	Type         string `json:"type"`
	Protocol     string `json:"protocol"`
	Format       string `json:"format"`
	Endpoint     string `json:"endpoint"`
	Path         string `json:"path"`
	ExternalPath string `json:"externalPath"`
	Method       string `json:"method"`
	TimeoutS     string `json:"timeoutS"`
}

func (*FP100099I) GetServiceKey() string {
	return "FP100099"
}
