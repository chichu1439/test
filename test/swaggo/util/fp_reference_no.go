package util

import (
	"strings"
	"sync"
	"time"
)

var flagNo int64 = 0
var lock sync.Mutex

const SEQMAX = 99999999

func GetClrSysRef() string {
	timeNow := time.Now()
	TranNo := "FRN" + FormatDate(timeNow, "yyyyMMdd") + GetRandomString(9) + FormatDate(timeNow, "HHmmss") + GetRandomString(9)
	return TranNo
}

//生成TXID
func GetTxId() string {
	timeNow := time.Now()
	TxId := "F" + FormatDate(timeNow, "yyyyMMdd") + GetRandomString(9)
	return TxId
}

//日期转字符串
func FormatDate(date time.Time, dateStyle DateStyle) string {
	layout := string(dateStyle)
	layout = strings.Replace(layout, "yyyy", "2006", 1)
	layout = strings.Replace(layout, "yy", "06", 1)
	layout = strings.Replace(layout, "MM", "01", 1)
	layout = strings.Replace(layout, "dd", "02", 1)
	layout = strings.Replace(layout, "HH", "15", 1)
	layout = strings.Replace(layout, "mm", "04", 1)
	layout = strings.Replace(layout, "ss", "05", 1)
	layout = strings.Replace(layout, "SSSS", "0000", -1)

	return date.Format(layout)
}

//日期格式：模仿java中的结构体
type DateStyle string

/*
	比较日期时间
	日期格式: yyyy-MM-dd
	time1 == time2 return 0
	time1 < time2 return -1
	time1 > time2 return 1
*/
func CompareDate(time1, time2 string) (int, error) {
	if time1 == time2 {
		return 0, nil
	}
	time1 = time1 + " 00:00:00"
	time2 = time2 + " 00:00:00"
	//先把时间字符串格式化成相同的时间类型
	t1, err := time.Parse("2006-01-02 15:04:05", time1)
	if nil != err {
		return 3, err
	}
	t2, err := time.Parse("2006-01-02 15:04:05", time2)
	if nil != err {
		return 3, err
	}
	if t1.Before(t2) {
		return -1, nil
	} else {
		return 1, nil
	}
}

// func FormatAddressingKey(AccountKeyType, AccountKeyNum string) string {
// 	return fmt.Sprintf("%s.%s.%s.%s", constant.FPSMemberId, constant.RedisPrefixAddressing, AccountKeyType, AccountKeyNum)
// }
// func FormatAddressingValue(BankCode, AccountNum string) string {
// 	return fmt.Sprintf("%s.%s", BankCode, AccountNum)
// }
// func SplitAddressingValue(addressingValue string) (BankCode, AccountNum string) {
// 	strArray := strings.Split(addressingValue, ".")
// 	if len(strArray) == 2 {
// 		return strArray[0], strArray[1]
// 	}
// 	return
// }
// func FormatNettingParamKey() string {
// 	return fmt.Sprintf("%s.%s", constant.FPSMemberId, constant.RedisPrefixNetId)
// }

// func FormatNettingParamValue(nettingId, nettingEnableStatus string) string {
// 	return fmt.Sprintf("%s.%s", nettingId, nettingEnableStatus)
// }
// func SplitNettingParamValue(nettingParamValue string) (nettingId, nettingEnableStatus string) {
// 	strArray := strings.Split(nettingParamValue, ".")
// 	if len(strArray) == 2 {
// 		return strArray[0], strArray[1]
// 	}
// 	return
// }
