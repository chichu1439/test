package controllers

import (
	"github.com/astaxie/beego"
)

// credit-transfer detail
type FP010006Controller struct {
	beego.Controller
}

// @Title
// @Description credit-transfer detail
// @Param body  body  models.FP010006I "Input parameter"
// @Success 200  {object} models.FP010006O  successful operation
// Failure 403 {"code":"error code","message":"error message","response":null}
// @router / [post]
func (c *FP010006Controller) Post()() {
}
