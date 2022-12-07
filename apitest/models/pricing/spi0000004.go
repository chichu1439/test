//Version: v0.0.1
package models

type PI000004I struct {
	InterestId string `json:"interestId" validate:"required" description:"Interest Id"`
}

type PI000004O struct {
}

func (*PI000004I) GetServiceKey() string {
	return "PI000004"
}
