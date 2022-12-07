package message

import (
	"apitest/models/iso20022/pain_v07"
	"encoding/xml"

	"apitest/models/iso20022/head_v01"
	"apitest/util"
)

type Pain014BizData struct {
	XmlHead *head_v01.BusinessApplicationHeaderV01 `xml:"AppHdr"`
	XmlDoc  FpsPain014Doc                          `xml:"Document"`
}

func (p Pain014BizData) Validate() error {
	return util.Validate(&p)
}

type FpsPain014Doc struct {
	XMLName xml.Name   `xml:"Document"`
	Attrs   []xml.Attr `xml:",any,attr,omitempty" json:",omitempty"`
	XmlDoc  *pain_v07.CreditorPaymentActivationRequestStatusReportV07
}

func (f FpsPain014Doc) Validate() error {
	return util.Validate(&f)
}

type FpsPain014 struct {
	XMLName xml.Name       `xml:"FpsMsg"`
	Attrs   []xml.Attr     `xml:",any,attr,omitempty" json:",omitempty"`
	BizData Pain014BizData `xml:"BizData"`
}

func (f FpsPain014) Validate() error {
	return util.Validate(&f)
}
