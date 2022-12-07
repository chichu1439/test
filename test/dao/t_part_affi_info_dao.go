package dao

type T_part_affi_info struct {
	AffiInfoPartId   string `orm:"column(affi_info_part_id);pk" description:"参与者id"`
	AffiInfoType     string `orm:"column(affi_info_type);size(10);null" description:"参与者类型"`
	AffiInfoTypeType string `orm:"column(affi_info_type_type);size(10);null" description:"AffiInfoTypeType"`
	AffiInfoContent  string `orm:"column(affi_info_content);size(40);null" description:"AffiInfoContent"`
	AffiInfoRelation string `orm:"column(affi_info_relation);size(10);null" description:"AffiInfoRelation"`
	AffiInfoStatus   string `orm:"column(affi_info_status);size(10);null" description:"AffiInfoStatus"`
	Remarks          string `orm:"column(remarks);size(40);null" description:"Remarks"`
	LastMaintDate    string `orm:"column(last_maint_date);type(date);null" description:"最后更新日期"`
	LastMaintTime    string `orm:"column(last_maint_time);type(time);null" description:"最后更新时间"`
	LastMaintBrno    string `orm:"column(last_maint_brno);size(20);null" description:"最后更新机构"`
	LastMaintTell    string `orm:"column(last_maint_tell);size(20);null" description:"最后更新柜员"`
	TccState         int    `orm:"column(tcc_state);size(10);null" description:"tcc状态"`
}

func (t *T_part_affi_info) TableName() string {
	return "t_part_affi_info"
}
