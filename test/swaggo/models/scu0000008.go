package models

type SCU0000008I struct {
	CustomerId string `json:"customerId" description:"Customer id" validate:"required"`
	ContractId string `json:"contractId" description:"Contract id" validate:"required"`
}

type SCU0000008O struct {
	ListContact []ListContact `json:"listContact" description:"contact List"`
	ListAddress []ListAddress `json:"listAddress" description:"address List"`
}

type ListContact struct {
	ContactId       string `json:"contactId" description:"Contact id"`
	ContactType     string `json:"contactType" description:"Contact Type:01- Home phone 02- Office phone 04- EMAIL 09- Mobile 99- other"`
	ContactNum      string `json:"contactNum" description:"Contact number"`
	Title           string `json:"title" description:"Title:Mr. Mrs. Miss"`
	EffectiveStatus string `json:"effectiveStatus" description:"Effective status"`
	UseType         string `json:"useType" description:"Use type"`
	DefaultFlag     string `json:"defaultFlag" description:"Default flag"`
}

type ListAddress struct {
	AddressId           string `json:"addressId" description:"Address id"`
	AddressType         string `json:"addressType" description:"Address Type:1-Mailing address 2-Household Registration"`
	CustomerId          string `json:"customerId" description:"Customer id"`
	Address             string `json:"address" description:"Address"`
	Address1            string `json:"address1" description:"Address1"`
	Address2            string `json:"address2" description:"Address2"`
	Address3            string `json:"address3" description:"Address3"`
	HouseResidentNumber string `json:"houseResidentNumber" description:"House resident number"`
	DefaultFlag         string `json:"defaultFlag" description:"Default flag"`
	AddressNation       string `json:"addressNation" description:"Address nation"`
	PostalCode          string `json:"postalCode" description:"Postal code"`
	EffectiveStatus     string `json:"effectiveStatus" description:"Effective status"`
	UseType             string `json:"useType" description:"Use type"`
}

type SCU0000008OCont struct {
	ContactRelationIdList []ContactRelationId `json:"contactRelationIdList" description:"Contact relation id list"`
}

type ContactRelationId struct {
	ContactRelationId string `json:"contactRelationId" description:"Contact relation id"`
}

func (*SCU0000008I) GetServiceKey() string {
	return "CU000008"
}
