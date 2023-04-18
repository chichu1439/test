package models

type SCU0000012I struct {
	CustomerId      string   `json:"customerId" description:"Customer id"`
	CustomerIdList  []string `json:"customerIdList" description:"Customer id list"`
	CustomerName    string   `json:"customerName" description:"Customer name"`
	IdNation        string   `json:"idNation" description:"Id Nation"`
	IdType          string   `json:"idType" description:"id Type:101-Citizen ID 104-Passport"`
	IdNum           string   `json:"idNum" description:"Id number"`
	CustomerStatus  []string `json:"customerStatus" description:"customer Status:0:Normal 4:Pending 9:Cancel "`
	PageNum         int      `json:"pageNum" description:"Page number"`
	PageRecordCount int      `json:"pageRecordCount" description:"Page record count"`
}

type SCU0000012O struct {
	PageNum         int            `json:"pageNum" description:"Page number"`
	PageRecordCount int            `json:"pageRecordCount" description:"Page record count"`
	TotalCount      int            `json:"totalCount" description:"Total count"`
	Records         []CustomerList `json:"records" description:"Records"`
}

type CustomerList struct {
	CustomerId     string `json:"customerId" description:"Customer id"`
	CustomerName   string `json:"customerName" description:"Customer name"`
	IdNation       string `json:"idNation" description:"Id Nation"`
	IdType         string `json:"idType" description:"Id type"`
	IdNum          string `json:"idNum" description:"Id number"`
	CustomerStatus string `json:"customerStatus" description:"Customer Status:0:Normal 4:Pending 9:Cancel "`
}

func (*SCU0000012I) GetServiceKey() string {
	return "CU000012"
}
