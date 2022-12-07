package models

import "encoding/xml"

type FpsAcmt023 struct {
	XMLName  xml.Name   `xml:"FpsMsg,omitempty"`
	NumOfMsg string     `xml:"NbOfMsgs,omitempty"`
	XMLHead  Head001    `xml:"BizData>AppHdr,omitempty"`
	Document Acmt023Doc `xml:"BizData>Document,omitempty"`
}

type Acmt023Doc struct {
	IdVrfctnReq struct {
		Assgnmt struct {
			MsgId   string `xml:"MsgId,omitempty" validate:"required,max=35"`
			CreDtTm string `xml:"CreDtTm,omitempty" validate:"required,max=35"`
			Cretr   struct {
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
									Cd    string `xml:"Cd,omitempty"  validate:"in=[ExternalOrganisationIdentification1Code]"`
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
				} `xml:"Pty,omitempty"`
				Agt struct {
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
			} `xml:"Cretr,omitempty"`
			FrstAgt struct {
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
			} `xml:"FrstAgt,omitempty"`
			Assgnr struct {
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
				} `xml:"Pty,omitempty"`
				Agt struct {
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
			} `xml:"Assgnr,omitempty"`
			Assgne struct {
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
				} `xml:"Pty,omitempty"`
				Agt struct {
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
			} `xml:"Assgne,omitempty"`
		} `xml:"Assgnmt,omitempty"`
		Vrfctn struct {
			Id           string `xml:"Id,omitempty"`
			PtyAndAcctId struct {
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
				} `xml:"Pty,omitempty"`
				Acct struct {
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
					Ccy  string `xml:"Ccy,omitempty" validate:"in=[CURRENCY]"`
					Nm   string `xml:"Nm,omitempty"`
					Prxy struct {
						Tp struct {
							Cd    string `xml:"Cd,omitempty" validate:"in=[ExternalProxyAccountType1Code]"`
							Prtry string `xml:"Prtry,omitempty"`
						} `xml:"Tp,omitempty"`
						Id string `xml:"Id,omitempty"`
					} `xml:"Prxy,omitempty"`
				} `xml:"Acct,omitempty"`
				Agt struct {
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
			} `xml:"PtyAndAcctId,omitempty"`
		} `xml:"Vrfctn,omitempty"`
		SplmtryData struct {
			PlcAndNm string `xml:"PlcAndNm,omitempty"`
			Envlp    struct {
				AnyElement string `xml:"AnyElement,omitempty"`
			} `xml:"Envlp,omitempty"`
		} `xml:"SplmtryData,omitempty"`
	} `xml:"IdVrfctnReq,omitempty"`
}
