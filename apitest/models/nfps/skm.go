package models

import "time"

type GetSkmSSLPemRequest struct {
	Algorithm string `json:"algorithm"` //rsa,sm2
}

type SkmPublicPemResponse struct {
	PublicPem string `json:"publicPem"`
}


const (
	DateFormat = "2006-01-02"
	TimeFormat = "2006-01-02 15:04:05"
)

type Date time.Time

func Now() Date {
	return Date(time.Now())
}

func (t *Date) UnmarshalJSON(data []byte) (err error) {
	if string(data) == "\"\"" {
		return
	}
	now, err := time.ParseInLocation(`"`+TimeFormat+`"`, string(data), time.Local)
	*t = Date(now)
	return
}

func (t Date) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(TimeFormat)+2)
	b = append(b, '"')
	b = time.Time(t).AppendFormat(b, TimeFormat)
	b = append(b, '"')
	return b, nil
}

func (t Date) String() string {
	return time.Time(t).Format(TimeFormat)
}

type GetKeyRequest struct {
	KeyType     string `validate:"required" json:"keyType"`
	ServiceId   string `validate:"required" json:"serviceId"`
	SystemType  string `validate:"required" json:"systemType"`
	EffectTime  Date   `json:"effectTime"`
	ExpireTime  Date   `json:"expireTime"`
	Operator    string `validate:"required" json:"operator"`
	RequestTime Date   `validate:"required" json:"requestTime"`
}

type RequestCrypto struct {
	Request         string `json:"request"`
	Key             string `json:"key"`
	CryptoAlgorithm string `json:"cryptoAlgorithm"` //sm2，rsa，default rsa
}

type ResponseCrypto struct {
	Response string `json:"response"`
}

type KeyInfo struct {
	ServiceId  string `json:"serviceId,omitempty"`
	SystemType string `json:"systemType,omitempty"`
	KeyType    string `json:"keyType,omitempty"`
	KeyId      string `json:"keyId,omitempty"`
	KeyVersion string `json:"keyVersion,omitempty"`
	KeyValue   string `json:"keyValue,omitempty"`
	State      *int   `json:"state,omitempty"`
	EffectTime string `json:"effectTime,omitempty"`
	ExpireTime string `json:"expireTime,omitempty"`
	Creator    string `json:"creator"`
	CreateTime string `json:"createTime"`
	Modifier   string `json:"modifier"`
	LastUpdate string `json:"lastUpdate"`
}

type KeyResult struct {
	ServiceId    string    `json:"serviceId"`
	ResponseTime string    `json:"responseTime"`
	Result       []KeyInfo `json:"result"`
}
