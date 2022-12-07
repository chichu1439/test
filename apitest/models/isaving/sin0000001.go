//Version: v1.0.0.0
package models

type IN000001I struct {
	FuncCode string          `json:"funcCode" description:"Func code" validate:"required"` //R-查询，C-新增，U-修改，D-删除
	Item     []IN000001IItem `json:"item" description:"Item"`
}

type IN000001IItem struct {
	Title             *string  `json:"title" description:"Title" validate:"required"`
	Seq               *string  `json:"seq" description:"Seq" validate:"required,len=4"`
	Name              *string  `json:"name,omitempty" description:"Name" validate:"omitempty"`
	Status            *string  `json:"status,omitempty" description:"Status" validate:"omitempty"`
	BalDirection      *string  `json:"balDirection,omitempty" description:"Bal direction" validate:"omitempty"`
	AccountingOrgAttr *string  `json:"accountingOrgAttr,omitempty" description:"Accounting org attr" validate:"omitempty"`
	CurrencyAttr      *string  `json:"currencyAttr,omitempty" description:"Currency attr" validate:"omitempty"`
	BalRedFlag        *string  `json:"balRedFlag,omitempty" description:"Bal red flag" validate:"omitempty"`
	OnAccountFlag     *string  `json:"onAccountFlag,omitempty" description:"On account flag" validate:"omitempty"`
	OnAccountType     *string  `json:"onAccountType,omitempty" description:"On account type" validate:"omitempty"`
	OnAccountMaxDays  *int     `json:"onAccountMaxDays,omitempty" description:"On account max days" validate:"omitempty"`
	OverdraftFlag     *string  `json:"overdraftFlag,omitempty" description:"Overdraft flag" validate:"omitempty"`
	AutoOpenFlag      *string  `json:"autoOpenFlag,omitempty" description:"Auto open flag" validate:"omitempty"`
	TimeLimit         *int     `json:"timeLimit,omitempty" description:"Time limit" validate:"omitempty"`
	TimeLimitUnit     *string  `json:"timeLimitUnit,omitempty" description:"Time limit unit" validate:"omitempty"`
	HotAccountFlag    *string  `json:"hotAccountFlag,omitempty" description:"Hot account flag" validate:"omitempty"`
	LimitAmount       *float64 `json:"limitAmount,omitempty" description:"Limit amount" validate:"omitempty"`
	EodBalCheck       *string  `json:"eodBalCheck,omitempty" description:"Eod bal check" validate:"omitempty"`
	EffectDate        *string  `json:"effectDate,omitempty" description:"Effect date" validate:"omitempty"`
	ExpireDate        *string  `json:"expireDate,omitempty" description:"Expire date" validate:"omitempty"`
}

type IN000001O struct {
	Item []IN000001OItem `json:"item" description:"Item"`
}

type IN000001OItem struct {
	Title             string  `json:"title" description:"Title"`
	Seq               string  `json:"seq" description:"Seq"`
	Name              string  `json:"name" description:"Name"`
	Status            string  `json:"status" description:"Status"`
	BalDirection      string  `json:"balDirection" description:"Bal direction"`
	AccountingOrgAttr string  `json:"accountingOrgAttr" description:"Accounting org attr"`
	CurrencyAttr      string  `json:"currencyAttr" description:"Currency attr"`
	BalRedFlag        string  `json:"balRedFlag" description:"Bal red flag"`
	OnAccountType     string  `json:"onAccountType" description:"On account type"`
	OnAccountFlag     string  `json:"onAccountFlag" description:"On account flag"`
	OnAccountMaxDays  int     `json:"onAccountMaxDays" description:"On account max days"`
	OverdraftFlag     string  `json:"overdraftFlag" description:"Overdraft flag"`
	AutoOpenFlag      string  `json:"autoOpenFlag" description:"Auto open flag"`
	HotAccountFlag    string  `json:"hotAccountFlag" description:"Hot account flag"`
	TimeLimit         int     `json:"timeLimit" description:"Time limit"`
	TimeLimitUnit     string  `json:"timeLimitUnit" description:"Time limit unit"`
	LimitAmount       float64 `json:"limitAmount" description:"Limit amount"`
	EodBalCheck       string  `json:"eodBalCheck" description:"Eod bal check"`
	EffectDate        string  `json:"effectDate" description:"Effect date"`
	ExpireDate        string  `json:"expireDate" description:"Expire date"`
}

func (*IN000001I) GetServiceKey() string {
	return "sin0000001"
}
