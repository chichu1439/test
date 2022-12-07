package models

import "encoding/xml"

type SMP1000024I struct {
	RequestDate   string `json:"requestDate"`
	RequestTime   string `json:"requestTime"`
	SystemRefNo   string `json:"systemRefNo"`
	ServiceMode   string `json:"serviceMode"`
	Status        string `json:"status"`
	Otp           string `json:"otp"`
	Username      string `json:"username"`
	Password      string `json:"password"`
	RequestNumber string `json:"requestNumber"`
	LogTypeId     string `json:"logTypeId"`
}

type SMP1000024O struct {
	ResponseDate  string `json:"responsedate" xml:"responsedate"`
	ResponseTime  string `json:"responsetime" xml:"responsetime"`
	RequestNumber string `json:"requestnumber" xml:"requestnumber"`
	ServiceMode   string `json:"serviceMode" xml:"service"`
	SuccessCode   string `json:"successcode" xml:"successcode"`
	SystemRefNo   string `json:"systemrefno" xml:"systemrefno"`
}

func (i *SMP1000024I) GetServiceKey() string {
	return "MP100024"
}

type SMP3000024I struct {
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
			SystemRefNo   string `xml:"otp:Systemrefno"`
			ServiceMode   string `xml:"otp:ServiceMode"`
			Status        string `xml:"otp:Status"`
			Otp           string `xml:"otp:OTP"`
		} `xml:"otp:verifyOtp"`
	} `xml:"soap:Body"`
}

type SMP3000024O struct {
	XMLName xml.Name `xml:"Envelope"`
	Text    string   `xml:",chardata"`
	Soapenv string   `xml:"soapenv,attr"`
	Body    struct {
		Text              string `xml:",chardata"`
		VerifyOtpResponse struct {
			Text   string `xml:",chardata"`
			Ns     string `xml:"ns,attr"`
			Return struct {
				Text  string `xml:",chardata"`
				Otpws struct {
					Text string `xml:",chardata"`
				} `xml:"otpws"`
			} `xml:"return"`
		} `xml:"verifyOtpResponse"`
	} `xml:"Body"`
}

func (w *SMP3000024I) GetServiceKey() string {
	return "MP300024"
}

type SMP3000024Otpws struct {
	XMLName xml.Name `xml:"otpws"`
	Text    string   `xml:",chardata"`
	Otp     struct {
		Text          string `xml:",chardata"`
		Service       string `xml:"service"`
		Responsedate  string `xml:"responsedate"`
		Responsetime  string `xml:"responsetime"`
		Requestnumber string `xml:"requestnumber"`
		Successcode   string `xml:"successcode"`
		Systemrefno   string `xml:"systemrefno"`
	} `xml:"otp"`
}
