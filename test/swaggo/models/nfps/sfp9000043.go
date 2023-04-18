package models

type FP900043I struct {
	MsgId   string `json:"msgId" validate:"required"`
	TransId string `json:"transId" validate:"required"`
}

type LogInsertInfo struct {
	DocTxId             string `json:"docTxId"`
	MessageDate         string `json:"messageDate"`         //IntrBkSttlmDt
	Currency            string `json:"currency"`            // <doc:IntrBkSttlmAmt Ccy="CNY">
	MessageType         string `json:"messageType"`         // MsgDefIdr
	ParticipantSent     string `json:"participantSent"`     // Fr
	ParticipantReceived string `json:"participantReceived"` //To
	MessageID           string `json:"messageId"`
	Msg                 string `json:"msg"`       //
	Timestamp           int64  `json:"timestamp"` //
	TraceId             string `json:"traceId"`   //
}

type AppLogService struct {
	Level    string `json:"level"`
	Time     string `json:"time"`
	Msg      string `json:"msg"`
	TraceId  string `json:"traceId"`
	ParentId string `json:"parentId"`
	SpanId   string `json:"spanId"`
}

//public class AgentEvent {
//private Map<String, Object> meta;
//private Map<String, String> fields;
//private String message;
//private String source;
//}

type AgentEvent struct {
	Meta    map[string]interface{} `json:"meta"`
	Fields  map[string]string      `json:"fields"`
	Message string                 `json:"message"`
	Source  string                 `json:"source"`
}

type LogAgentEventMetric struct {
	Meta    map[string]interface{} `json:"meta"`
	Fields  map[string]string      `json:"fields"`
	Message AppLogService          `json:"message"`
	Source  string                 `json:"source"`
}

func (*FP900043I) GetServiceKey() string {
	return "asmLogApp"
}

type BaseInfo struct {
	Workspace string `json:"workspace"`
	Env       string `json:"env"`
	Org       string `json:"org"`
	Az        string `json:"az"`
	Su        string `json:"su"`
	NodeID    string `json:"node_id"`
}

type GetKeySfp910032ServiceRequest struct {
	Services []GetKeySfp910032Request `json:"services"`
}

type GetKeySfp910032Request struct {
	KeyType     string `validate:"required" json:"keyType"`
	ServiceId   string `validate:"required" json:"serviceId"`
	SystemType  string `validate:"required" json:"systemType"`
	Operator    string `validate:"required" json:"operator"`
	RequestTime Date   `validate:"required" json:"requestTime"`
}
