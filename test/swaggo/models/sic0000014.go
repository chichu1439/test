package models

type IC000014I struct {
	Month        int    `json:"month" description:"Month"`
	SourceSystem string `json:"sourceSystem" description:"Source System"`
	ProductId    string `json:"productId" description:"Product id"`
}

type IC000014O struct {
	Records []CustStatisticsBean `json:"records" description:"Records"`
}

type CustStatisticsBean struct {
	StatisticsDate          string `json:"statisticsDate" description:"Statistics Date"`
	CustomerTotalNumber     int    `json:"customerTotalNumber" description:"Customer Total Number"`
	CurrentMonthNewCustomer int    `json:"currentMonthNewCustomer" description:"Current Month New Customer"`
	TodayNewCustomer        int    `json:"todayNewCustomer" description:"Today New Customer"`
}

func (*IC000014I) GetServiceKey() string {
	return "IC000014"
}
