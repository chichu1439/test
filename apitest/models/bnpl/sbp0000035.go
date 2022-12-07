//Version: v1
package models

type BP000035Request struct {
	CustomerId string `json:"customerId" description:"Customer Id" validate:"required"`
}

type BP000035Response struct {
	BindAcctList []BindAcctInfo `json:"bindAcctList" description:"Bind Acct List"`
}

type BindAcctInfo struct {
	BindAcctNo      string `json:"bindAcctNo" description:"Bind Acct No"`
	BindAcctType    string `json:"bindAcctType" description:"Bind Acct Type"`
	BindAcctOweBank string `json:"bindAcctOweBank" description:"Bind Acct Owe Bank"`
	BindDate        string `json:"bindDate" description:"Bind Date"`
	BindTime        string `json:"bindTime" description:"Bind Time"`
	BindStatus      string `json:"bindStatus" description:"Bind Status"`
	Priority        string `json:"priority" description:"Priority"`
}

func (*BP000035Request) GetServiceKey() string {
	return "BP000035"
}
