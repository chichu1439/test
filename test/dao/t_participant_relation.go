package dao

type TPartRelation struct {
	SeqNo                    string `json:"seq_no"`
	DirectPartClearingCode   string `json:"direct_part_clearing_code"`
	DirectPartName           string `json:"direct_part_name"`
	IndirectPartClearingCode string `json:"indirect_part_clearing_code"`
	IndirectPartName         string `json:"indirect_part_name"`
	Currency                 string `json:"currency"`
	Remark                   string `json:"remark"`
	RelationStatus           string `json:"relation_status"`
	CreateTime               string `json:"create_time"`
	UpdateTime               string `json:"update_time"`
}

func (t *TPartRelation) TableName() string {
	return "t_part_relation"
}
