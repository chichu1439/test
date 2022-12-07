package dao

type T_sys_status_control struct {
	KeprcdNo        string `orm:"column(keprcd_no);pk" description:"记录编号"`
	SysStusCd       string `orm:"column(sys_stus_cd);size(1);null" description:"系统状态代码  O-日间服务F-日终服务"`
	SysCtofModeCd   string `orm:"column(sys_ctof_mode_cd);size(2);null" description:"系统日切模式代码  N-自然日切C-受控日切"`
	OnlineBizDt     string `orm:"column(online_biz_dt);type(date);null" description:"联机业务日期"`
	BatBizDt        string `orm:"column(bat_biz_dt);type(date);null" description:"批量业务日期"`
	NxtoneBizDt     string `orm:"column(nxtone_biz_dt);type(date);null" description:"下一业务日期"`
	LstoneBizDt     string `orm:"column(lstone_biz_dt);type(date);null" description:"上一业务日期 "`
	CtofTm          string `orm:"column(ctof_tm);type(timestamp);null" description:"日切时间  描述事件过程长短和发生顺序的度量,格式:HH:MM:SS,"`
	TranOnlineTm    string `orm:"column(tran_online_tm);type(timestamp);null" description:"转联机时间  描述事件过程长短和发生顺序的度量,格式:HH:MM:SS,"`
	FinlModfyOrgNo  string `orm:"column(finl_modfy_org_no);size(4);null" description:"最后修改机构号  记录事件的最后修改人从属机构号,"`
	FinlModfyTelrNo string `orm:"column(finl_modfy_telr_no);size(6);null" description:"最后修改柜员号  记录最后修改事件的柜员编号,柜员一般包括有交易柜员,复核柜员,授权柜员等不同角色,"`
	TccState        int    `orm:"column(tcc_state);null"`
}

func (t *T_sys_status_control) TableName() string {
	return "t_sys_status_control"
}
