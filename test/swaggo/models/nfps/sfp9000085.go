package models

import "time"

type FP900085I struct {
	Data []FileMessage `json:"data"`
}

type FileMessage struct {
	FileName           string    `json:"file_name"`
	FileSize           int64     `json:"file_size"`
	FileCreateTime     time.Time `json:"file_create_time"`
	MessageArrivalTime time.Time `json:"message_arrival_time"`
	FileType           string    `json:"file_type"`
	FileStatus         string    `json:"file_status"`
	FileContentHash    string    `json:"file_content_hash"`
	BankCode           string    `json:"bank_code"`
	Remark             string    `json:"remark"`
	InFileName         string    `json:"inFileName"`
	MessageType        string    `json:"message_type"`
	Relation           string    `json:"relation"`
}

type FP900085O struct {
	Status string `json:"status"`
}

func (*FP900085I) GetServiceKey() string {
	return "FP900085"
}
