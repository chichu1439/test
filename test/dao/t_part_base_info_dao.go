package dao

type T_part_base_info struct {
	PartId          string `orm:"column(part_id);pk" description:"参与者id"`
	PartType        string `orm:"column(part_type);size(10);null" description:"参与者类型"`
	PartStatus      string `orm:"column(part_status);size(10);null" description:"参与者状态"`
	PartCnName      string `orm:"column(part_cn_name);size(30);null" description:"参与者母语名称"`
	PartCnNameAbbre string `orm:"column(part_cn_name_abbre);size(20);null" description:"参与者母语名称简称"`
	PartCountryCode string `orm:"column(part_country_code);size(10);null" description:"参与者国家代码"`
	PartAreaCode    string `orm:"column(part_area_code);size(10);null" description:"参与者地区代码"`
	PaymentPart     string `orm:"column(payment_part);size(30);null" description:"支付系统银行行号"`
	SwiftCode       string `orm:"column(swift_code);size(10);null" description:"swift代码"`
	PartAddr        string `orm:"column(part_addr);size(10);null" description:"参与者地址"`
	PartEmail       string `orm:"column(part_email);size(10);null" description:"参与者邮箱"`
	PartBic         string `orm:"column(part_bic);size(10);null" description:"参与者bic"`
	PartPostcode    string `orm:"column(part_postcode);size(10);null" description:"参与者邮编"`
	LastMaintDate   string `orm:"column(last_maint_date);size(10);null" description:"最后更新日期"`
	LastMaintTime   string `orm:"column(last_maint_time);size(10);null" description:"最后更新时间"`
	LastMaintBrno   string `orm:"column(last_maint_brno);size(10);null" description:"最后更新机构"`
	LastMaintTell   string `orm:"column(last_maint_tell);size(10);null" description:"最后更新柜员"`
	TccState        int    `orm:"column(tcc_state);size(10);null" description:"tcc状态"`
}

func (t *T_part_base_info) TableName() string {
	return "t_part_base_info"
}
