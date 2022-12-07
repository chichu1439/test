package aes

import "testing"

//var aesKey = "23YOpknd5HIy1CC+l2cQ45HSgrT8gQBRPoBNbXzErLo="
//
var aesKey = "kMEf8JqzuGGSkCbErnatjWSSwSodEpuhnv35lZ75jzQ="

//var iv = "l8DrqXtoOTRZup1G"

var iv = "MTQ2MzdhMzAtZDYwNC00Yjg1LTg2MjQtYzhhYTkzNzQ4NzM2"
//var iv = "OTM4MkFF"

func TestInitAesGcm(t *testing.T) {
	var a = AesGcm{
		padding:        "",
		mode:           "",
		key:            nil,
		ivOrNonce:      nil,
		isFixIvOrNonce: false,
	}
	gcm, i, err := a.InitAesGcm(aesKey, iv)
	if err != nil {
		t.Errorf("InitAesGcm err %++v", err)
		return
	}
	encrypt, err := a.Encrypt(gcm, i, "111111")
	//encrypt, err := a.Encrypt(gcm, i, "This is a plain text which need to be encrypted by Java AES 256 GCM Encryption Algorithm")
	if err != nil {
		t.Errorf("Encrypt err %++v", err)
		return
	}
	t.Logf("encrypt: %++v", encrypt)
	decrypt, err := a.Decrypt(gcm, i, encrypt)
	if err != nil {
		t.Errorf("Decrypt err %++v", err)
		return
	}
	t.Logf("decrypt: %++v", decrypt)

}
