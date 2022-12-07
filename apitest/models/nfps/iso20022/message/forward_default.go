package message

import (
	"encoding/xml"

	"apitest/models/iso20022/head_v01"
)

type FpsForwardDefault struct {
	XMLName xml.Name       `xml:"FpsMsg"`
	Attrs   []xml.Attr     `xml:",any,attr,omitempty" json:",omitempty"`
	BizData ForwardBizData `xml:"BizData"`
}

type ForwardBizData struct {
	XmlHead *head_v01.BusinessApplicationHeaderV01 `xml:"AppHdr"`
	XmlDoc  XmlDoc                                 `xml:"Document"`
}
type XmlDoc struct {
	Attrs  []xml.Attr `xml:",any,attr,omitempty" json:",omitempty"`
	XmlDoc string     `xml:",innerxml"`
}
type FpsForwardDefault2 struct {
	XMLName  xml.Name       `xml:"FpsMsg"`
	NbOfMsgs string         `xml:"NbOfMsgs,omitempty"`
	Attrs    []xml.Attr     `xml:",any,attr,omitempty" json:",omitempty"`
	BizData  ForwardBizData `xml:"FpsPylds>BizData"`
}
