package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
	"net/url"
)

const (
	key = "AFzdlyUswrRTXZ6UUgoRHBNsmslX42EnNKHfKVMs8K3JxwpLfx"
	iv  = "EnNKHfKVMs8K3JxwpLfx"
)

//var es = "t/M/0kAU15EgHLn38Ci4pkaq5LwdsIsvpVKhBhEaxEoP7nK+MLEMN1Weyy+pfOzk"
var es = "t%2FM%2F0kAU15EgHLn38Ci4pkaq5LwdsIsvpVKhBhEaxEoP7nK%2BMLEMN1Weyy%2BpfOzk"

//var es = "8fe8yYpxTdU8rwZ%2BWIo%2Fr1a0P6EkaUm4gn8WxfKJenx1eHvEVALE0CIo%2FHXPUA4S"
var err error

func main() {
	//str := "fjaskldfjlasjfdlasdflasdf"
	//es, _ = AesEncrypt(str, []byte(key[:32]))
	//fmt.Println(es)

	ds, err := AesDecrypt(es, []byte(key[:32]))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("ds:", string(ds))
	dss, err := CBCDecrypt(es, key[:32])
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("dss:", string(dss))

}

func AesEncrypt(encodeStr string, key []byte) (string, error) {
	encodeBytes := []byte(encodeStr)
	//根据key 生成密文
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	blockSize := block.BlockSize()
	encodeBytes = PKCS5Padding(encodeBytes, blockSize)

	blockMode := cipher.NewCBCEncrypter(block, []byte(iv[0:16]))
	crypted := make([]byte, len(encodeBytes))
	blockMode.CryptBlocks(crypted, encodeBytes)

	return base64.StdEncoding.EncodeToString(crypted), nil
}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	//填充
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)

	return append(ciphertext, padtext...)
}

func AesDecrypt(decodeStr string, key []byte) ([]byte, error) {
	queryUnescape, err := url.QueryUnescape(decodeStr)
	fmt.Println("queryUnescape:", queryUnescape)
	if err != nil {
		return nil, err
	}
	//先解密base64
	decodeBytes, err := base64.StdEncoding.DecodeString(queryUnescape)
	if err != nil {
		return nil, err
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockMode := cipher.NewCBCDecrypter(block, []byte(iv[:16]))
	origData := make([]byte, len(decodeBytes))

	blockMode.CryptBlocks(origData, decodeBytes)
	origData = PKCS5UnPadding(origData)
	return origData, nil
}

func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

// CBCDecrypt AES-CBC 解密
func CBCDecrypt(ciphertext string, key string) (string, error) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("cbc decrypt err:", err)
		}
	}()

	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}
	queryUnescape, err := url.QueryUnescape(ciphertext)
	if err != nil {
		return "", err
	}
	ciphercode, err := base64.StdEncoding.DecodeString(queryUnescape)
	if err != nil {
		return "", err
	}
	//ciphercode := []byte(ciphertext)
	iv := ciphercode[:aes.BlockSize]        // 密文的前 16 个字节为 iv
	ciphercode = ciphercode[aes.BlockSize:] // 正式密文

	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(ciphercode, ciphercode)

	plaintext := string(ciphercode) // ↓ 减去 padding
	return plaintext[:len(plaintext)-int(plaintext[len(plaintext)-1])], nil
}
