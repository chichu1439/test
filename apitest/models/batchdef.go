package models

import (
	//model "git.sirius.io/banking/common/models/customer"
	"time"
)

type BatchDefine struct {
	Id         int                 `json:"id"`
	ProjId     int                 `json:"projId"`
	FlowId     string              `json:"flowId"`
	ExecId     int                 `json:"execId"`
	JobId      string              `json:"jobId"`
	ReplyTopic string              `json:"replyTopic"`
	RetCode    string              `json:"retCode"`
	RetMsg     string              `json:"retMsg"`
	PubMap     map[string]string   `json:"pubMap"`
	EnvMap     map[string]string   `json:"envMap"`
	ArrayMap   map[string][]string `json:"arrayMap"`
}

type BatchArgs struct {
	PubMap   map[string]string   `json:"pubMap"`
	EnvMap   map[string]string   `json:"envMap"`
	ArrayMap map[string][]string `json:"arrayMap"`
}

type ReplyBody struct {
	ReturnCode    string      `json:"returnCode" description:"Return code"`
	ReturnMessage string      `json:"returnMessage" description:"Return message"`
	Data          interface{} `json:"data"`
}

//Result文件名及后缀：Batch code_Batch id_YYYYMMDD.csv 以半角逗号（即，）作分隔符，UTF-8
//文件内容：Batch code，Batch id，DCN，处理总条目，成功条目，成功金额，失败条目，失败金额，处理状态，失败原因，实际开始时间，实际结束时间

type ResultDefine struct {
	BatchCode       string    `json:"batchCode"`
	BatchId         string    `json:"batchId"`
	GroupDcn        string    `json:"groupDcn"`
	TotalNumber     int       `json:"totalNumber"`
	SuccNumber      int       `json:"succNumber"`
	SuccAmount      float64   `json:"succAmount"`
	FailNumber      int       `json:"failNumber"`
	FailAmount      float64   `json:"failAmount"`
	ProcessState    string    `json:"processState"`
	FailReason      string    `json:"failReason"`
	ActionStartTime time.Time `json:"actionStartTime"`
	ActionEndTime   time.Time `json:"actionEndTime"`
}

type CM001021RequestForm struct {
	BatchId string            `json:"batchId" description:"Batch id"`
	EnvMap  map[string]string `json:"EnvMap" description:"Env map"`
}

/*
这是取参数，下载文件，上传文件的postman示例和job配置示例：
1. DACM1060传入参数为：
{
    "BIZ_batch_code": "001"
}
出参为t_common_batch_control表中的信息：
{
        "BIZ_batch_cutoff_time": "09:00:00",
        "BIZ_batch_dcn_type": "OW-DN-CM",
        "BIZ_batch_description": "ACCOUNTING",
        "BIZ_batch_end_time": "10:00:00",
        "BIZ_batch_service_path": "XXX",
        "BIZ_batch_service_topic": "XXX",
        "BIZ_batch_service_type": "XXX",
        "BIZ_batch_start_time": "00:00:00",
        "BIZ_local_file_flag": "X",
        "BIZ_local_sftp_ip": "XX",
        "BIZ_local_sftp_path": "XX",
        "BIZ_local_sftp_port": "XX",
        "BIZ_rec_mail_addr_1": "XX",
        "BIZ_rec_mail_addr_2": "XX",
        "BIZ_rec_mail_addr_3": "XX",
        "BIZ_rec_mail_addr_4": "XX",
        "BIZ_rec_mail_addr_5": "XX",
        "BIZ_snd_mail_addr": "XX",
        "BIZ_snd_mail_agm": "XX",
        "BIZ_snd_mail_flag": "X",
        "BIZ_snd_mail_ip": "XX",
        "BIZ_snd_mail_name": "XX",
        "BIZ_snd_mail_port": "XX",
        "BIZ_snd_mail_pwd": "XX",
        "BIZ_target_sftp_acc": "XX",
        "BIZ_target_sftp_flag": "X",
        "BIZ_target_sftp_ip": "XX",
        "BIZ_target_sftp_path": "/data",
        "BIZ_target_sftp_port": "XX",
        "BIZ_target_sftp_pwd": "XX"
    }
2. CM001020为下载文件到dos 传入参数为
{
        "BIZ_target_sftp_path": "/data",
        "BIZ_target_sftp_file_name":"front.yaml"
    }
返回为dos的fileId, version:
{
        "BIZ_target_sftp_file_id": "front.yaml",
        "BIZ_target_sftp_file_name": "front.yaml",
        "BIZ_target_sftp_file_version": "",
        "BIZ_target_sftp_path": "/data"
    }
3. CM001021 为上传文件到sftp 传入参数为：
{
        "BIZ_target_sftp_path": "/data/han",
        "BIZ_target_sftp_file_name":"front.yaml",
        "BIZ_target_sftp_file_id":"front.yaml",
        "BIZ_target_sftp_file_version": ""
    }
成功返回入参，失败报错
*/
