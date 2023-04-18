package message

import (
	"encoding/xml"
	"fmt"

	"isaving-swagger/models/nfps/iso20022"
	"isaving-swagger/models/nfps/iso20022/camt_v06"
	"isaving-swagger/models/nfps/iso20022/camt_v07"
	"isaving-swagger/models/nfps/iso20022/head_v01"
	"isaving-swagger/models/nfps/iso20022/head_v02"
	"isaving-swagger/models/nfps/iso20022/pacs_v06"
	"isaving-swagger/models/nfps/iso20022/pacs_v07"
	"isaving-swagger/models/nfps/iso20022/pacs_v08"
	"isaving-swagger/models/nfps/iso20022/pacs_v09"
	"isaving-swagger/models/nfps/iso20022/pacs_v10"
	"isaving-swagger/models/nfps/iso20022/pacs_v11"
	"isaving-swagger/models/nfps/iso20022/pain_v05"
	"isaving-swagger/models/nfps/iso20022/pain_v07"
	"isaving-swagger/models/nfps/iso20022/pain_v08"
	"isaving-swagger/util"
)

const (
	XmlDefaultNamespace = "xmlns"
	FpsNamespace        = "urn:CentralBank:fps:xsd:fps.envelope.01"
)

type messageDummy struct {
	XmlName xml.Name   `xml:"FpsMsg"`
	Attrs   []xml.Attr `xml:",any,attr,omitempty" json:",omitempty"`
	XmlHead headDummy  `xml:"BizData>AppHdr"`
	XmlDoc  docDummy   `xml:"BizData>Document"`
}

type headDummy struct {
	XmlName xml.Name   `xml:"AppHdr"`
	Attrs   []xml.Attr `xml:",any,attr,omitempty" json:",omitempty"`
}

type docDummy struct {
	XmlName xml.Name   `xml:"Document"`
	Attrs   []xml.Attr `xml:",any,attr,omitempty" json:",omitempty"`
}

func (d messageDummy) NameSpace() string {

	for _, attr := range d.Attrs {
		if attr.Name.Space == XmlDefaultNamespace || attr.Name.Local == XmlDefaultNamespace {
			return attr.Value
		}
	}

	return ""
}

func (d messageDummy) HeadNameSpace() string {

	for _, attr := range d.XmlHead.Attrs {
		if attr.Name.Space == XmlDefaultNamespace || attr.Name.Local == XmlDefaultNamespace {
			return attr.Value
		}
	}

	return ""
}

func (d messageDummy) DocNameSpace() string {

	for _, attr := range d.XmlDoc.Attrs {
		if attr.Name.Space == XmlDefaultNamespace || attr.Name.Local == XmlDefaultNamespace {
			return attr.Value
		}
	}

	return ""
}

func NewFpsMessage(headNamespace, docNamespace string) (msg FpsCombineMessage, err error) {

	headConstructor := headMessageConstructor[headNamespace]
	if headConstructor == nil {
		return nil, fmt.Errorf("the namespace %s of head is not support", headNamespace)
	}
	docConstructor := docMessageConstructor[docNamespace]
	if docConstructor == nil {
		return nil, fmt.Errorf("the namespace %s of document is not support", headNamespace)
	}

	return &FpsCombineMessageObject{
		BizData: FpsBizData{
			XmlHead: headConstructor(),
			XmlDoc: FpsDoc{
				XmlDoc: docConstructor(),
			},
		},
	}, nil
}

type FpsCombineMessage interface {
	// Validate will be process validation check of document
	Validate() error

	// NameSpace check will return xmlns of document
	NameSpace() string

	// GetXmlName returns xml name of document
	GetXmlName() *xml.Name

	// GetAttrs returns attributes of document
	GetAttrs() []xml.Attr

	// HeadMessage return header
	HeadMessage() Iso20022Message

	// DocMessage returns document
	DocMessage() Iso20022Message
}

// Iso20022Message is element interface for ISO 20022
type Iso20022Message interface {
	// Validate will be process validation check of document

	Validate() error
}

type constructorFunc func() Iso20022Message

var (
	headMessageConstructor = map[string]constructorFunc{
		iso20022.DocumentHead00100101NameSpace: func() Iso20022Message { return &head_v01.BusinessApplicationHeaderV01{} },
		iso20022.DocumentHead00100102NameSpace: func() Iso20022Message { return &head_v02.BusinessApplicationHeaderV02{} },
	}
	docMessageConstructor = map[string]constructorFunc{
		iso20022.DocumentPacs00800106NameSpace: func() Iso20022Message { return &pacs_v06.FIToFICustomerCreditTransferV06{} },
		iso20022.DocumentPacs00800108NameSpace: func() Iso20022Message { return &pacs_v08.FIToFICustomerCreditTransferV08{} },
		iso20022.DocumentPacs00800109NameSpace: func() Iso20022Message { return &pacs_v09.FIToFICustomerCreditTransferV09{} },

		iso20022.DocumentPacs00200107NameSpace: func() Iso20022Message { return &pacs_v07.FIToFIPaymentStatusReportV07{} },
		iso20022.DocumentPacs00200108NameSpace: func() Iso20022Message { return &pacs_v08.FIToFIPaymentStatusReportV08{} },
		iso20022.DocumentPacs00200110NameSpace: func() Iso20022Message { return &pacs_v10.FIToFIPaymentStatusReportV10{} },
		iso20022.DocumentPacs00200111NameSpace: func() Iso20022Message { return &pacs_v11.FIToFIPaymentStatusReportV11{} },

		iso20022.DocumentCamt02300107NameSpace: func() Iso20022Message { return &camt_v07.BackupPaymentV07{} },
		iso20022.DocumentCamt02400106NameSpace: func() Iso20022Message { return &camt_v06.ModifyStandingOrderV06{} },

		iso20022.DocumentPain01300105NameSpace: func() Iso20022Message { return &pain_v05.CreditorPaymentActivationRequestV05{} },
		iso20022.DocumentPain01300107NameSpace: func() Iso20022Message { return &pain_v07.CreditorPaymentActivationRequestV07{} },
		iso20022.DocumentPain01300108NameSpace: func() Iso20022Message { return &pain_v08.CreditorPaymentActivationRequestV08{} },
	}
)

type FpsDoc struct {
	XmlName xml.Name        `xml:"Document"`
	Attrs   []xml.Attr      `xml:",any,attr,omitempty" json:",omitempty"`
	XmlDoc  Iso20022Message `xml:",any"`
}

type FpsBizData struct {
	XmlHead Iso20022Message `xml:"AppHdr"`
	XmlDoc  FpsDoc          `xml:"Document"`
}

type FpsCombineMessageObject struct {
	XmlName        xml.Name   `xml:"FpsMsg"`
	Attrs          []xml.Attr `xml:",any,attr,omitempty" json:",omitempty"`
	XmlCombineName xml.Name   `xml:",any"`
	XmlHead        xml.Name   `xml:">AppHdr"`
	BizData        FpsBizData `xml:"BizData"`
}

func (f FpsCombineMessageObject) Validate() error {
	if len(f.NameSpace()) == 0 {
		return util.Validate(&f)
	}

	for _, attr := range f.Attrs {
		if attr.Name.Local == XmlDefaultNamespace && f.NameSpace() == attr.Value {
			return util.Validate(&f)
		}
	}

	return util.NewErrInvalidNameSpace()
}

func (f FpsCombineMessageObject) NameSpace() string {
	for _, attr := range f.Attrs {
		if attr.Name.Space == XmlDefaultNamespace || attr.Name.Local == XmlDefaultNamespace {
			return attr.Value
		}
	}
	return ""
}

func (f FpsCombineMessageObject) GetXmlName() *xml.Name {
	return &f.XmlName
}

func (f FpsCombineMessageObject) GetAttrs() []xml.Attr {
	return f.Attrs
}

func (f FpsCombineMessageObject) HeadMessage() Iso20022Message {
	return f.BizData.XmlHead
}

func (f FpsCombineMessageObject) DocMessage() Iso20022Message {
	return f.BizData.XmlDoc.XmlDoc
}

/*func (f FpsCombineMessageObject) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	a := struct {
		XmlName xml.Name   `xml:"FpsMsg"`
		Attrs   []xml.Attr `xml:",any,attr,omitempty" json:",omitempty"`
		BizData FpsBizData `xml:"BizData"`
	}(f)

	updatingStartElement(&start, f.Attrs, f.XmlName)
	return e.EncodeElement(&a, start)
}

func updatingStartElement(start *xml.StartElement, attrs []xml.Attr, name xml.Name) {
	for _, attr := range attrs {
		if attr.Name.Local == XmlDefaultNamespace {
			name.Space = ""
		}
	}
	if len(name.Local) > 0 {
		start.Name.Local = name.Local
	}
	start.Name.Space = name.Space
}*/

func ParseFpsCombineMessage(buf []byte) (FpsCombineMessage, error) {

	var dummy messageDummy
	var err error

	err = xml.Unmarshal(buf, &dummy)
	if err != nil {
		return nil, err
	}

	namespace := dummy.NameSpace()
	if namespace == "" {
		return nil, fmt.Errorf("not found fps namespace in fps message")
	}

	if namespace != FpsNamespace {
		return nil, fmt.Errorf("not found fps namespace in fps message")
	}
	headNamespace := dummy.HeadNameSpace()
	if headNamespace == "" {
		return nil, fmt.Errorf("not found head namespace in fps message")
	}
	docNamespace := dummy.DocNameSpace()
	if docNamespace == "" {
		return nil, fmt.Errorf("not found document namespace in fps message")
	}

	headConstructor := headMessageConstructor[headNamespace]
	if headConstructor == nil {
		return nil, fmt.Errorf("the namespace %s of head is not support", headNamespace)
	}
	docConstructor := docMessageConstructor[docNamespace]
	if docConstructor == nil {
		return nil, fmt.Errorf("the namespace %s of document is not support", headNamespace)
	}

	msg := &FpsCombineMessageObject{
		BizData: FpsBizData{
			XmlHead: headConstructor(),
			XmlDoc: FpsDoc{
				XmlDoc: docConstructor(),
			},
		},
	}

	err = xml.Unmarshal(buf, msg)
	if err != nil {
		return nil, err
	}

	return msg, nil
}
