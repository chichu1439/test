//Version: v0.0.1
package models

type SSV0000004I struct {
	AgreementID              string `json:"agreementID" validate:"required,max=30" description:"contract ID"`             //合约号
	AgreementType            string `json:"agreementType" validate:"required,max=5" description:"contract type"`          //合约类型 10001-个人活期账户;10002-个人定期账户;10003-个人活期保证金账户;10004-个人定期保证金账户;10005-对公活期账户;10006-对公定期账户;10007-对公活期保证金账户;10008-对公定期保证金账户;10009-内部账户;20001-个人贷款账户;20002-对公贷款账户;30001-协议存款;30002-保证金;30003-同业存款;30004-智能通知存款;30005-法人账户透支;30006-集团账户;30007-协定存款;30008-收息账号;30009-活期存款转定期约定;30010-网银开户签约;30011-贷款;30012-贷款收费;30013-普通定期存款;30014-活期保证金存款;30015-定期保证金存款;30016-普通活期存款;40001-客户对账单;-
	AccountName              string `json:"accountName" validate:"max=120" description:"account name"`                    //账户名称
	CustomerContactTelephone string `json:"customerContactTelephone" validate:"max=20" description:"mobile phone number"` //手机号码
	Address                  string `json:"address" validate:"max=200" description:"contact address"`                     //联系地址
	CustomerContactEmail     string `json:"customerContactEmail" validate:"max=200" description:"mailbox"`                //邮箱
	PostalCode               string `json:"postalCode" description:"Postal code"`
	NationCode               string `json:"nationCode" description:"Nation code"`
	Address1                 string `json:"address1" description:"Address1"`
	Address2                 string `json:"address2" description:"Address2"`
	Address3                 string `json:"address3" description:"Address3"`
}

type SSV0000004O struct {
	//Status string
}

func (*SSV0000004I) GetServiceKey() string {
	return "ssv0000004"
}
