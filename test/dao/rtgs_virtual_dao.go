package dao

type RtgsVirtual struct {
	PartId          string `orm:"column(part_id);pk"`
	PartType        string `orm:"column(part_type);size(10)"`
	PartStatus      string `orm:"column(part_status);size(10)"`
	PartCnName      string `orm:"column(part_cn_name);size(30)"`
	PartCnNameAbbre string `orm:"column(part_cn_name_abbre);size(20)"`
	PartEnName      string `orm:"column(part_en_name);size(30)"`
	PartEnAbbre     string `orm:"column(part_en_abbre);size(20)"`
	PartCountryCode string `orm:"column(part_country_code);size(10)"`
	PartAreaCode    string `orm:"column(part_area_code);size(10)"`
	PaymentPart     string `orm:"column(payment_part);size(30);null"`
	SwiftCode       string `orm:"column(swift_code);size(30);null"`
	PartAddr        string `orm:"column(part_addr);size(40)"`
	PartEmail       string `orm:"column(part_email);size(30)"`
	PartBic         string `orm:"column(part_bic);size(30)"`
	PartPostcode    string `orm:"column(part_postcode);size(20)"`
	BuildDate       string `orm:"column(build_date);type(date);null"`
	PartCertiType1  string `orm:"column(part_certi_type1);size(2)"`
	PartCertiId1    string `orm:"column(part_certi_id1);size(50)"`
	PartCertiType2  string `orm:"column(part_certi_type2);size(2)"`
	PartCertiId2    string `orm:"column(part_certi_id2);size(50)"`
	PartIp          string `orm:"column(part_ip);size(50)"`
	StmiAccountNo   string `orm:"column(stmi_account_no);size(30)"`
	LastMaintDate   string `orm:"column(last_maint_date);type(date);null"`
	LastMaintTime   string `orm:"column(last_maint_time);type(time);null"`
	LastMaintBrno   string `orm:"column(last_maint_brno);size(30);null"`
	LastMaintTell   string `orm:"column(last_maint_tell);size(30);null"`
	TccState        int    `orm:"column(tcc_state)"`
}

func (t *RtgsVirtual) TableName() string {
	return "rtgs_virtual"
}
