package main

import (
	"encoding/base64"
	"fmt"
	"github.com/wumansgy/goEncrypt"
)

// 密钥 24字节
const secretKey = "hjdo843k920s6um4kk8sk3s8"

// 三重DES加密
func TripleDesEncrypt(encryptedString string) (string, error) {
	cryptText, err := goEncrypt.TripleDesEncrypt([]byte(encryptedString), []byte(secretKey), nil)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(cryptText), nil
}

// 三重DES解密
func TripleDesDecrypt(decryptString string) (string, error) {
	decryptBytes, err := base64.StdEncoding.DecodeString(decryptString)
	if err != nil {
		return "", err
	}
	cryptText, err := goEncrypt.TripleDesDecrypt(decryptBytes, []byte(secretKey), nil)
	if err != nil {
		return "", err
	}

	return string(cryptText), nil
}
func main() {
	tripleDesEncrypt, err := TripleDesEncrypt("1234567890123456")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(tripleDesEncrypt)
}
