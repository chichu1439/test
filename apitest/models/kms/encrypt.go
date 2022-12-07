package models

import (
	"math/rand"
	"os"
	"time"
)

func FileLoad(filepath string) ([]byte, error) {

	privatefile, err := os.Open(filepath)
	defer privatefile.Close()

	if err != nil {
		return nil, err
	}

	privateKey := make([]byte, 20480)
	number, err := privatefile.Read(privateKey)

	if nil != err {
		return nil, err
	}

	return privateKey[:number], nil
}

//Krand . 随机字符串
func Krand(size int, kind int) []byte {
	ikind, kinds, result := kind, [][]int{{10, 48}, {26, 97}, {26, 65}}, make([]byte, size)
	isAll := kind > 2 || kind < 0
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		if isAll { // random ikind
			ikind = rand.Intn(3)
		}
		scope, base := kinds[ikind][0], kinds[ikind][1]
		result[i] = uint8(base + rand.Intn(scope))
	}
	return result
}
