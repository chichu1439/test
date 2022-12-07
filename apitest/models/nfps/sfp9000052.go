//Version: v1
package models

type FP900052Request struct {
	ScenarioId         int64 `json:"scenarioId" validate:"required"`
	ScenarioStepId     int64 `json:"scenarioStepId" validate:"required"`
	//ApiId              int `json:"apiId" validate:"required"`
	SenderHeaderId     int64 `json:"senderHeaderId"`
	SenderRequestId    int64 `json:"senderRequestId"`
	SenderResponseId   int64 `json:"senderResponseId"`
	ReceiverHeaderId   int64 `json:"receiverHeaderId"`
	ReceiverRequestId  int64 `json:"receiverRequestId"`
	ReceiverResponseId int64 `json:"receiverResponseId"`
}

type FP900052Response struct {
	SenderHeaderFields     []ScenarioFieldValue `json:"senderHeaderFields"`
	SenderRequestFields    []ScenarioFieldValue `json:"senderRequestFields"`
	SenderResponseFields   []ScenarioFieldValue `json:"senderResponseFields"`
	ReceiverHeaderFields   []ScenarioFieldValue `json:"receiverHeaderFields"`
	ReceiverRequestFields  []ScenarioFieldValue `json:"receiverRequestFields"`
	ReceiverResponseFields []ScenarioFieldValue `json:"receiverResponseFields"`
}

func (*FP900052Request) GetServiceKey() string {
	return "FP900052"
}
