package models


type BatchRequest struct {
 Id         int   `json:"id"`
 ProjId     int   `json:"projId"`
 FlowId     string   `json:"flowId"`
 ExecId     int   `json:"execId"`
 JobId      string   `json:"jobId"`
 ReplyTopic string   `json:"replyTopic"`
 PubMap     map[string]string   `json:"pubMap"`
 EnvMap     map[string]string   `json:"envMap"`
}


type BatchResponse struct {
 Id         int   `json:"id"`
 ProjId     int   `json:"projId"`
 FlowId     string   `json:"flowId"`
 ExecId     int   `json:"execId"`
 JobId      string   `json:"jobId"`
 ReplyTopic string   `json:"replyTopic"`
 RetCode    string   `json:"retCode"`
 RetMsg     string   `json:"retMsg"`
 PubMap     map[string]string   `json:"pubMap"`
 EnvMap     map[string]string   `json:"envMap"`
}

func (*BatchResponse) GetServiceKey() string {
	return "bcReceiveEvent"
}
