package message

import (
	"apitest/models/nfps/iso20022/acmt_v03"
	"encoding/xml"

	"apitest/models/nfps/iso20022/head_v01"
	"apitest/util"
)

type Acmt023BizData struct {
	XmlHead *head_v01.BusinessApplicationHeaderV01 `xml:"AppHdr"`
	XmlDoc  FpsAcmt023Doc                          `xml:"Document"`
}

func (p Acmt023BizData) Validate() error {
	return util.Validate(&p)
}

type FpsAcmt023Doc struct {
	XMLName xml.Name   `xml:"Document"`
	Attrs   []xml.Attr `xml:",any,attr,omitempty" json:",omitempty"`
	XmlDoc  *acmt_v03.IdentificationVerificationRequestV03
}

func (f FpsAcmt023Doc) Validate() error {
	return util.Validate(&f)
}

type FpsAcmt023 struct {
	XMLName xml.Name       `xml:"FpsMsg"`
	Attrs   []xml.Attr     `xml:",any,attr,omitempty" json:",omitempty"`
	BizData Acmt023BizData `xml:"BizData"`
}

func (f FpsAcmt023) Validate() error {
	return util.Validate(&f)
}
