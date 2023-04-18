/**
 * @Author: fangwen
 * @Description:
 * @File: sfp9000099
 * @Version: 1.0.0
 * @Date: 2022/3/27 1:19 PM
 */

package models

type FP900099I struct {
	TestScenarioId int64  `json:"testScenarioId"`
	ScenarioStepId int64  `json:"scenarioStepId"`
	TcpStatus      string `json:"tcpStatus"`
	Type           string `json:"type"`
}

type FP900099O struct {
	HeaderFields  []ScenarioFieldValue `json:"headerFields"`
	HeaderName    string               `json:"headerName"`
	RequestFields []ScenarioFieldValue `json:"requestFields"`
	RequestName   string               `json:"requestName"`
	//SenderResponseFields   []ScenarioFieldValue `json:"senderResponseFields"`
	//ReceiverHeaderFields  []ScenarioFieldValue `json:"receiverHeaderFields"`
	//ReceiverRequestFields []ScenarioFieldValue `json:"receiverRequestFields"`
	// ReceiverResponseFields []ScenarioFieldValue `json:"receiverResponseFields"`
	SedTemplate SedTemplate `json:"sedTemplate"`
}

type SedTemplate struct {
	SenderRequestTemplate   string `json:"sendRequestTemplate" xorm:"'send_request_template'"`
	ReceiverRequestTemplate string `json:"receiverRequestTemplate" xorm:"'receiver_request_template'"`
	ApiId                   int64  `json:"apiId"  xorm:"'api_id'"`
	Type                    string `json:"type" xorm:"'type'"`
	Simulation              bool   `json:"simulation" xorm:"'simulation'"`
}

func (*FP900099I) GetServiceKey() string {
	return "FP900099"
}
