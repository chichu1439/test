//
// Copyright 2019 Shenzhen Forms Syntron Information Co., Ltd. All rights reserved.
//

package util

import (
	"encoding/base64"
	"errors"
	"fmt"
	"git.multiverse.io/framework/common/json"

	"git.forms.io/universe/common/crypto/aes"
	commRsa "git.forms.io/universe/common/crypto/rsa"
	"git.forms.io/universe/solapp-sdk/log"
	"github.com/astaxie/beego"
	. "math/rand"
	"time"
)

const (
	DateFormat = "2006-01-02"
	TimeFormat = "2006-01-02 15:04:05"
)

type Date time.Time

func Now() Date {
	return Date(time.Now())
}

func (t *Date) UnmarshalJSON(data []byte) (err error) {
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

// KmsConfig is used to store Kms related configuration information.
type KmsConfig struct {
	KeySpanDuration   uint
	KeyEffectDuration uint
	KeyVersion        string
}

// KmsServiceKey is used to store information related to the service key value.
type KmsServiceKey struct {
	ServiceId string
	Version   string
	KeyId     string
	KeyValue  string
	KeyState  int
}

// KeyType is alias of string
type KeyType string

// KmsKey is used to store the details of each key
type KmsKey struct {
	ServiceId   string
	SystemType  string
	KeyType     string
	KeyId       string
	KeyVersion  int
	KeyValue    string
	State       int
	EffectTime  string
	DeffectTime string
	Creator     string
	CreateTime  string
	Modifier    string
	LastUpdate  string
}

// A type, typically a collection, that satisfies sort.Interface can be
// sorted by the routines in this package. The methods require that the
// elements of the collection be enumerated by an integer index.
type KmsKeySlice []KmsKey

// Len is the number of elements in the collection.
func (s KmsKeySlice) Len() int { return len(s) }

// Swap swaps the elements with indexes i and j.
func (s KmsKeySlice) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

// KmsKeyResult is used to store key information
type KmsKeyResult struct {
	Code      int
	Error     string
	ServiceId string
	Result    KmsKeySlice
}

// KmsOperKeyResult is used to store the results of kms operations
type KmsOperKeyResult struct {
	Code   int
	Error  string
	Result bool
}

// KmsServiceRequest is used to store the request structure that the client requests
// to obtain a single service key information from the Kms adapter.
type KmsServiceRequest struct {
	KeyType    KeyType
	ServiceId  string
	SystemType string
	//EffectTime  string
	DeffectTime string
	Operator    string
	RequestTime string
}

// KmsServicesRequest is used to store the request structure that
// the client requests to obtain multiple service key information from the Kms adapter.
type KmsServicesRequest struct {
	KeyType    KeyType
	ServiceIds []string
}

// KmsKeyRequest is a structure for obtaining a single key information based on the key id
type KmsKeyRequest struct {
	KeyId string
}

// KmsKeyRequest is a structure for obtaining multiple key information according to the key id array
type KmsKeysRequest struct {
	KeyId []string
}

// KmsKeyIdResponse stores a single key id
type KmsKeyIdResponse struct {
	KeyId string
}

// A request structure for storing a client requesting a Kms adapter
// to obtain a plurality of different key types and different service key information.
type KmsKeyValuesRequest struct {
	Services []KmsServiceRequest
}

//KmsKeysResultResponse stores the execution result of multiple kms operations
type KmsKeysResultResponse struct {
	Total  int
	Succ   int
	Result []KmsOperKeyResult
}

// KmsKeysRootResponse stores a single service key information
type KmsKeysRootResponse struct {
	Code    int
	Message string
	Data    interface{}
}

// KmsKeysResponse is used to store multiple service key information
type KmsKeysResponse struct {
	Total  int
	Succ   int
	Result []KmsKeyResult
}

// KmsKeyValuesResponse is used to store information related to multiple service key values.
type KmsKeyValuesResponse struct {
	Total     int
	Succ      int
	KeyValues []KmsServiceKey
}

// KmsDestroyKeysRequest is used to store the request message structure for destroying multiple service keys.
type KmsDestroyKeysRequest struct {
	Services []KmsServiceRequest
}

// KmsServiceKeyUpdateNotice is used by KMS Adapter
// to notify related services for key update
//type KmsServiceKeyUpdateNotice struct {
//	ServiceId string
//	Key string
//	RequestTime  time.Time
//}

// KmsRequestCrypto for all request that need to be encrypted.
type KmsRequestCrypto struct {
	Request string
	Key     string
}

// KmsResponseCrypto for all response that need to be decrypted
type KmsResponseCrypto struct {
	Response string
}

func PackRequest(readRequest interface{}, K1 []byte) ([]byte, error) {
	kmsReq, err := EncryptKmsRequest(readRequest, K1)
	if nil != err {
		return nil, err
	}

	dataByte, err := json.Marshal(kmsReq)
	//log.Debug("req json:" + string(dataByte))
	if nil != err {
		return nil, err
	}
	return dataByte, nil
}

func EncryptKmsRequest(data interface{}, K1 []byte) (KmsRequestCrypto, error) {
	kmsRequestCrypto := KmsRequestCrypto{}
	//rsa
	cryptoAlgorithm := "rsa"
	kmsPubKey, err := FileLoad(beego.AppConfig.String("kms_key::keyPath"))
	if err != nil {
		return kmsRequestCrypto, err
	}
	cryptoIns := aes.NewAES("PKCS5", "GCM", K1)
	rsaInstance := commRsa.NewRSA(nil, kmsPubKey)
	s2, err := rsaInstance.Encrypt(K1)
	if nil != err {
		return kmsRequestCrypto, err
	}

	//2. 加密请求数据
	b, err := json.Marshal(data)
	if err != nil {
		return kmsRequestCrypto, err
	}
	cipherText, err := cryptoIns.Encrypt(b)
	if nil != err {
		return kmsRequestCrypto, err
	}
	M1 := base64.StdEncoding.EncodeToString(cipherText)
	M2 := base64.StdEncoding.EncodeToString(s2)

	req := KmsRequestCrypto{
		Request: M1,
		Key:     M2,
	}
	fmt.Println("Request: " + M1)
	fmt.Println("Key: " + M2)
	fmt.Println("CryptoAlgorithm: " + cryptoAlgorithm)
	return req, nil
}

func DecryptKmsResponse(rspBody []byte, K1 []byte) (dateByte []byte, err error) {
	log.Debugf("SendRequestReplyMsg Succ result=%v\n", json.Get(rspBody, "Response").ToString())
	responseStr := json.Get(rspBody, "Response").ToString()

	//写入返回结果中Response的值  responseStr := ""
	cryptoIns := aes.NewAES("PKCS5", "GCM", K1)
	//cryptoIns, err := sm4.NewSM4("CBC", SMK1)

	preResponse, err := base64.StdEncoding.DecodeString(responseStr)
	if nil != err {
		return nil, err
	}
	originResponse, err := cryptoIns.Decrypt(preResponse)
	if nil != err {
		return nil, err
	}

	kmsKeysRootResponse := KmsKeysRootResponse{}
	err = json.Unmarshal(originResponse, &kmsKeysRootResponse)
	if nil != err {
		return nil, err
	}
	log.Debug(kmsKeysRootResponse)

	if kmsKeysRootResponse.Code != 0 {
		return nil, errors.New(kmsKeysRootResponse.Message)
	}

	dataByte, err := json.Marshal(kmsKeysRootResponse.Data)
	if nil != err {
		return nil, err
	}

	return dataByte, nil
}

func UnpackGetKeyWithCrypto(rspBody []byte, K1 []byte) (rsp []KmsKey, err error) {
	//K1 = []byte("1234567890ABCDEF1234567890ABCDEF")
	dataByte, err := DecryptKmsResponse(rspBody, K1)
	if nil != err {
		return
	}

	kmsKeyResult := KmsKeyResult{}
	err = json.Unmarshal(dataByte, &kmsKeyResult)
	if nil != err {
		return
	}
	//log.Debug(getKeyWithCryptoRsp.Result)

	return kmsKeyResult.Result, nil
}

//Krand . 随机字符串
func Krand(size int, kind int) []byte {
	ikind, kinds, result := kind, [][]int{[]int{10, 48}, []int{26, 97}, []int{26, 65}}, make([]byte, size)
	isAll := kind > 2 || kind < 0
	Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		if isAll { // random ikind
			ikind = Intn(3)
		}
		scope, base := kinds[ikind][0], kinds[ikind][1]
		result[i] = uint8(base + Intn(scope))
	}
	return result
}
