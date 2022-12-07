package util

import (
	"encoding/base64"
	"git.multiverse.io/framework/common/crypto/aes"
	"git.multiverse.io/framework/common/crypto/rsa"
)

// @Desc aes Decrypt
func AesDecrypt(aesEncrypt []byte, aeskey []byte) (aesDecrypt []byte, err error) {
	cryptoIns := aes.NewAES("PKCS5", "ECB", aeskey)
	cryptoBytes, err := cryptoIns.Decrypt(aesEncrypt)
	if nil != err {
		return nil, err
	}
	return (cryptoBytes), nil
}

// @Desc aes Encrypt
func AesEncrypt(originData []byte, aeskey []byte) (aesEncrypt []byte, err error) {
	cryptoIns := aes.NewAES("PKCS5", "ECB", aeskey)
	cryptoBytes, err := cryptoIns.Encrypt(originData)
	if nil != err {
		return nil, err
	}
	return (cryptoBytes), nil
}

// @Desc rsa Decrypt
func RsaDecrypt(rsaEncrypt []byte, rsaPrikey []byte) (rsaDecrypt []byte, err error) {
	rsaIns := rsa.NewRSA(rsaPrikey, nil)
	pretxt, err := rsaIns.Decrypt(rsaEncrypt)
	if nil != err {
		return nil, err
	}
	return pretxt, nil
}

// @Desc rsa Decrypt
func RsaEncrypt(originData []byte, rsaPubkey []byte) (rsaEncrypt []byte, err error) {
	rsaIns := rsa.NewRSA(nil, rsaPubkey)
	pretxt, err := rsaIns.Encrypt(originData)
	if nil != err {
		return nil, err
	}
	return pretxt, nil
}

// @Desc decode base64
func BaseDecode(baseEncode string) (baseDecode []byte, err error) {
	r, err := base64.StdEncoding.DecodeString(baseEncode)
	if nil != err {
		return nil, err
	}
	return r, nil
}

// @Desc encode base64
func BaseEncode(originData []byte) (baseEncode string) {
	return base64.StdEncoding.EncodeToString(originData)
}
