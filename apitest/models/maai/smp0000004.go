package models

// swagger:parameters Request
type SMP0000004I struct {
	AgreementType string `json:"agreementType" description:"agreement type" validate:"required,max=10"`
	CustomerNo    string `json:"customerNo" description:"customer no" validate:"max=50"`
	Language      string `json:"language" description:"agreement content language" validate:"max=10"`
}

//type SMP0000004O struct {
//	SMP9CM0004O
//}
type SMP0000004O SMP9CM0004O

func (i SMP0000004I) GetServiceKey() string {
	return "MP000004"
}
