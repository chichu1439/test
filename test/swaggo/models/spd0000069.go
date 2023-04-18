package models

type PD000069I struct {
	SystemId     string `json:"systemId" validate:"required" description:"system Id:systemId(LN-icredit;SV-isaving)"`
	ProductType  string `json:"productType" description:"product Type:Saving 0-Current Deposit 1-Fixed Deposit *Loan 0-Micro Loan"`
	CustomerType string `json:"customerType" description:"customer Type:0-personal;1-corporate"`
}

type PD000069O struct {
	ListService []ProdServicePool `json:"listService" description:"List service"`
}

type ProdServicePool struct {
	ServiceId       string               `json:"serviceId" description:"Service id"`
	ServiceName     string               `json:"serviceName" description:"service Name:01-Approval Date 02-Account Opening Day 03-Loan Date 04-Loan Date plus Interest-free period "`
	SystemId        string               `json:"systemId" description:"system Id;SV-Saving LN-Loan"`
	ProductType     string               `json:"productType" description:"product Type:*Saving 0-Current Deposit 1-Fixed Deposit *Loan 0-Micro Loan "`
	CustomerType    string               `json:"customerType" description:" customer Type:0-Personal 1-Corporate"`
	ServiceType     string               `json:"serviceType" description:"service Type:0-Opening 1-Closing 2-Management 3-Query 4-Finance Trans 5-Notification"`
	TransactionType string               `json:"transactionType" description:"transaction Type:0-Online 1-Batch""`
	EffectiveDate   string               `json:"effectiveDate" description:"Effective date"`
	ExpireDate      string               `json:"expireDate" description:"Expire date"`
	ListFeature     []ProdServiceFeature `json:"listFeature" description:"List feature"`
}

type ProdServiceFeature struct {
	FeatureId   string `json:"featureId" description:"Feature Id"`
	FeatureName string `json:"featureName" description:"Feature Name"`
	FeatureType string `json:"featureType" description:"Feature Type"`
	FeatureCode string `json:"featureCode" description:"Feature Code"`
	//SystemId     String `json:"systemId" description:"SV-Saving LN-Loan FP-Fast payment "`
	//ProductType  String `json:"productType" description:"Saving 0-Current Deposit 1-Fixed Deposit *Loan 0-Micro Loan "`
	//CustomerType String `json:"customerType" description:"0-Personal 1-Corporate "`
	//ServiceType  String `json:"serviceType" description:"0-Opening 1-Closing 2-Management 3-Query 4-Finance Trans 5-Notification "`
	//Status       String `json:"status" description:"0-normal 1-deleted "`
}

func (*PD000069I) GetServiceKey() string {
	return "PD000069"
}
