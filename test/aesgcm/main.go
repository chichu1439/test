// scrypt.go
package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"log"
	"time"

	"golang.org/x/crypto/scrypt"
)

func Encrypt(key, data []byte) ([]byte, error) {
	//key, salt, err := DeriveKey(key, nil)
	//if err != nil {
	//	return nil, err
	//}
	aesGcmKey, err := base64.StdEncoding.DecodeString("23YOpknd5HIy1CC+l2cQ45HSgrT8gQBRPoBNbXzErLo=")
	if err != nil {
		return nil, err
	}
	blockCipher, err := aes.NewCipher(aesGcmKey)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(blockCipher)
	if err != nil {
		return nil, err
	}
	nonce, err := base64.StdEncoding.DecodeString("l8DrqXtoOTRZup1G")
	if err != nil {
		return nil,err
	}
	//nonce := make([]byte, gcm.NonceSize())
	//if _, err = rand.Read(nonce); err != nil {
	//	return nil, err
	//}

	ciphertext := gcm.Seal(nonce, nonce, data, nil)

	//ciphertext = append(ciphertext, salt...)

	return ciphertext, nil
}

func Decrypt(key, data []byte) ([]byte, error) {
	salt, data := data[len(data)-32:], data[:len(data)-32]

	key, _, err := DeriveKey(key, salt)
	if err != nil {
		return nil, err
	}

	blockCipher, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(blockCipher)
	if err != nil {
		return nil, err
	}

	nonce, ciphertext := data[:gcm.NonceSize()], data[gcm.NonceSize():]

	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, err
	}

	return plaintext, nil
}

func DeriveKey(password, salt []byte) ([]byte, []byte, error) {
	saltstr := `fd3f63485dc0be26941a188b46768e90a747c1ed985b54e9c2294d9ad3aff5a8`
	start := time.Now()
	salt, _ = hex.DecodeString(saltstr)
	if salt == nil {
		salt = make([]byte, 32)
		if _, err := rand.Read(salt); err != nil {
			return nil, nil, err
		}
	}
	fmt.Printf("salt: %s\n", hex.EncodeToString(salt))

	key, err := scrypt.Key(password, salt, 1048576, 8, 1, 32)
	if err != nil {
		return nil, nil, err
	}
	fmt.Printf("key: %s\n", hex.EncodeToString(key))
	end := time.Now()
	cost := end.Sub(start)
	fmt.Printf("cost: %s\n", (cost))

	return key, salt, nil
}
func hash256(message string) string {
	start := time.Now()
	bytes2 := sha256.Sum256([]byte(message))   //计算哈希值，返回一个长度为32的数组
	hashcode2 := hex.EncodeToString(bytes2[:]) //将数组转换成切片，转换成16进制，返回字符串
	fmt.Printf("hashcode2: %s\n", hashcode2)
	end := time.Now()
	cost := end.Sub(start)
	fmt.Printf("hash256 cost: %s\n", (cost))
	return hashcode2
}
func main() {
	var (
		password = []byte("mysecretpassword")
		data     = []byte("")
	)

	ciphertext, err := Encrypt(password, data)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("ciphertext: %s\n", hex.EncodeToString(ciphertext))

	plaintext, err := Decrypt(password, ciphertext)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("hash256: %s\n", hash256("idno"))
	fmt.Printf("plaintext: %s\n", plaintext)
}
