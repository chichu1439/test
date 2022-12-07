/*
 * Copyright (c) 2020. Shenzhen Forms Syntron Information Co., Ltd. All rights reserved.
 * Platform universe is a distributed platform of agile operation, which has an open business model.
 * Truly distributed Architecture capability with latest techs such as Event Mesh, Containerized Micro services.
 * Overhead enables light-weight deployment and Edge Computing Mode.
 * Centric with BCP and cost efficiency at the heart of the design principle.
 * With model driven operation and development that further enhanced the overall agility.
 */

package encry

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"strings"
)

var (
	KmsCache *KmsKey
	ServiceConf      = &ServicesConfig{}
)
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
	EncryType   string
}
type ServicesConfig struct {
	Organization        string
	GroupDcn            string
	DataCenterNode      string
	DcnWithoutFirstByte string
	InstanceID          string
	ReqFormId           string
	RespFormId          string
	AppName             string
	TopicId             string
	Langs               string

	ParamTopic string
	ParamGroup string
	TxnParam   map[string]string

	EncryDlsKeyFlag bool
	DecryPubKeyFlag bool
	EncryLogInfoFlag bool

	KmsPath        string
	KmsKeyType     string
	KmsServiceId   string
	KmsSystemType  string
	KmsDeffectTime string
	KmsOperator    string
}

//var remoteCall = comm.NewRemoteClientFactory().CreateClient()
type KmsService struct {
}

//加密 && 解密
func AesGcmEncrypt(plaintext []byte, key []byte, encType string) ([]byte, error) {
	switch strings.ToUpper(strings.Trim(encType, " ")) {
	case "GCM":
		return aesEncryptGCM(plaintext, key), nil
	case "CFB":
		return aesEncryptCFB(plaintext, key), nil
	case "ECB":
		return aesEncryptECB(plaintext, key), nil
	case "CBC":
		return aesEncryptCBC(plaintext, key), nil
	default:
		panic(fmt.Errorf("Not Support Encry Type[%s]", encType))
	}
}
func AesGcmDecrypt(ciphertext []byte, key []byte, decType string) ([]byte, error) {
	switch strings.ToUpper(strings.Trim(decType, " ")) {
	case "GCM":
		return aesDecryptGCM(ciphertext, key), nil
	case "CFB":
		return aesDecryptCFB(ciphertext, key), nil
	case "ECB":
		return aesDecryptECB(ciphertext, key), nil
	case "CBC":
		return aesDecryptCBC(ciphertext, key), nil
	default:
		panic(fmt.Errorf("Not Support Deccry Type[%s]", decType))
	}
}
func aesEncryptGCM(origData []byte, key []byte) (encrypted []byte) {
	c, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	gcm, err :=  cipher.NewGCMWithNonceSize(c, 0)
	if err != nil {
		panic(err)
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err)
	}
	return gcm.Seal(nonce, nonce, origData, nil)
}
func aesDecryptGCM(origData []byte, key []byte) (decrypted []byte) {
	c, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	gcm, err := cipher.NewGCMWithNonceSize(c, 0)
	if err != nil {
		panic(err)
	}

	nonceSize := gcm.NonceSize()
	if len(origData) < nonceSize {
		panic(fmt.Errorf("aesDecryptGCM Faild, data Length is sorter than nonceSize"))
	}

	nonce, ciphertext := origData[:nonceSize], origData[nonceSize:]
	decrypted, err = gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		panic(err)
	}
	return
}
func aesEncryptCFB(origData []byte, key []byte) (encrypted []byte) {
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	encrypted = make([]byte, aes.BlockSize+len(origData))
	iv := encrypted[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}
	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(encrypted[aes.BlockSize:], origData)
	return encrypted
}
func aesDecryptCFB(origData []byte, key []byte) (decrypted []byte) {
	block, _ := aes.NewCipher(key)
	if len(origData) < aes.BlockSize {
		panic("ciphertext too short")
	}
	iv := origData[:aes.BlockSize]
	decrypted = origData[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(decrypted, decrypted)
	return
}
func aesEncryptECB(origData []byte, key []byte) (encrypted []byte) {
	cipher, _ := aes.NewCipher(generateKey(key))
	length := (len(origData) + aes.BlockSize) / aes.BlockSize
	plain := make([]byte, length*aes.BlockSize)
	copy(plain, origData)
	pad := byte(len(plain) - len(origData))
	for i := len(origData); i < len(plain); i++ {
		plain[i] = pad
	}
	encrypted = make([]byte, len(plain))
	// 分组分块加密
	for bs, be := 0, cipher.BlockSize(); bs <= len(origData); bs, be = bs+cipher.BlockSize(), be+cipher.BlockSize() {
		cipher.Encrypt(encrypted[bs:be], plain[bs:be])
	}

	return encrypted
}
func aesDecryptECB(origData []byte, key []byte) (decrypted []byte) {
	cipher, _ := aes.NewCipher(generateKey(key))
	decrypted = make([]byte, len(origData))
	//
	for bs, be := 0, cipher.BlockSize(); bs < len(origData); bs, be = bs+cipher.BlockSize(), be+cipher.BlockSize() {
		cipher.Decrypt(decrypted[bs:be], origData[bs:be])
	}

	trim := 0
	if len(decrypted) > 0 {
		trim = len(decrypted) - int(decrypted[len(decrypted)-1])
	}

	return decrypted[:trim]
}
func aesEncryptCBC(origData []byte, key []byte) (encrypted []byte) {
	// 分组秘钥
	// NewCipher该函数限制了输入k的长度必须为16, 24或者32
	block, _ := aes.NewCipher(key)
	blockSize := block.BlockSize()                              // 获取秘钥块的长度
	origData = pkcs5Padding(origData, blockSize)                // 补全码
	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize]) // 加密模式
	encrypted = make([]byte, len(origData))                     // 创建数组
	blockMode.CryptBlocks(encrypted, origData)                  // 加密
	return encrypted
}
func aesDecryptCBC(origData []byte, key []byte) (decrypted []byte) {
	// 分组秘钥
	// NewCipher该函数限制了输入k的长度必须为16, 24或者32
	block, _ := aes.NewCipher(key)                              // 分组秘钥
	blockSize := block.BlockSize()                              // 获取秘钥块的长度
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize]) // 加密模式
	decrypted = make([]byte, len(origData))                    // 创建数组
	blockMode.CryptBlocks(decrypted, origData)                 // 解密
	decrypted = pkcs5UnPadding(decrypted)                       // 去除补全码
	return
}
func pkcs5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}
func pkcs5UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}
func generateKey(key []byte) (genKey []byte) {
	genKey = make([]byte, 16)
	copy(genKey, key)
	for i := 16; i < len(key); {
		for j := 0; j < 16 && i < len(key); j, i = j+1, i+1 {
			genKey[j] ^= key[i]
		}
	}
	return genKey
}

//通过kms加密
func KmsEncrypt(plaintext string) (ciphertext string) {
	var encoded = plaintext
	if ServiceConf.EncryDlsKeyFlag {
		//service :=KmsService{}
		if KmsCache == nil {
			panic(fmt.Errorf("get kms failed! kmsCache is null"))
		}
		key := KmsCache.KeyValue
		bytes, err := AesGcmEncrypt([]byte(plaintext), []byte(key), KmsCache.EncryType)
		if err != nil {
			fmt.Printf("Aes Gcm Encrpy failed!%s", err)
			panic(err)
		}
		encoded = base64.StdEncoding.EncodeToString(bytes)
		if ServiceConf.EncryLogInfoFlag {
			//fmt.Printf("KmsEncrypt[%s] Result[%s]", dlscfg.Desensitization(plaintext), dlscfg.Desensitization(encoded))
		}
	}

	return encoded
}

//通过kms解密
func KmsDecrypt(ciphertext string) (plaintext string) {
	plaintext = ciphertext
	if len(ciphertext) == 0 {
		return
	}
	if ServiceConf.DecryPubKeyFlag {
		decoded, err := base64.StdEncoding.DecodeString(ciphertext)
		if KmsCache == nil {
			panic(fmt.Errorf("get kms failed! kmsCache is null"))
		}
		key := KmsCache.KeyValue
		bytes, err := AesGcmDecrypt([]byte(decoded), []byte(key), KmsCache.EncryType)
		if err != nil {
			fmt.Printf("Aes Gcm Decrpy failed!%s", err)
		} else {
			plaintext = string(bytes)
		}
	}

	return plaintext
}
