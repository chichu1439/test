package main

import (
	"bufio"
	"fmt"
	"os"
)

var dpCnt = 100
var ipCntPerDp = 10

func main() {
	s, err := os.OpenFile("txt", os.O_SYNC|os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		panic(err)
	}
	sql, err := os.OpenFile("sql", os.O_SYNC|os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		panic(err)
	}
	//defer s.Sync()
	defer s.Close()
	defer sql.Close()
	wt := bufio.NewWriter(s)
	wsql := bufio.NewWriter(sql)

	for i := 0; i < dpCnt; i++ {
		d := i % 6
		d = d + 1
		a := fmt.Sprintf("SADD REL.DP10000%04d.156 DP10000%04d\n", i, i)
		_, e := wt.Write([]byte(a))
		if e != nil {
			fmt.Println(e)
		}
		bal := fmt.Sprintf("SET BAL.DP10000%04d.156 100000000\n", i)
		_, e = wt.Write([]byte(bal))
		if e != nil {
			fmt.Println(e)
		}
		sqlBal := fmt.Sprintf("('DP10000%04d', '156', 'DP', NULL, 'DP100000000', 100000000.00, 100000000.00, 'N', 0.00, 0.00, '2021-06-16 14:39:18', '2021-06-16 14:39:18'),\n", i)
		//sqlBal := fmt.Sprintf("\"FPSUMG\",\"MSG\",\"MSG\",\"DP10000%04d\",\"FPSUMG%02d\",\"1\"\n", i, d)
		//sqlBal := fmt.Sprintf("hset MSG.MSG.DP10000%04d FPSUMG FPSUMG%02d\n", i, d)
		_, e = wsql.Write([]byte(sqlBal))
		if e != nil {
			fmt.Println(e)
		}
		for j := 0; j < ipCntPerDp; j++ {
			b := fmt.Sprintf("SADD REL.DP10000%04d.156 IP10000%04d\n", i, i*ipCntPerDp+j)
			_, e := wt.Write([]byte(b))
			if e != nil {
				fmt.Println(e)
			}
			bal := fmt.Sprintf("SET BAL.IP10000%04d.156 100000\n", i*ipCntPerDp+j)
			_, e = wt.Write([]byte(bal))
			if e != nil {
				fmt.Println(e)
			}
			sqlBal := fmt.Sprintf("('IP10000%04d', '156', 'IP', 'DP10000%04d', 'DP10000%04d156', 100000.00, 100000.00, 'N', 0.00, 0.00, '2021-06-16 14:39:18', '2021-06-16 14:39:18'),\n", i*ipCntPerDp+j, i, i)
			//sqlBal := fmt.Sprintf("\"FPSUMG\",\"MSG\",\"MSG\",\"IP10000%04d\",\"FPSUMG%02d\",\"1\"\n", i*ipCntPerDp+j, d)
			//sqlBal := fmt.Sprintf("hset MSG.MSG.IP10000%04d FPSUMG FPSUMG%02d\n", i*ipCntPerDp+j, d)
			_, e = wsql.Write([]byte(sqlBal))
			if e != nil {
				fmt.Println(e)
			}
		}
		wt.Flush()
		wsql.Flush()
	}
	wt.Flush()
	wsql.Flush()
}
