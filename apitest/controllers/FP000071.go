package controllers

import (
	"github.com/astaxie/beego"
)

// tcpSending
type FP000071Controller struct {
	beego.Controller
}

// @Title
// @Description tcpSending
// @Param body  body  models.FP000071I "Input parameter"
// @Success 200  {object} models.FP000071O  successful operation
// Failure 403 {"code":"error code","message":"error message","response":null}
// @router / [post]
func (c *FP000071Controller) Post()() {
}
