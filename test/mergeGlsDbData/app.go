package main

import (
	"bufio"

	//"dlsapp/dlssql/impl/hash"
	"flag"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"io"
	"os"
	"strconv"
	"strings"
)

type DlsTreeSlice struct {
	//DbHashring *hash.HashRing
	DbMaps   map[string]int
	DbSlices map[string]DlsSliceDatabase
}
type DlsSliceElement struct {
	Id          int
	Class       string
	Name        string
	Db          string
	Tbl         string
	Parent      int
	VirtualType int
	SaltValue   string
	Top         string
	Bottom      string
	Pos         int
	Slice       int
	Weight_rule string
	Effect      int
}
type DlsSliceDatabase struct {
	//TblHashring *hash.HashRing
	DbSlice DlsSliceElement
	TblMaps map[string]int
	//Children    *DlsTreeSlice
}

func main() {
	// main args
	var sql string
	var aftsql string
	var cnt string

	flag.StringVar(&sql, "sql", "data.sql", "Import sql File Name")
	flag.StringVar(&aftsql, "aftsql", "data_aft.sql", "Export Sql File Name")
	flag.StringVar(&cnt, "cnt", "1001", "Number of combined records ")
	cntInt, _ := strconv.ParseInt(cnt, 10, 64)
	if cntInt == 0 {
		cntInt = 1000
	}
	flag.Parse()
	details := make(map[string][]string, 0)

	c, err := os.OpenFile(sql, os.O_SYNC|os.O_RDONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer c.Close()
	s, err := os.OpenFile(aftsql, os.O_SYNC|os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		panic(err)
	}
	//defer s.Sync()
	defer s.Close()

	rd := bufio.NewReader(c)
	wt := bufio.NewWriter(s)
	for {
		line, err := rd.ReadString('\n')
		if err != nil && len(line) < 10 {
			break
		}
		line = strings.ReplaceAll(line, "\n", "")
		arg := strings.Split(line, "(")
		if len(arg) < 3 {
			continue
		}
		values := strings.Split(arg[2], ";")
		if details[arg[0]] == nil {
			details[arg[0]] = append(details[arg[0]], arg[0]+"("+arg[1])
			details[arg[0]] = append(details[arg[0]], "("+values[0])
			//details[arg[0]] = fmt.Sprintf("%s%s", details[arg[0]], arg[1])
		} else {
			details[arg[0]] = append(details[arg[0]], ",("+values[0])
		}
		//values := strings.Split(arg[2], ";")
		//details[arg[0]] = append(details[arg[0]], ",("+values[0])
		//details[arg[0]] = fmt.Sprintf("%s,%s", details[arg[0]], values[0])
		if len(details[arg[0]]) == int(cntInt) {
			str := fmt.Sprintf("%s %s %s", details[arg[0]], ";", "\n")
			str = strings.ReplaceAll(str, "[", "")
			str = strings.ReplaceAll(str, "]", "")
			_, e := wt.Write([]byte(str))
			if e != nil && e != err {
				panic(e)
			}
			details[arg[0]] = nil
			wt.Flush()
		}
		if io.EOF == err {
			break
		}
	}
	fmt.Println("read eof")
	for _, rm := range details {
		if rm != nil {
			str := fmt.Sprintf("%s %s %s", rm, ";", "\n")
			str = strings.ReplaceAll(str, "[", "")
			str = strings.ReplaceAll(str, "]", "")
			_, e := wt.Write([]byte(str))
			if e != err {
				panic(e)
			}
			wt.Flush()
		}
	}
	fmt.Println("write eof")

	//for _, slice := range dbs {
	//	wt.Write([]byte(fmt.Sprintf("USE %s; COMMIT; \n", slice)))
	//	wt.Flush()
	//}
	wt.Write([]byte(fmt.Sprintf("COMMIT; \n")))
	wt.Flush()
}

func preCsvField(fd string) string {
	return strings.Trim(strings.ReplaceAll(fd, "\"", ""), " ")
}
