package main

//
//import (
//	"crypto/cipher"
//	"crypto/des"
//	"encoding/hex"
//	"fmt"
//)
//
//func main() {
//	var key = []byte(`4b22ae95fb3cfe7569ae12eff2d1f40a`)
//	//var org = "1234567890123456"
//	var org = "0000066800054336"
//	//dst := "98aac47d4265caa8"
//	//var src []byte
//	src, err := hex.DecodeString(org)
//	if err != nil {
//		return
//	}
//	fmt.Printf("src:%++v\n", string(src))
//	cipher, err := des.NewTripleDESCipher(key[:24])
//	if err != nil {
//		fmt.Printf("%++v\n", err)
//		return
//	}
//	en := encrypt(cipher, src)
//	_ = decrypt(cipher, en)
//	dst := make([]byte, len(src))
//	cipher.Encrypt(dst, src)
//	fmt.Printf("dst:%x\n", dst)
//	//"3fc7ef779be3af7c0000000000000000"
//	//encode, err := encoding.ASCIIToHex.Encode(dst)
//	//if err != nil {
//	//	fmt.Printf("%++v\n", err)
//	//	return
//	//}
//	//fmt.Printf("encode:%x\n", encode)
//
//	xx := make([]byte,hex.EncodedLen(len(dst))) // 计算转换为16进制后的长度
//	hex.Encode(xx,dst) // 转码
//	fmt.Printf("%s\n",xx) // 打印 616263313233，a的ascii码为97,97转换为16进制为61，后面的bc123同理
//
//}
//func encrypt(cipher cipher.Block, src []byte) []byte {
//	dst := make([]byte, len(src))
//	cipher.Encrypt(dst, src)
//	fmt.Printf("encrypt %x\n", dst)
//	return dst
//}
//func decrypt(cipher cipher.Block, src []byte) []byte {
//	dst := make([]byte, len(src))
//	cipher.Encrypt(dst, src)
//	fmt.Printf("decrypt %x\n", dst)
//	return dst
//}
