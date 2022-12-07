package models

type FP120001I struct {
	CreDtTm                string `json:"grpHdrCreDtTm" validate:"required,max=20"`          // time now
	AdrId                  string `json:"adrId" validate:"max=35"`                           // AccountKeyNum
	AdrTp                  string `json:"adrTp" validate:"required,max=35"`                  // AccountKetType
	AdrCusId               string `json:"adrCusId" validate:"required,max=35"`               // CustomerId
	AdrSts                 string `json:"adrSts" validate:"required,max=4"`                  // NWRG(new register) ACCT(success) RJCT(fail) CNCL(delete) AMND(update)
	AdrDflt                string `json:"adrDflt" validate:"max=4"`                          // DefaultIndicator
	AdrStsUpdTm            string `json:"adrStsUpdTm" validate:"required,max=20"`            // time now
	AdrInstgAgtClrSysMmbId string `json:"adrInstgAgtClrSysMmbId" validate:"required,max=35"` // ParticipantClearingCode
	AccountNum             string `json:"accountNum"`
	BankCode               string `json:"bankCode"`
}
type FP120001O struct {
}

func (*FP120001I) GetServiceKey() string {
	return "FP120001"
}
