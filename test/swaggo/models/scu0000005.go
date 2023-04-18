package models

type SCU0000005I struct {
	CustomerId   string          `json:"customerId" description:"Customer id" validate:"required"`
	ContractId   string          `json:"contractId" description:"Contract id" validate:"required"`
	ContractType string          `json:"contractType" validate:"required" description:"Contract Type:01- Home phone 02- Office phone 04- EMAIL 09- Mobile 99- other"`
	AListContact []ListContact02 `json:"aListContact" description:"contact list"`
	AListAddress []ListAddress02 `json:"aListAddress" description:"address list"`
}

type ListContact02 struct {
	ContactType     string `json:"contactType" description:"Contact Type:01- Home phone 02- Office phone 04- EMAIL 09- Mobile 99- other"`
	ContactNum      string `json:"contactNum" description:"Contact number"`
	Title           string `json:"title" description:"Title:Mr. Mrs. Miss"`
	EffectiveStatus string `json:"effectiveStatus" description:"Effective status"`
	UseType         string `json:"useType" description:"Use type"`
	DefaultFlag     string `json:"defaultFlag" description:"Default flag"`
}

type ListAddress02 struct {
	AddressType         string `json:"addressType"  description:"address Type:1-Mailing address 2-Household Registration"`
	CustomerId          string `json:"customerId" description:"Customer id"`
	Address             string `json:"address" description:"Address" `
	Address1            string `json:"address1" description:"Address1" `
	Address2            string `json:"address2" description:"Address2" `
	HouseResidentNumber string `json:"houseResidentNumber" description:"House resident number"`
	Address3            string `json:"address3" description:"Address3" `
	DefaultFlag         string `json:"defaultFlag" description:"Default flag"`
	AddressNation       string `json:"addressNation" description:"Address nation"`
	PostalCode          string `json:"postalCode" description:"Postal code"`
	EffectiveStatus     string `json:"effectiveStatus" description:"Effective status"`
	UseType             string `json:"useType" description:"Use type"`
}

type SCU0000005O struct {
}

func (*SCU0000005I) GetServiceKey() string {
	return "CU000005"
}
