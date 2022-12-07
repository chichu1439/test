//Version: v1.0.0.0
package models

// swagger:parameters SMP9CM0002I
type SMP9CM0002I struct {
	Language    string   `json:"language"`
	TextIdList  []string `json:"textIdList"`
	AgreementId string   `xorm:"VARCHAR(45)"`
}

// swagger:parameters SMP9CM0002O
type SMP9CM0002O struct {
	TextList []*CmText `json:"textList"`
}

type CmText struct {
	TextId   string `json:"textId"`
	Language string `json:"language"`
	Content  string `json:"content"`
}

func (*SMP9CM0002I) GetServiceKey() string {
	return "MP9CM002"
}
