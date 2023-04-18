package models

import (
	"encoding/xml"
	"testing"
	"time"
)

func TestFPSpacs008(t *testing.T) {

	pacs008 := &FPSPacs008{
		NumOfMsg: "1",
		XMLHead: Head001{
			FromMemberId:   "123",
			ToMemberId:     "555",
			BizMsgId:       "111",
			MsgDefId:       "pacs.008",
			BizSvc:         "xxx",
			CreateDatetime: time.Now().Format(time.RFC3339),
		},
	}

	pacs008Body, err := xml.Marshal(pacs008)
	if err != nil {
		t.Error(err)
	} else {
		t.Logf("%s", pacs008Body)
	}

}
