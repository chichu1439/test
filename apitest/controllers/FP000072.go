package controllers

import (
	"github.com/astaxie/beego"
)

// tcpReceive
type FP000072Controller struct {
	beego.Controller
}

// @Title
// @Description tcpReceive
// @Param body  body  models.FP000072I "Input parameter"
// @Success 200  {object} models.FP000072O  successful operation
// Failure 403 {"code":"error code","message":"error message","response":null}
// @router / [post]
func (c *FP000072Controller) Post()() {
}
