package models

type FP900083I struct {
	PartyCode   string `json:"partyCode"`
	FileDate    string `json:"fileDate"`
	FileName    string `json:"fileName"`
	TypeId      string `json:"typeId"`
	MessageType string `json:"messageTypeType"`
	FileStatus  string `json:"fileStatus"`
	PageNum     int    `json:"pageNum"`
	PageSize    int    `json:"pageSize"`
}

type FP900083O struct {
	//Data       []FileEnquiryList `json:"data"`
	SendData   []SendEnquiryList `json:"sendData"`
	TotalCount int64             `json:"totalCount"`
}

func (*FP900083I) GetServiceKey() string {
	return "FP900083"
}

type SendEnquiryList struct {
	Id               int64            `json:"id"`
	FileName         string           `json:"fileName"`
	FileRecveiveDate int64            `json:"fileRecveiveDate"`
	FileSendDateTime int64            `json:"fileSendDateTime"`
	FileType         string           `json:"fileType"`
	MessageType      string           `json:"messageType"`
	FileStatus       string           `json:"fileStatus"`
	Children         []RevEnquiryList `json:"children"`
}

type RevEnquiryList struct {
	Id               int64        `json:"id"`
	FileName         string       `json:"fileName"`
	FileRecveiveDate int64        `json:"fileRecveiveDate"`
	FileSendDateTime int64        `json:"fileSendDateTime"`
	FileType         string       `json:"fileType"`
	MessageType      string       `json:"messageType"`
	FileStatus       string       `json:"fileStatus"`
	Children         []RevEnquiry `json:"children"`
}

type RevEnquiry struct {
	Id               int64  `json:"id"`
	FileName         string `json:"fileName"`
	FileRecveiveDate int64  `json:"fileRecveiveDate"`
	FileSendDateTime int64  `json:"fileSendDateTime"`
	FileType         string `json:"fileType"`
	MessageType      string `json:"messageType"`
	FileStatus       string `json:"fileStatus"`
}
