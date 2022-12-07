package controllers

import (
	"github.com/astaxie/beego"
)

// 8583 lookup and account verify
type FP010016Controller struct {
	beego.Controller
}

// @Title
// @Description 8583 lookup and account verify
// @Param body  body  models.FP010016I "Input parameter"
// @Success 200  {object} models.FP010016O  successful operation
// Failure 403 {"code":"error code","message":"error message","response":null}
// @router / [post]
func (c *FP010016Controller) Post() {
}
