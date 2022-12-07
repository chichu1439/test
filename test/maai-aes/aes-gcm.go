package aes

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"git.multiverse.io/framework/common/crypto/padding"
)

type Interface interface {
	InitAesGcm(key, iv string) (cipher.AEAD, []byte, error)
	Encrypt(gcm cipher.AEAD, nonce []byte, data string) (string, error)
	Decrypt(gcm cipher.AEAD, nonce []byte, data string) (string, error)
}

type AesGcm struct {
	padding        string
	mode           string
	key            []byte
	ivOrNonce      []byte
	isFixIvOrNonce bool
	blockCipher    cipher.Block
}

func (a *AesGcm) InitAesGcm(key, iv string) (cipher.AEAD, []byte, error) {
	aesGcmKey, err := base64.StdEncoding.DecodeString(key)
	if err != nil {
		return nil, nil, err
	}
	blockCipher, err := aes.NewCipher(aesGcmKey)
	if err != nil {
		return nil, nil, err
	}
	a.blockCipher = blockCipher
	nonce, err := base64.StdEncoding.DecodeString(iv)
	if err != nil {
		return nil, nil, err
	}
	gcm, err := cipher.NewGCMWithNonceSize(blockCipher, len(nonce))
	if err != nil {
		return nil, nil, err
	}
	return gcm, nonce, nil
}

func (a *AesGcm) Encrypt(gcm cipher.AEAD, nonce []byte, data string) (string, error) {
	var plainText []byte
	plainText, err := padding.Padding([]byte(data), a.blockCipher.BlockSize(), a.padding)
	if nil != err {
		return "", err
	}
	ciphertext := gcm.Seal(nil, nonce, plainText, nil)

	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

func (a *AesGcm) Decrypt(gcm cipher.AEAD, nonce []byte, data string) (string, error) {

	ciphertext, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return "", err
	}
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", err
	}
	plaintext = padding.Unpadding(plaintext, a.padding)
	return string(plaintext), nil
}
