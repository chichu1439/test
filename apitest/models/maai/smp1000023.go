package models

import (
	"encoding/xml"
)

type SMP1000023I struct {
	RequestDate   string `json:"requestDate"`
	RequestTime   string `json:"requestTime"`
	PhoneNo       string `json:"phoneNo"`
	Email         string `json:"email"`
	TemplateId    string `json:"templateId"`
	SendType      string `json:"sendtype"`
	Lang          string `json:"lang"`
	ServiceMode   string `json:"serviceMode"`
	Username      string `json:"username"`
	Password      string `json:"password"`
	RequestNumber string `json:"requestNumber"`
}
type SMP1000023O struct {
	ResponseDate   string `json:"responsedate"`
	ResponseTime   string `json:"responsetime"`
	RequestNumber  string `json:"requestnumber"`
	Service        string `json:"service"`
	SuccessCode    string `json:"successcode"`
	SystemRefNo    string `json:"systemrefno"`
	OtpPrefix      string `json:"otpprefix"`
	SendType       string `json:"sendtype"`
	PhoneNo        string `json:"phoneno"`
	Email          string `json:"email"`
	Lang           string `json:"lang"`
	TemplateId     string `json:"templateid"`
	GenerationTime string `json:"generationtime"`
	ExpiryTime     string `json:"expirytime"`
}

func (i *SMP1000023I) GetServiceKey() string {
	return "MP100023"
}

type SMP3000023I struct {
	XMLName xml.Name `xml:"soap:Envelope"`
	Text    string   `xml:",chardata"`
	Soap    string   `xml:"xmlns:soap,attr"`
	Otp     string   `xml:"xmlns:otp,attr"`
	Header  string   `xml:"soap:Header"`
	Body    struct {
		Text                 string `xml:",chardata"`
		GenerateOtpbyPhoneNo struct {
			Text          string `xml:",chardata"`
			RequestDATE   string `xml:"otp:RequestDATE"`
			RequestTIME   string `xml:"otp:RequestTIME"`
			RequestNUMBER string `xml:"otp:RequestNUMBER"`
			Username      string `xml:"otp:Username"`
			Password      string `xml:"otp:Password"`
			PhoneNo       string `xml:"otp:PhoneNo"`
			Email         string `xml:"otp:Email"`
			TempleteId    string `xml:"otp:TempleteId"`
			Sendtype      string `xml:"otp:Sendtype"`
			Lang          string `xml:"otp:Lang"`
			ServiceMode   string `xml:"otp:ServiceMode"`
		} `xml:"otp:generateOtpbyPhoneNo"`
	} `xml:"soap:Body"`
}

type SMP3000023O struct {
	XMLName xml.Name `xml:"Envelope"`
	Text    string   `xml:",chardata"`
	Soapenv string   `xml:"soapenv,attr"`
	Body    struct {
		Text                         string `xml:",chardata"`
		GenerateOtpbyPhoneNoResponse struct {
			Text   string `xml:",chardata"`
			Ns     string `xml:"ns,attr"`
			Return struct {
				Text  string `xml:",chardata"`
				Otpws struct {
					Text string `xml:",chardata"`
				} `xml:"otpws"`
			} `xml:"return"`
		} `xml:"generateOtpbyPhoneNoResponse"`
	} `xml:"Body"`
}

func (i *SMP3000023I) GetServiceKey() string {
	return "MP300023"
}

type SMP3000023Otpws struct {
	XMLName xml.Name `xml:"otpws"`
	Text    string   `xml:",chardata"`
	Otp     struct {
		Text string `xml:",chardata"`
		SMP1000023O
	} `xml:"otp"`
}
