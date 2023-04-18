package models

import (
	"encoding/xml"
)

type FpsPain014 struct {
	XMLName  xml.Name   `xml:"FpsMsg,omitempty,omitempty"`
	NumOfMsg string     `xml:"NbOfMsgs,omitempty,omitempty"`
	XMLHead  Head001    `xml:"BizData>AppHdr,omitempty,omitempty"`
	Document Pain014Doc `xml:"BizData>Document,omitempty,omitempty"`
}
type Pain014Doc struct {
	CdtrPmtActvtnReqStsRpt struct {
		GrpHdr struct {
			MsgId    string `xml:"MsgId,omitempty"`
			CreDtTm  string `xml:"CreDtTm,omitempty"`
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
			DbtrAgt struct {
				FinInstnId struct {
					BICFI       string `xml:"BICFI,omitempty"`
					ClrSysMmbId struct {
						ClrSysId struct {
							Cd    string `xml:"Cd,omitempty"`
							Prtry string `xml:"Prtry,omitempty"`
						} `xml:"ClrSysId,omitempty"`
						MmbId string `xml:"MmbId,omitempty"`
					} `xml:"ClrSysMmbId,omitempty"`
					LEI     string `xml:"LEI,omitempty"`
					Nm      string `xml:"Nm,omitempty"`
					PstlAdr struct {
						AdrTp struct {
							Cd    string `xml:"Cd,omitempty"`
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
							Cd    string `xml:"Cd,omitempty"`
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
							Cd    string `xml:"Cd,omitempty"`
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
			CdtrAgt struct {
				FinInstnId struct {
					BICFI       string `xml:"BICFI,omitempty"`
					ClrSysMmbId struct {
						ClrSysId struct {
							Cd    string `xml:"Cd,omitempty"`
							Prtry string `xml:"Prtry,omitempty"`
						} `xml:"ClrSysId,omitempty"`
						MmbId string `xml:"MmbId,omitempty"`
					} `xml:"ClrSysMmbId,omitempty"`
					LEI     string `xml:"LEI,omitempty"`
					Nm      string `xml:"Nm,omitempty"`
					PstlAdr struct {
						AdrTp struct {
							Cd    string `xml:"Cd,omitempty"`
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
							Cd    string `xml:"Cd,omitempty"`
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
							Cd    string `xml:"Cd,omitempty"`
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
		} `xml:"GrpHdr,omitempty"`
		OrgnlGrpInfAndSts struct {
			OrgnlMsgId   string  `xml:"OrgnlMsgId,omitempty"`
			OrgnlMsgNmId string  `xml:"OrgnlMsgNmId,omitempty"`
			OrgnlCreDtTm string  `xml:"OrgnlCreDtTm,omitempty"`
			OrgnlNbOfTxs string  `xml:"OrgnlNbOfTxs,omitempty"`
			OrgnlCtrlSum float64 `xml:"OrgnlCtrlSum,omitempty"`
			GrpSts       string  `xml:"GrpSts,omitempty"`
			StsRsnInf    struct {
				Orgtr struct {
					Nm      string `xml:"Nm,omitempty"`
					PstlAdr struct {
						AdrTp struct {
							Cd    string `xml:"Cd,omitempty"`
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
									Cd    string `xml:"Cd,omitempty"`
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
									Cd    string `xml:"Cd,omitempty"`
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
				} `xml:"Orgtr,omitempty"`
				Rsn struct {
					Cd    string `xml:"Cd,omitempty"`
					Prtry string `xml:"Prtry,omitempty"`
				} `xml:"Rsn,omitempty"`
				AddtlInf string `xml:"AddtlInf,omitempty"`
			} `xml:"StsRsnInf,omitempty"`
			NbOfTxsPerSts struct {
				DtldNbOfTxs string  `xml:"DtldNbOfTxs,omitempty"`
				DtldSts     string  `xml:"DtldSts,omitempty"`
				DtldCtrlSum float64 `xml:"DtldCtrlSum,omitempty"`
			} `xml:"NbOfTxsPerSts,omitempty"`
		} `xml:"OrgnlGrpInfAndSts,omitempty"`
		OrgnlPmtInfAndSts struct {
			OrgnlPmtInfId string  `xml:"OrgnlPmtInfId,omitempty"`
			OrgnlNbOfTxs  string  `xml:"OrgnlNbOfTxs,omitempty"`
			OrgnlCtrlSum  float64 `xml:"OrgnlCtrlSum,omitempty"`
			PmtInfSts     string  `xml:"PmtInfSts,omitempty"`
			StsRsnInf     struct {
				Orgtr struct {
					Nm      string `xml:"Nm,omitempty"`
					PstlAdr struct {
						AdrTp struct {
							Cd    string `xml:"Cd,omitempty"`
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
									Cd    string `xml:"Cd,omitempty"`
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
									Cd    string `xml:"Cd,omitempty"`
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
				} `xml:"Orgtr,omitempty"`
				Rsn struct {
					Cd    string `xml:"Cd,omitempty"`
					Prtry string `xml:"Prtry,omitempty"`
				} `xml:"Rsn,omitempty"`
				AddtlInf string `xml:"AddtlInf,omitempty"`
			} `xml:"StsRsnInf,omitempty"`
			NbOfTxsPerSts struct {
				DtldNbOfTxs string  `xml:"DtldNbOfTxs,omitempty"`
				DtldSts     string  `xml:"DtldSts,omitempty"`
				DtldCtrlSum float64 `xml:"DtldCtrlSum,omitempty"`
			} `xml:"NbOfTxsPerSts,omitempty"`
			TxInfAndSts struct {
				StsId           string `xml:"StsId,omitempty"`
				OrgnlInstrId    string `xml:"OrgnlInstrId,omitempty"`
				OrgnlEndToEndId string `xml:"OrgnlEndToEndId,omitempty"`
				OrgnlUETR       string `xml:"OrgnlUETR,omitempty"`
				TxSts           string `xml:"TxSts,omitempty"`
				StsRsnInf       struct {
					Orgtr struct {
						Nm      string `xml:"Nm,omitempty"`
						PstlAdr struct {
							AdrTp struct {
								Cd    string `xml:"Cd,omitempty"`
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
										Cd    string `xml:"Cd,omitempty"`
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
										Cd    string `xml:"Cd,omitempty"`
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
					} `xml:"Orgtr,omitempty"`
					Rsn struct {
						Cd    string `xml:"Cd,omitempty"`
						Prtry string `xml:"Prtry,omitempty"`
					} `xml:"Rsn,omitempty"`
					AddtlInf string `xml:"AddtlInf,omitempty"`
				} `xml:"StsRsnInf,omitempty"`
				PmtCondSts struct {
					AccptdAmt int  `xml:"AccptdAmt,omitempty"`
					GrntedPmt bool `xml:"GrntedPmt,omitempty"`
					EarlyPmt  bool `xml:"EarlyPmt,omitempty"`
				} `xml:"PmtCondSts,omitempty"`
				ChrgsInf struct {
					Amt int `xml:"Amt,omitempty"`
					Agt struct {
						FinInstnId struct {
							BICFI       string `xml:"BICFI,omitempty"`
							ClrSysMmbId struct {
								ClrSysId struct {
									Cd    string `xml:"Cd,omitempty"`
									Prtry string `xml:"Prtry,omitempty"`
								} `xml:"ClrSysId,omitempty"`
								MmbId string `xml:"MmbId,omitempty"`
							} `xml:"ClrSysMmbId,omitempty"`
							LEI     string `xml:"LEI,omitempty"`
							Nm      string `xml:"Nm,omitempty"`
							PstlAdr struct {
								AdrTp struct {
									Cd    string `xml:"Cd,omitempty"`
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
									Cd    string `xml:"Cd,omitempty"`
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
									Cd    string `xml:"Cd,omitempty"`
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
					} `xml:"Agt,omitempty"`
				} `xml:"ChrgsInf,omitempty"`
				DbtrDcsnDtTm string `xml:"DbtrDcsnDtTm,omitempty"`
				AccptncDtTm  string `xml:"AccptncDtTm,omitempty"`
				AcctSvcrRef  string `xml:"AcctSvcrRef,omitempty"`
				ClrSysRef    string `xml:"ClrSysRef,omitempty"`
				OrgnlTxRef   struct {
					Amt struct {
						InstdAmt int `xml:"InstdAmt,omitempty"`
						EqvtAmt  struct {
							Amt      int    `xml:"Amt,omitempty"`
							CcyOfTrf string `xml:"CcyOfTrf,omitempty"`
						} `xml:"EqvtAmt,omitempty"`
					} `xml:"Amt,omitempty"`
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
							Amt  int `xml:"Amt,omitempty"`
							Rate int `xml:"Rate,omitempty"`
						} `xml:"ImdtPmtRbt,omitempty"`
						GrntedPmtReqd bool `xml:"GrntedPmtReqd,omitempty"`
					} `xml:"PmtCond,omitempty"`
					PmtTpInf struct {
						InstrPrty string `xml:"InstrPrty,omitempty"`
						SvcLvl    struct {
							Cd    string `xml:"Cd,omitempty"`
							Prtry string `xml:"Prtry,omitempty"`
						} `xml:"SvcLvl,omitempty"`
						LclInstrm struct {
							Cd    string `xml:"Cd,omitempty"`
							Prtry string `xml:"Prtry,omitempty"`
						} `xml:"LclInstrm,omitempty"`
						CtgyPurp struct {
							Cd    string `xml:"Cd,omitempty"`
							Prtry string `xml:"Prtry,omitempty"`
						} `xml:"CtgyPurp,omitempty"`
					} `xml:"PmtTpInf,omitempty"`
					PmtMtd string `xml:"PmtMtd,omitempty"`
					RmtInf struct {
						Ustrd string `xml:"Ustrd,omitempty"`
						Strd  struct {
							RfrdDocInf struct {
								Tp struct {
									CdOrPrtry struct {
										Cd    string `xml:"Cd,omitempty"`
										Prtry string `xml:"Prtry,omitempty"`
									} `xml:"CdOrPrtry,omitempty"`
									Issr string `xml:"Issr,omitempty"`
								} `xml:"Tp,omitempty"`
								Nb       string `xml:"Nb,omitempty"`
								RltdDt   string `xml:"RltdDt,omitempty"`
								LineDtls struct {
									Id struct {
										Tp struct {
											CdOrPrtry struct {
												Cd    string `xml:"Cd,omitempty"`
												Prtry string `xml:"Prtry,omitempty"`
											} `xml:"CdOrPrtry,omitempty"`
											Issr string `xml:"Issr,omitempty"`
										} `xml:"Tp,omitempty"`
										Nb     string `xml:"Nb,omitempty"`
										RltdDt string `xml:"RltdDt,omitempty"`
									} `xml:"Id,omitempty"`
									Desc string `xml:"Desc,omitempty"`
									Amt  struct {
										DuePyblAmt   int `xml:"DuePyblAmt,omitempty"`
										DscntApldAmt struct {
											Tp struct {
												Cd    string `xml:"Cd,omitempty"`
												Prtry string `xml:"Prtry,omitempty"`
											} `xml:"Tp,omitempty"`
											Amt int `xml:"Amt,omitempty"`
										} `xml:"DscntApldAmt,omitempty"`
										CdtNoteAmt int `xml:"CdtNoteAmt,omitempty"`
										TaxAmt     struct {
											Tp struct {
												Cd    string `xml:"Cd,omitempty"`
												Prtry string `xml:"Prtry,omitempty"`
											} `xml:"Tp,omitempty"`
											Amt int `xml:"Amt,omitempty"`
										} `xml:"TaxAmt,omitempty"`
										AdjstmntAmtAndRsn struct {
											Amt       int    `xml:"Amt,omitempty"`
											CdtDbtInd string `xml:"CdtDbtInd,omitempty"`
											Rsn       string `xml:"Rsn,omitempty"`
											AddtlInf  string `xml:"AddtlInf,omitempty"`
										} `xml:"AdjstmntAmtAndRsn,omitempty"`
										RmtdAmt int `xml:"RmtdAmt,omitempty"`
									} `xml:"Amt,omitempty"`
								} `xml:"LineDtls,omitempty"`
							} `xml:"RfrdDocInf,omitempty"`
							RfrdDocAmt struct {
								DuePyblAmt   int `xml:"DuePyblAmt,omitempty"`
								DscntApldAmt struct {
									Tp struct {
										Cd    string `xml:"Cd,omitempty"`
										Prtry string `xml:"Prtry,omitempty"`
									} `xml:"Tp,omitempty"`
									Amt int `xml:"Amt,omitempty"`
								} `xml:"DscntApldAmt,omitempty"`
								CdtNoteAmt int `xml:"CdtNoteAmt,omitempty"`
								TaxAmt     struct {
									Tp struct {
										Cd    string `xml:"Cd,omitempty"`
										Prtry string `xml:"Prtry,omitempty"`
									} `xml:"Tp,omitempty"`
									Amt int `xml:"Amt,omitempty"`
								} `xml:"TaxAmt,omitempty"`
								AdjstmntAmtAndRsn struct {
									Amt       int    `xml:"Amt,omitempty"`
									CdtDbtInd string `xml:"CdtDbtInd,omitempty"`
									Rsn       string `xml:"Rsn,omitempty"`
									AddtlInf  string `xml:"AddtlInf,omitempty"`
								} `xml:"AdjstmntAmtAndRsn,omitempty"`
								RmtdAmt int `xml:"RmtdAmt,omitempty"`
							} `xml:"RfrdDocAmt,omitempty"`
							CdtrRefInf struct {
								Tp struct {
									CdOrPrtry struct {
										Cd    string `xml:"Cd,omitempty"`
										Prtry string `xml:"Prtry,omitempty"`
									} `xml:"CdOrPrtry,omitempty"`
									Issr string `xml:"Issr,omitempty"`
								} `xml:"Tp,omitempty"`
								Ref string `xml:"Ref,omitempty"`
							} `xml:"CdtrRefInf,omitempty"`
							Invcr struct {
								Nm      string `xml:"Nm,omitempty"`
								PstlAdr struct {
									AdrTp struct {
										Cd    string `xml:"Cd,omitempty"`
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
												Cd    string `xml:"Cd,omitempty"`
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
												Cd    string `xml:"Cd,omitempty"`
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
							} `xml:"Invcr,omitempty"`
							Invcee struct {
								Nm      string `xml:"Nm,omitempty"`
								PstlAdr struct {
									AdrTp struct {
										Cd    string `xml:"Cd,omitempty"`
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
												Cd    string `xml:"Cd,omitempty"`
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
												Cd    string `xml:"Cd,omitempty"`
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
							} `xml:"Invcee,omitempty"`
							TaxRmt struct {
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
								UltmtDbtr struct {
									TaxId   string `xml:"TaxId,omitempty"`
									RegnId  string `xml:"RegnId,omitempty"`
									TaxTp   string `xml:"TaxTp,omitempty"`
									Authstn struct {
										Titl string `xml:"Titl,omitempty"`
										Nm   string `xml:"Nm,omitempty"`
									} `xml:"Authstn,omitempty"`
								} `xml:"UltmtDbtr,omitempty"`
								AdmstnZone      string `xml:"AdmstnZone,omitempty"`
								RefNb           string `xml:"RefNb,omitempty"`
								Mtd             string `xml:"Mtd,omitempty"`
								TtlTaxblBaseAmt int    `xml:"TtlTaxblBaseAmt,omitempty"`
								TtlTaxAmt       int    `xml:"TtlTaxAmt,omitempty"`
								Dt              string `xml:"Dt,omitempty"`
								SeqNb           int    `xml:"SeqNb,omitempty"`
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
										Rate         int `xml:"Rate,omitempty"`
										TaxblBaseAmt int `xml:"TaxblBaseAmt,omitempty"`
										TtlAmt       int `xml:"TtlAmt,omitempty"`
										Dtls         struct {
											Prd struct {
												Yr     string `xml:"Yr,omitempty"`
												Tp     string `xml:"Tp,omitempty"`
												FrToDt struct {
													FrDt string `xml:"FrDt,omitempty"`
													ToDt string `xml:"ToDt,omitempty"`
												} `xml:"FrToDt,omitempty"`
											} `xml:"Prd,omitempty"`
											Amt int `xml:"Amt,omitempty"`
										} `xml:"Dtls,omitempty"`
									} `xml:"TaxAmt,omitempty"`
									AddtlInf string `xml:"AddtlInf,omitempty"`
								} `xml:"Rcrd,omitempty"`
							} `xml:"TaxRmt,omitempty"`
							GrnshmtRmt struct {
								Tp struct {
									CdOrPrtry struct {
										Cd    string `xml:"Cd,omitempty"`
										Prtry string `xml:"Prtry,omitempty"`
									} `xml:"CdOrPrtry,omitempty"`
									Issr string `xml:"Issr,omitempty"`
								} `xml:"Tp,omitempty"`
								Grnshee struct {
									Nm      string `xml:"Nm,omitempty"`
									PstlAdr struct {
										AdrTp struct {
											Cd    string `xml:"Cd,omitempty"`
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
													Cd    string `xml:"Cd,omitempty"`
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
													Cd    string `xml:"Cd,omitempty"`
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
								} `xml:"Grnshee,omitempty"`
								GrnshmtAdmstr struct {
									Nm      string `xml:"Nm,omitempty"`
									PstlAdr struct {
										AdrTp struct {
											Cd    string `xml:"Cd,omitempty"`
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
													Cd    string `xml:"Cd,omitempty"`
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
													Cd    string `xml:"Cd,omitempty"`
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
								} `xml:"GrnshmtAdmstr,omitempty"`
								RefNb             string `xml:"RefNb,omitempty"`
								Dt                string `xml:"Dt,omitempty"`
								RmtdAmt           int    `xml:"RmtdAmt,omitempty"`
								FmlyMdclInsrncInd bool   `xml:"FmlyMdclInsrncInd,omitempty"`
								MplyeeTermntnInd  bool   `xml:"MplyeeTermntnInd,omitempty"`
							} `xml:"GrnshmtRmt,omitempty"`
							AddtlRmtInf string `xml:"AddtlRmtInf,omitempty"`
						} `xml:"Strd,omitempty"`
					} `xml:"RmtInf,omitempty"`
					NclsdFile struct {
						Tp struct {
							Cd    string `xml:"Cd,omitempty"`
							Prtry struct {
								Id      string `xml:"Id,omitempty"`
								SchmeNm string `xml:"SchmeNm,omitempty"`
								Issr    string `xml:"Issr,omitempty"`
							} `xml:"Prtry,omitempty"`
						} `xml:"Tp,omitempty"`
						Id     string `xml:"Id,omitempty"`
						IsseDt struct {
							Dt   string `xml:"Dt,omitempty"`
							DtTm string `xml:"DtTm,omitempty"`
						} `xml:"IsseDt,omitempty"`
						Nm     string `xml:"Nm,omitempty"`
						LangCd string `xml:"LangCd,omitempty"`
						Frmt   struct {
							Cd    string `xml:"Cd,omitempty"`
							Prtry struct {
								Id      string `xml:"Id,omitempty"`
								SchmeNm string `xml:"SchmeNm,omitempty"`
								Issr    string `xml:"Issr,omitempty"`
							} `xml:"Prtry,omitempty"`
						} `xml:"Frmt,omitempty"`
						FileNm    string `xml:"FileNm,omitempty"`
						DgtlSgntr struct {
							Pty struct {
								Nm      string `xml:"Nm,omitempty"`
								PstlAdr struct {
									AdrTp struct {
										Cd    string `xml:"Cd,omitempty"`
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
												Cd    string `xml:"Cd,omitempty"`
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
												Cd    string `xml:"Cd,omitempty"`
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
							} `xml:"Pty,omitempty"`
							Sgntr struct {
								AnyElement string `xml:"AnyElement,omitempty"`
							} `xml:"Sgntr,omitempty"`
						} `xml:"DgtlSgntr,omitempty"`
						Nclsr string `xml:"Nclsr,omitempty"`
					} `xml:"NclsdFile,omitempty"`
					UltmtDbtr struct {
						Nm      string `xml:"Nm,omitempty"`
						PstlAdr struct {
							AdrTp struct {
								Cd    string `xml:"Cd,omitempty"`
								Prtry struct {
									Id      string `xml:"Id,omitempty"`
									Issr    string `xml:"Issr,omitempty"`
									SchmeNm string `xml:"SchmeNm,omitempty"`
								} `xml:"Prtry,omitempty"`
							} `xml:"AdrTp,omitempty"`
						} `xml:"PstlAdr,omitempty"`
					} `xml:"UltmtDbtr,omitempty"`
					CdtrAgt struct {
						FinInstnId string `xml:"FinInstnId,omitempty"`
					} `xml:"CdtrAgt,omitempty"`
					Cdtr string `xml:"Cdtr,omitempty"`
				} `xml:"OrgnlTxRef,omitempty"`
			} `xml:"TxInfAndSts,omitempty"`
		} `xml:"OrgnlPmtInfAndSts,omitempty"`
	} `xml:"CdtrPmtActvtnReqStsRpt,omitempty"`
}
