package main

import (
	"fmt"

	//models "git.sirius.io/banking/test/model"
	"time"
)

func main() {
	//var a []byte
	//a = []byte{60, 63, 120, 109, 108, 32, 118, 101, 114, 115, 105, 111, 110, 61, 34, 49, 46, 48, 34, 32, 101, 110, 99, 111, 100, 105, 110, 103, 61, 34, 73, 83, 79, 45, 56, 56, 53, 57, 45, 49, 34, 63, 62, 10, 60, 100, 101, 99, 114, 121, 112, 116, 101, 100, 82, 101, 115, 117, 108, 116, 62, 60, 47, 100, 101, 99, 114, 121, 112, 116, 101, 100, 82, 101, 115, 117, 108, 116, 62, 10, 13, 10, 13, 10, 13, 10}
	//fmt.Println(string(a))
	//s, err := os.OpenFile("txt.csv", os.O_SYNC|os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	//if err != nil {
	//	panic(err)
	//}
	//defer s.Close()
	//wt := csv.NewWriter(s)
	//var a string
	//for i := 1; i <= 200; i++ {
	//	a = "<fps:FpsMsg xmlns:fps=\"urn:Central Bank:fps:xsd:fps.envelope.01\" xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xsi:schemaLocation=\"urn:Central Bank:fps:xsd:fps.envelope.01 fps.envelope.01.xsd\"><fps:NbOfMsgs> </fps:NbOfMsgs><fps:FpsPylds> <fps:BizData>     <ah:AppHdr xmlns:ah=\"urn:iso:std:iso:20022:tech:xsd:head.001.001.01\" ID=\"1657781189432886612\"><ah:Fr>    <ah:FIId>        <ah:FinInstnId>            <ah:ClrSysMmbId>                <ah:MmbId>401</ah:MmbId>            </ah:ClrSysMmbId>        </ah:FinInstnId>    </ah:FIId></ah:Fr><ah:To>    <ah:FIId>        <ah:FinInstnId>            <ah:ClrSysMmbId>                <ah:MmbId>FPS</ah:MmbId>            </ah:ClrSysMmbId>        </ah:FinInstnId>    </ah:FIId></ah:To><ah:BizMsgIdr>20220711344626900007191020166021</ah:BizMsgIdr><ah:MsgDefIdr>pacs.008.001.06</ah:MsgDefIdr><ah:BizSvc>PAYC01</ah:BizSvc><ah:CreDt>2022-07-16T14:46:26Z</ah:CreDt>     <ds:Signature xmlns:ds=\"http://www.w3.org/2000/09/xmldsig#\"><ds:SignedInfo><ds:CanonicalizationMethod Algorithm=\"http://www.w3.org/2006/12/xml-c14n11\"/><ds:SignatureMethod Algorithm=\"http://www.w3.org/2001/04/xmldsig-more#rsa-sha256\"/><ds:Reference URI=\"#1657781189432886612\"><ds:Transforms><ds:Transform Algorithm=\"http://www.w3.org/2000/09/xmldsig#enveloped-signature\"/><ds:Transform Algorithm=\"http://www.w3.org/2006/12/xml-c14n11\"/></ds:Transforms><ds:DigestMethod Algorithm=\"http://www.w3.org/2001/04/xmlenc#sha256\"/><ds:DigestValue>kRnEdWeP2xiEjlUNSHqbw0DfT+n3loZy2I6DYO1ofJw=</ds:DigestValue></ds:Reference></ds:SignedInfo><ds:SignatureValue>PZgwXFH17tIW7hMLv6nkFKCnfk5MVnfKJfPKGUrlHyzAA9e+2m3M4xqhCCvCun2Bl4SeiH18dkMD6n8f7L+XKlVxLFbSfZ5PyAMLAyj4DNPS7yUZdyRLNMwULPRG5ex1JMfLmKBWBS/1hglqtXs8rHlwrSAlGLWpVNNP7caca/jJ/EAWO99Toaklm0WPz8PfhuVpx10qbg76mtLYV0QkFz1BL/Syiud1JsD2YjVvfAhW9PxCZl5YfmuzLKgPC/7s00JbBiemH7NOowNtJT1P1FSxhg7BDf5NUsajacbPnR5fsUlZvbaIVqaYcDsRdpsWLOEBoKh+Ac6ayd0eNb+byA==</ds:SignatureValue><ds:KeyInfo><ds:X509Data><ds:X509Certificate>MIIDfTCCAmWgAwIBAgIQAOmEL+w4CnBPzoQEJ2IwuzANBgkqhkiG9w0BAQsFADAqMSgwJgYDVQQDDB9odHRwOi8vd3d3LmNvb2tjb2RlLmNjL3NlbGZzaWduMB4XDTIyMDYyMDA5MTkwOVoXDTIzMDYyMDA5MTkwOVowKjEoMCYGA1UEAwwfaHR0cDovL3d3dy5jb29rY29kZS5jYy9zZWxmc2lnbjCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBAN9boOcCZK/OFzcqDY7GH9tKx6bRAjcxny+9gCGfnLdC3L9hmsbBnK0KD2Zr+7b/ztKd4v9gPozlsAIhz+CD96cuKustUHjxQvRHRyykww1mjlhoWpH08i6qzkGb5R3KlQW+HRIAlNfmGmLEccU/89+DBZRkQ6r9NMnvAT0z4zKxAIHHacMWlFDQ+dLwykPhy7oA/oFhgJqEk1Ioo7kTtsk9qUalrT8lw5qUHeXAI66RwUvZW4l85oKVSHfOZOWr2sLRPQAcL+qqZH0LHRXHeTuOZl9tyLTdHK30KzXufeLP4Co+giK63A1bsw6AhIzUNqSdQBFgMwmibONljTGSxh8CAwEAAaOBnjCBmzAdBgNVHQ4EFgQUbmygNhjFTt2+zsq8oHfRWb+e9cEwDgYDVR0PAQH/BAQDAgSwMAwGA1UdEwEB/wQCMAAwOwYDVR0lBDQwMgYIKwYBBQUHAwIGCCsGAQUFBwMBBggrBgEFBQcDAwYIKwYBBQUHAwQGCCsGAQUFBwMIMB8GA1UdIwQYMBaAFG5soDYYxU7dvs7KvKB30Vm/nvXBMA0GCSqGSIb3DQEBCwUAA4IBAQBr8Cdmk+JnNvE0JhVq/OyBFeaXFZovbTO3FPx2+8zsRFonzpHKsNdxhhKRrD4F6KoNKvGTDnQs6LJGgoQOZ0cdvP2A9Pg8p4PWZcznJlYAkrC2MuVgj7sU1KgAF3/i6AoCIU4H2leXuTDatp4ps3chf21OKCAQyp9DtrgzBNkwMXauL32k5ZSn4srBzCQ3irnwkmh2d/okFdoGBCLdncSpZnNO/rYOEZWK/GV0G8oD70D8HzNK5ZAa7HGzumfHaavvvGfxhaaF6s31V9FNIgetaWzbOhlAxecApUtZNPkx/Qa61k35CtOhp4fogkpXwYhzeEFRxnTC/wNbwq9RT4wE</ds:X509Certificate></ds:X509Data></ds:KeyInfo></ds:Signature></ah:AppHdr>     <doc:Document xmlns:doc=\"urn:iso:std:iso:20022:tech:xsd:pacs.008.001.06\"><doc:FIToFICstmrCdtTrf>    <doc:GrpHdr>        <doc:MsgId>20220711344626900007191020166021</doc:MsgId>        <doc:CreDtTm>2022-07-18T14:46:26.909</doc:CreDtTm>        <doc:IntrBkSttlmDt>2022-07-18</doc:IntrBkSttlmDt>        <doc:TtlIntrBkSttlmAmt Ccy=\"THB\">9</doc:TtlIntrBkSttlmAmt>        <doc:NbOfTxs>1</doc:NbOfTxs>        <doc:SttlmInf>            <doc:SttlmMtd>CLRG</doc:SttlmMtd>            <doc:ClrSys>                <doc:Prtry>FPS</doc:Prtry>                <doc:Cd>FPS</doc:Cd>            </doc:ClrSys>        </doc:SttlmInf>    </doc:GrpHdr>    <doc:CdtTrfTxInf>        <doc:PmtId>            <doc:EndToEndId>1009150719150012221021</doc:EndToEndId>            <doc:InstrId>1009150719250012221021</doc:InstrId>            <doc:TxId>1009150719250012221021</doc:TxId>        </doc:PmtId>        <doc:PmtTpInf>            <doc:LclInstrm>                <doc:Prtry>SKIP_PYE_VRF</doc:Prtry>            </doc:LclInstrm>            <doc:CtgyPurp>                <doc:Prtry>CXMRCH</doc:Prtry>            </doc:CtgyPurp>        </doc:PmtTpInf>        <doc:Purp>            <doc:Cd/>        </doc:Purp>        <doc:IntrBkSttlmAmt Ccy=\"THB\">9</doc:IntrBkSttlmAmt>        <doc:IntrBkSttlmDt>2022-07-18</doc:IntrBkSttlmDt>        <doc:ChrgBr>SLEV</doc:ChrgBr>        <doc:Dbtr>            <doc:Nm>Ho Mai Li</doc:Nm>        </doc:Dbtr>        <doc:DbtrAcct>            <doc:Id>                <doc:Othr>                    <doc:Id>0856791321987</doc:Id>                    <doc:SchmeNm>                        <doc:Prtry>BBAN</doc:Prtry>                    </doc:SchmeNm>                </doc:Othr>            </doc:Id>        </doc:DbtrAcct>        <doc:DbtrAgt>            <doc:FinInstnId>                <doc:ClrSysMmbId>                    <doc:MmbId>401</doc:MmbId>                </doc:ClrSysMmbId>            </doc:FinInstnId>        </doc:DbtrAgt>        <doc:CdtrAgt>            <doc:FinInstnId>                <doc:ClrSysMmbId>                    <doc:MmbId>402</doc:MmbId>                </doc:ClrSysMmbId>            </doc:FinInstnId>        </doc:CdtrAgt>        <doc:Cdtr>            <doc:Nm>Generic Kindergarten</doc:Nm>        </doc:Cdtr>        <doc:CdtrAcct>            <doc:Id>                <doc:Othr>                    <doc:Id>0854791321987</doc:Id>                    <doc:SchmeNm>                        <doc:Prtry>BBAN</doc:Prtry>                    </doc:SchmeNm>                </doc:Othr>            </doc:Id>            <doc:RecvTaxld>1212</doc:RecvTaxld>            <doc:SendTaxId>55565</doc:SendTaxId>        </doc:CdtrAcct>    </doc:CdtTrfTxInf>    <doc:SplmtryData>        <doc:Envlp>            <doc:AdditionalData>                <doc:PrcCd>38000</doc:PrcCd>                <doc:MTI>0200</doc:MTI>            </doc:AdditionalData>        </doc:Envlp>        <doc:Agent>            <doc:FinInstnId>                <doc:ClrSysMmbId>                    <doc:MmbId>002</doc:MmbId>                </doc:ClrSysMmbId>            </doc:FinInstnId>        </doc:Agent>        <doc:OrgMsg>            <doc:MchTp>002</doc:MchTp>            <doc:TTp>02</doc:TTp>            <doc:Ota>02</doc:Ota>            <doc:OrgnlInstrId>02</doc:OrgnlInstrId>            <doc:TpRecv>2</doc:TpRecv>            <doc:TpSend>2</doc:TpSend>            <doc:MchBillId>02</doc:MchBillId>            <doc:BillDisNmTHA>02</doc:BillDisNmTHA>            <doc:BillDisNmENG>02</doc:BillDisNmENG>            <doc:CustDisNm>02</doc:CustDisNm>        </doc:OrgMsg>    </doc:SplmtryData></doc:FIToFICstmrCdtTrf>     </doc:Document> </fps:BizData></fps:FpsPylds><Signature xmlns=\"http://www.w3.org/2000/09/xmldsig#\"> <SignedInfo>     <CanonicalizationMethod Algorithm=\"http://www.w3.org/2001/10/xml-exc-c14n#\"/>     <SignatureMethod Algorithm=\"http://www.w3.org/2001/04/xmldsig-more#rsa-sha256\"/>     <Reference URI=\"\"><Transforms>    <Transform Algorithm=\"http://www.w3.org/2000/09/xmldsig#enveloped-signature\"/></Transforms><DigestMethod Algorithm=\"http://www.w3.org/2001/04/xmlenc#sha256\"/><DigestValue>HMwpTa6Fcp9UnAdoY0rj90Z5mf0vHAVJvtN4hDryPSQ= \t\t\t\t</DigestValue>     </Reference> </SignedInfo> <SignatureValue>YWH+Wjf2uZ8AzNcp35dgLgOROfuLZHFwRvDw4mB6u5Csiu+QoMunAPpSvx/QNPzFXgZCt2Pg Kg5I \t\t</SignatureValue> <KeyInfo>     <X509Data><X509Certificate>MIICrjCCAZagAwIBAgIEV7PBijANBgkqhkiG9w0BAQsFADAZMRcwFQYDVQQDEw5nZW5lcml jIGJh \t\t\t\t</X509Certificate>     </X509Data> </KeyInfo></Signature></fps:FpsMsg>"
	//	date1 := time.Now().Format("20060102")
	//	date := time.Now().Format("2006-01-02")
	//	msgId := fmt.Sprintf("%s%s%03d", date1, "344626900007191020166021", i)
	//	endToEnd := fmt.Sprintf("%s%03d", "1009150719150012221021", i)
	//	instrId := fmt.Sprintf("%s%03d", "1009150719250012221021", i)
	//	a = strings.ReplaceAll(a, "20220711344626900007191020166021", msgId)
	//	a = strings.ReplaceAll(a, "2022-07-18", date)
	//	a = strings.ReplaceAll(a, "1009150719150012221021", endToEnd)
	//	a = strings.ReplaceAll(a, "1009150719250012221021", instrId)
	//	e := wt.Write([]string{a})
	//	if e != nil {
	//		fmt.Println(e)
	//	}
	//	//wt.Flush()
	//}
	//wt.Flush()
	//	var a []byte
	//	a = []byte(`<?xml version="1.0" encoding="UTF-8"?>
	//<fps:FpsMsg
	//    xmlns:fps="urn:Central Bank:fps:xsd:fps.envelope.01"
	//    xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xsi:schemaLocation="urn:Central Bank:fps:xsd:fps.envelope.01 fps.envelope.01.xsd">
	//    <fps:NbOfMsgs>1</fps:NbOfMsgs>
	//    <fps:FpsPylds>
	//        <fps:BizData>
	//            <ah:AppHdr
	//                xmlns:ah="urn:iso:std:iso:20022:tech:xsd:head.001.001.01">
	//                <ah:Fr>
	//                    <ah:FIId>
	//                        <ah:FinInstnId>
	//                            <ah:ClrSysMmbId>
	//                                <ah:MmbId>{{fromMem}}</ah:MmbId>
	//                            </ah:ClrSysMmbId>
	//                        </ah:FinInstnId>
	//                    </ah:FIId>
	//                </ah:Fr>
	//                <ah:To>
	//                    <ah:FIId>
	//                        <ah:FinInstnId>
	//                            <ah:ClrSysMmbId>
	//                                <ah:MmbId>FPS</ah:MmbId>
	//                            </ah:ClrSysMmbId>
	//                        </ah:FinInstnId>
	//                    </ah:FIId>
	//                </ah:To>
	//                <ah:BizMsgIdr>20221009151701990012221150060{{seqno}}</ah:BizMsgIdr>
	//                <ah:MsgDefIdr>pacs.008.001.06</ah:MsgDefIdr>
	//                <ah:BizSvc>PAYC01</ah:BizSvc>
	//                <ah:CreDt>{{today}}T06:25:11Z</ah:CreDt>
	//            </ah:AppHdr>
	//            <doc:Document
	//                xmlns:doc="urn:iso:std:iso:20022:tech:xsd:pacs.008.001.06">
	//                <IdVrctnReq>
	//                    <Assgnmt>
	//                        <MsgId>CA-0023896</MsgId>
	//                        <CreDtTm>2012-04-07T11:30:00</CreDtTm>
	//                        <Cretr>
	//                            <Pty>
	//                                <Nm>Corporate Actions Inc.</Nm>
	//                            </Pty>
	//                        </Cretr>
	//                        <Assgnr>
	//                            <Pty>
	//                                <Nm>Corporate Actions Inc.</Nm>
	//                            </Pty>
	//                        </Assgnr>
	//                        <Assgne>
	//                            <Agt>
	//                                <FinInstnId>
	//                                    <BICFI>CCCCFRTT</BICFI>
	//                                </FinInstnId>
	//                            </Agt>
	//                        </Assgne>
	//                    </Assgnmt>
	//                    <Vrfctn>
	//                        <Id>VER-12365</Id>
	//                        <PtyAndAcctId>
	//                            <Acct>
	//<Id>
	//                                <IBAN>FR430041020030300015M02606</IBAN>
	//</Id>
	//                            </Acct>
	//                        </PtyAndAcctId>
	//                    </Vrfctn>
	//                </IdVrctnReq>
	//            </doc:Document>
	//        </fps:BizData>
	//    </fps:FpsPylds>
	//</fps:FpsMsg>
	//`)
	//req := &models.Acmt023{}
	//err := xml.Unmarshal(a, req)
	//if err != nil {
	//	fmt.Println("err=", err)
	//}
	//currentTime := time.Now()
	//fmt.Println(currentTime.UTC().Format("2006-01-02T15:04:05Z"))
	//fmt.Println(currentTime.Format("2006-01-02T15:04:05.000Z"))
	//fmt.Println(currentTime.UTC().Format("2006-01-02T15:04:05Z07:00"))
	//fmt.Println(currentTime.In(time.Local).Format("2006-01-02T15:04:05Z07:00"))
	//fmt.Println(currentTime.In(time.Local))
	//fmt.Println([]byte("aa"))
	//fmt.Println(req)
	//
	//xmlDoc := `<?xml version="1.0" encoding="UTF-8"?>
	//            <note>
	//              <to>Tove</to>
	//              <from>Jani</from>
	//              <heading>Reminder</heading>
	//              <body>Don't forget me this weekend!</body>
	//            </note>`
	//type xmlStruct struct {
	//	XMLName xml.Name `xml:"note"`
	//	To      string   `xml:"to"`
	//	From    string   `xml:"from"`
	//	Heading string   `xml:"heading"`
	//	Body    string   `xml:"body"`
	//}
	//x := xmlStruct{}
	//err = xml.Unmarshal([]byte(xmlDoc), &x)
	//if err != nil {
	//	fmt.Println(err.Error())
	//} else {
	//	fmt.Println(x.Body)
	//}

	//bytes, err := xml.MarshalIndent(&x, "", "    ")
	//bytes = append([]byte(xml.Header), bytes...)
	//if err != nil {
	//	fmt.Println(err.Error())
	//} else {
	//	fmt.Println(string(bytes))
	//}

	//获取当前时间
	t := time.Now() //2018-07-11 15:07:51.8858085 +0800 CST m=+0.004000001
	fmt.Println(t)

	//获取当前时间戳
	fmt.Println(t.Unix()) //1531293019

	//获得当前的时间
	//fmt.Println(t.Unix().Format("2006-01-02 15:04:05"))  //2018-7-15 15:23:00

	//时间 to 时间戳
	loc, _ := time.LoadLocation("Asia/Shanghai")                                     //设置时区
	tt, _ := time.ParseInLocation("2006-01-02 15:04:05", "2018-07-11 15:07:51", loc) //2006-01-02 15:04:05是转换的格式如php的"Y-m-d H:i:s"
	fmt.Println(tt.Unix())                                                           //1531292871
	fmt.Println(tt.Unix())                                                           //1531292871

	//时间戳 to 时间
	tm := time.Unix(1531293019, 0)
	fmt.Println(tm.Format("2006-01-02 15:04:05")) //2018-07-11 15:10:19
	j := 0
	for i := 0; i < 10; i++ {
		j += 1
		fmt.Println(j)
	}
}
