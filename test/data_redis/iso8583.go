package main

import (
	"fmt"

	"github.com/ideazxy/iso8583"
)

type Data struct {
	No   *iso8583.Numeric      `field:"3" length:"6" encode:"bcd"`    // bcd value encoding
	Oper *iso8583.Numeric      `field:"26" length:"2" encode:"ascii"` // ascii value encoding
	Ret  *iso8583.Alphanumeric `field:"39" length:"2"`
	Sn   *iso8583.Llvar        `field:"45" length:"23" encode:"bcd,ascii"` // bcd length encoding, ascii value encoding
	Info *iso8583.Lllvar       `field:"46" length:"42" encode:"bcd,ascii"`
	Mac  *iso8583.Binary       `field:"64" length:"8"`
}

func main() {
	data := &Data{
		No:   iso8583.NewNumeric("001111"),
		Oper: iso8583.NewNumeric("22"),
		Ret:  iso8583.NewAlphanumeric("ok"),
		Sn:   iso8583.NewLlvar([]byte("abc001")),
		Info: iso8583.NewLllvar([]byte("你好 golang!")),
		Mac:  iso8583.NewBinary([]byte("a1s2d3f4")),
	}
	msg := iso8583.Message{
		Mti:          "0110",
		MtiEncode:    iso8583.ASCII,
		SecondBitmap: false,
		Data:         data,
	}
	b, err := msg.Bytes()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("% x\n", b)

	////------
	data2 := &Data{
		No:   iso8583.NewNumeric(""),
		Oper: iso8583.NewNumeric(""),
		Ret:  iso8583.NewAlphanumeric(""),
		Sn:   iso8583.NewLlvar(nil),
		Info: iso8583.NewLllvar(nil),
		Mac:  iso8583.NewBinary(nil),
	}
	iso := iso8583.Message{
		Mti:          "",
		MtiEncode:    iso8583.ASCII,
		SecondBitmap: true,
		Data:         data2,
	}
	err = iso.Load(b)
	if err != nil {
		fmt.Println(err.Error())
	}
	result := iso.Data.(*Data)
	fmt.Println("result=", result.Mac)
}

//func newDataIso() *Data {
//	return &Data{
//		No:   iso8583.NewNumeric(""),
//		Oper: iso8583.NewNumeric(""),
//		Ret:  iso8583.NewAlphanumeric(""),
//		Sn:   iso8583.NewLlvar(nil),
//		Info: iso8583.NewLllvar(nil),
//		Mac:  iso8583.NewBinary(nil),
//	}
//}

type name interface {
}
