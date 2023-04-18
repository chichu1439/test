package message

import (
	"encoding/xml"
	"testing"
)

var pacs008Str = `<?xml version="1.0" encoding="UTF-8"?>
<fps:FpsMsg xmlns:fps="urn:CentralBank:fps:xsd:fps.envelope.01">
  <fps:BizData>
    <ah:AppHdr xmlns:ah="urn:iso:std:iso:20022:tech:xsd:head.001.001.01">
      <ah:Fr>
        <ah:FIId>
          <ah:FinInstnId>
            <ah:ClrSysMmbId>
              <ah:MmbId>{{fromMem}}</ah:MmbId>
            </ah:ClrSysMmbId>
          </ah:FinInstnId>
        </ah:FIId>
      </ah:Fr>
      <ah:To>
        <ah:FIId>
          <ah:FinInstnId>
            <ah:ClrSysMmbId>
              <ah:MmbId>FPS</ah:MmbId>
            </ah:ClrSysMmbId>
          </ah:FinInstnId>
        </ah:FIId>
      </ah:To>
      <ah:BizMsgIdr>20211009151701990012221150060{{seqno}}</ah:BizMsgIdr>
      <ah:MsgDefIdr>pacs.008.001.06</ah:MsgDefIdr>
      <ah:BizSvc>PAYC01</ah:BizSvc>
      <ah:CreDt>2022-08-01T06:25:11Z</ah:CreDt>
      <ah:Sgntr>
        <Signature xmlns="http://www.w3.org/2000/09/xmldsig#">
          <SignedInfo>
            <CanonicalizationMethod Algorithm="http://www.w3.org/2001/10/xml-exc-c14n#"/>
            <SignatureMethod Algorithm="http://www.w3.org/2001/04/xmldsig-more#rsa-sha256"/>
            <Reference URI="">
              <Transforms>
                <Transform Algorithm="http://www.w3.org/2000/09/xmldsig#enveloped-signature"/>
              </Transforms>
              <DigestMethod Algorithm="http://www.w3.org/2001/04/xmlenc#sha256"/>
              <DigestValue>HMwpTa6Fcp9UnAdoY0rj90Z5mf0vHAVJvtN4hDryPSQ=
              </DigestValue>
            </Reference>
          </SignedInfo>
          <SignatureValue>YWH+Wjf2uZ8AzNcp35dgLgOROfuLZHFwRvDw4mB6u5Csiu+QoMunAPpSvx/QNPzFXgZCt2Pg Kg5I&#13;
          </SignatureValue>
          <KeyInfo>
            <X509Data>
              <X509Certificate>MIICrjCCAZagAwIBAgIEV7PBijANBgkqhkiG9w0BAQsFADAZMRcwFQYDVQQDEw5nZW5lcml
                                jIGJh&#13;
              </X509Certificate>
            </X509Data>
          </KeyInfo>
        </Signature>
      </ah:Sgntr>
    </ah:AppHdr>
    <doc:Document xmlns:doc="urn:iso:std:iso:20022:tech:xsd:pacs.008.001.06">
      <doc:FIToFICstmrCdtTrf>
        <doc:GrpHdr>
          <doc:MsgId>20211009151701990012221150060{{seqno}}</doc:MsgId>
          <doc:CreDtTm>2022-08-01T14:25:11.001</doc:CreDtTm>
          <doc:IntrBkSttlmDt>2022-08-01</doc:IntrBkSttlmDt>
          <doc:TtlIntrBkSttlmAmt Ccy="THB">0.01</doc:TtlIntrBkSttlmAmt>
          <doc:NbOfTxs>1</doc:NbOfTxs>
          <doc:SttlmInf>
            <doc:SttlmMtd>CLRG</doc:SttlmMtd>
            <doc:ClrSys>
              <doc:Prtry>FPS</doc:Prtry>
              <doc:Cd>FPS</doc:Cd>
            </doc:ClrSys>
          </doc:SttlmInf>
        </doc:GrpHdr>
        <doc:CdtTrfTxInf>
          <doc:PmtId>
            <doc:EndToEndId>1008151701900012221150{{relseqno}}</doc:EndToEndId>
            <doc:InstrId>1009151701900012221150{{relseqno}}</doc:InstrId>
            <doc:TxId>1008151701900012221150{{relseqno}}</doc:TxId>
          </doc:PmtId>
          <doc:PmtTpInf>
            <doc:LclInstrm>
              <doc:Prtry>SKIP_PYE_VRF</doc:Prtry>
            </doc:LclInstrm>
            <doc:CtgyPurp>
              <doc:Prtry></doc:Prtry>
            </doc:CtgyPurp>
          </doc:PmtTpInf>
          <doc:IntrBkSttlmAmt Ccy="THB">0.01</doc:IntrBkSttlmAmt>
          <doc:IntrBkSttlmDt>2022-08-01</doc:IntrBkSttlmDt>
          <doc:ChrgBr>SLEV</doc:ChrgBr>
          <doc:Dbtr>
            <doc:Nm>Ho Mai Li</doc:Nm>
          </doc:Dbtr>
          <doc:DbtrAcct>
            <doc:Id>
              <doc:Othr>
                <doc:Id>0856791321987</doc:Id>
                <doc:SchmeNm>
                  <doc:Prtry>BBAN</doc:Prtry>
                </doc:SchmeNm>
              </doc:Othr>
            </doc:Id>
          </doc:DbtrAcct>
          <doc:DbtrAgt>
            <doc:FinInstnId>
              <doc:ClrSysMmbId>
                <doc:MmbId>{{fromMem}}</doc:MmbId>
              </doc:ClrSysMmbId>
            </doc:FinInstnId>
          </doc:DbtrAgt>
          <doc:CdtrAgt>
            <doc:FinInstnId>
              <doc:ClrSysMmbId>
                <doc:MmbId>{{toMem}}</doc:MmbId>
              </doc:ClrSysMmbId>
            </doc:FinInstnId>
          </doc:CdtrAgt>
          <doc:Cdtr>
            <doc:Nm>Generic Kindergarten</doc:Nm>
          </doc:Cdtr>
          <doc:CdtrAcct>
            <doc:Id>
              <doc:Othr>
                <doc:Id>0854791321987</doc:Id>
                <doc:SchmeNm>
                  <doc:Prtry>BBAN</doc:Prtry>
                </doc:SchmeNm>
              </doc:Othr>
            </doc:Id>
            <doc:RecvTaxld>1212</doc:RecvTaxld>
            <doc:SendTaxId>55565</doc:SendTaxId>
          </doc:CdtrAcct>
        </doc:CdtTrfTxInf>
        <doc:SplmtryData>
          <doc:Envlp>
            <doc:AdditionalData>
              <doc:PrcCd>38000</doc:PrcCd>
              <doc:MTI>0200</doc:MTI>
              <doc:Agent>
                <doc:FinInstnId>
                  <doc:ClrSysMmbId>
                    <doc:MmbId>002</doc:MmbId>
                  </doc:ClrSysMmbId>
                </doc:FinInstnId>
              </doc:Agent>
              <doc:OrgMsg>
                <doc:MchTp>002</doc:MchTp>
                <doc:TTp>02</doc:TTp>
                <doc:Ota>02</doc:Ota>
                <doc:OrgnlInstrId>02</doc:OrgnlInstrId>
                <doc:TpRecv>2</doc:TpRecv>
                <doc:TpSend>2</doc:TpSend>
                <doc:MchBillId>02</doc:MchBillId>
                <doc:BillDisNmTHA>02</doc:BillDisNmTHA>
                <doc:BillDisNmENG>02</doc:BillDisNmENG>
                <doc:CustDisNm>02</doc:CustDisNm>
              </doc:OrgMsg>
            </doc:AdditionalData>
          </doc:Envlp>
        </doc:SplmtryData>
      </doc:FIToFICstmrCdtTrf>
    </doc:Document>
  </fps:BizData>
</fps:FpsMsg>`

func TestParseFpsCombineMessage(t *testing.T) {

	var combineMessage FpsCombineMessage
	var err error
	combineMessage, err = ParseFpsCombineMessage([]byte(pacs008Str))

	if err != nil {
		t.Error(err)
	} else {
		t.Logf("%++v\n", combineMessage)
	}

	t.Logf("-----------------------------\n")

	if buf, err := xml.MarshalIndent(combineMessage, "", "\t"); err != nil {
		t.Error(err)
	} else {
		t.Logf("%s\n", buf)
	}

}
