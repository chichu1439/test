package models

import "time"

// response of GetAllOrgKeyType
type ResGetAllOrgKeyType struct {
	Data []OrgKeyList `json:"data"`
}

type OrgKeyList struct {
	OrganisationId string       `json:"orgId"`
	KeyList        []OrgKeyInfo `json:"keyList"`
}

type OrgKeyInfo struct {
	KeyID       string `json:"keyId"`
	ServiceID   string `json:"serviceId"` //创建key时的组装规则：$orgID+'-SWK', orgID+'-RWK', 例如 NxBgZuIaA3JcLYLrdHLyrp-SWK,当前SKM的serviceID （DB）不能超过50个符
	KeyType     string `json:"keyType"`   //now support DES-,3DES-256, 3DES-168, 3DES-112,  ...
	KeyLength   int    `json:"keyLength"`
	Version     string `json:"version"`
	Usage       string `json:"usage"` // sending-key, receiving-key
	Description string `json:"description"`
	EncryptMode string `json:"encMode"`
	//ECB模式 全称Electronic Codebook模式，译为电子密码本模式
	//CBC模式 全称Cipher Block Chaining模式，译为密文分组链接模式
	//CFB模式 全称Cipher FeedBack模式，译为密文反馈模式
	//OFB模式 全称Output Feedback模式，译为输出反馈模式。
	//CTR模式 全称Counter模式，译为计数器模式。

	Padding string `json:"padding"`
	//NoPadding
	//PKCS5Padding
	//PKCS7Padding
	//ZerosPadding

	ExpireTime time.Time `json:"expireTime"`
	Creator    string    `json:"creator"`
	Updater    string    `json:"updater"`
	CreatedAt  time.Time `json:"createdTime"`
	UpdatedAt  time.Time `json:"updatedTime"`
}

//request of GetOrgKeyType, get given participants' key info.
type ReqGetOrgKeyType struct {
	OrganisationIds []string `json:"orgIds,required"`
}

type ResGetOrgKeyType struct {
	Data []OrgKeyList `json:"data"`
}
