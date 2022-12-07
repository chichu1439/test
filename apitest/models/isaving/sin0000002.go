//Version: v1.0.0.0
package models

type IN000002I struct {
	Title    string `json:"title" description:"Title" validate:"required"`
	Seq      string `json:"seq" description:"Seq" validate:"required,len=4"`
	Currency string `json:"currency" description:"Currency" validate:"required"`
}

type IN000002O struct {
	InternalAccount string `json:"internalAccount" description:"Internal account"`
}

func (*IN000002I) GetServiceKey() string {
	return "sin0000002"
}
