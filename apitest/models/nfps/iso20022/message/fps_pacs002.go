package message

import (
	"encoding/xml"

	"apitest/models/iso20022/head_v01"
	"apitest/models/iso20022/pacs_v08"
	"apitest/util"
)

type Pacs002BizData struct {
	XmlHead *head_v01.BusinessApplicationHeaderV01 `xml:"AppHdr"`
	XmlDoc  FpsPacs002Doc                          `xml:"Document"`
}

func (p Pacs002BizData) Validate() error {
	return util.Validate(&p)
}

type FpsPacs002Doc struct {
	XMLName xml.Name   `xml:"Document"`
	Attrs   []xml.Attr `xml:",any,attr,omitempty" json:",omitempty"`
	XmlDoc  *pacs_v08.FIToFIPaymentStatusReportV08
}

func (f FpsPacs002Doc) Validate() error {
	return util.Validate(&f)
}

type FpsPacs002 struct {
	XMLName xml.Name       `xml:"FpsMsg"`
	Attrs   []xml.Attr     `xml:",any,attr,omitempty" json:",omitempty"`
	BizData Pacs002BizData `xml:"BizData"`
}

func (f FpsPacs002) Validate() error {
	return util.Validate(&f)
}
