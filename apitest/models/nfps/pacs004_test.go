package models

import (
	"encoding/xml"
	"testing"
)

var pacs004Raw = []byte(`<?xml version="1.0" encoding="UTF-8"?>
<fps:FpsMsg xmlns:fps="urn:Central Bank:fps:xsd:fps.envelope.01"
  xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xsi:schemaLocation="urn:Central Bank:fps:xsd:fps.envelope.01 fps.envelope.01.xsd">
  <fps:NbOfMsgs>1</fps:NbOfMsgs>
  <fps:FpsPylds>
    <fps:BizData>
      <ah:AppHdr xmlns:ah="urn:iso:std:iso:20022:tech:xsd:head.001.001.01">
        <ah:Fr>
          <ah:FIId>
            <ah:FinInstnId>
              <ah:ClrSysMmbId>
                <ah:MmbId>part00011</ah:MmbId>
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
        <ah:BizMsgIdr>M555201701050000001331</ah:BizMsgIdr>
        <ah:MsgDefIdr>pacs.004.001.07</ah:MsgDefIdr>
        <ah:BizSvc>PAYR01</ah:BizSvc>
        <ah:CreDt>2017-01-06T09:25:11.001Z</ah:CreDt>
      </ah:AppHdr>
      <doc:Document xmlns:doc="urn:iso:std:iso:20022:tech:xsd:pacs.004.001.07">
        <doc:PmtRtr>
          <doc:GrpHdr>
            <doc:MsgId>M555201701050000001331</doc:MsgId>
            <doc:CreDtTm>2017-01-06T17:25:11</doc:CreDtTm>
            <doc:NbOfTxs>1</doc:NbOfTxs>
            <doc:SttlmInf>
              <doc:SttlmMtd>CLRG</doc:SttlmMtd>
              <doc:ClrSys>
                <doc:Prtry>FPS</doc:Prtry>
              </doc:ClrSys>
            </doc:SttlmInf>
          </doc:GrpHdr>
          <doc:TxInf>
            <doc:RtrId>AB57F0D7E5DA4995AFF89F2CC50E95D5</doc:RtrId>
            <doc:OrgnlEndToEndId>CROPS/SX-25T/2015-10-13</doc:OrgnlEndToEndId>
            <doc:OrgnlTxId>7A0AD19BD76444979280D5A80F36176D</doc:OrgnlTxId>
            <doc:OrgnlClrSysRef>FRN20160801000000002</doc:OrgnlClrSysRef>
            <doc:RtrdIntrBkSttlmAmt Ccy="CNY">0.01</doc:RtrdIntrBkSttlmAmt>
            <doc:IntrBkSttlmDt>2017-01-05</doc:IntrBkSttlmDt>
            <doc:RtrdInstdAmt Ccy="CNY">0.01</doc:RtrdInstdAmt>
            <doc:ChrgBr>SLEV</doc:ChrgBr>
            <doc:RtrRsnInf>
              <doc:Rsn>
                <doc:Prtry>NARR</doc:Prtry>
              </doc:Rsn>
              <doc:AddtlInf>/CATPRP/RPRFND</doc:AddtlInf>
              <doc:AddtlInf>/REASON/Business Refund</doc:AddtlInf>
            </doc:RtrRsnInf>
            <doc:OrgnlTxRef>
              <doc:IntrBkSttlmAmt Ccy="CNY">0.01</doc:IntrBkSttlmAmt>
              <doc:IntrBkSttlmDt>2017-01-05</doc:IntrBkSttlmDt>
              <doc:PmtTpInf>
                <doc:CtgyPurp>
                  <doc:Prtry>CXSALA</doc:Prtry>
                </doc:CtgyPurp>
              </doc:PmtTpInf>
              <doc:RmtInf>
                <doc:Ustrd>Allowance Paid #12345</doc:Ustrd>
              </doc:RmtInf>
              <doc:Dbtr>
                <doc:Nm>MISC HOLDING LTD.</doc:Nm>
                <doc:Id>
                  <doc:OrgId>
                    <doc:Othr>
                      <doc:Id>0050092135555</doc:Id>
                      <doc:SchmeNm>
                        <doc:Cd>CUST</doc:Cd>
                      </doc:SchmeNm>
                    </doc:Othr>
                  </doc:OrgId>
                </doc:Id>
                <doc:CtctDtls>
                  <doc:MobNb>+852-22223333</doc:MobNb>
                  <doc:EmailAdr>susan.wong@abc.com</doc:EmailAdr>
                </doc:CtctDtls>
              </doc:Dbtr>
              <doc:DbtrAcct>
                <doc:Id>
                  <doc:Othr>
                    <doc:Id>0050092135555</doc:Id>
                    <doc:SchmeNm>
                      <doc:Prtry>BBAN</doc:Prtry>
                    </doc:SchmeNm>
                  </doc:Othr>
                </doc:Id>
              </doc:DbtrAcct>
              <doc:DbtrAgt>
                <doc:FinInstnId>
                  <doc:ClrSysMmbId>
                    <doc:MmbId>part00010</doc:MmbId>
                  </doc:ClrSysMmbId>
                </doc:FinInstnId>
              </doc:DbtrAgt>
              <doc:CdtrAgt>
                <doc:FinInstnId>
                  <doc:ClrSysMmbId>
                    <doc:MmbId>part00011</doc:MmbId>
                  </doc:ClrSysMmbId>
                </doc:FinInstnId>
              </doc:CdtrAgt>
              <doc:Cdtr>
                <doc:Nm>Susan Wong</doc:Nm>
                <doc:Id>
                  <doc:PrvtId>
                    <doc:Othr>
                      <doc:Id>ID884223344</doc:Id>
                      <doc:SchmeNm>
                        <doc:Cd>CUST</doc:Cd>
                      </doc:SchmeNm>
                    </doc:Othr>
                  </doc:PrvtId>
                </doc:Id>
                <doc:CtctDtls>
                  <doc:MobNb>+852-23456789</doc:MobNb>
                </doc:CtctDtls>
              </doc:Cdtr>
              <doc:CdtrAcct>
                <doc:Id>
                  <doc:Othr>
                    <doc:Id>0200015869999</doc:Id>
                    <doc:SchmeNm>
                      <doc:Prtry>BBAN</doc:Prtry>
                    </doc:SchmeNm>
                  </doc:Othr>
                </doc:Id>
              </doc:CdtrAcct>
            </doc:OrgnlTxRef>
          </doc:TxInf>
        </doc:PmtRtr>
      </doc:Document>
    </fps:BizData>
  </fps:FpsPylds>
  <Signature xmlns="http://www.w3.org/2000/09/xmldsig#">
    <SignedInfo>
      <CanonicalizationMethod Algorithm="http://www.w3.org/2001/10/xml-exc-c14n#" />
      <SignatureMethod Algorithm="http://www.w3.org/2001/04/xmldsig-more#rsa-sha256" />
      <Reference URI="">
        <Transforms>
          <Transform Algorithm="http://www.w3.org/2000/09/xmldsig#enveloped-signature" />
        </Transforms>
        <DigestMethod Algorithm="http://www.w3.org/2001/04/xmlenc#sha256" />
        <DigestValue>HMwpTa6Fcp9UnAdoY0rj90Z5mf0vHAVJvtN4hDryPSQ=
        </DigestValue>
      </Reference>
    </SignedInfo>
    <SignatureValue>YWH+Wjf2uZ8AzNcp35dgLgOROfuLZHFwRvDw4mB6u5Csiu+QoMunAPpSvx/QNPzFXgZCt2Pg Kg5I&#13;
    </SignatureValue>
    <KeyInfo>
      <X509Data>
        <X509Certificate>MIICrjCCAZagAwIBAgIEV7PBijANBgkqhkiG9w0BAQsFADAZMRcwFQYDVQQDEw5nZW5lcml jIGJh&#13;
        </X509Certificate>
      </X509Data>
    </KeyInfo>
  </Signature>
</fps:FpsMsg>
`)

func TestFPSPacs004(t *testing.T) {

	pacs004 := &FPSPacs004{}
	if err := xml.Unmarshal(pacs004Raw, pacs004); err != nil {
		t.Error(err)
	} else {
		t.Logf("%++v",pacs004)
	}


}