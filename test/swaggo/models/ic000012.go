package models

type IC000012I struct {
	CustomerId  string `json:"customerId" description:"Customer id" validate:"required"` // 客户号
	IsSendEmail string `json:"isSendEmail" description:"Is send email:Y-send N-no send"` // Y- send N-no send
}

type IC000012O struct {
	CustomerId  string `json:"customerId" description:"Customer id"`
	PdfFileName string `json:"pdfFileName" description:"Pdf file name"`
	PdfFileId   string `json:"pdfFileId" description:"Pdf file id"`
}

func (*IC000012I) GetServiceKey() string {
	return "ic000012"
}
