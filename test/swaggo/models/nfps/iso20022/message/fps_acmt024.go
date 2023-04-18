package message

import (
	"encoding/xml"
	"isaving-swagger/models/nfps/iso20022/acmt_v03"

	"isaving-swagger/models/nfps/iso20022/head_v01"
	"isaving-swagger/util"
)

type Acmt024BizData struct {
	XmlHead *head_v01.BusinessApplicationHeaderV01 `xml:"AppHdr"`
	XmlDoc  FpsAcmt024Doc                          `xml:"Document"`
}

func (p Acmt024BizData) Validate() error {
	return util.Validate(&p)
}

type FpsAcmt024Doc struct {
	XMLName xml.Name   `xml:"Document"`
	Attrs   []xml.Attr `xml:",any,attr,omitempty" json:",omitempty"`
	XmlDoc  *acmt_v03.IdentificationVerificationReportV03
}

func (f FpsAcmt024Doc) Validate() error {
	return util.Validate(&f)
}

type FpsAcmt024 struct {
	XMLName xml.Name       `xml:"FpsMsg"`
	Attrs   []xml.Attr     `xml:",any,attr,omitempty" json:",omitempty"`
	BizData Acmt024BizData `xml:"BizData"`
}

func (f FpsAcmt024) Validate() error {
	return util.Validate(&f)
}
