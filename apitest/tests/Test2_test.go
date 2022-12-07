package test

import (
	"fmt"
	"testing"
	"time"
)
import "github.com/360EntSecGroup-Skylar/excelize"

func Test_test2(t *testing.T) {
	mapFile := make(map[string][]string, 0)
	fmt.Println(len(mapFile["1"]))
	fmt.Println(mapFile["1"])
}

func Test_test1(t *testing.T) {
	Sheet := "New Loan Report"
	//创建excel文件
	xlsx := excelize.NewFile()
	//创建新表单
	xlsx.SetSheetName("Sheet1", Sheet)
	index := xlsx.GetSheetIndex(Sheet)
	//写入数据

	data := map[string]string{
		"A1":  "Seq No.",
		"B1":  "CIF",
		"C1":  "Customer Name",
		"D1":  "Loan Reference Number",
		"E1":  "Loan Status",
		"F1":  "ID Type",
		"G1":  "Certificate ID",
		"H1":  "Customer Phone Number",
		"I1":  "Product ID",
		"J1":  "Product Name",
		"K1":  "Drawdown Date",
		"L1":  "Maturity Date",
		"M1":  "Amount",
		"N1":  "Interest Rate (%)",
		"O1":  "Fee",
		"P1":  "Tax Fees",
		"Q1":  "Repayment Date",
		"R1":  "Repayment Method",
		"S1":  "Currency",
		"T1":  "Tenor",
		"U1":  "Repayment Frequency",
		"V1":  "Disbursement Channel",
		"W1":  "Bank Name",
		"X1":  "Bank Account Name",
		"Y1":  "Bank Account Number",
		"Z1":  "Guarantee Method",
		"AA1": "Arrangement Purpose",
	}
	for k, v := range data {
		//设置单元格的值
		xlsx.SetCellValue(Sheet, k, v)
	}

	xlsx.InsertRow(Sheet, 0)
	xlsx.InsertRow(Sheet, 0)
	xlsx.InsertRow(Sheet, 0)
	xlsx.InsertRow(Sheet, 0)
	xlsx.MergeCell(Sheet, "A1", "AA1")
	xlsx.SetCellValue(Sheet, "M1", "Internal Success disbursement report")
	xlsx.MergeCell(Sheet, "A2", "AA2")
	xlsx.SetCellValue(Sheet, "A2", fmt.Sprintf("Report date:%s", time.Now().Format("2006-01-02")))
	xlsx.MergeCell(Sheet, "A3", "AA3")
	xlsx.SetCellValue(Sheet, "A3", fmt.Sprintf("Report frequency:%s", "daily"))
	xlsx.MergeCell(Sheet, "A4", "AA4")
	xlsx.SetCellValue(Sheet, "A4", fmt.Sprintf("Statistics date:%s", time.Now().Format("2006-01-02")))
	styleTitle, _ := xlsx.NewStyle(`{"font":{"bold":true},"alignment":{"horizontal": "center"}}`)
	styleRow, _ := xlsx.NewStyle(`{"alignment":{"horizontal": "left"}}`)
	xlsx.SetCellStyle(Sheet, "A1", "AA1", styleTitle)
	xlsx.SetCellStyle(Sheet, "A2", "AA2", styleRow)
	xlsx.SetCellStyle(Sheet, "A3", "AA3", styleRow)
	xlsx.SetCellStyle(Sheet, "A4", "AA4", styleRow)

	//设置默认打开的表单
	xlsx.SetActiveSheet(index)

	//保存文件到指定路径
	xlsx.SaveAs("./daily new loan report1.xlsx")
}

func hello(pipline chan string)  {
	fmt.Println(<-pipline)
}

func TestAbc(t *testing.T) {
	//pipline := make(chan string,1)
	////go hello(pipline)
	//pipline <- "hello world"
	//fmt.Println(<-pipline)

	//pipline := make(chan string)
	//go func() {
	//	pipline<-"hello world"
	//	pipline<-"hello China"
	//	close(pipline)
	//}()
	//for i := range pipline {
	//	fmt.Println(i)
	//}
	quit = make(chan int)
	for a:=0;a<10;a++ {
		go foo(a)
	}

	for i:=0;i<10;i++{
		fmt.Println(i,"-end-",<-quit)
	}
}

var quit chan int
func foo(id int){
	fmt.Println(id)
	quit<-0
}