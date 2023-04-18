package models

type SCU0000009I struct {
	CustomerId    string        `json:"customerId" validate:"required,max=20" description:"Customer id"` //客户号 Y
	ContractId    string        `json:"contractId" validate:"required" description:"Contract ID"`        //合约号 Y
	ListContact   []ContactInfo `json:"listContact" description:"Contact list"`                          //联系信息数组 array 合约号必输
	ListAddress   []AddressInfo `json:"listAddress" description:"Address list"`                          //地址信息数组 array 合约号必输
	SpouseInfo    SpouseInfo    `json:"spouseInfo" description:"Spouse information"`
	EmergencyInfo EmergencyInfo `json:"emergencyInfo" description:"Emergency information"`
	ContactList   []ContactInfo `json:"contactList" description:"Contact list"` // 客户联系信息
}

//contact    array

type ContactInfo struct {
	ContactId       string `json:"contactId" description:"Contact id"`
	ContactType     string `json:"contactType" description:"Contact Type:01- Home phone 02- Office phone 04- EMAIL 09- Mobile 99- other"`
	ContactNum      string `json:"contactNum" description:"Contact number"`
	CustomerId      string `json:"customerId" description:"Customer id"`
	DefaultFlag     string `json:"defaultFlag" description:"Default flag"`
	EffectiveStatus string `json:"effectiveStatus" description:"Effective status"`
	Title           string `json:"title" description:"Title:Mr. Mrs. Miss"`
	UseType         string `json:"useType" description:"Use type"`
}

// address array

type AddressInfo struct {
	AddressId           string `json:"addressId" description:"Address id"`
	AddressType         string `json:"addressType" description:"1-Mailing address 2-Household Registration"`
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

type SCU0000009O struct {
}

func (*SCU0000009I) GetServiceKey() string {
	return "CU000009"
}
