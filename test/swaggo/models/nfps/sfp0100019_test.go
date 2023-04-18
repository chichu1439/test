package models

import (
	"testing"

	"github.com/moov-io/iso8583"
)

func TestFP010019I(t *testing.T) {

	msg := iso8583.NewMessage(Spec0200RequestToPay)
	msg.MTI("0200")
	msg.Field(2, "1001000000000000000")
	msg.Field(3, "380000")
	msg.Field(4, "000000000001")
	msg.Field(7, "0928140000")
	msg.Field(11, "000001")
	msg.Field(12, "140000")
	msg.Field(13, "0928")
	msg.Field(15, "0928")
	msg.Field(18, "6011")
	msg.Field(22, "021")
	msg.Field(32, "006")
	msg.Field(33, "006")
	msg.Field(37, "127114100002")
	msg.Field(43, "BBLAT001B006B001P00130000101000101    TH")
	msg.Field(48, "00000000000000000000000000000000000000610010000000000000001                    zcc                                               0066800054336000                                                                                                                MSISDN      005                                                                                                                                            P                                                  R006200100000001005                                                                                                                             NATID                                                                                                                                                                                       0101202112010101")
	msg.Field(49, "764")
	msg.Field(52, "2aaddca6dfc48243")
	msg.Field(62, "000000000000")
	// msg.Field(102, "10010000000000000001")
	// msg.Field(103, "10010000000000000002")
	msgBytes, err := msg.Pack()
	if err != nil {
		t.Error(err)
	} else {
		t.Logf("%s\n", msgBytes)
	}
	jsonBytes, err := msg.MarshalJSON()
	if err != nil {
		t.Error(err)
	} else {
		t.Logf("%s\n", jsonBytes)
	}
	msgBytes2 := []byte(msgBytes)
	msg2 := iso8583.NewMessage(Spec0200RequestToPay)
	iso85830210Data := &ISO85830210Data{}
	if err = msg2.SetData(iso85830210Data); err != nil {
		t.Error(err)
	}
	if err = msg2.Unpack(msgBytes2); err != nil {
		t.Error(err)
	}

	t.Logf("%++v\n", iso85830210Data)
	//t.Logf("%s\n", iso85830210Data.F39.Value)

	//msgBytes3 := []byte("0230723844018A219004191001000000000000000380000000000000001092814000000000114000009286011021030060300212711400001300BBLAT001B006B001P00130000101000101    TH76600000000000000000000000000000000000000210010000000000000001                    zcc                                               0066800054336                                                                                                                   MSISDN                                                                                                                                                     P                                                                  0                                                                                                                                                                                                                                                                                            C 76446464646464646464646464646464646012000000000000")
	msgBytes3 := []byte("0210723A44018A2190041910010000000000000003800000000000000010928140000000001140000092809286011021030060300612711410000200BBLAT001B006B001P00130000101000101    TH81100000000000000000000000000000000000000610010000000000000001                    zcc                                               0066800054336000                                                                                                                MSISDN      005                                                                                                                                            P                                                  R006200100000001005                                                                                                                             NATID                                                                                                                                                                                       010120211201010176432616164646361366466633438323433012000000000000")
	msg3 := iso8583.NewMessage(Spec0200RequestToPay)
	iso85830210Data3 := &ISO85830210Data{}
	if err = msg3.SetData(iso85830210Data3); err != nil {
		t.Error(err)
	}
	if err = msg3.Unpack(msgBytes3); err != nil {
		t.Error(err)
	}
	msgJson3, _ := msg3.MarshalJSON()
	t.Logf("%s\n", msgJson3)
}
