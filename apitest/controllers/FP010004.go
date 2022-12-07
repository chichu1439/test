package controllers

import (
	"github.com/astaxie/beego"
)

// payment return
type FP010004Controller struct {
	beego.Controller
}

// @Title
// @Description payment return
// @Param body  body  models.FPSPacs004 "Input parameter"
// @Success 200  {object} models.FPSPacs002  successful operation
// Failure 403 {"code":"error code","message":"error message","response":null}
// @router / [post]
func (c *FP010004Controller) Post() {
}
