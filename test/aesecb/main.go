package main

import (
	"encoding/base64"
	"fmt"
	"git.multiverse.io/framework/common/crypto/aes"
	"git.multiverse.io/framework/common/crypto/rsa"
	"git.multiverse.io/framework/log"
)

var aesKey = "abcdefghijk01234"
var aesEnstr = "oldKF0KSHDCEQoj0uTWZy2gLEdMawt5Ru9tuNVL8NgS5ulXTm6VqM16Ay7viZ7VMVAEZiSRO0Q0K3tTrx8dCsQ8nTfmNSoHhLeA5L8mf+kBgVAf7LyPATArafjWAD7FutYEUQzAe3GkQaSsp9YwtAzyZ+Fo7CjHSsiN9bXFrZ97CZwGIjNrI1T10/WxbKmffaeoPa7BTxtvTC/KtZ4i9E51EawFqI9sl7Y+m7xqb6B21Urq1lr059SbkFogniVeZjqsstz6tvtwsoE7l2sQ0wuCpBiGgGdx9MCs7h02f/bQq0jsJD5IN5lNpBp2Kvkmtezi7wY1mlINvkDQWl0qYWMZheMIJCQx55F8sj+LZ95HT+xEHqsx3s67Pm/3VwQjbCuxq4qJqUfWTOJFry6FJEw=="

func main() {
	decrypt, err := aesDecrypt(aesKey, aesEnstr)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(decrypt)
}
func aesDecrypt(aeskey string, aesDncrypt string) (decrypt string, err error) {
	log.Debugs("Start aesDecrypt!")
	baseDecode, errs := base64.StdEncoding.DecodeString(aesDncrypt)
	if nil != errs {
		return "", errs
	}
	log.Debugs("Try base64_DecodeString Successful!")
	aesDecrypt, errs := AesDecrypt(baseDecode, []byte(aeskey))
	if nil != errs {
		return "", errs
	}

	log.Debugs("Try AES_Decrypt Successful! aesDecrypt = %s", aesDecrypt)
	return string(aesDecrypt), err
}

func aesEncrypt(aeskey string, toBeEncrypt string) (encrypt string, err error) {
	log.Debugs("Start aesEncrypt!")
	baseDecode, errs := AesEncrypt([]byte(toBeEncrypt), []byte(aeskey))
	if nil != errs {
		return "", errs
	}
	log.Debugsf("Try RSA_Encrypt Successful!")
	rsaEncrypt := base64.StdEncoding.EncodeToString(baseDecode)
	log.Debugsf("Try rsaEncrypt Successful! rsaEncrypt = %s", rsaEncrypt)
	return rsaEncrypt, nil
}

// @Desc aes Decrypt
func AesDecrypt(aesEncrypt []byte, aeskey []byte) (aesDecrypt []byte, err error) {
	cryptoIns := aes.NewAES("PKCS7", "ECB", aeskey)
	cryptoBytes, err := cryptoIns.Decrypt(aesEncrypt)
	if nil != err {
		return nil, err
	}
	return (cryptoBytes), nil
}

// @Desc aes Encrypt
func AesEncrypt(originData []byte, aeskey []byte) (aesEncrypt []byte, err error) {
	cryptoIns := aes.NewAES("PKCS7", "ECB", aeskey)
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
