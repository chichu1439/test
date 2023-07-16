package main

// import (
// 	"bytes"
// 	"crypto/tls"
// 	"encoding/json"
// 	"flag"
// 	"fmt"
// 	"git.sirius.io/banking/test/dao"
// 	"git.sirius.io/banking/test/model"
// 	_ "github.com/go-sql-driver/mysql"
// 	"github.com/xormplus/xorm"
// 	"io/ioutil"
// 	"net/http"
// 	"time"
// )

// //var Db map[string]config.Db
// var O *xorm.Engine
// var serveraddr string
// var addr string
// var user string
// var pwd string
// var database string

// var destaddr string

// var maxIdleConns int
// var maxOpenConns int
// var defaultQueryLimit int
// var maxLimitValue int
// var maxLifeValue int

// func main() {

// 	flag.StringVar(&serveraddr, "serveraddr", "127.0.0.1:8081", "server Url")
// 	flag.StringVar(&destaddr, "destaddr", "127.0.0.1:8082", "destination Url")
// 	//flag.StringVar(&destaddr, "destaddr", "https://10.20.5.28:28089/uploadTransaction", "destination Url")
// 	flag.StringVar(&addr, "address", "127.0.0.1:3306", " Url")
// 	flag.StringVar(&user, "user", "root", "user")
// 	flag.StringVar(&pwd, "password", "root", "Password")
// 	flag.StringVar(&database, "database", "payment", "database")
// 	flag.IntVar(&maxIdleConns, "maxIdleConns", 50, "maxIdleConns")
// 	flag.IntVar(&maxOpenConns, "maxOpenConns", 50, "maxOpenConns")
// 	flag.IntVar(&defaultQueryLimit, "defaultQueryLimit", 30, "defaultQueryLimit")
// 	flag.IntVar(&maxLimitValue, "maxLimitValue", 50000, "maxLimitValue")
// 	flag.IntVar(&maxLifeValue, "maxLifeValue", 540, "maxLifeValue")

// 	flag.Parse()
// 	//
// 	dsn := user + ":" + pwd + "@tcp(" + addr + ")/" + database

// 	engine, err := xorm.NewEngine("mysql", dsn)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	engine.SetMaxIdleConns(maxIdleConns)
// 	engine.SetMaxOpenConns(maxOpenConns)
// 	engine.SetConnMaxLifetime(time.Duration(maxLifeValue) * time.Second)
// 	engine.ShowSQL(true)
// 	O = engine
// 	http.HandleFunc("/", sendBankData)
// 	server := &http.Server{
// 		Addr: serveraddr,
// 	}
// 	err = server.ListenAndServe()
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// }

// type ReqData struct {
// 	Date string
// }

// func sendBankData(w http.ResponseWriter, req *http.Request) {
// 	var err error
// 	defer func() {
// 		if err != nil {
// 			w.WriteHeader(400)
// 			w.Write([]byte(err.Error()))
// 		}
// 	}()
// 	reqBody, err := ioutil.ReadAll((req.Body))
// 	var reqData ReqData

// 	if err := json.Unmarshal([]byte(reqBody), &reqData); err != nil {
// 		fmt.Println(err)
// 	}
// 	settlementFlows := make([]dao.TSlSettlementFlow, 0)
// 	err = O.Table("t_mock_sl_settlement_flow").Find(&settlementFlows, dao.TSlSettlementFlow{
// 		StlmtDate: reqData.Date,
// 	})
// 	transactions := make([]models.BankTransaction, 0)
// 	for _, info := range settlementFlows {
// 		TransDirection := "D"
// 		if info.TransDirection == "CR" {
// 			TransDirection = "C"
// 		}
// 		transaction := models.BankTransaction{
// 			TransactionId:     info.TransId,
// 			MessageId:         info.TransId,
// 			Currency:          info.Currency,
// 			TransactionAmount: info.TransAmount,
// 			DateTime:          info.StlmtDate + " " + info.StlmtTime,
// 			RecordType:        TransDirection,
// 			TransactionType:   info.TransType,
// 			BankId:            info.TransPart,
// 		}
// 		transactions = append(transactions, transaction)
// 	}
// 	data := &models.BankTransactionData{
// 		Transaction: transactions,
// 	}
// 	if len(data.Transaction) > 0 {
// 		fmt.Printf("data [%++v]\n", data)
// 		jsonStr, err := json.Marshal(data)
// 		if err != nil {
// 			return
// 		}
// 		newRequest, err := http.NewRequest("POST", destaddr, bytes.NewBuffer(jsonStr))
// 		if err != nil {
// 			fmt.Println(err)
// 		}
// 		newRequest.Header.Set("Content-Type", "application/json")
// 		client := &http.Client{Transport: &http.Transport{
// 			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
// 		}}
// 		newResponse, err := client.Do(newRequest)
// 		if err != nil {
// 			fmt.Println(err)
// 		}
// 		defer newResponse.Body.Close()
// 		fmt.Println("response Status:", newResponse.Status)
// 		fmt.Println("response Headers:", newResponse.Header)
// 		body, _ := ioutil.ReadAll(newResponse.Body)
// 		fmt.Println("response Body:", string(body))

// 		w.WriteHeader(200)
// 		w.Write([]byte("success"))
// 	} else {
// 		w.WriteHeader(201)
// 		w.Write([]byte("no data"))
// 	}
// }
