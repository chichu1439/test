package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"github.com/google/uuid"
	"strings"
)

func GenerateHAMCsha256(content, secret string) string {
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(content))
	//sha := hex.EncodeToString(h.Sum(nil))
	//fmt.Println(sha)
	return base64.StdEncoding.EncodeToString([]byte(h.Sum(nil)))
}

type name struct {
	CU, SC string
}

func main() {
	h := name{
		//CU: "dkDsmDHT6rQxx22tqYlyUL0SlmJe5FQX",
		//SC: "aee9bUGE5ArZXjV9Ts-lD7UwW6NNXSC8",
		//CU: "xNfHzf3WtmGZKFdl2ve7ZVbcKq87bohs",
		//SC: "EGJ2X1bwIcZMeQB0vEk8Cpymb7AV39K6",
		CU: "TKhgyWdc424cocraTVBb48RcyR1S8Ym2",
		SC: "Wh6mMmJ2N6PGyTuU5dqE8EpggOxqDIkO",
	}
	nonce := uuid.New()
	nonceStr := strings.ReplaceAll(nonce.String(), "-", "")
	//nonceStr := "ypGrBEKXNNbpN4WMit8yx7PtYAtAe9CC"
	signature := fmt.Sprintf("%s%s", h.CU, nonceStr)
	header := make(map[string]string)
	header["XSignature"] = fmt.Sprintf("%s %s %s", h.CU, nonceStr, GenerateHAMCsha256(signature, h.SC))
	fmt.Println(header["XSignature"])
}
