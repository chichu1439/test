//Version: v0.0.1
package models

type PI000007I struct {
	StrategyId string `json:"strategyId" description:"Strategy Id"`
}

type PI000007O struct {
	StrategyId    string          `json:"strategyId" description:"Strategy Id"`
	InterestRate  float64         `json:"interestRate" description:"Interest rate"`
	ListParmMatch []ListParmMatch `json:"listParmMatch" description:"List parm match"`
	ListParmFloat []ListParmFloat `json:"listParmFloat" description:"List parm float"`
}

func (*PI000007I) GetServiceKey() string {
	return "PI000007"
}
