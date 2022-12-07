package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func main() {
	client := &http.Client{}
	url := "http://www.baidu.com"
	var r http.Request
	err := r.ParseForm()

	r.Form.Add("merchantId", "900000158")
	r.Form.Add("merchantApiId", "ktbapi")
	r.Form.Add("password", "test1234")
	r.Form.Add("memberId", "w.req.ContId") //required
	r.Form.Add("actionType", "w.req.ActionType")
	wl2fpra0I := strings.TrimSpace(r.Form.Encode())

	req, err := http.NewRequest("GET", url, bytes.NewBuffer([]byte(wl2fpra0I))) //GET大写
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%++v", req)
	rep, err := client.Do(req) //发起请求
	if err != nil {
		log.Fatal(err)
	}
	data, err := ioutil.ReadAll(rep.Body)
	rep.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%s", data)
}
