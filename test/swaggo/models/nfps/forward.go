package models

import (
	"encoding/xml"
)

type FpsForwardDefault struct {
	XMLName xml.Name       `xml:"FpsMsg"`
	Attrs   []xml.Attr     `xml:",any,attr,omitempty" json:",omitempty"`
	BizData ForwardBizData `xml:"BizData"`
}

type ForwardBizData struct {
	XmlHead Head001 `xml:"AppHdr"`
	XmlDoc  XmlDoc  `xml:"Document"`
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
