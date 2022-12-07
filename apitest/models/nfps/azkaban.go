package models

type BatchRequest struct {
	Id         int
	ProjId     int
	FlowId     string
	ExecId     int
	JobId      string
	ReplyTopic string
	PubMap     map[string]string
	EnvMap     map[string]string
}

type BatchResponse struct {
	Id         int
	ProjId     int
	FlowId     string
	ExecId     int
	JobId      string
	ReplyTopic string
	RetCode    string
	RetMsg     string
	PubMap     map[string]string
	EnvMap     map[string]string
}

func (*BatchResponse) GetServiceKey() string {
	return "bcReceiveEvent"
}
