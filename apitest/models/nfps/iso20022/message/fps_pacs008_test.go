package message

import (
	"encoding/xml"
	"testing"
	"time"

	"apitest/models"
)

func case1() {
	pacs008 := &FpsPacs008{}
	err := xml.Unmarshal([]byte(pacs008Str), pacs008)
	if err != nil {
		return
	}
}

func case2() {
	pacs008 := &models.FPSPacs008{}
	err := xml.Unmarshal([]byte(pacs008Str), pacs008)
	if err != nil {
		return
	}
}

func BenchmarkCase1(b *testing.B) {
	for n := 0; n < b.N; n++ {
		case1()
	}
}

func BenchmarkCase2(b *testing.B) {
	for n := 0; n < b.N; n++ {
		case2()
	}
}

func TestUnmarshal(t *testing.T) {

	startTime := time.Now()

	counter := 10000
	for i := 0; i < counter; i++ {
		func() {
			pacs008 := &FpsPacs008{}
			err := xml.Unmarshal([]byte(pacs008Str), pacs008)
			if err != nil {
				t.Error(err)
			}
		}()
	}

	t.Logf("case1 time: %f s", time.Now().Sub(startTime).Seconds())

	startTime2 := time.Now()

	for i := 0; i < counter; i++ {
		func() {
			pacs008 := &models.FPSPacs008{}
			err := xml.Unmarshal([]byte(pacs008Str), pacs008)
			if err != nil {
				t.Error(err)
			}
		}()
	}
	t.Logf("case2 time: %f s", time.Now().Sub(startTime2).Seconds())
	//pacs008 := &FpsPacs008{}
	//err := xml.Unmarshal([]byte(pacs008Str), pacs008)
	//if err != nil {
	//	t.Error(err)
	//} else {
	//	t.Logf("%++v\n", pacs008)
	//}
	//
	//buf, err := xml.MarshalIndent(pacs008, "", "\t")
	//if err != nil {
	//	t.Error(err)
	//} else {
	//	t.Logf("%s\n", buf)
	//}
	//
	//if err := pacs008.Validate(); err != nil {
	//	t.Error(err)
	//} else {
	//	t.Logf("validate pacs008 success")
	//}

}
