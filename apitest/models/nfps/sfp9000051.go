//Version: v1
package models

type FP900051I struct {
	ScenarioId int64 `json:"scenarioId"`
}

type FP900051O struct {
	ScenarioSteps []ScenarioStep `json:"scenarioSteps"`
}
type ScenarioStep struct {
	Id                  int64  `json:"id"`
	Name                string `json:"name"`
	Description         string `json:"description"`
	Step                int64  `json:"step"`
	ApiId               int64  `json:"apiId"`
	ApiName             string  `json:"apiName"`
	EventId             string  `json:"eventId"`
	Result              bool   `json:"result"`
	SenderHeaderId      string `json:"senderHeaderId"`
	SenderRequestId     string `json:"senderRequestId"`
	SenderResponseId    string `json:"senderResponseId"`
	ReceiverHeaderId    string `json:"receiverHeaderId"`
	ReceiverRequestId   string `json:"receiverRequestId"`
	ReceiverResponseId  string `json:"receiverResponseId"`
	PreRequestScript    string `json:"preRequestScript"`
	PostResponseScript  string `json:"postResponseScript"`
	TestAssertionScript string `json:"testAssertionScript"`
	Types               string `json:"type"`
	MessageName         string `json:"messageName"`
}

func (*FP900051I) GetServiceKey() string {
	return "FP900051"
}
