package models

type FP900086I struct {
	FileName        string `json:"fileName"`
	StartTime       int64  `json:"startTime"`
	EndTime         int64  `json:"endTime"`
	FileType        string `json:"file_type"`
	FileStatus      string `json:"file_status"`
	FileContentHash string `json:"file_content_hash"`
	BankCode        string `json:"bank_code"`
	PageSize        int    `json:"page_size"`
	PageNum         int    `json:"page_num"`
}

type FP900086O struct {
	Data       []FileInfoData
	TotalCount int64 `json:"totalCount"`
}

type FileInfoData struct {
	Id                 int64  `json:"id"`
	FileName           string `json:"file_name"`
	FileSize           int64  `json:"file_size"`
	FileCreateTime     int64  `json:"file_create_time"`
	MessageArrivalTime int64  `json:"message_arrival_time"`
	FileType           string `json:"file_type"`
	FileStatus         string `json:"file_status"`
	FileContentHash    string `json:"file_content_hash"`
	BankCode           string `json:"bank_code"`
	Remark             string `json:"remark"`
}

func (*FP900086I) GetServiceKey() string {
	return "FP900086"
}
