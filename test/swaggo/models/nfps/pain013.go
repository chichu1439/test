package models

import (
	"encoding/xml"
)

type FPSPain013 struct {
	XMLName  xml.Name `xml:"FpsMsg" json:"xmlName"`
	NumOfMsg string   `xml:"NbOfMsgs" json:"numOfMsg"`
	XMLHead  Head001  `xml:"FpsPylds>BizData>AppHdr" json:"xmlHead"`
	Document Pain013  `xml:"FpsPylds>BizData>Document"`
}

type Pain013 struct {
	MsgId          string `xml:"CdtrPmtActvtnReq_Proxy>GrpHdr>MsgId,omitempty" json:"msgId" validate:"required,max=35"`
	CreateDatetime string `xml:"CdtrPmtActvtnReq_Proxy>GrpHdr>CreDtTm,omitempty" validate:"required,max=35"`
	NumOfTrans     string `xml:"CdtrPmtActvtnReq_Proxy>GrpHdr>NbOfTxs,omitempty" validate:"required,max=15"`
	// Requester / Biller Bank
	SendBank string `xml:"CdtrPmtActvtnReq_Proxy>GrpHdr>InitgPty>Id>OrgId>Othr>Id,omitempty" validate:"required,max=3"`
	// YYYYMMDDDHHMMSSNNNNNBBBYDDDHHNNnnnn
	PaymentId string `xml:"CdtrPmtActvtnReq_Proxy>PmtInf>PmtInfId,omitempty" validate:"required,max=35"`
	// Specifies the means of payment that will be used to move the amount of money.
	PaymentMethod string `xml:"CdtrPmtActvtnReq_Proxy>PmtInf>PmtMtd,omitempty" validate:"required,max=3"`
	// Date at which the initiating party requests the clearing agent to process the payment.
	PaymentExcDate string `xml:"CdtrPmtActvtnReq_Proxy>PmtInf>ReqdExctnDt,omitempty" validate:"required,max=10"`
	// Customer Account Name
	DrName string `xml:"CdtrPmtActvtnReq_Proxy>PmtInf>Dbtr>Nm,omitempty" validate:"max=50"`
	// Identification of the account, this could be a tokenised account number or an account number in the clear.帐户号码，也可以是清算中的帐户号码
	DrAcctNo string `xml:"CdtrPmtActvtnReq_Proxy>PmtInf>DbtrAcct>Id>Othr>Id,omitempty" validate:"required,max=35"`
	// Identification of a member of a clearing system 清算系统成员的识别
	DrMemberId string `xml:"CdtrPmtActvtnReq_Proxy>PmtInf>DbtrAgt>FinInstnId>ClrSysMmbId>MmbId,omitempty" validate:"max=3"`
	// RBBBydddnnnnnnnn
	InstructionId string   `xml:"CdtrPmtActvtnReq_Proxy>PmtInf>CdtTrfTx>PmtId>InstrId,omitempty" validate:"required,max=35"`
	EndToEndId    string   `xml:"CdtrPmtActvtnReq_Proxy>PmtInf>CdtTrfTx>PmtId>EndToEndId,omitempty" validate:"required,max=35"`
	SvcLvlCode    string   `xml:"CdtrPmtActvtnReq_Proxy>PmtInf>CdtTrfTx>PmtTpInf>SvcLvl>Cd,omitempty" validate:"required,max=35"`
	TransAmount   AmtField `xml:"CdtrPmtActvtnReq_Proxy>PmtInf>CdtTrfTx>Amt>InstdAmt,omitempty" validate:"required,max=24"`
	// Specifies which party/parties will bear the charges associated with the processing of the payment transaction.
	ChargeBearer string `xml:"CdtrPmtActvtnReq_Proxy>PmtInf>ChrgBr,omitempty" validate:"required,max=4"`
	// Identification of a member of a clearing system.清算系统成员的识别
	CrMemberId string `xml:"CdtrPmtActvtnReq_Proxy>PmtInf>CdtrAgt>FinInstnId>ClrSysMmbId>MmbId,omitempty" validate:"max=3"`
	// Biller / Request Account Name
	CrName string `xml:"CdtrPmtActvtnReq_Proxy>PmtInf>Cdtr>Nm,omitempty" validate:"max=140"`
	// Biller / Requester Account No
	CrAcctNo    string `xml:"CdtrPmtActvtnReq_Proxy>PmtInf>CdtrAcct>Id>Othr>Id,omitempty" validate:"required,max=35"`
	RemInf      string `xml:"CdtrPmtActvtnReq_Proxy>PmtInf>RmtInf>Ustrd,omitempty" validate:"max=140"`
	ProcessCode string `xml:"CdtrPmtActvtnReq_Proxy>SplmtryData>Envlp>AdditionalData>PrcCd,omitempty" validate:"required,max=6"`
	MessageType string `xml:"CdtrPmtActvtnReq_Proxy>SplmtryData>Envlp>AdditionalData>MTI,omitempty" validate:"required,max=4"`
	// TEPA Code or (Customer Bank Code)
	CustMemberId string `xml:"CdtrPmtActvtnReq_Proxy>SplmtryData>Agent>FinInstnId>ClrSysMmbId>MmbId,omitempty" validate:"max=3"`
	OrigMsg      string `xml:"CdtrPmtActvtnReq_Proxy>SplmtryData>OrgMsg,omitempty"`
}
type Pain013OriMsg struct {
	// This field identifies the type of merchant. 6012–POS, 6050–Teller, 6007–Internet, 6070–CDM etc.
	MerchantType string `xml:"CdtrPmtActvtnReq_Proxy>SplmtryData>OrgMsg>MchTp,omitempty" validate:"required,max=4"`
	// The terminal type code as specified for Channel Transactions:20 – Kiosk  40 – EDC/POS 60 – Internet 80 – Mobile
	TermType    string `xml:"CdtrPmtActvtnReq_Proxy>SplmtryData>OrgMsg>TTp,omitempty" validate:"required,max=2"`
	OneTimeAuth string `xml:"CdtrPmtActvtnReq_Proxy>SplmtryData>OrgMsg>Ota,omitempty" validate:"max=16"`
	// Type of Merchant / Biller
	RecvType string `xml:"CdtrPmtActvtnReq_Proxy>SplmtryData>OrgMsg>TpRecv,omitempty" validate:"required,max=1"`
	// Merchant / Biller Tax ID
	RecvTaxId      string `xml:"CdtrPmtActvtnReq_Proxy>SplmtryData>OrgMsg>RecvTaxId,omitempty" validate:"max=13"`
	MerchantBillId string `xml:"CdtrPmtActvtnReq_Proxy>SplmtryData>OrgMsg>MchBillId,omitempty" validate:"required,max=35"`
	BillDisNameTha string `xml:"CdtrPmtActvtnReq_Proxy>SplmtryData>OrgMsg>BillDisNmTHA,omitempty" validate:"required,max=50"`
	BillDisNameEng string `xml:"CdtrPmtActvtnReq_Proxy>SplmtryData>OrgMsg>BillDisNmENG,omitempty" validate:"required,max=50"`
	BillRef1       string `xml:"CdtrPmtActvtnReq_Proxy>SplmtryData>OrgMsg>Ref1,omitempty" validate:"required,max=20"`
	BillRef2       string `xml:"CdtrPmtActvtnReq_Proxy>SplmtryData>OrgMsg>Ref2,omitempty" validate:"max=20"`
	BillRef3       string `xml:"CdtrPmtActvtnReq_Proxy>SplmtryData>OrgMsg>Ref3,omitempty" validate:"max=20"`
	AddiNote       string `xml:"CdtrPmtActvtnReq_Proxy>SplmtryData>OrgMsg>AddiNote,omitempty" validate:"max=50"`
	// For cross border, Local amount at payment initiate
	LocalAmount     AmtField `xml:"CdtrPmtActvtnReq_Proxy>SplmtryData>OrgMsg>LcAmt,omitempty" validate:"max=24"`
	CountryCode     string   `xml:"CdtrPmtActvtnReq_Proxy>SplmtryData>OrgMsg>CtryCd,omitempty" validate:"max=2"`
	VatRate         string   `xml:"CdtrPmtActvtnReq_Proxy>SplmtryData>OrgMsg>VatRate,omitempty" validate:"max=35"`
	Vat             string   `xml:"CdtrPmtActvtnReq_Proxy>SplmtryData>OrgMsg>Vat,omitempty" validate:"max=35"`
	IncomeType      string   `xml:"CdtrPmtActvtnReq_Proxy>SplmtryData>OrgMsg>TpInc,omitempty" validate:"max=35"`
	WithHoldTaxRate string   `xml:"CdtrPmtActvtnReq_Proxy>SplmtryData>OrgMsg>WhtTaxRate,omitempty" validate:"max=35"`
	WithHoldTax     string   `xml:"CdtrPmtActvtnReq_Proxy>SplmtryData>OrgMsg>WhtTax,omitempty" validate:"max=35"`
	WithHoldTaxCond string   `xml:"CdtrPmtActvtnReq_Proxy>SplmtryData>OrgMsg>WhtTaxCon,omitempty" validate:"max=35"`
}

type FpsPain013 struct {
	XMLName  xml.Name   `xml:"FpsMsg" json:"xmlName"`
	NumOfMsg string     `xml:"NbOfMsgs" json:"numOfMsg"`
	XMLHead  Head001    `xml:"BizData>AppHdr" json:"xmlHead"`
	Document Pain013Doc `xml:"BizData>Document"`
}

type Pain013Doc struct {
	CdtrPmtActvtnReq struct {
		GrpHdr struct {
			MsgId    string  `xml:"MsgId,omitempty" validate:"required,max=35"`
			CreDtTm  string  `xml:"CreDtTm,omitempty" validate:"required,max=35"`
			NbOfTxs  string  `xml:"NbOfTxs,omitempty" validate:"required,max=15"`
			CtrlSum  float64 `xml:"CtrlSum,omitempty"`
			InitgPty struct {
				Nm      string `xml:"Nm,omitempty"`
				PstlAdr struct {
					AdrTp struct {
						Cd    string `xml:"Cd,omitempty" validate:"in=[AddressType2Code]"`
						Prtry struct {
							Id      string `xml:"Id,omitempty"`
							Issr    string `xml:"Issr,omitempty"`
							SchmeNm string `xml:"SchmeNm,omitempty"`
						} `xml:"Prtry,omitempty"`
					} `xml:"AdrTp,omitempty"`
					Dept        string `xml:"Dept,omitempty"`
					SubDept     string `xml:"SubDept,omitempty"`
					StrtNm      string `xml:"StrtNm,omitempty"`
					BldgNb      string `xml:"BldgNb,omitempty"`
					BldgNm      string `xml:"BldgNm,omitempty"`
					Flr         string `xml:"Flr,omitempty"`
					PstBx       string `xml:"PstBx,omitempty"`
					Room        string `xml:"Room,omitempty"`
					PstCd       string `xml:"PstCd,omitempty"`
					TwnNm       string `xml:"TwnNm,omitempty"`
					TwnLctnNm   string `xml:"TwnLctnNm,omitempty"`
					DstrctNm    string `xml:"DstrctNm,omitempty"`
					CtrySubDvsn string `xml:"CtrySubDvsn,omitempty"`
					Ctry        string `xml:"Ctry,omitempty"`
					AdrLine     string `xml:"AdrLine,omitempty"`
				} `xml:"PstlAdr,omitempty"`
				Id struct {
					OrgId struct {
						AnyBIC string `xml:"AnyBIC,omitempty"`
						LEI    string `xml:"LEI,omitempty"`
						Othr   struct {
							Id      string `xml:"Id,omitempty"`
							SchmeNm struct {
								Cd    string `xml:"Cd,omitempty" validate:"in=[ExternalOrganisationIdentification1Code]"`
								Prtry string `xml:"Prtry,omitempty"`
							} `xml:"SchmeNm,omitempty"`
							Issr string `xml:"Issr,omitempty"`
						} `xml:"Othr,omitempty"`
					} `xml:"OrgId,omitempty"`
					PrvtId struct {
						DtAndPlcOfBirth struct {
							BirthDt     string `xml:"BirthDt,omitempty"`
							PrvcOfBirth string `xml:"PrvcOfBirth,omitempty"`
							CityOfBirth string `xml:"CityOfBirth,omitempty"`
							CtryOfBirth string `xml:"CtryOfBirth,omitempty"`
						} `xml:"DtAndPlcOfBirth,omitempty"`
						Othr struct {
							Id      string `xml:"Id,omitempty"`
							SchmeNm struct {
								Cd    string `xml:"Cd,omitempty" validate:"in=[ExternalPersonIdentification1Code]"`
								Prtry string `xml:"Prtry,omitempty"`
							} `xml:"SchmeNm,omitempty"`
							Issr string `xml:"Issr,omitempty"`
						} `xml:"Othr,omitempty"`
					} `xml:"PrvtId,omitempty"`
				} `xml:"Id,omitempty"`
				CtryOfRes string `xml:"CtryOfRes,omitempty"`
				CtctDtls  struct {
					NmPrfx    string `xml:"NmPrfx,omitempty"`
					Nm        string `xml:"Nm,omitempty"`
					PhneNb    string `xml:"PhneNb,omitempty"`
					MobNb     string `xml:"MobNb,omitempty"`
					FaxNb     string `xml:"FaxNb,omitempty"`
					EmailAdr  string `xml:"EmailAdr,omitempty"`
					EmailPurp string `xml:"EmailPurp,omitempty"`
					JobTitl   string `xml:"JobTitl,omitempty"`
					Rspnsblty string `xml:"Rspnsblty,omitempty"`
					Dept      string `xml:"Dept,omitempty"`
					Othr      struct {
						ChanlTp string `xml:"ChanlTp,omitempty"`
						Id      string `xml:"Id,omitempty"`
					} `xml:"Othr,omitempty"`
					PrefrdMtd string `xml:"PrefrdMtd,omitempty"`
				} `xml:"CtctDtls,omitempty"`
			} `xml:"InitgPty,omitempty"`
		} `xml:"GrpHdr,omitempty"`
		PmtInf struct {
			PmtInfId   string `xml:"PmtInfId,omitempty"`
			PmtMtd     string `xml:"PmtMtd,omitempty"`
			ReqdAdvcTp struct {
				CdtAdvc struct {
					Cd    string `xml:"Cd,omitempty"`
					Prtry string `xml:"Prtry,omitempty"`
				} `xml:"CdtAdvc,omitempty"`
				DbtAdvc struct {
					Cd    string `xml:"Cd,omitempty"`
					Prtry string `xml:"Prtry,omitempty"`
				} `xml:"DbtAdvc,omitempty"`
			} `xml:"ReqdAdvcTp,omitempty"`
			PmtTpInf struct {
				InstrPrty string `xml:"InstrPrty,omitempty"`
				SvcLvl    struct {
					Cd    string `xml:"Cd,omitempty" validate:"in=[ExternalServiceLevel1Code]"`
					Prtry string `xml:"Prtry,omitempty"`
				} `xml:"SvcLvl,omitempty"`
				LclInstrm struct {
					Cd    string `xml:"Cd,omitempty" validate:"in=[ExternalLocalInstrument1Code]"`
					Prtry string `xml:"Prtry,omitempty"`
				} `xml:"LclInstrm,omitempty"`
				CtgyPurp struct {
					Cd    string `xml:"Cd,omitempty" validate:"in=[ExternalCategoryPurpose1Code]"`
					Prtry string `xml:"Prtry,omitempty"`
				} `xml:"CtgyPurp,omitempty"`
			} `xml:"PmtTpInf,omitempty"`
			ReqdExctnDt struct {
				Dt   string `xml:"Dt,omitempty"`
				DtTm string `xml:"DtTm,omitempty"`
			} `xml:"ReqdExctnDt,omitempty"`
			XpryDt struct {
				Dt   string `xml:"Dt,omitempty"`
				DtTm string `xml:"DtTm,omitempty"`
			} `xml:"XpryDt,omitempty"`
			PmtCond struct {
				AmtModAllwd   bool   `xml:"AmtModAllwd,omitempty"`
				EarlyPmtAllwd bool   `xml:"EarlyPmtAllwd,omitempty"`
				DelyPnlty     string `xml:"DelyPnlty,omitempty"`
				ImdtPmtRbt    struct {
					Amt  AmtField `xml:"Amt,omitempty"`
					Rate float64  `xml:"Rate,omitempty"`
				} `xml:"ImdtPmtRbt,omitempty"`
				GrntedPmtReqd bool `xml:"GrntedPmtReqd,omitempty"`
			} `xml:"PmtCond,omitempty"`
			Dbtr struct {
				Nm      string `xml:"Nm,omitempty"`
				PstlAdr struct {
					AdrTp struct {
						Cd    string `xml:"Cd,omitempty" validate:"in=[AddressType2Code]"`
						Prtry struct {
							Id      string `xml:"Id,omitempty"`
							Issr    string `xml:"Issr,omitempty"`
							SchmeNm string `xml:"SchmeNm,omitempty"`
						} `xml:"Prtry,omitempty"`
					} `xml:"AdrTp,omitempty"`
					Dept        string `xml:"Dept,omitempty"`
					SubDept     string `xml:"SubDept,omitempty"`
					StrtNm      string `xml:"StrtNm,omitempty"`
					BldgNb      string `xml:"BldgNb,omitempty"`
					BldgNm      string `xml:"BldgNm,omitempty"`
					Flr         string `xml:"Flr,omitempty"`
					PstBx       string `xml:"PstBx,omitempty"`
					Room        string `xml:"Room,omitempty"`
					PstCd       string `xml:"PstCd,omitempty"`
					TwnNm       string `xml:"TwnNm,omitempty"`
					TwnLctnNm   string `xml:"TwnLctnNm,omitempty"`
					DstrctNm    string `xml:"DstrctNm,omitempty"`
					CtrySubDvsn string `xml:"CtrySubDvsn,omitempty"`
					Ctry        string `xml:"Ctry,omitempty"`
					AdrLine     string `xml:"AdrLine,omitempty"`
				} `xml:"PstlAdr,omitempty"`
				Id struct {
					OrgId struct {
						AnyBIC string `xml:"AnyBIC,omitempty"`
						LEI    string `xml:"LEI,omitempty"`
						Othr   struct {
							Id      string `xml:"Id,omitempty"`
							SchmeNm struct {
								Cd    string `xml:"Cd,omitempty" validate:"in=[ExternalOrganisationIdentification1Code]"`
								Prtry string `xml:"Prtry,omitempty"`
							} `xml:"SchmeNm,omitempty"`
							Issr string `xml:"Issr,omitempty"`
						} `xml:"Othr,omitempty"`
					} `xml:"OrgId,omitempty"`
					PrvtId struct {
						DtAndPlcOfBirth struct {
							BirthDt     string `xml:"BirthDt,omitempty"`
							PrvcOfBirth string `xml:"PrvcOfBirth,omitempty"`
							CityOfBirth string `xml:"CityOfBirth,omitempty"`
							CtryOfBirth string `xml:"CtryOfBirth,omitempty"`
						} `xml:"DtAndPlcOfBirth,omitempty"`
						Othr struct {
							Id      string `xml:"Id,omitempty"`
							SchmeNm struct {
								Cd    string `xml:"Cd,omitempty" validate:"in=[ExternalPersonIdentification1Code]"`
								Prtry string `xml:"Prtry,omitempty"`
							} `xml:"SchmeNm,omitempty"`
							Issr string `xml:"Issr,omitempty"`
						} `xml:"Othr,omitempty"`
					} `xml:"PrvtId,omitempty"`
				} `xml:"Id,omitempty"`
				CtryOfRes string `xml:"CtryOfRes,omitempty"`
				CtctDtls  struct {
					NmPrfx    string `xml:"NmPrfx,omitempty"`
					Nm        string `xml:"Nm,omitempty"`
					PhneNb    string `xml:"PhneNb,omitempty"`
					MobNb     string `xml:"MobNb,omitempty"`
					FaxNb     string `xml:"FaxNb,omitempty"`
					EmailAdr  string `xml:"EmailAdr,omitempty"`
					EmailPurp string `xml:"EmailPurp,omitempty"`
					JobTitl   string `xml:"JobTitl,omitempty"`
					Rspnsblty string `xml:"Rspnsblty,omitempty"`
					Dept      string `xml:"Dept,omitempty"`
					Othr      struct {
						ChanlTp string `xml:"ChanlTp,omitempty"`
						Id      string `xml:"Id,omitempty"`
					} `xml:"Othr,omitempty"`
					PrefrdMtd string `xml:"PrefrdMtd,omitempty"`
				} `xml:"CtctDtls,omitempty"`
			} `xml:"Dbtr,omitempty"`
			DbtrAcct struct {
				Id struct {
					IBAN string `xml:"IBAN,omitempty"`
					Othr struct {
						Id      string `xml:"Id,omitempty"`
						SchmeNm struct {
							Cd    string `xml:"Cd,omitempty" validate:"in=[ExternalAccountIdentification1Code]"`
							Prtry string `xml:"Prtry,omitempty"`
						} `xml:"SchmeNm,omitempty"`
						Issr string `xml:"Issr,omitempty"`
					} `xml:"Othr,omitempty"`
				} `xml:"Id,omitempty"`
				Tp struct {
					Cd    string `xml:"Cd,omitempty" validate:"in=[ExternalCashAccountType1Code]"`
					Prtry string `xml:"Prtry,omitempty"`
				} `xml:"Tp,omitempty"`
				Ccy  string `xml:"Ccy,omitempty"`
				Nm   string `xml:"Nm,omitempty"`
				Prxy struct {
					Tp struct {
						Cd    string `xml:"Cd,omitempty" validate:"in=[ExternalProxyAccountType1Code]"`
						Prtry string `xml:"Prtry,omitempty"`
					} `xml:"Tp,omitempty"`
					Id string `xml:"Id,omitempty"`
				} `xml:"Prxy,omitempty"`
			} `xml:"DbtrAcct,omitempty"`
			DbtrAgt struct {
				FinInstnId struct {
					BICFI       string `xml:"BICFI,omitempty"`
					ClrSysMmbId struct {
						ClrSysId struct {
							Cd    string `xml:"Cd,omitempty" validate:"in=[ExternalClearingSystemIdentification1Code]"`
							Prtry string `xml:"Prtry,omitempty"`
						} `xml:"ClrSysId,omitempty"`
						MmbId string `xml:"MmbId,omitempty"`
					} `xml:"ClrSysMmbId,omitempty"`
					LEI     string `xml:"LEI,omitempty"`
					Nm      string `xml:"Nm,omitempty"`
					PstlAdr struct {
						AdrTp struct {
							Cd    string `xml:"Cd,omitempty" validate:"in=[AddressType2Code]"`
							Prtry struct {
								Id      string `xml:"Id,omitempty"`
								Issr    string `xml:"Issr,omitempty"`
								SchmeNm string `xml:"SchmeNm,omitempty"`
							} `xml:"Prtry,omitempty"`
						} `xml:"AdrTp,omitempty"`
						Dept        string `xml:"Dept,omitempty"`
						SubDept     string `xml:"SubDept,omitempty"`
						StrtNm      string `xml:"StrtNm,omitempty"`
						BldgNb      string `xml:"BldgNb,omitempty"`
						BldgNm      string `xml:"BldgNm,omitempty"`
						Flr         string `xml:"Flr,omitempty"`
						PstBx       string `xml:"PstBx,omitempty"`
						Room        string `xml:"Room,omitempty"`
						PstCd       string `xml:"PstCd,omitempty"`
						TwnNm       string `xml:"TwnNm,omitempty"`
						TwnLctnNm   string `xml:"TwnLctnNm,omitempty"`
						DstrctNm    string `xml:"DstrctNm,omitempty"`
						CtrySubDvsn string `xml:"CtrySubDvsn,omitempty"`
						Ctry        string `xml:"Ctry,omitempty"`
						AdrLine     string `xml:"AdrLine,omitempty"`
					} `xml:"PstlAdr,omitempty"`
					Othr struct {
						Id      string `xml:"Id,omitempty"`
						SchmeNm struct {
							Cd    string `xml:"Cd,omitempty" validate:"in=[ExternalFinancialInstitutionIdentification1Code]"`
							Prtry string `xml:"Prtry,omitempty"`
						} `xml:"SchmeNm,omitempty"`
						Issr string `xml:"Issr,omitempty"`
					} `xml:"Othr,omitempty"`
				} `xml:"FinInstnId,omitempty"`
				BrnchId struct {
					Id      string `xml:"Id,omitempty"`
					LEI     string `xml:"LEI,omitempty"`
					Nm      string `xml:"Nm,omitempty"`
					PstlAdr struct {
						AdrTp struct {
							Cd    string `xml:"Cd,omitempty" validate:"in=[AddressType2Code]"`
							Prtry struct {
								Id      string `xml:"Id,omitempty"`
								Issr    string `xml:"Issr,omitempty"`
								SchmeNm string `xml:"SchmeNm,omitempty"`
							} `xml:"Prtry,omitempty"`
						} `xml:"AdrTp,omitempty"`
						Dept        string `xml:"Dept,omitempty"`
						SubDept     string `xml:"SubDept,omitempty"`
						StrtNm      string `xml:"StrtNm,omitempty"`
						BldgNb      string `xml:"BldgNb,omitempty"`
						BldgNm      string `xml:"BldgNm,omitempty"`
						Flr         string `xml:"Flr,omitempty"`
						PstBx       string `xml:"PstBx,omitempty"`
						Room        string `xml:"Room,omitempty"`
						PstCd       string `xml:"PstCd,omitempty"`
						TwnNm       string `xml:"TwnNm,omitempty"`
						TwnLctnNm   string `xml:"TwnLctnNm,omitempty"`
						DstrctNm    string `xml:"DstrctNm,omitempty"`
						CtrySubDvsn string `xml:"CtrySubDvsn,omitempty"`
						Ctry        string `xml:"Ctry,omitempty"`
						AdrLine     string `xml:"AdrLine,omitempty"`
					} `xml:"PstlAdr,omitempty"`
				} `xml:"BrnchId,omitempty"`
			} `xml:"DbtrAgt,omitempty"`
			UltmtDbtr struct {
				Nm      string `xml:"Nm,omitempty"`
				PstlAdr struct {
					AdrTp struct {
						Cd    string `xml:"Cd,omitempty" validate:"in=[AddressType2Code]"`
						Prtry struct {
							Id      string `xml:"Id,omitempty"`
							Issr    string `xml:"Issr,omitempty"`
							SchmeNm string `xml:"SchmeNm,omitempty"`
						} `xml:"Prtry,omitempty"`
					} `xml:"AdrTp,omitempty"`
					Dept        string `xml:"Dept,omitempty"`
					SubDept     string `xml:"SubDept,omitempty"`
					StrtNm      string `xml:"StrtNm,omitempty"`
					BldgNb      string `xml:"BldgNb,omitempty"`
					BldgNm      string `xml:"BldgNm,omitempty"`
					Flr         string `xml:"Flr,omitempty"`
					PstBx       string `xml:"PstBx,omitempty"`
					Room        string `xml:"Room,omitempty"`
					PstCd       string `xml:"PstCd,omitempty"`
					TwnNm       string `xml:"TwnNm,omitempty"`
					TwnLctnNm   string `xml:"TwnLctnNm,omitempty"`
					DstrctNm    string `xml:"DstrctNm,omitempty"`
					CtrySubDvsn string `xml:"CtrySubDvsn,omitempty"`
					Ctry        string `xml:"Ctry,omitempty"`
					AdrLine     string `xml:"AdrLine,omitempty"`
				} `xml:"PstlAdr,omitempty"`
				Id struct {
					OrgId struct {
						AnyBIC string `xml:"AnyBIC,omitempty"`
						LEI    string `xml:"LEI,omitempty"`
						Othr   struct {
							Id      string `xml:"Id,omitempty"`
							SchmeNm struct {
								Cd    string `xml:"Cd,omitempty" validate:"in=[ExternalOrganisationIdentification1Code]"`
								Prtry string `xml:"Prtry,omitempty"`
							} `xml:"SchmeNm,omitempty"`
							Issr string `xml:"Issr,omitempty"`
						} `xml:"Othr,omitempty"`
					} `xml:"OrgId,omitempty"`
					PrvtId struct {
						DtAndPlcOfBirth struct {
							BirthDt     string `xml:"BirthDt,omitempty"`
							PrvcOfBirth string `xml:"PrvcOfBirth,omitempty"`
							CityOfBirth string `xml:"CityOfBirth,omitempty"`
							CtryOfBirth string `xml:"CtryOfBirth,omitempty"`
						} `xml:"DtAndPlcOfBirth,omitempty"`
						Othr struct {
							Id      string `xml:"Id,omitempty"`
							SchmeNm struct {
								Cd    string `xml:"Cd,omitempty" validate:"in=[ExternalPersonIdentification1Code]"`
								Prtry string `xml:"Prtry,omitempty"`
							} `xml:"SchmeNm,omitempty"`
							Issr string `xml:"Issr,omitempty"`
						} `xml:"Othr,omitempty"`
					} `xml:"PrvtId,omitempty"`
				} `xml:"Id,omitempty"`
				CtryOfRes string `xml:"CtryOfRes,omitempty"`
				CtctDtls  struct {
					NmPrfx    string `xml:"NmPrfx,omitempty"`
					Nm        string `xml:"Nm,omitempty"`
					PhneNb    string `xml:"PhneNb,omitempty"`
					MobNb     string `xml:"MobNb,omitempty"`
					FaxNb     string `xml:"FaxNb,omitempty"`
					EmailAdr  string `xml:"EmailAdr,omitempty"`
					EmailPurp string `xml:"EmailPurp,omitempty"`
					JobTitl   string `xml:"JobTitl,omitempty"`
					Rspnsblty string `xml:"Rspnsblty,omitempty"`
					Dept      string `xml:"Dept,omitempty"`
					Othr      struct {
						ChanlTp string `xml:"ChanlTp,omitempty"`
						Id      string `xml:"Id,omitempty"`
					} `xml:"Othr,omitempty"`
					PrefrdMtd string `xml:"PrefrdMtd,omitempty"`
				} `xml:"CtctDtls,omitempty"`
			} `xml:"UltmtDbtr,omitempty"`
			ChrgBr   string `xml:"ChrgBr,omitempty"`
			CdtTrfTx struct {
				PmtId struct {
					InstrId    string `xml:"InstrId,omitempty"`
					EndToEndId string `xml:"EndToEndId,omitempty"`
					UETR       string `xml:"UETR,omitempty"`
				} `xml:"PmtId,omitempty"`
				PmtTpInf struct {
					InstrPrty string `xml:"InstrPrty,omitempty"`
					SvcLvl    struct {
						Cd    string `xml:"Cd,omitempty" validate:"in=[ExternalServiceLevel1Code]"`
						Prtry string `xml:"Prtry,omitempty"`
					} `xml:"SvcLvl,omitempty"`
					LclInstrm struct {
						Cd    string `xml:"Cd,omitempty" validate:"in=[ExternalLocalInstrument1Code]"`
						Prtry string `xml:"Prtry,omitempty"`
					} `xml:"LclInstrm,omitempty"`
					CtgyPurp struct {
						Cd    string `xml:"Cd,omitempty" validate:"in=[ExternalCategoryPurpose1Code]"`
						Prtry string `xml:"Prtry,omitempty"`
					} `xml:"CtgyPurp,omitempty"`
				} `xml:"PmtTpInf,omitempty"`
				PmtCond struct {
					AmtModAllwd   bool   `xml:"AmtModAllwd,omitempty"`
					EarlyPmtAllwd bool   `xml:"EarlyPmtAllwd,omitempty"`
					DelyPnlty     string `xml:"DelyPnlty,omitempty"`
					ImdtPmtRbt    struct {
						Amt  AmtField `xml:"Amt,omitempty"`
						Rate float64  `xml:"Rate,omitempty"`
					} `xml:"ImdtPmtRbt,omitempty"`
					GrntedPmtReqd bool `xml:"GrntedPmtReqd,omitempty"`
				} `xml:"PmtCond,omitempty"`
				Amt struct {
					InstdAmt AmtField `xml:"InstdAmt,omitempty"`
					EqvtAmt  struct {
						Amt      AmtField `xml:"Amt,omitempty"`
						CcyOfTrf string   `xml:"CcyOfTrf,omitempty"`
					} `xml:"EqvtAmt,omitempty"`
				} `xml:"Amt,omitempty"`
				ChrgBr      string `xml:"ChrgBr,omitempty"`
				MndtRltdInf struct {
					MndtId string `xml:"MndtId,omitempty"`
					Tp     struct {
						SvcLvl struct {
							Cd    string `xml:"Cd,omitempty" validate:"in=[ExternalServiceLevel1Code]"`
							Prtry string `xml:"Prtry,omitempty"`
						} `xml:"SvcLvl,omitempty"`
						LclInstrm struct {
							Cd    string `xml:"Cd,omitempty" validate:"in=[ExternalLocalInstrument1Code]"`
							Prtry string `xml:"Prtry,omitempty"`
						} `xml:"LclInstrm,omitempty"`
						CtgyPurp struct {
							Cd    string `xml:"Cd,omitempty" validate:"in=[ExternalCategoryPurpose1Code]"`
							Prtry string `xml:"Prtry,omitempty"`
						} `xml:"CtgyPurp,omitempty"`
						Clssfctn struct {
							Cd    string `xml:"Cd,omitempty"`
							Prtry string `xml:"Prtry,omitempty"`
						} `xml:"Clssfctn,omitempty"`
					} `xml:"Tp,omitempty"`
					DtOfSgntr    string `xml:"DtOfSgntr,omitempty"`
					DtOfVrfctn   string `xml:"DtOfVrfctn,omitempty"`
					ElctrncSgntr string `xml:"ElctrncSgntr,omitempty"`
					FrstPmtDt    string `xml:"FrstPmtDt,omitempty"`
					FnlPmtDt     string `xml:"FnlPmtDt,omitempty"`
					Frqcy        struct {
						Tp  string `xml:"Tp,omitempty"`
						Prd struct {
							Tp        string `xml:"Tp,omitempty"`
							CntPerPrd int    `xml:"CntPerPrd,omitempty"`
						} `xml:"Prd,omitempty"`
						PtInTm struct {
							Tp     string `xml:"Tp,omitempty"`
							PtInTm string `xml:"PtInTm,omitempty"`
						} `xml:"PtInTm,omitempty"`
					} `xml:"Frqcy,omitempty"`
					Rsn struct {
						Cd    string `xml:"Cd,omitempty"`
						Prtry string `xml:"Prtry,omitempty"`
					} `xml:"Rsn,omitempty"`
				} `xml:"MndtRltdInf,omitempty"`
				ChqInstr struct {
					ChqTp string `xml:"ChqTp,omitempty"`
					ChqNb string `xml:"ChqNb,omitempty"`
					ChqFr struct {
						Nm  string `xml:"Nm,omitempty"`
						Adr struct {
							AdrTp struct {
								Cd    string `xml:"Cd,omitempty" validate:"in=[AddressType2Code]"`
								Prtry struct {
									Id      string `xml:"Id,omitempty"`
									Issr    string `xml:"Issr,omitempty"`
									SchmeNm string `xml:"SchmeNm,omitempty"`
								} `xml:"Prtry,omitempty"`
							} `xml:"AdrTp,omitempty"`
							Dept        string `xml:"Dept,omitempty"`
							SubDept     string `xml:"SubDept,omitempty"`
							StrtNm      string `xml:"StrtNm,omitempty"`
							BldgNb      string `xml:"BldgNb,omitempty"`
							BldgNm      string `xml:"BldgNm,omitempty"`
							Flr         string `xml:"Flr,omitempty"`
							PstBx       string `xml:"PstBx,omitempty"`
							Room        string `xml:"Room,omitempty"`
							PstCd       string `xml:"PstCd,omitempty"`
							TwnNm       string `xml:"TwnNm,omitempty"`
							TwnLctnNm   string `xml:"TwnLctnNm,omitempty"`
							DstrctNm    string `xml:"DstrctNm,omitempty"`
							CtrySubDvsn string `xml:"CtrySubDvsn,omitempty"`
							Ctry        string `xml:"Ctry,omitempty"`
							AdrLine     string `xml:"AdrLine,omitempty"`
						} `xml:"Adr,omitempty"`
					} `xml:"ChqFr,omitempty"`
					DlvryMtd struct {
						Cd    string `xml:"Cd,omitempty"`
						Prtry string `xml:"Prtry,omitempty"`
					} `xml:"DlvryMtd,omitempty"`
					DlvrTo struct {
						Nm  string `xml:"Nm,omitempty"`
						Adr struct {
							AdrTp struct {
								Cd    string `xml:"Cd,omitempty" validate:"in=[AddressType2Code]"`
								Prtry struct {
									Id      string `xml:"Id,omitempty"`
									Issr    string `xml:"Issr,omitempty"`
									SchmeNm string `xml:"SchmeNm,omitempty"`
								} `xml:"Prtry,omitempty"`
							} `xml:"AdrTp,omitempty"`
							Dept        string `xml:"Dept,omitempty"`
							SubDept     string `xml:"SubDept,omitempty"`
							StrtNm      string `xml:"StrtNm,omitempty"`
							BldgNb      string `xml:"BldgNb,omitempty"`
							BldgNm      string `xml:"BldgNm,omitempty"`
							Flr         string `xml:"Flr,omitempty"`
							PstBx       string `xml:"PstBx,omitempty"`
							Room        string `xml:"Room,omitempty"`
							PstCd       string `xml:"PstCd,omitempty"`
							TwnNm       string `xml:"TwnNm,omitempty"`
							TwnLctnNm   string `xml:"TwnLctnNm,omitempty"`
							DstrctNm    string `xml:"DstrctNm,omitempty"`
							CtrySubDvsn string `xml:"CtrySubDvsn,omitempty"`
							Ctry        string `xml:"Ctry,omitempty"`
							AdrLine     string `xml:"AdrLine,omitempty"`
						} `xml:"Adr,omitempty"`
					} `xml:"DlvrTo,omitempty"`
					InstrPrty   string `xml:"InstrPrty,omitempty"`
					ChqMtrtyDt  string `xml:"ChqMtrtyDt,omitempty"`
					FrmsCd      string `xml:"FrmsCd,omitempty"`
					MemoFld     string `xml:"MemoFld,omitempty"`
					RgnlClrZone string `xml:"RgnlClrZone,omitempty"`
					PrtLctn     string `xml:"PrtLctn,omitempty"`
					Sgntr       string `xml:"Sgntr,omitempty"`
				} `xml:"ChqInstr,omitempty"`
				UltmtDbtr struct {
					Nm      string `xml:"Nm,omitempty"`
					PstlAdr struct {
						AdrTp struct {
							Cd    string `xml:"Cd,omitempty" validate:"in=[AddressType2Code]"`
							Prtry struct {
								Id      string `xml:"Id,omitempty"`
								Issr    string `xml:"Issr,omitempty"`
								SchmeNm string `xml:"SchmeNm,omitempty"`
							} `xml:"Prtry,omitempty"`
						} `xml:"AdrTp,omitempty"`
						Dept        string `xml:"Dept,omitempty"`
						SubDept     string `xml:"SubDept,omitempty"`
						StrtNm      string `xml:"StrtNm,omitempty"`
						BldgNb      string `xml:"BldgNb,omitempty"`
						BldgNm      string `xml:"BldgNm,omitempty"`
						Flr         string `xml:"Flr,omitempty"`
						PstBx       string `xml:"PstBx,omitempty"`
						Room        string `xml:"Room,omitempty"`
						PstCd       string `xml:"PstCd,omitempty"`
						TwnNm       string `xml:"TwnNm,omitempty"`
						TwnLctnNm   string `xml:"TwnLctnNm,omitempty"`
						DstrctNm    string `xml:"DstrctNm,omitempty"`
						CtrySubDvsn string `xml:"CtrySubDvsn,omitempty"`
						Ctry        string `xml:"Ctry,omitempty"`
						AdrLine     string `xml:"AdrLine,omitempty"`
					} `xml:"PstlAdr,omitempty"`
					Id struct {
						OrgId struct {
							AnyBIC string `xml:"AnyBIC,omitempty"`
							LEI    string `xml:"LEI,omitempty"`
							Othr   struct {
								Id      string `xml:"Id,omitempty"`
								SchmeNm struct {
									Cd    string `xml:"Cd,omitempty" validate:"in=[ExternalOrganisationIdentification1Code]"`
									Prtry string `xml:"Prtry,omitempty"`
								} `xml:"SchmeNm,omitempty"`
								Issr string `xml:"Issr,omitempty"`
							} `xml:"Othr,omitempty"`
						} `xml:"OrgId,omitempty"`
						PrvtId struct {
							DtAndPlcOfBirth struct {
								BirthDt     string `xml:"BirthDt,omitempty"`
								PrvcOfBirth string `xml:"PrvcOfBirth,omitempty"`
								CityOfBirth string `xml:"CityOfBirth,omitempty"`
								CtryOfBirth string `xml:"CtryOfBirth,omitempty"`
							} `xml:"DtAndPlcOfBirth,omitempty"`
							Othr struct {
								Id      string `xml:"Id,omitempty"`
								SchmeNm struct {
									Cd    string `xml:"Cd,omitempty" validate:"in=[ExternalPersonIdentification1Code]"`
									Prtry string `xml:"Prtry,omitempty"`
								} `xml:"SchmeNm,omitempty"`
								Issr string `xml:"Issr,omitempty"`
							} `xml:"Othr,omitempty"`
						} `xml:"PrvtId,omitempty"`
					} `xml:"Id,omitempty"`
					CtryOfRes string `xml:"CtryOfRes,omitempty"`
					CtctDtls  struct {
						NmPrfx    string `xml:"NmPrfx,omitempty"`
						Nm        string `xml:"Nm,omitempty"`
						PhneNb    string `xml:"PhneNb,omitempty"`
						MobNb     string `xml:"MobNb,omitempty"`
						FaxNb     string `xml:"FaxNb,omitempty"`
						EmailAdr  string `xml:"EmailAdr,omitempty"`
						EmailPurp string `xml:"EmailPurp,omitempty"`
						JobTitl   string `xml:"JobTitl,omitempty"`
						Rspnsblty string `xml:"Rspnsblty,omitempty"`
						Dept      string `xml:"Dept,omitempty"`
						Othr      struct {
							ChanlTp string `xml:"ChanlTp,omitempty"`
							Id      string `xml:"Id,omitempty"`
						} `xml:"Othr,omitempty"`
						PrefrdMtd string `xml:"PrefrdMtd,omitempty"`
					} `xml:"CtctDtls,omitempty"`
				} `xml:"UltmtDbtr,omitempty"`
				IntrmyAgt1 struct {
					FinInstnId struct {
						BICFI       string `xml:"BICFI,omitempty"`
						ClrSysMmbId struct {
							ClrSysId struct {
								Cd    string `xml:"Cd,omitempty" validate:"in=[ExternalClearingSystemIdentification1Code]"`
								Prtry string `xml:"Prtry,omitempty"`
							} `xml:"ClrSysId,omitempty"`
							MmbId string `xml:"MmbId,omitempty"`
						} `xml:"ClrSysMmbId,omitempty"`
						LEI     string `xml:"LEI,omitempty"`
						Nm      string `xml:"Nm,omitempty"`
						PstlAdr struct {
							AdrTp struct {
								Cd    string `xml:"Cd,omitempty" validate:"in=[AddressType2Code]"`
								Prtry struct {
									Id      string `xml:"Id,omitempty"`
									Issr    string `xml:"Issr,omitempty"`
									SchmeNm string `xml:"SchmeNm,omitempty"`
								} `xml:"Prtry,omitempty"`
							} `xml:"AdrTp,omitempty"`
							Dept        string `xml:"Dept,omitempty"`
							SubDept     string `xml:"SubDept,omitempty"`
							StrtNm      string `xml:"StrtNm,omitempty"`
							BldgNb      string `xml:"BldgNb,omitempty"`
							BldgNm      string `xml:"BldgNm,omitempty"`
							Flr         string `xml:"Flr,omitempty"`
							PstBx       string `xml:"PstBx,omitempty"`
							Room        string `xml:"Room,omitempty"`
							PstCd       string `xml:"PstCd,omitempty"`
							TwnNm       string `xml:"TwnNm,omitempty"`
							TwnLctnNm   string `xml:"TwnLctnNm,omitempty"`
							DstrctNm    string `xml:"DstrctNm,omitempty"`
							CtrySubDvsn string `xml:"CtrySubDvsn,omitempty"`
							Ctry        string `xml:"Ctry,omitempty"`
							AdrLine     string `xml:"AdrLine,omitempty"`
						} `xml:"PstlAdr,omitempty"`
						Othr struct {
							Id      string `xml:"Id,omitempty"`
							SchmeNm struct {
								Cd    string `xml:"Cd,omitempty" validate:"in=[ExternalFinancialInstitutionIdentification1Code]"`
								Prtry string `xml:"Prtry,omitempty"`
							} `xml:"SchmeNm,omitempty"`
							Issr string `xml:"Issr,omitempty"`
						} `xml:"Othr,omitempty"`
					} `xml:"FinInstnId,omitempty"`
					BrnchId struct {
						Id      string `xml:"Id,omitempty"`
						LEI     string `xml:"LEI,omitempty"`
						Nm      string `xml:"Nm,omitempty"`
						PstlAdr struct {
							AdrTp struct {
								Cd    string `xml:"Cd,omitempty" validate:"in=[AddressType2Code]"`
								Prtry struct {
									Id      string `xml:"Id,omitempty"`
									Issr    string `xml:"Issr,omitempty"`
									SchmeNm string `xml:"SchmeNm,omitempty"`
								} `xml:"Prtry,omitempty"`
							} `xml:"AdrTp,omitempty"`
							Dept        string `xml:"Dept,omitempty"`
							SubDept     string `xml:"SubDept,omitempty"`
							StrtNm      string `xml:"StrtNm,omitempty"`
							BldgNb      string `xml:"BldgNb,omitempty"`
							BldgNm      string `xml:"BldgNm,omitempty"`
							Flr         string `xml:"Flr,omitempty"`
							PstBx       string `xml:"PstBx,omitempty"`
							Room        string `xml:"Room,omitempty"`
							PstCd       string `xml:"PstCd,omitempty"`
							TwnNm       string `xml:"TwnNm,omitempty"`
							TwnLctnNm   string `xml:"TwnLctnNm,omitempty"`
							DstrctNm    string `xml:"DstrctNm,omitempty"`
							CtrySubDvsn string `xml:"CtrySubDvsn,omitempty"`
							Ctry        string `xml:"Ctry,omitempty"`
							AdrLine     string `xml:"AdrLine,omitempty"`
						} `xml:"PstlAdr,omitempty"`
					} `xml:"BrnchId,omitempty"`
				} `xml:"IntrmyAgt1,omitempty"`
				IntrmyAgt2 struct {
					FinInstnId struct {
						BICFI       string `xml:"BICFI,omitempty"`
						ClrSysMmbId struct {
							ClrSysId struct {
								Cd    string `xml:"Cd,omitempty" validate:"in=[ExternalClearingSystemIdentification1Code]"`
								Prtry string `xml:"Prtry,omitempty"`
							} `xml:"ClrSysId,omitempty"`
							MmbId string `xml:"MmbId,omitempty"`
						} `xml:"ClrSysMmbId,omitempty"`
						LEI     string `xml:"LEI,omitempty"`
						Nm      string `xml:"Nm,omitempty"`
						PstlAdr struct {
							AdrTp struct {
								Cd    string `xml:"Cd,omitempty" validate:"in=[AddressType2Code]"`
								Prtry struct {
									Id      string `xml:"Id,omitempty"`
									Issr    string `xml:"Issr,omitempty"`
									SchmeNm string `xml:"SchmeNm,omitempty"`
								} `xml:"Prtry,omitempty"`
							} `xml:"AdrTp,omitempty"`
							Dept        string `xml:"Dept,omitempty"`
							SubDept     string `xml:"SubDept,omitempty"`
							StrtNm      string `xml:"StrtNm,omitempty"`
							BldgNb      string `xml:"BldgNb,omitempty"`
							BldgNm      string `xml:"BldgNm,omitempty"`
							Flr         string `xml:"Flr,omitempty"`
							PstBx       string `xml:"PstBx,omitempty"`
							Room        string `xml:"Room,omitempty"`
							PstCd       string `xml:"PstCd,omitempty"`
							TwnNm       string `xml:"TwnNm,omitempty"`
							TwnLctnNm   string `xml:"TwnLctnNm,omitempty"`
							DstrctNm    string `xml:"DstrctNm,omitempty"`
							CtrySubDvsn string `xml:"CtrySubDvsn,omitempty"`
							Ctry        string `xml:"Ctry,omitempty"`
							AdrLine     string `xml:"AdrLine,omitempty"`
						} `xml:"PstlAdr,omitempty"`
						Othr struct {
							Id      string `xml:"Id,omitempty"`
							SchmeNm struct {
								Cd    string `xml:"Cd,omitempty" validate:"in=[ExternalFinancialInstitutionIdentification1Code]"`
								Prtry string `xml:"Prtry,omitempty"`
							} `xml:"SchmeNm,omitempty"`
							Issr string `xml:"Issr,omitempty"`
						} `xml:"Othr,omitempty"`
					} `xml:"FinInstnId,omitempty"`
					BrnchId struct {
						Id      string `xml:"Id,omitempty"`
						LEI     string `xml:"LEI,omitempty"`
						Nm      string `xml:"Nm,omitempty"`
						PstlAdr struct {
							AdrTp struct {
								Cd    string `xml:"Cd,omitempty" validate:"in=[AddressType2Code]"`
								Prtry struct {
									Id      string `xml:"Id,omitempty"`
									Issr    string `xml:"Issr,omitempty"`
									SchmeNm string `xml:"SchmeNm,omitempty"`
								} `xml:"Prtry,omitempty"`
							} `xml:"AdrTp,omitempty"`
							Dept        string `xml:"Dept,omitempty"`
							SubDept     string `xml:"SubDept,omitempty"`
							StrtNm      string `xml:"StrtNm,omitempty"`
							BldgNb      string `xml:"BldgNb,omitempty"`
							BldgNm      string `xml:"BldgNm,omitempty"`
							Flr         string `xml:"Flr,omitempty"`
							PstBx       string `xml:"PstBx,omitempty"`
							Room        string `xml:"Room,omitempty"`
							PstCd       string `xml:"PstCd,omitempty"`
							TwnNm       string `xml:"TwnNm,omitempty"`
							TwnLctnNm   string `xml:"TwnLctnNm,omitempty"`
							DstrctNm    string `xml:"DstrctNm,omitempty"`
							CtrySubDvsn string `xml:"CtrySubDvsn,omitempty"`
							Ctry        string `xml:"Ctry,omitempty"`
							AdrLine     string `xml:"AdrLine,omitempty"`
						} `xml:"PstlAdr,omitempty"`
					} `xml:"BrnchId,omitempty"`
				} `xml:"IntrmyAgt2,omitempty"`
				IntrmyAgt3 struct {
					FinInstnId struct {
						BICFI       string `xml:"BICFI,omitempty"`
						ClrSysMmbId struct {
							ClrSysId struct {
								Cd    string `xml:"Cd,omitempty" validate:"in=[ExternalClearingSystemIdentification1Code]"`
								Prtry string `xml:"Prtry,omitempty"`
							} `xml:"ClrSysId,omitempty"`
							MmbId string `xml:"MmbId,omitempty"`
						} `xml:"ClrSysMmbId,omitempty"`
						LEI     string `xml:"LEI,omitempty"`
						Nm      string `xml:"Nm,omitempty"`
						PstlAdr struct {
							AdrTp struct {
								Cd    string `xml:"Cd,omitempty" validate:"in=[AddressType2Code]"`
								Prtry struct {
									Id      string `xml:"Id,omitempty"`
									Issr    string `xml:"Issr,omitempty"`
									SchmeNm string `xml:"SchmeNm,omitempty"`
								} `xml:"Prtry,omitempty"`
							} `xml:"AdrTp,omitempty"`
							Dept        string `xml:"Dept,omitempty"`
							SubDept     string `xml:"SubDept,omitempty"`
							StrtNm      string `xml:"StrtNm,omitempty"`
							BldgNb      string `xml:"BldgNb,omitempty"`
							BldgNm      string `xml:"BldgNm,omitempty"`
							Flr         string `xml:"Flr,omitempty"`
							PstBx       string `xml:"PstBx,omitempty"`
							Room        string `xml:"Room,omitempty"`
							PstCd       string `xml:"PstCd,omitempty"`
							TwnNm       string `xml:"TwnNm,omitempty"`
							TwnLctnNm   string `xml:"TwnLctnNm,omitempty"`
							DstrctNm    string `xml:"DstrctNm,omitempty"`
							CtrySubDvsn string `xml:"CtrySubDvsn,omitempty"`
							Ctry        string `xml:"Ctry,omitempty"`
							AdrLine     string `xml:"AdrLine,omitempty"`
						} `xml:"PstlAdr,omitempty"`
						Othr struct {
							Id      string `xml:"Id,omitempty"`
							SchmeNm struct {
								Cd    string `xml:"Cd,omitempty" validate:"in=[ExternalFinancialInstitutionIdentification1Code]"`
								Prtry string `xml:"Prtry,omitempty"`
							} `xml:"SchmeNm,omitempty"`
							Issr string `xml:"Issr,omitempty"`
						} `xml:"Othr,omitempty"`
					} `xml:"FinInstnId,omitempty"`
					BrnchId struct {
						Id      string `xml:"Id,omitempty"`
						LEI     string `xml:"LEI,omitempty"`
						Nm      string `xml:"Nm,omitempty"`
						PstlAdr struct {
							AdrTp struct {
								Cd    string `xml:"Cd,omitempty" validate:"in=[AddressType2Code]"`
								Prtry struct {
									Id      string `xml:"Id,omitempty"`
									Issr    string `xml:"Issr,omitempty"`
									SchmeNm string `xml:"SchmeNm,omitempty"`
								} `xml:"Prtry,omitempty"`
							} `xml:"AdrTp,omitempty"`
							Dept        string `xml:"Dept,omitempty"`
							SubDept     string `xml:"SubDept,omitempty"`
							StrtNm      string `xml:"StrtNm,omitempty"`
							BldgNb      string `xml:"BldgNb,omitempty"`
							BldgNm      string `xml:"BldgNm,omitempty"`
							Flr         string `xml:"Flr,omitempty"`
							PstBx       string `xml:"PstBx,omitempty"`
							Room        string `xml:"Room,omitempty"`
							PstCd       string `xml:"PstCd,omitempty"`
							TwnNm       string `xml:"TwnNm,omitempty"`
							TwnLctnNm   string `xml:"TwnLctnNm,omitempty"`
							DstrctNm    string `xml:"DstrctNm,omitempty"`
							CtrySubDvsn string `xml:"CtrySubDvsn,omitempty"`
							Ctry        string `xml:"Ctry,omitempty"`
							AdrLine     string `xml:"AdrLine,omitempty"`
						} `xml:"PstlAdr,omitempty"`
					} `xml:"BrnchId,omitempty"`
				} `xml:"IntrmyAgt3,omitempty"`
				CdtrAgt struct {
					FinInstnId struct {
						BICFI       string `xml:"BICFI,omitempty"`
						ClrSysMmbId struct {
							ClrSysId struct {
								Cd    string `xml:"Cd,omitempty" validate:"in=[ExternalClearingSystemIdentification1Code]"`
								Prtry string `xml:"Prtry,omitempty"`
							} `xml:"ClrSysId,omitempty"`
							MmbId string `xml:"MmbId,omitempty"`
						} `xml:"ClrSysMmbId,omitempty"`
						LEI     string `xml:"LEI,omitempty"`
						Nm      string `xml:"Nm,omitempty"`
						PstlAdr struct {
							AdrTp struct {
								Cd    string `xml:"Cd,omitempty" validate:"in=[AddressType2Code]"`
								Prtry struct {
									Id      string `xml:"Id,omitempty"`
									Issr    string `xml:"Issr,omitempty"`
									SchmeNm string `xml:"SchmeNm,omitempty"`
								} `xml:"Prtry,omitempty"`
							} `xml:"AdrTp,omitempty"`
							Dept        string `xml:"Dept,omitempty"`
							SubDept     string `xml:"SubDept,omitempty"`
							StrtNm      string `xml:"StrtNm,omitempty"`
							BldgNb      string `xml:"BldgNb,omitempty"`
							BldgNm      string `xml:"BldgNm,omitempty"`
							Flr         string `xml:"Flr,omitempty"`
							PstBx       string `xml:"PstBx,omitempty"`
							Room        string `xml:"Room,omitempty"`
							PstCd       string `xml:"PstCd,omitempty"`
							TwnNm       string `xml:"TwnNm,omitempty"`
							TwnLctnNm   string `xml:"TwnLctnNm,omitempty"`
							DstrctNm    string `xml:"DstrctNm,omitempty"`
							CtrySubDvsn string `xml:"CtrySubDvsn,omitempty"`
							Ctry        string `xml:"Ctry,omitempty"`
							AdrLine     string `xml:"AdrLine,omitempty"`
						} `xml:"PstlAdr,omitempty"`
						Othr struct {
							Id      string `xml:"Id,omitempty"`
							SchmeNm struct {
								Cd    string `xml:"Cd,omitempty" validate:"in=[ExternalFinancialInstitutionIdentification1Code]"`
								Prtry string `xml:"Prtry,omitempty"`
							} `xml:"SchmeNm,omitempty"`
							Issr string `xml:"Issr,omitempty"`
						} `xml:"Othr,omitempty"`
					} `xml:"FinInstnId,omitempty"`
					BrnchId struct {
						Id      string `xml:"Id,omitempty"`
						LEI     string `xml:"LEI,omitempty"`
						Nm      string `xml:"Nm,omitempty"`
						PstlAdr struct {
							AdrTp struct {
								Cd    string `xml:"Cd,omitempty" validate:"in=[AddressType2Code]"`
								Prtry struct {
									Id      string `xml:"Id,omitempty"`
									Issr    string `xml:"Issr,omitempty"`
									SchmeNm string `xml:"SchmeNm,omitempty"`
								} `xml:"Prtry,omitempty"`
							} `xml:"AdrTp,omitempty"`
							Dept        string `xml:"Dept,omitempty"`
							SubDept     string `xml:"SubDept,omitempty"`
							StrtNm      string `xml:"StrtNm,omitempty"`
							BldgNb      string `xml:"BldgNb,omitempty"`
							BldgNm      string `xml:"BldgNm,omitempty"`
							Flr         string `xml:"Flr,omitempty"`
							PstBx       string `xml:"PstBx,omitempty"`
							Room        string `xml:"Room,omitempty"`
							PstCd       string `xml:"PstCd,omitempty"`
							TwnNm       string `xml:"TwnNm,omitempty"`
							TwnLctnNm   string `xml:"TwnLctnNm,omitempty"`
							DstrctNm    string `xml:"DstrctNm,omitempty"`
							CtrySubDvsn string `xml:"CtrySubDvsn,omitempty"`
							Ctry        string `xml:"Ctry,omitempty"`
							AdrLine     string `xml:"AdrLine,omitempty"`
						} `xml:"PstlAdr,omitempty"`
					} `xml:"BrnchId,omitempty"`
				} `xml:"CdtrAgt,omitempty"`
				Cdtr struct {
					Nm      string `xml:"Nm,omitempty"`
					PstlAdr struct {
						AdrTp struct {
							Cd    string `xml:"Cd,omitempty" validate:"in=[AddressType2Code]"`
							Prtry struct {
								Id      string `xml:"Id,omitempty"`
								Issr    string `xml:"Issr,omitempty"`
								SchmeNm string `xml:"SchmeNm,omitempty"`
							} `xml:"Prtry,omitempty"`
						} `xml:"AdrTp,omitempty"`
						Dept        string `xml:"Dept,omitempty"`
						SubDept     string `xml:"SubDept,omitempty"`
						StrtNm      string `xml:"StrtNm,omitempty"`
						BldgNb      string `xml:"BldgNb,omitempty"`
						BldgNm      string `xml:"BldgNm,omitempty"`
						Flr         string `xml:"Flr,omitempty"`
						PstBx       string `xml:"PstBx,omitempty"`
						Room        string `xml:"Room,omitempty"`
						PstCd       string `xml:"PstCd,omitempty"`
						TwnNm       string `xml:"TwnNm,omitempty"`
						TwnLctnNm   string `xml:"TwnLctnNm,omitempty"`
						DstrctNm    string `xml:"DstrctNm,omitempty"`
						CtrySubDvsn string `xml:"CtrySubDvsn,omitempty"`
						Ctry        string `xml:"Ctry,omitempty"`
						AdrLine     string `xml:"AdrLine,omitempty"`
					} `xml:"PstlAdr,omitempty"`
					Id struct {
						OrgId struct {
							AnyBIC string `xml:"AnyBIC,omitempty"`
							LEI    string `xml:"LEI,omitempty"`
							Othr   struct {
								Id      string `xml:"Id,omitempty"`
								SchmeNm struct {
									Cd    string `xml:"Cd,omitempty" validate:"in=[ExternalOrganisationIdentification1Code]"`
									Prtry string `xml:"Prtry,omitempty"`
								} `xml:"SchmeNm,omitempty"`
								Issr string `xml:"Issr,omitempty"`
							} `xml:"Othr,omitempty"`
						} `xml:"OrgId,omitempty"`
						PrvtId struct {
							DtAndPlcOfBirth struct {
								BirthDt     string `xml:"BirthDt,omitempty"`
								PrvcOfBirth string `xml:"PrvcOfBirth,omitempty"`
								CityOfBirth string `xml:"CityOfBirth,omitempty"`
								CtryOfBirth string `xml:"CtryOfBirth,omitempty"`
							} `xml:"DtAndPlcOfBirth,omitempty"`
							Othr struct {
								Id      string `xml:"Id,omitempty"`
								SchmeNm struct {
									Cd    string `xml:"Cd,omitempty" validate:"in=[ExternalPersonIdentification1Code]"`
									Prtry string `xml:"Prtry,omitempty"`
								} `xml:"SchmeNm,omitempty"`
								Issr string `xml:"Issr,omitempty"`
							} `xml:"Othr,omitempty"`
						} `xml:"PrvtId,omitempty"`
					} `xml:"Id,omitempty"`
					CtryOfRes string `xml:"CtryOfRes,omitempty"`
					CtctDtls  struct {
						NmPrfx    string `xml:"NmPrfx,omitempty"`
						Nm        string `xml:"Nm,omitempty"`
						PhneNb    string `xml:"PhneNb,omitempty"`
						MobNb     string `xml:"MobNb,omitempty"`
						FaxNb     string `xml:"FaxNb,omitempty"`
						EmailAdr  string `xml:"EmailAdr,omitempty"`
						EmailPurp string `xml:"EmailPurp,omitempty"`
						JobTitl   string `xml:"JobTitl,omitempty"`
						Rspnsblty string `xml:"Rspnsblty,omitempty"`
						Dept      string `xml:"Dept,omitempty"`
						Othr      struct {
							ChanlTp string `xml:"ChanlTp,omitempty"`
							Id      string `xml:"Id,omitempty"`
						} `xml:"Othr,omitempty"`
						PrefrdMtd string `xml:"PrefrdMtd,omitempty"`
					} `xml:"CtctDtls,omitempty"`
				} `xml:"Cdtr,omitempty"`
				CdtrAcct struct {
					Id struct {
						IBAN string `xml:"IBAN,omitempty"`
						Othr struct {
							Id      string `xml:"Id,omitempty"`
							SchmeNm struct {
								Cd    string `xml:"Cd,omitempty" validate:"in=[ExternalAccountIdentification1Code]"`
								Prtry string `xml:"Prtry,omitempty"`
							} `xml:"SchmeNm,omitempty"`
							Issr string `xml:"Issr,omitempty"`
						} `xml:"Othr,omitempty"`
					} `xml:"Id,omitempty"`
					Tp struct {
						Cd    string `xml:"Cd,omitempty" validate:"in=[ExternalCashAccountType1Code]"`
						Prtry string `xml:"Prtry,omitempty"`
					} `xml:"Tp,omitempty"`
					Ccy  string `xml:"Ccy,omitempty"`
					Nm   string `xml:"Nm,omitempty"`
					Prxy struct {
						Tp struct {
							Cd    string `xml:"Cd,omitempty" validate:"in=[ExternalProxyAccountType1Code]"`
							Prtry string `xml:"Prtry,omitempty"`
						} `xml:"Tp,omitempty"`
						Id string `xml:"Id,omitempty"`
					} `xml:"Prxy,omitempty"`
				} `xml:"CdtrAcct,omitempty"`
				UltmtCdtr struct {
					Nm      string `xml:"Nm,omitempty"`
					PstlAdr struct {
						AdrTp struct {
							Cd    string `xml:"Cd,omitempty" validate:"in=[AddressType2Code]"`
							Prtry struct {
								Id      string `xml:"Id,omitempty"`
								Issr    string `xml:"Issr,omitempty"`
								SchmeNm string `xml:"SchmeNm,omitempty"`
							} `xml:"Prtry,omitempty"`
						} `xml:"AdrTp,omitempty"`
						Dept        string `xml:"Dept,omitempty"`
						SubDept     string `xml:"SubDept,omitempty"`
						StrtNm      string `xml:"StrtNm,omitempty"`
						BldgNb      string `xml:"BldgNb,omitempty"`
						BldgNm      string `xml:"BldgNm,omitempty"`
						Flr         string `xml:"Flr,omitempty"`
						PstBx       string `xml:"PstBx,omitempty"`
						Room        string `xml:"Room,omitempty"`
						PstCd       string `xml:"PstCd,omitempty"`
						TwnNm       string `xml:"TwnNm,omitempty"`
						TwnLctnNm   string `xml:"TwnLctnNm,omitempty"`
						DstrctNm    string `xml:"DstrctNm,omitempty"`
						CtrySubDvsn string `xml:"CtrySubDvsn,omitempty"`
						Ctry        string `xml:"Ctry,omitempty"`
						AdrLine     string `xml:"AdrLine,omitempty"`
					} `xml:"PstlAdr,omitempty"`
					Id struct {
						OrgId struct {
							AnyBIC string `xml:"AnyBIC,omitempty"`
							LEI    string `xml:"LEI,omitempty"`
							Othr   struct {
								Id      string `xml:"Id,omitempty"`
								SchmeNm struct {
									Cd    string `xml:"Cd,omitempty" validate:"in=[ExternalOrganisationIdentification1Code]"`
									Prtry string `xml:"Prtry,omitempty"`
								} `xml:"SchmeNm,omitempty"`
								Issr string `xml:"Issr,omitempty"`
							} `xml:"Othr,omitempty"`
						} `xml:"OrgId,omitempty"`
						PrvtId struct {
							DtAndPlcOfBirth struct {
								BirthDt     string `xml:"BirthDt,omitempty"`
								PrvcOfBirth string `xml:"PrvcOfBirth,omitempty"`
								CityOfBirth string `xml:"CityOfBirth,omitempty"`
								CtryOfBirth string `xml:"CtryOfBirth,omitempty"`
							} `xml:"DtAndPlcOfBirth,omitempty"`
							Othr struct {
								Id      string `xml:"Id,omitempty"`
								SchmeNm struct {
									Cd    string `xml:"Cd,omitempty" validate:"in=[ExternalPersonIdentification1Code]"`
									Prtry string `xml:"Prtry,omitempty"`
								} `xml:"SchmeNm,omitempty"`
								Issr string `xml:"Issr,omitempty"`
							} `xml:"Othr,omitempty"`
						} `xml:"PrvtId,omitempty"`
					} `xml:"Id,omitempty"`
					CtryOfRes string `xml:"CtryOfRes,omitempty"`
					CtctDtls  struct {
						NmPrfx    string `xml:"NmPrfx,omitempty"`
						Nm        string `xml:"Nm,omitempty"`
						PhneNb    string `xml:"PhneNb,omitempty"`
						MobNb     string `xml:"MobNb,omitempty"`
						FaxNb     string `xml:"FaxNb,omitempty"`
						EmailAdr  string `xml:"EmailAdr,omitempty"`
						EmailPurp string `xml:"EmailPurp,omitempty"`
						JobTitl   string `xml:"JobTitl,omitempty"`
						Rspnsblty string `xml:"Rspnsblty,omitempty"`
						Dept      string `xml:"Dept,omitempty"`
						Othr      struct {
							ChanlTp string `xml:"ChanlTp,omitempty"`
							Id      string `xml:"Id,omitempty"`
						} `xml:"Othr,omitempty"`
						PrefrdMtd string `xml:"PrefrdMtd,omitempty"`
					} `xml:"CtctDtls,omitempty"`
				} `xml:"UltmtCdtr,omitempty"`
				InstrForCdtrAgt struct {
					Cd       string `xml:"Cd,omitempty" validate:"in=[ExternalCreditorAgentInstruction1Code]"`
					InstrInf string `xml:"InstrInf,omitempty"`
				} `xml:"InstrForCdtrAgt,omitempty"`
				Purp struct {
					Cd    string `xml:"Cd,omitempty" validate:"in=[ExternalPurpose1Code]"`
					Prtry string `xml:"Prtry,omitempty"`
				} `xml:"Purp,omitempty"`
				RgltryRptg struct {
					DbtCdtRptgInd string `xml:"DbtCdtRptgInd,omitempty"`
					Authrty       struct {
						Nm   string `xml:"Nm,omitempty"`
						Ctry string `xml:"Ctry,omitempty"`
					} `xml:"Authrty,omitempty"`
					Dtls struct {
						Tp   string   `xml:"Tp,omitempty"`
						Dt   string   `xml:"Dt,omitempty"`
						Ctry string   `xml:"Ctry,omitempty"`
						Cd   string   `xml:"Cd,omitempty"`
						Amt  AmtField `xml:"Amt,omitempty"`
						Inf  string   `xml:"Inf,omitempty"`
					} `xml:"Dtls,omitempty"`
				} `xml:"RgltryRptg,omitempty"`
				Tax struct {
					Cdtr struct {
						TaxId  string `xml:"TaxId,omitempty"`
						RegnId string `xml:"RegnId,omitempty"`
						TaxTp  string `xml:"TaxTp,omitempty"`
					} `xml:"Cdtr,omitempty"`
					Dbtr struct {
						TaxId   string `xml:"TaxId,omitempty"`
						RegnId  string `xml:"RegnId,omitempty"`
						TaxTp   string `xml:"TaxTp,omitempty"`
						Authstn struct {
							Titl string `xml:"Titl,omitempty"`
							Nm   string `xml:"Nm,omitempty"`
						} `xml:"Authstn,omitempty"`
					} `xml:"Dbtr,omitempty"`
					AdmstnZone      string   `xml:"AdmstnZone,omitempty"`
					RefNb           string   `xml:"RefNb,omitempty"`
					Mtd             string   `xml:"Mtd,omitempty"`
					TtlTaxblBaseAmt AmtField `xml:"TtlTaxblBaseAmt,omitempty"`
					TtlTaxAmt       AmtField `xml:"TtlTaxAmt,omitempty"`
					Dt              string   `xml:"Dt,omitempty"`
					SeqNb           int      `xml:"SeqNb,omitempty"`
					Rcrd            struct {
						Tp       string `xml:"Tp,omitempty"`
						Ctgy     string `xml:"Ctgy,omitempty"`
						CtgyDtls string `xml:"CtgyDtls,omitempty"`
						DbtrSts  string `xml:"DbtrSts,omitempty"`
						CertId   string `xml:"CertId,omitempty"`
						FrmsCd   string `xml:"FrmsCd,omitempty"`
						Prd      struct {
							Yr     string `xml:"Yr,omitempty"`
							Tp     string `xml:"Tp,omitempty"`
							FrToDt struct {
								FrDt string `xml:"FrDt,omitempty"`
								ToDt string `xml:"ToDt,omitempty"`
							} `xml:"FrToDt,omitempty"`
						} `xml:"Prd,omitempty"`
						TaxAmt struct {
							Rate         float64  `xml:"Rate,omitempty"`
							TaxblBaseAmt AmtField `xml:"TaxblBaseAmt,omitempty"`
							TtlAmt       AmtField `xml:"TtlAmt,omitempty"`
							Dtls         struct {
								Prd struct {
									Yr     string `xml:"Yr,omitempty"`
									Tp     string `xml:"Tp,omitempty"`
									FrToDt struct {
										FrDt string `xml:"FrDt,omitempty"`
										ToDt string `xml:"ToDt,omitempty"`
									} `xml:"FrToDt,omitempty"`
								} `xml:"Prd,omitempty"`
								Amt AmtField `xml:"Amt,omitempty"`
							} `xml:"Dtls,omitempty"`
						} `xml:"TaxAmt,omitempty"`
						AddtlInf string `xml:"AddtlInf,omitempty"`
					} `xml:"Rcrd,omitempty"`
				} `xml:"Tax,omitempty"`
				RltdRmtInf struct {
					RmtId       string `xml:"RmtId,omitempty"`
					RmtLctnDtls struct {
						Mtd        string `xml:"Mtd,omitempty"`
						ElctrncAdr string `xml:"ElctrncAdr,omitempty"`
						PstlAdr    struct {
							Nm  string `xml:"Nm,omitempty"`
							Adr struct {
								AdrTp struct {
									Cd    string `xml:"Cd,omitempty" validate:"in=[AddressType2Code]"`
									Prtry struct {
										Id      string `xml:"Id,omitempty"`
										Issr    string `xml:"Issr,omitempty"`
										SchmeNm string `xml:"SchmeNm,omitempty"`
									} `xml:"Prtry,omitempty"`
								} `xml:"AdrTp,omitempty"`
								Dept        string `xml:"Dept,omitempty"`
								SubDept     string `xml:"SubDept,omitempty"`
								StrtNm      string `xml:"StrtNm,omitempty"`
								BldgNb      string `xml:"BldgNb,omitempty"`
								BldgNm      string `xml:"BldgNm,omitempty"`
								Flr         string `xml:"Flr,omitempty"`
								PstBx       string `xml:"PstBx,omitempty"`
								Room        string `xml:"Room,omitempty"`
								PstCd       string `xml:"PstCd,omitempty"`
								TwnNm       string `xml:"TwnNm,omitempty"`
								TwnLctnNm   string `xml:"TwnLctnNm,omitempty"`
								DstrctNm    string `xml:"DstrctNm,omitempty"`
								CtrySubDvsn string `xml:"CtrySubDvsn,omitempty"`
								Ctry        string `xml:"Ctry,omitempty"`
								AdrLine     string `xml:"AdrLine,omitempty"`
							} `xml:"Adr,omitempty"`
						} `xml:"PstlAdr,omitempty"`
					} `xml:"RmtLctnDtls,omitempty"`
				} `xml:"RltdRmtInf,omitempty"`
				RmtInf struct {
					Ustrd string `xml:"Ustrd,omitempty"`
					Strd  struct {
						RfrdDocInf struct {
							Tp struct {
								CdOrPrtry struct {
									Cd    string `xml:"Cd,omitempty" validate:"in=[DocumentType6Code]"`
									Prtry string `xml:"Prtry,omitempty"`
								} `xml:"CdOrPrtry,omitempty"`
								Issr string `xml:"Issr,omitempty"`
							} `xml:"Tp,omitempty"`
							Nb       string `xml:"Nb,omitempty"`
							RltdDt   string `xml:"RltdDt,omitempty"`
							LineDtls struct {
								Id string `xml:"Id,omitempty"`
							} `xml:"LineDtls,omitempty"`
						} `xml:"RfrdDocInf,omitempty"`
					} `xml:"Strd,omitempty"`
				} `xml:"RmtInf,omitempty"`
			} `xml:"CdtTrfTx,omitempty"`
		} `xml:"PmtInf,omitempty"`
	} `xml:"CdtrPmtActvtnReq,omitempty"`
}
