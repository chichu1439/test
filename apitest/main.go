package main

import (
	_ "apitest/routers"
	"github.com/astaxie/beego"
)

// @Title getStaticBlock
// @Description get all the staticblock by key
// @Param   key     models.IL166682I  true        "The email for login"
// @Success 200 {object} models.IL166682O
// @Failure 400 Invalid email supplied
// @Failure 404 User not found
// @router /staticblock/:key [get]
func main() {
	beego.BConfig.WebConfig.DirectoryIndex = true
	beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	beego.Run()
}
