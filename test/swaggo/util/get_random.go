package util

import (
	"math/rand"
	"time"
)

// generate message id

func GetMsgID(ClsCode string) string {
	timeNow := time.Now()
	MsgID := "M" + ClsCode + timeNow.Format("20060102") + GetRandomString(8) + timeNow.Format("150405") + GetRandomString(8)
	return MsgID
}

// Randomly generate a combination of upper case letters and numbers with specified digits

func GetRandomString(l int) string {
	str := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}
