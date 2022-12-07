package dao

type TNettingRecordInfo struct {
	Id              int64  `json:"id" xorm:"not null pk autoincr INT(11)"`
	NettingId       string `json:"netting_id"`
	NettingDate     string
	NettingSequence int
	EffectiveDate   string `json:"effective_date"`
	ExpireDate      string `json:"expire_date"`
	//Version       int    `json:"version"`
	CreateTime string `json:"create_time" xorm:"comment('create time') DATETIME created"`
	UpdateTime string `json:"update_time" xorm:"comment('update time') DATETIME updated"`
}

func (*TNettingRecordInfo) TableName() string {
	return "t_netting_record_info"
}
