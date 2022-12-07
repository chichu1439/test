package main

import (
	"log"

	"github.com/dongri/emv-qrcode/emv/mpm"
)

func main() {
	var err error
	//// MPM Encode
	emvqr := new(mpm.EMVQR)
	//emvqr.SetPayloadFormatIndicator("01")
	//emvqr.SetPointOfInitiationMethod("12") // 11 is static qrcode
	//merchantAccountInformationJCB := new(mpm.MerchantAccountInformation)
	//merchantAccountInformationJCB.SetGloballyUniqueIdentifier("D123456")
	//merchantAccountInformationJCB.AddPaymentNetworkSpecific("13", "JCB1234567890")
	//emvqr.AddMerchantAccountInformation(mpm.ID("29"), merchantAccountInformationJCB)
	//
	//merchantAccountInformationMaster := new(mpm.MerchantAccountInformation)
	//merchantAccountInformationMaster.SetGloballyUniqueIdentifier("M123456")
	//merchantAccountInformationMaster.AddPaymentNetworkSpecific("04", "MASTER1234567890")
	//emvqr.AddMerchantAccountInformation(mpm.ID("31"), merchantAccountInformationMaster)
	//
	//emvqr.SetMerchantCategoryCode("5311")
	//emvqr.SetTransactionCurrency("392")
	//emvqr.SetTransactionAmount("999.123")
	//emvqr.SetCountryCode("JP")
	//emvqr.SetMerchantName("DONGRI")
	//emvqr.SetMerchantCity("TOKYO")
	//additionalTemplate := new(mpm.AdditionalDataFieldTemplate)
	//additionalTemplate.SetBillNumber("hoge")
	//additionalTemplate.SetReferenceLabel("fuga")
	//additionalTemplate.SetTerminalLabel("piyo")
	//emvqr.SetAdditionalDataFieldTemplate(additionalTemplate)
	//code, err := mpm.Encode(emvqr)
	//if err != nil {
	//	log.Println(err.Error())
	//	return
	//}
	//log.Println(code) // 0002010102121313JCB12345678900416MASTER12345678905204531153033925407999.1235802JP5906DONGRI6005TOKYO62240104hoge0504fuga0704piyo6304C343

	// MPM Decode
	//emvqr, err = mpm.Decode("00020101021229300012D156000000000510A93FO3230Q31280012D15600000001030812345678520441115802CN5914BEST TRANSPORT6007BEIJING64200002ZH0104最佳运输0202北京540523.7253031565502016233030412340603***0708A60086670902ME91320016A0112233449988770708123456786304A13A")
	emvqr, err = mpm.Decode("00020101021202164225631530009888")
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(emvqr)

	//// Print Raw Data
	//raw := emvqr.RawData()
	//log.Println("\n" + raw)
	//
	//// Print Binary Data
	//binary := emvqr.BinaryData()
	//log.Println("\n" + binary)

	// Print JSON
	json := emvqr.JSON()
	log.Println(json)

}
