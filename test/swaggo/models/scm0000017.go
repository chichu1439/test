package models

type CM000017I struct {
}

type CM000017O struct {
	RecordNum                string `json:"recordNum" description:"Record number"`                                                                                      // 记录编号
	SystemStatusCode         string `json:"systemStatusCode" description:"System status code(O - day service; F - end of day service)"`                                 // 系统状态代码 O-日间服务 F-日终服务
	SystemSwitchDateModeCode string `json:"systemSwitchDateModeCode" description:"System switch date mode code(N-natural diurnal shear; C-controlled daily switching)"` // 系统日切模式代码 N-自然日切 C-受控日切
	OnlineBusinessDate       string `json:"onlineBusinessDate" description:"Online business date"`                                                                      // 联机业务日期
	BatchBusinessDate        string `json:"batchBusinessDate" description:"Batch business date"`                                                                        // 批量业务日期
	NextoneBusinessDate      string `json:"nextoneBusinessDate" description:"Nextone business date"`                                                                    // 下一业务日期
	LastoneBusinessDate      string `json:"lastoneBusinessDate" description:"Lastone business date"`                                                                    // 上一业务日期
	SwitchDateTime           string `json:"switchDateTime" description:"Switch date time"`                                                                              // 日切时间
	TransOnlineTime          string `json:"transOnlineTime" description:"Trans online time"`                                                                            // 转联机时间
}

func (*CM000017I) GetServiceKey() string {
	return "CM000017"
}
