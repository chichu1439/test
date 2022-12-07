package util

import (
	"encoding/base64"
	"testing"
)

var AESKey = `abcdefghijk01234`
var idNo = `11`
var id = `BK+kR55uTiB3/NOTPQASgA==`
var idx = ``

func TestName(t *testing.T) {
	//id = EncryptECB([]byte(idNo), []byte(AESKey))
	////if nil != err {
	////	t.Error(err)
	////}
	//t.Log(id)
	//
	//idx := DecryptECB(id, []byte(AESKey))
	////if nil != err {
	////	t.Error(err)
	////}
	//t.Log(idx)
	var err error
	id, err = aesEncrypt((AESKey), (idNo))
	if nil != err {
		t.Error(err)
	}
	t.Log(id)
	idx, err = aesDecrypt((AESKey), id)
	if nil != err {
		t.Error(err)
	}
	t.Log(idx)

	id, err = aesEncryptBaseUrl((AESKey), (idNo))
	if nil != err {
		t.Error(err)
	}
	t.Log(id)
	idx, err = aesDecryptBaseUrl((AESKey), id)
	if nil != err {
		t.Error(err)
	}
	t.Log(idx)
}

func aesEncrypt(aeskey string, toBeEncrypt string) (encrypt string, err error) {
	baseDecode, err := AesEncrypt([]byte(toBeEncrypt), []byte(aeskey))
	if nil != err {
		return "", err
	}
	rsaEncrypt := base64.StdEncoding.EncodeToString(baseDecode)
	if nil != err {
		return "", err
	}

	return rsaEncrypt, nil
}
func aesDecrypt(aeskey string, aesDncrypt string) (decrypt string, err error) {
	baseDecode, errs := base64.StdEncoding.DecodeString(aesDncrypt)
	if nil != errs {
		return "", errs
	}
	aesDecrypt, errs := AesDecrypt(baseDecode, []byte(aeskey))
	if nil != errs {
		return "", errs
	}

	return string(aesDecrypt), err
}

func aesEncryptBaseUrl(aeskey string, toBeEncrypt string) (encrypt string, err error) {
	baseDecode, err := AesEncrypt([]byte(toBeEncrypt), []byte(aeskey))
	if nil != err {
		return "", err
	}
	rsaEncrypt := base64.URLEncoding.EncodeToString(baseDecode)
	if nil != err {
		return "", err
	}

	return rsaEncrypt, nil
}
func aesDecryptBaseUrl(aeskey string, aesDncrypt string) (decrypt string, err error) {
	baseDecode, errs := base64.URLEncoding.DecodeString(aesDncrypt)
	if nil != errs {
		return "", errs
	}
	aesDecrypt, errs := AesDecrypt(baseDecode, []byte(aeskey))
	if nil != errs {
		return "", errs
	}

	return string(aesDecrypt), err
}
