package util

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"io"
	"os"
)

var kmsCache *KmsKey

//加密
func AesGcmEncrypt(plaintext []byte, key []byte) ([]byte, error) {
	c, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}
	//nonce := []byte{98, 72, 98, 52, 54, 49, 116, 50, 76, 119, 79, 121}
	return gcm.Seal(nonce, nonce, plaintext, nil), nil
}

//解密
func AesGcmDecrypt(ciphertext []byte, key []byte) ([]byte, error) {
	c, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return nil, err
	}

	nonceSize := gcm.NonceSize()
	if len(ciphertext) < nonceSize {
		return nil, errors.New("ciphertext too short")
	}

	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
	return gcm.Open(nil, nonce, ciphertext, nil)
}

//通过kms加密
func KmsEncrypt(plaintext string, kms KmsKey) (ciphertext string, err error) {
	kmsCache = &kms
	if err != nil {
		return ciphertext, err
	}
	key := kmsCache.KeyValue
	bytes, err := AesGcmEncrypt([]byte(plaintext), []byte(key))
	encoded := base64.StdEncoding.EncodeToString(bytes)
	return encoded, err
}

//通过kms解密
func KmsDecrypt(ciphertext string, kms KmsKey) (plaintext string, err error) {
	decoded, err := base64.StdEncoding.DecodeString(ciphertext)
	kmsCache = &kms
	if err != nil {
		return plaintext, err
	}
	key := kmsCache.KeyValue
	bytes, err := AesGcmDecrypt([]byte(decoded), []byte(key))
	return string(bytes), err
}

func FileLoad(filepath string) ([]byte, error) {

	privatefile, err := os.Open(filepath)
	defer privatefile.Close()

	if err != nil {
		return nil, err
	}

	privateKey := make([]byte, 20480)
	num, err := privatefile.Read(privateKey)

	if nil != err {
		return nil, err
	}

	return privateKey[:num], nil
}

func Hash256(message string) string {
	bytes2 := sha256.Sum256([]byte(message)) //calculate the hash value ，returns an array of length 32
	return hex.EncodeToString(bytes2[:])     //convert the array into slices, convert into hexadecimal, and return a string
}
