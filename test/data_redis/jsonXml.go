package main
//
//import (
//	"encoding/xml"
//	"fmt"
//	"git.sirius.io/banking/nfps/common/models"
//)
//
//func main() {
//	camt052 := models.FPSAdmi004{
//		XMLName:  xml.Name{},
//		NumOfMsg: "1",
//		XMLHead: models.Head001{
//			FromMemberId:   "1",
//			ToMemberId:     "1",
//			BizMsgId:       "1",
//			MsgDefId:       "1",
//			BizSvc:         "1",
//			CreateDatetime: "1",
//		},
//		Document: models.Admi004{
//			EvtCd:    "2",
//			EvtParam: "2",
//			EvtDesc:  "2",
//			EvtTm:    "2",
//		},
//	}
//	bytes, err := xml.Marshal(camt052)
//	if err != nil {
//		fmt.Println(err)
//	}
//	fmt.Println(string(bytes))
//}
