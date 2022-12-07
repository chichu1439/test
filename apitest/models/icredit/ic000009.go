package models

type IC000009I struct {
	CustomerId    string `json:"customerId" description:"Customer id" validate:"required"`
	ChangeComment string `json:"changeComment" description:"Change Comment"`
}

type IC000009O struct {
}

func (*IC000009I) GetServiceKey() string {
	return "ic000009"
}
