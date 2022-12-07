/*
 * Copyright (c) 2020. Shenzhen Forms Syntron Information Co., Ltd. All rights reserved.
 * Platform universe is a distributed platform of agile operation, which has an open business model.
 * Truly distributed Architecture capability with latest techs such as Event Mesh, Containerized Micro services.
 * Overhead enables light-weight deployment and Edge Computing Mode.
 * Centric with BCP and cost efficiency at the heart of the design principle.
 * With model driven operation and development that further enhanced the overall agility.
 */

package encry

import (
	"fmt"
	"testing"
)

func TestEncryDecryData(t *testing.T) {
	ServiceConf.EncryDlsKeyFlag = true
	ServiceConf.DecryPubKeyFlag = true
	KmsCache = &KmsKey{
		KeyValue:    "YjVkZjUzMDdiYmIwMjZjYTI2MDhkMjBh",
	}
	var srcData = "abc道可道_e32#@!"
	var encData, decData string

	KmsCache.EncryType = "CBC"
	fmt.Println("Src=", srcData)
	encData = KmsEncrypt(srcData)
	fmt.Println("Enc=", encData)
	decData = KmsDecrypt(encData)
	fmt.Println("Dec=", decData)
	if decData != srcData || srcData == encData {
		t.Errorf("Encry Type[%s] Data[%s] Faild! encData[%s] decData[%s]", KmsCache.EncryType, srcData, encData, decData)
	}

	KmsCache.EncryType = "ECB"
	fmt.Println("Src=", srcData)
	encData = KmsEncrypt(srcData)
	fmt.Println("Enc=", encData)
	decData = KmsDecrypt(encData)
	fmt.Println("Dec=", decData)
	if decData != srcData || srcData == encData {
		t.Errorf("Encry Type[%s] Data[%s] Faild! encData[%s] decData[%s]", KmsCache.EncryType, srcData, encData, decData)
	}

	KmsCache.EncryType = "CFB"
	fmt.Println("Src=", srcData)
	encData = KmsEncrypt(srcData)
	fmt.Println("Enc=", encData)
	decData = KmsDecrypt(encData)
	fmt.Println("Dec=", decData)
	if decData != srcData || srcData == encData {
		t.Errorf("Encry Type[%s] Data[%s] Faild! encData[%s] decData[%s]", KmsCache.EncryType, srcData, encData, decData)
	}

	KmsCache.EncryType = "GCM"
	fmt.Println("Src=", srcData)
	encData = KmsEncrypt(srcData)
	fmt.Println("Enc=", encData)
	decData = KmsDecrypt(encData)
	fmt.Println("Dec=", decData)
	if decData != srcData || srcData == encData {
		t.Errorf("Encry Type[%s] Data[%s] Faild! encData[%s] decData[%s]", KmsCache.EncryType, srcData, encData, decData)
	}
}
