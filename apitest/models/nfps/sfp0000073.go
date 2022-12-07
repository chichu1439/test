package models

import "github.com/go-playground/validator/v10"

type FP000073I struct {
	FileName    string `json:"fileName" validate:"required"`
	FileContent string `json:"fileContent" validate:"required"`
}

type FP000073O struct {
	Status string `json:"status"`
}

func (*FP000073I) GetServiceKey() string {
	return "FP000073"
}

func (d *FP000073I) Validate() error {
	validate := validator.New()
	return validate.Struct(d)
}

type FP000074I struct {
	PartyCode   string `json:"partyCode"`
	FileDate    string `json:"fileDate"`
	FileName    string `json:"fileName"`
	MessageType string `json:"messageType"`
	FileStatus  string `json:"fileStatus"`
	PageNum     int    `json:"pageNum"`
	PageSize    int    `json:"pageSize"`
}

type FP000074O struct {
	Data       []SendEnquiryList `json:"data"`
	TotalCount int64             `json:"totalCount"`
}

type FileEnquiryList struct {
	Id               int64  `json:"id"`
	FileName         string `json:"fileName"`
	FileRecveiveDate int64  `json:"fileRecveiveDate"`
	FileSendDateTime int64  `json:"fileSendDateTime"`
	FileType         string `json:"fileType"`
	MessageType      string `json:"messageType"`
	FileStatus       string `json:"fileStatus"`
}

func (*FP000074I) GetServiceKey() string {
	return "FP000074"
}

func (d *FP000074I) Validate() error {
	validate := validator.New()
	return validate.Struct(d)
}

type FP000075I struct {
	FileName string `json:"fileName"`
	BankCode string `json:"bankCode"`
	FileType string `json:"fileType"`
}

type FP000075O struct {
	Data []FileInfo `json:"data"`
}

type FileInfo struct {
	FileName       string `json:"fileName"`
	BankCode       string `json:"bankCode"`
	FileType       string `json:"fileType"`
	FilePath       string `json:"filePath"`
	FileSize       int64  `json:"fileSize"`
	FileCreateTime int64  `json:"fileCreateTime"`
}

type FP000076I struct {
	FileName string `json:"fileName" validate:"required"`
	BankCode string `json:"bankCode" validate:"required"`
	FilePath string `json:"filePath" validate:"required"`
}

type FP000076O struct {
	Status string `json:"status"`
}

func (d *FP000076I) Validate() error {
	validate := validator.New()
	return validate.Struct(d)
}

type FP000077I struct {
	PartClearingCode string `json:"partClearingCode" validate:"required"`
	Currency         string `json:"currency"`
	PageNo           int    `json:"pageNo"`       //查询页数
	PageRecCount     int    `json:"pageRecCount"` //每页记录数
}

type FP000077O struct {
	TotalRecCount int             `json:"totalRecCount"`
	TotalPage     int             `json:"totalPage"`
	Records       []FP000077OItem `json:"records"`
}

type FP000077OItem struct {
	PartClearingCode string  `json:"partClearingCode"`
	PartName         string  `json:"partName"`
	PartType         string  `json:"partType"`
	Currency         string  `json:"currency"`
	Balance          float64 `json:"balance"`
	AlertBalance     float64 `json:"alertBalance"`
	AlertMinuteTime  int     `json:"alertMinuteTime"`
	Status           string  `json:"status"`
	HightWaterMark   float64 `json:"hightWaterMark"`
	LowWaterMark     float64 `json:"lowWaterMark"`
}

func (d *FP000077I) Validate() error {
	validate := validator.New()
	return validate.Struct(d)
}
func (*FP000077I) GetServiceKey() string {
	return "FP000077"
}

type FP000078I struct {
	PartClearingCode string   `json:"partClearingCode"`
	AlertBalance     *float64 `json:"alertBalance"`
	AlertMinuteTime  *int     `json:"alertMinuteTime"`
	Currency         string   `json:"currency"`
	HightWaterMark   *float64 `json:"hightWaterMark"`
	LowWaterMark     *float64 `json:"lowWaterMark"`
}

type FP00078O struct {
	AlertBalance   float64 `json:"alertBalance"`
	HightWaterMark float64 `json:"hightWaterMark"`
	LowWaterMark   float64 `json:"lowWaterMark"`
}

func (d *FP000078I) Validate() error {
	validate := validator.New()
	return validate.Struct(d)
}
func (*FP000078I) GetServiceKey() string {
	return "FP000078"
}

type FP000079I struct {
	FileType string `json:"fileType"`
	FileName string `json:"fileName"`
}

type FP00079O struct {
}

func (d *FP000079I) Validate() error {
	validate := validator.New()
	return validate.Struct(d)
}
func (*FP000079I) GetServiceKey() string {
	return "FP000079"
}
