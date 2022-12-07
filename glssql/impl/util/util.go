package util

import (
	"bufio"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"os"
	"strings"
	"time"
)

func InputKeyValue(echo string) []string {
	result := make([]string, 0)
	reader := bufio.NewReader(os.Stdin)
	for true {
		fmt.Println(echo)
		s, e := reader.ReadString('\n')
		s = strings.Trim(s, "\n")
		if e != nil || strings.ToLower(strings.Trim(s, " ")) == ":q" {
			return result
		}
		result = append(result, strings.Trim(s, " "))
	}
	return result
}


type InvitationCode struct {
	Token   string
	TokenExpire time.Time
}

func getOneTimeSalt() string {
	buf := make([]byte, 20)
	_, err := rand.Read(buf)
	if err != nil {
		fmt.Printf("getOneTimeSalt invoke rand.Read failed %s", err.Error())
	}
	return string(buf)
}

func (ic *InvitationCode) Generate(expire int) {
	salt := getOneTimeSalt()
	uuid := uuid.NewV4().String()
	token := fmt.Sprintf("%x", sha256.Sum256([]byte(uuid+salt)))
	loc, _ := time.LoadLocation("Local")
	now := time.Now().In(loc)
	if expire <= 0 {
		expire = 600
	}
	expireTime := now.Add(time.Duration(expire) * time.Second)
	ic.Token = token
	ic.TokenExpire = expireTime
	fmt.Println(fmt.Sprintf("%v, %v", ic.Token, ic.TokenExpire))
}
