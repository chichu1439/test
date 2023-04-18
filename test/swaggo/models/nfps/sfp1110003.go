package models

type FP111003I struct {
	Data              []byte   `json:"data"`
	InFileName        string   `json:"inFileName"`
	FileName          string   `json:"fileName"`
	BankCode          string   `json:"bankCode"`
	CatalogueType     string   `json:"catalogueType"`
	SaveFileType      string   `json:"saveFileType"`
	MessagePrefix     string   `json:"messagePrefix"`
	FileStatus        string   `json:"fileStatus"`
	PartPath          string   `json:"partPath"`
	RelevantFile      []string `json:"relevantFile"`
	NotSaveFileStatus bool     `json:"saveFileStatus"`
}

type FP111003O struct {
	Status string `json:"status"`
}

func (*FP111003I) GetServiceKey() string {
	return "FP111003"
}
