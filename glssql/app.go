package main

import (
	"bufio"
	"dlsapp/glssql/impl/dlsredis"
	"dlsapp/glssql/impl/encry"
	"dlsapp/glssql/impl/hash"
	//"dlsapp/mutask/muti"
	"flag"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"io"
	"os"
	"strings"
)

type DlsTreeSlice struct {
	DbHashring *hash.HashRing
	DbMaps     map[string]int
	DbSlices   map[string]DlsSliceDatabase
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
	TblHashring *hash.HashRing
	DbSlice     DlsSliceElement
	TblMaps     map[string]int
	Children    *DlsTreeSlice
}

func main() {
	// main args
	var csv string
	var sql string
	var dburl string
	var dbusr string
	var dbpwd string
	var encryKey string
	var rdflag bool
	var rdtype string
	var rdaddr string
	var rdpwd string
	var dls_slice_spots uint = 100
	var cifdbnum = 3
	var deleteflag = false
	var encryType = "CBC"
	var encryVector = true


	flag.StringVar(&csv, "csv", "data.csv", "Import Csv File Name")
	flag.StringVar(&sql, "sql", "data.sql", "Export Sql File Name")
	flag.StringVar(&dburl, "url", "127.0.0.1:3306", "MySql Database Url")
	flag.StringVar(&dbusr, "usr", "root", "MySql Database User")
	flag.StringVar(&dbpwd, "pwd", "123456", "MySql Database Password")
	flag.StringVar(&encryKey, "key", "", "Encry Key")
	flag.StringVar(&encryType, "enc", "CBC", "Encry Type")
	flag.BoolVar(&encryVector, "vec", false, "Encry Vectory(true=enc/false=dec)")
	flag.StringVar(&rdtype, "redistype", "single", "Redis Type(single/cluster)")
	flag.StringVar(&rdaddr, "redisaddress", "127.0.0.1:6379", "Redis Url")
	flag.StringVar(&rdpwd, "redispwd", "123456", "Redis Password")
	flag.BoolVar(&rdflag, "redisflag", false, "Load Redis Flag(true/false)")
	flag.BoolVar(&deleteflag, "delete", false, "Delete Redis and Sql(true/false)")

	flag.Parse()

	// encry
	if len(strings.Trim(encryKey, " ")) > 0 {
		encry.KmsCache = &encry.KmsKey{KeyValue: strings.Trim(encryKey, " "), EncryType: encryType}
		if encryVector {
			encry.ServiceConf.EncryDlsKeyFlag = true
		} else {
			encry.ServiceConf.DecryPubKeyFlag = true
		}
	}

	var dbList = make([]string, 0)
	var dbs = make(map[string]string, 0)
	for i := 1; i <= cifdbnum; i++ {
		dbList = append(dbList, fmt.Sprintf("DLS_SLICE_CIF_DB_%d|%s|dls_%d|%s|%s|10|dls_cif_|gls_data_", i, dburl, i, dbusr, dbpwd))
		dbs[fmt.Sprintf("DLS_SLICE_CIF_DB_%d", i)] = fmt.Sprintf("gls_%d", i)
	}
	var tree = DlsTreeSlice{
		DbMaps: make(map[string]int, 0),
		DbSlices: make(map[string]DlsSliceDatabase, 0),
	}
	tree.DbHashring = hash.NewHashRing(len(dbList) * (int)(dls_slice_spots))
	for _, str := range dbList {
		ss := strings.Split(str, "|")

		tree.DbMaps[ss[0]] = 1
		slc := 10 //muti.Atoi(ss[5])
		tbm := make(map[string] int, 0)
		for i := 0; i < slc; i++ {
			tbm[fmt.Sprintf("%s%d", ss[6], i)] = 1
		}
		hs := hash.NewHashRing(slc * (int)(dls_slice_spots))
		hs.AddNodes(tbm)
		tree.DbSlices[ss[0]] = DlsSliceDatabase{
			TblHashring: hs,
			DbSlice:     DlsSliceElement{},
			TblMaps:     tbm,
			Children:    nil,
		}
	}
	tree.DbHashring.AddNodes(tree.DbMaps)

	c, err := os.OpenFile(csv, os.O_SYNC|os.O_RDONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer c.Close()
	s, err := os.OpenFile(sql, os.O_SYNC|os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		panic(err)
	}
	//defer s.Sync()
	defer s.Close()

	//var pipe *redis.Pipeline
	if rdflag {
		// regist redis
		var rds = dlsredis.DlsQueryConfig{
			SessionName:     "",
			HostOrg:         "",
			HostDcn:         "",
			HostTopicId:     "",
			QueryExist:      "",
			RedisType:       rdtype,
			RedisAddr:       rdaddr,
			RedisPoolNum:    10,
			RedisPassword:   rdpwd,
			RedisMultiFlag:  false,
			RedisReadonly:   false,
			TopicDcnTitle:   "",
		}

		dlsredis.DlsConfig = rds

		err = dlsredis.RegisterRedis()
		if err != nil {
			panic(err)
		}
		//pipe = dlsredis.GetPipe()
		//defer dlsredis.DeferPipe(pipe)
	}

	rd := bufio.NewReader(c)
	wt := bufio.NewWriter(s)
	cnt := 0
	for {
		line, err := rd.ReadString('\n')
		if err != nil && len(line) < 10 {
			break
		}
		if len(line) < 10 {
			continue
		}
		line = strings.ReplaceAll(line, "\n", "")
		arg := strings.Split(line, ",")
		if len(arg) < 8 {
			continue
		}
		var dlsKey = arg[6]
		if encry.ServiceConf.EncryDlsKeyFlag {
			dlsKey = encry.KmsEncrypt(dlsKey)
		} else if encry.ServiceConf.DecryPubKeyFlag {
			dlsKey = encry.KmsDecrypt(dlsKey)
		}
		key := fmt.Sprintf("%s.%s.%s.%s.%s.%s.%s",
			preCsvField(arg[0]),
			preCsvField(arg[1]),
			preCsvField(arg[2]),
			preCsvField(arg[3]),
			preCsvField(arg[4]),
			preCsvField(arg[5]),
			preCsvField(dlsKey))
		db, tbl := DlsSliceDbTable(&tree, key)
		str := ""
		//fmt.Println(str)
		if deleteflag {
			str = fmt.Sprintf("DELETE FROM %s.%s WHERE DCN_TYPE = '%s' AND C_TYPE = '%s' AND C_CLASS = '%s' AND C_KEY = '%s'; \n",
				dbs[db], tbl, preCsvField(arg[0]), preCsvField(arg[1]), preCsvField(arg[2]), preCsvField(dlsKey))
		} else {
			str = fmt.Sprintf("INSERT INTO %s.%s(tenant, workspace, environment, sutype, gls_type, gls_class, gls_key, gls_value, effect) VALUES('%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', %s); \n",
				dbs[db], tbl, preCsvField(arg[0]), preCsvField(arg[1]), preCsvField(arg[2]), preCsvField(arg[3]), preCsvField(arg[4]), preCsvField(arg[5]), preCsvField(dlsKey), preCsvField(arg[7]), preCsvField(arg[8]))
		}
		if len(strings.Trim(str, " ")) > 10 {
			_, e := wt.Write([]byte(str))
			if e != nil {
				fmt.Println(e)
			}
		}
		wt.Flush()
		if io.EOF == err {
			break
		}
		//if rdflag {
		//	if deleteflag {
		//		deleteElementRedis(client.DlsElement{
		//			ElementType:  preCsvField(arg[1]),
		//			ElementClass: preCsvField(arg[2]),
		//			ElementId:    preCsvField(arg[3]),
		//		}, preCsvField(arg[0]))
		//	} else {
		//		writeElementRedis(pipe, client.DlsElement{
		//			ElementType:  preCsvField(arg[1]),
		//			ElementClass: preCsvField(arg[2]),
		//			ElementId:    preCsvField(arg[3]),
		//		}, preCsvField(arg[0]), preCsvField(arg[4]))
		//	}
		//}
		cnt += 1
		if cnt % 10000 == 0 {
			for slice, _ := range tree.DbMaps {
				wt.Write([]byte(fmt.Sprintf("USE %s; COMMIT; \n", slice)))
				wt.Flush()
			}
			fmt.Println("Sql Count Line Number ", cnt)
			if rdflag {
				//pipe.Exec()
			}
		}
	}
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

func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func HashSlice(hash *hash.HashRing, arg string) string {
	node := strings.Split(hash.GetNode(arg), ".")
	return node[0]
}

func DlsSliceDbTable(tree *DlsTreeSlice, key string) (db string, tbl string) {
	// 获取Hash值并分配数据库分片
	db = HashSlice(tree.DbHashring, key)
	if len(db) <= 0 {
		return
	}
	ds := tree.DbSlices[db]
	gra := Reverse(key)
	tbl = HashSlice(ds.TblHashring, gra)
	return
}
//
//func writeElementRedis(pipe *redis.Pipeline, element client.DlsElement, dcntype, dcngroup string) int {
//	err := dlsredis.SetPipeValue(pipe, fmt.Sprintf("%s.%s.%s", element.ElementType, element.ElementClass, element.ElementId), dcntype, dcngroup)
//	if err != nil {
//		fmt.Println(err)
//		return 0
//	}
//	return 1
//}
//
//func deleteElementRedis(element client.DlsElement, dcntype string) int {
//	err := dlsredis.HDelKey(fmt.Sprintf("%s.%s.%s", element.ElementType, element.ElementClass, element.ElementId), dcntype)
//	if err != nil {
//		fmt.Println(err)
//		return 0
//	}
//	return 1
//}