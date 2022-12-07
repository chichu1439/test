package message

import (
	"apitest/models/iso20022/pain_v07"
	"encoding/xml"

	"apitest/models/iso20022/head_v01"
	"apitest/util"
)

type Pain013BizData struct {
	XmlHead *head_v01.BusinessApplicationHeaderV01 `xml:"AppHdr"`
	XmlDoc  FpsPain013Doc                          `xml:"Document"`
}

func (p Pain013BizData) Validate() error {
	return util.Validate(&p)
}

type FpsPain013Doc struct {
	XMLName xml.Name   `xml:"Document"`
	Attrs   []xml.Attr `xml:",any,attr,omitempty" json:",omitempty"`
	XmlDoc  *pain_v07.CreditorPaymentActivationRequestV07
}

func (f FpsPain013Doc) Validate() error {
	return util.Validate(&f)
}

type FpsPain013 struct {
	XMLName xml.Name       `xml:"FpsMsg"`
	Attrs   []xml.Attr     `xml:",any,attr,omitempty" json:",omitempty"`
	BizData Pain013BizData `xml:"BizData"`
}

func (f FpsPain013) Validate() error {
	return util.Validate(&f)
}
