package models

import "encoding/xml"

type SMP1000025I struct {
	RequestDate   string `json:"requestDate"  validate:"required"`
	RequestNumber string `json:"requestNumber"  validate:"required"`
	RequestTime   string `json:"requestTime"  validate:"required"`
	SystemRefNo   string `json:"systemRefNo"  validate:"required"`
	TemplateId    string `json:"templateId"  validate:"required"`
	SendType      string `json:"sendType"  validate:"required"`
	Lang          string `json:"lang"  validate:"required"`
	ServiceMode   string `json:"serviceMode"  validate:"required"`
	Username      string `json:"username"  validate:"required"`
	Password      string `json:"password"  validate:"required"`
}
type SMP1000025O struct {
	ResponseDate  string `json:"responsedate" xml:"responsedate"`
	ResponseTime  string `json:"responsetime" xml:"responsetime"`
	RequestNumber string `json:"requestnumber" xml:"requestnumber"`
	Service       string `json:"service" xml:"service"`
	SuccessCode   string `json:"successcode" xml:"successcode"`
	ServiceMode   string `json:"serviceMode" xml:"serviceMode"`
	SystemRefNo   string `json:"systemrefno" xml:"systemrefno"`
	OtpPrefix     string `json:"otpprefix" xml:"otpprefix"`
	SendType      string `json:"sendtype" xml:"sendtype"`
	Retry         int    `json:"retry" xml:"retry"`
}

func (i *SMP1000025I) GetServiceKey() string {
	return "MP100025"
}

type SMP3000025I struct {
	XMLName xml.Name `xml:"soap:Envelope"`
	Text    string   `xml:",chardata"`
	Soap    string   `xml:"xmlns:soap,attr"`
	Otp     string   `xml:"xmlns:otp,attr"`
	Header  string   `xml:"soap:Header"`
	Body    struct {
		Text    string `xml:",chardata"`
		Request struct {
			Text          string `xml:",chardata"`
			RequestDate   string `xml:"otp:RequestDATE"`
			RequestTime   string `xml:"otp:RequestTIME"`
			RequestNumber string `xml:"otp:RequestNUMBER"`
			Username      string `xml:"otp:Username"`
			Password      string `xml:"otp:Password"`
			MerchantName  string `xml:"otp:MerchantName"`
			Amount        string `xml:"otp:Amount"`
			Currency      string `xml:"otp:Currency"`
			CardNumber    string `xml:"otp:CardNumber"`
			SystemRefNo   string `xml:"otp:Systemrefno"`
			SendType      string `xml:"otp:Sendtype"`
			Lang          string `xml:"otp:Lang"`
			TemplateId    string `xml:"otp:TempleteId"`
			ServiceMode   string `xml:"otp:ServiceMode"`
		} `xml:"otp:resendOtp"`
	} `xml:"soap:Body"`
}

type SMP3000025O struct {
	XMLName xml.Name `xml:"Envelope"`
	Text    string   `xml:",chardata"`
	Soapenv string   `xml:"soapenv,attr"`
	Body    struct {
		Text              string `xml:",chardata"`
		ResendOtpResponse struct {
			Text   string `xml:",chardata"`
			Ns     string `xml:"ns,attr"`
			Return struct {
				Text  string `xml:",chardata"`
				Otpws struct {
					Text string `xml:",chardata"`
				} `xml:"otpws"`
			} `xml:"return"`
		} `xml:"resendOtpResponse"`
	} `xml:"Body"`
}

func (w *SMP3000025I) GetServiceKey() string {
	return "MP300025"
}

type SMP3000025Otpws struct {
	Text string `xml:",chardata"`
	Otp  struct {
		Text           string `xml:",chardata"`
		Responsedate   string `xml:"responsedate"`
		Responsetime   string `xml:"responsetime"`
		Requestnumber  string `xml:"requestnumber"`
		Service        string `xml:"service"`
		Successcode    string `xml:"successcode"`
		Systemrefno    string `xml:"systemrefno"`
		Otpprefix      string `xml:"otpprefix"`
		Retry          string `xml:"retry"`
		Sendtype       string `xml:"sendtype"`
		Sendto         string `xml:"sendto"`
		Lang           string `xml:"lang"`
		Templeteid     string `xml:"templeteid"`
		Generationtime string `xml:"generationtime"`
		Expirytime     string `xml:"expirytime"`
	} `xml:"otp"`
}
