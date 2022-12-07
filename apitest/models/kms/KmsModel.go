package models

//* GetSkmSSLPem

type GetSkmSSLPem struct {
 Algorithm string `json:"algorithm" description:"Algorithm"` 
}

type GetSkmSSLPemResponse struct {
 PublicPem string `json:"publicPem" description:"Public pem"` 
}

//*
//*GetRotateKeyWithCrypto

type GetRotateKeyWithCrypto struct {
 KeyType     string `json:"keyType" description:"Key type"` 
 ServiceId   string `json:"serviceId" description:"Service id"` 
 SystemType  string `json:"systemType" description:"System type"` 
 EffectTime  string `json:"effectTime" description:"Effect time"` 
 ExpireTime  string `json:"expireTime" description:"Expire time"` 
 Operator    string `json:"operator" description:"Operator"` 
 RequestTime string `json:"requestTime" description:"Request time"` 
}

type KmsRequestCrypto struct {
 Request         string `json:"request" description:"Request"` 
 Key             string `json:"key" description:"Key"` 
 CryptoAlgorithm string `json:"cryptoAlgorithm" description:"Crypto algorithm"` 
}

//
type CommResp struct {
// ErrorCode string      `json:"errorCode" description:"Error code"` 
// ErrorMsg  string      `json:"errorMsg" description:"Error msg"` 
//	Response  interface{} `json:"response"`
//}

//Response

type GetRotateKeyWithCryptoResponse struct {
 Response string   `json:"response"`
}

type KmsKeysRootResponse struct {
 ResponseTime string                      `json:"responseTime" description:"Response time"` 
 ServiceID    string                      `json:"serviceId" description:"Service i d"` 
 Result       []KmsKeysRootResponseResult `json:"result" description:"Result"` 
}


type KmsKeysRootResponseResult struct {
 ServiceID   string `json:"ServiceId" description:"Service i d"` 
 SystemType  string `json:"SystemType" description:"System type"` 
 KeyType     string `json:"KeyType" description:"Key type"` 
 KeyID       string `json:"KeyId" description:"Key i d"` 
 KeyVersion  string `json:"KeyVersion" description:"Key version"` 
 KeyValue    string `json:"KeyValue" description:"Key value"` 
 State       int    `json:"State" description:"State"` 
 EffectTime  string `json:"EffectTime" description:"Effect time"` 
 DeffectTime string `json:"DeffectTime" description:"Deffect time"` 
 Creator     string `json:"Creator" description:"Creator"` 
 CreateTime  string `json:"CreateTime" description:"Create time"` 
 Modifier    string `json:"Modifier" description:"Modifier"` 
 LastUpdate  string `json:"LastUpdate" description:"Last update"` 
}

//
//func PackRequestK1(readRequest interface{}, K1 []byte) ([]byte, error) {
//	kmsReq, err := EncryptKmsRequest(readRequest, K1)
//	if nil != err {
//		return nil, err
//	}
//
//	dataByte, err := json.Marshal(kmsReq)
//	//log.Debug("req json:" + string(dataByte))
//	if nil != err {
//		return nil, err
//	}
//	return dataByte, nil
//}
//
//func EncryptKmsRequest(data interface{}, K1 []byte) (KmsRequestCrypto, error) {
//	kmsRequestCrypto := KmsRequestCrypto{}
//	//rsa
//	cryptoAlgorithm := "rsa"
//	kmsPubKey, _ := FileLoad(beego.AppConfig.String("kms_key::keyPath"))
//	cryptoIns := aes.NewAES("PKCS5", "GCM", K1)
//	rsaInstance := commRsa.NewRSA(nil, kmsPubKey)
//	s2, err := rsaInstance.Encrypt(K1)
//	if nil != err {
//		return kmsRequestCrypto, err
//	}
//
//	//2. 加密请求数据
//	b, _ := json.Marshal(data)
//	cipherText, err := cryptoIns.Encrypt(b)
//	if nil != err {
//		return kmsRequestCrypto, err
//	}
//	M1 := base64.StdEncoding.EncodeToString(cipherText)
//	M2 := base64.StdEncoding.EncodeToString(s2)
//
//	req := KmsRequestCrypto{
//		Request: M1,
//		Key:     M2,
//	}
//	fmt.Println("Request: " + M1)
//	fmt.Println("Key: " + M2)
//	fmt.Println("CryptoAlgorithm: " + cryptoAlgorithm)
//	return req, nil
//}
//
////func DecryptKmsResponse(rspBody []byte, K1 []byte) (dateByte []byte, err error) {
////	log.Debug("SendRequestReplyMsg Succ result=%v\n", json.Get(rspBody, "Response").ToString())
////	responseStr := json.Get(rspBody, "Response").ToString()
////
////	//写入返回结果中Response的值  responseStr := ""
////	cryptoIns := aes.NewAES("PKCS5", "GCM", K1)
////	//cryptoIns, err := sm4.NewSM4("CBC", SMK1)
////
////	preResponse, err := base64.StdEncoding.DecodeString(responseStr)
////	if nil != err {
////		return nil, err
////	}
////	originResponse, err := cryptoIns.Decrypt(preResponse)
////	if nil != err {
////		return nil, err
////	}
////
////	kmsKeysRootResponse := KmsKeysRootResponse{}
////	err = json.Unmarshal(originResponse, &kmsKeysRootResponse)
////	if nil != err {
////		return nil, err
////	}
////	log.Debug(kmsKeysRootResponse)
////
////	if kmsKeysRootResponse.Code != 0 {
////		return nil, errors.New(kmsKeysRootResponse.Message)
////	}
////
////	dataByte, err := json.Marshal(kmsKeysRootResponse.Data)
////	if nil != err {
////		return nil, err
////	}
////
////	return dataByte, nil
////}
