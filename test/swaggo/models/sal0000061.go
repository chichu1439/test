package models

type AL000061I struct {
	Id         string `json:"id" description:"Id" validate:"required"`
	HashValue  string `json:"hashValue" description:"Hash value" validate:"required"`
	RandomCode string `json:"randomCode" description:"Random code"`
}

type AL000061O struct {
	HashValue    string       `json:"hashValue" description:"Hash value"`
	OrderId      string       `json:"orderId" description:"Order id"`
	Version      string       `json:"version" description:"Version"`
	MerchantId   string       `json:"merchantId" description:"Merchant id"`
	Currency     string       `json:"currency" description:"Currency"`
	ResultUrl1   string       `json:"resultUrl1" description:"Result url1"`
	ResultUrl2   string       `json:"resultUrl2" description:"Result url2"`
	Amount       string       `json:"amount" description:"Amount"`
	TokenRspData TokenRspData `json:"tokenRspData" description:"Token response data"`
}

type TokenRspData struct {
	WebPaymentUrl string `json:"webPaymentUrl" description:"Web payment url"`
	PaymentToken  string `json:"paymentToken" description:"Payment token"`
	RespCode      string `json:"respCode" description:"Response code"`
	RespDesc      string `json:"respDesc" description:"Response desc"`
}

func (*AL000061I) GetServiceKey() string {
	return "AL000061"
}
