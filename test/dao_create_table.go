package main
// // package main
// import (
// 	"fmt"
// 	"git.sirius.io/banking/test/dao"
// 	_ "github.com/go-sql-driver/mysql"
// 	"github.com/xormplus/xorm"
// 	//"github.com/astaxie/beego/orm"
// )

// func main() {
// 	engine, err := xorm.NewEngine("mysql",
// 		"root:root@/payment?charset=utf8")
// 	if err != nil {
// 		fmt.Println(err.Error())
// 		return
// 	}
// 	err = engine.Sync(new(dao.TNettingRecordInfo))
// 	//err = engine.Sync(new(dao.TPartBalanceFlow))

// 	if err != nil {
// 		fmt.Println(err.Error())
// 	}
// }
