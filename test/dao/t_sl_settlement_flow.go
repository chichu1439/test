package dao

type TSlSettlementFlow struct {
	ID                 int     `json:"id" xorm:"'id' not null pk autoincr INT"`
	NettingId          string  `json:"netting_id"`
	TransId            string  `json:"trans_id" xorm:"not null comment('transaction id') VARCHAR(35)"`
	ClrSysRef          string  `json:"clr_sys_ref" xorm:"not null comment('cleraing system reference number') VARCHAR(35)"`
	TransType          string  `json:"trans_type" xorm:"not null comment('01-Credit Transfer; 02-Direct Debit; 03-Return / Refund; 04-Balance Sweeping Notification; 05-IP Balance Adjustment Notification; 06-Payment Status Enquiry; 07-Account Balance Enquiry;') VARCHAR(2)"`
	PaymentPurposeCode string  `json:"payment_purpose_code" xorm:"not null comment('CXSALA-Salary and Benefits Payment; CXMRCH-Merchant Payment; CXBSNS-General Business Payment; CXPSNL-General Personal Payment; DDBILL-Bills Payment; DDTOPU-Account Top-up Payment; DDECOM-E-commerce Payment; DDOTHR-Others (General Direct Debit Payment); RPRTRN-Payment Return; RPRFND-Refund; SWAUTO-Auto Account Balance Sweeping; SWMANU-Manual Account Balance Sweeping; SIDRCR-CB Direct Debit/Credit; CPBADJ-IP Balance Adjustment; ') VARCHAR(10)"`
	TransDirection     string  `json:"trans_direction" xorm:"not null comment('DR-debit, CR-credit') VARCHAR(2)"`
	TransPart          string  `json:"trans_part" xorm:"not null comment('transaction participant') VARCHAR(35)"`
	TransAcctNo        string  `json:"trans_acct_no" xorm:"default 'NULL' comment('transaction account no') VARCHAR(30)"`
	CounterpartyPart   string  `json:"counterparty_part" xorm:"default 'NULL' comment('counterparty participant') VARCHAR(35)"`
	CounterpartyAcctNo string  `json:"counterparty_acct_no" xorm:"default 'NULL' comment('counterparty account no') VARCHAR(30)"`
	Currency           string  `json:"currency" xorm:"not null comment('CNY; JPY; THB; GBP; USD; EUR;') VARCHAR(3)"`
	TransAmount        float64 `json:"trans_amount" xorm:"not null comment('transaction amount') DECIMAL(18,5)"`
	AccountingType     string  `json:"accounting_type" xorm:"not null comment('00-No need to net,01-db balance netting,02-redis/db balance netting') VARCHAR(2)"`
	AccountingStatus   string  `json:"accounting_status" xorm:"comment('00-complete; 01-incomplete';02-processing) VARCHAR(2)"`
	StlmtDate          string  `json:"stlmt_date" xorm:"not null comment('settlement date') DATE"`
	StlmtTime          string  `json:"stlmt_time" xorm:"not null comment('settlement time') TIME"`
	SettlementUnixTime string  `json:"settlement_unix_time" xorm:"not null comment('settlement unit time') BIGINT created"`
	CreateTime         string  `json:"create_time" xorm:"comment('create time') DATETIME created"`
	UpdateTime         string  `json:"update_time" xorm:"comment('update time') DATETIME updated"`
}
